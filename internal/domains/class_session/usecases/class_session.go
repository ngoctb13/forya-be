package usecases

import (
	"context"
	"errors"
	"fmt"

	"github.com/shopspring/decimal"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	classRepo "github.com/ngoctb13/forya-be/internal/domains/class/repos"
	classSessionRepo "github.com/ngoctb13/forya-be/internal/domains/class_session/repos"
	courseStudentRepo "github.com/ngoctb13/forya-be/internal/domains/course_student/repos"
	"github.com/ngoctb13/forya-be/internal/domains/inputs"
	supplyBatchRepo "github.com/ngoctb13/forya-be/internal/domains/supply_batch/repos"
	supplyUsageRepo "github.com/ngoctb13/forya-be/internal/domains/supply_usage/repos"
)

type ClassSession struct {
	classSessionRepo        classSessionRepo.IClassSession
	classRepo               classRepo.IClassRepo
	classSessionAttendances classSessionRepo.IClassSessionAttendance
	courseStudentRepo       courseStudentRepo.ICourseStudentRepo
	supplyBatchRepo         supplyBatchRepo.ISupplyBatch
	supplyUsageRepo         supplyUsageRepo.ISupplyUsage
}

type supplyUsageRequest struct {
	courseStudentID   string
	supplyID          string
	quantity          int
	hasUnitPrice      bool
	unitPriceOverride decimal.Decimal
}

func NewClassSession(
	classSessionRepo classSessionRepo.IClassSession,
	classRepo classRepo.IClassRepo,
	attendanceRepo classSessionRepo.IClassSessionAttendance,
	courseStudentRepo courseStudentRepo.ICourseStudentRepo,
	supplyBatchRepo supplyBatchRepo.ISupplyBatch,
	supplyUsageRepo supplyUsageRepo.ISupplyUsage,
) *ClassSession {
	return &ClassSession{
		classSessionRepo:        classSessionRepo,
		classRepo:               classRepo,
		classSessionAttendances: attendanceRepo,
		courseStudentRepo:       courseStudentRepo,
		supplyBatchRepo:         supplyBatchRepo,
		supplyUsageRepo:         supplyUsageRepo,
	}
}

func (cl *ClassSession) CreateClassSession(ctx context.Context, input *inputs.CreateClassSessionInput) error {
	session := &models.ClassSession{
		Name:    input.Name,
		ClassID: input.ClassID,
		HeldAt:  input.HeldAt,
	}

	return cl.classSessionRepo.Create(ctx, session)
}

func (cl *ClassSession) ListClassSessions(ctx context.Context, input *inputs.ListClassSessionsInput) ([]*models.ClassSession, *models.Pagination, error) {
	queries := make(map[string]interface{})
	if input.ClassID != nil {
		queries["class_id"] = *input.ClassID
	}
	if input.StartTime != nil {
		queries["start_time"] = *input.StartTime
	}
	if input.EndTime != nil {
		queries["end_time"] = *input.EndTime
	}

	pagination := models.NewPagination(input.Page, input.Limit)

	csArr, p, err := cl.classSessionRepo.List(ctx, queries, pagination)
	if err != nil {
		return nil, nil, err
	}

	classIDSet := make(map[string]bool)
	for _, cs := range csArr {
		if cs.ClassID != "" {
			classIDSet[cs.ClassID] = true
		}
	}

	classIDs := make([]string, 0, len(classIDSet))
	for id := range classIDSet {
		classIDs = append(classIDs, id)
	}

	if len(classIDs) > 0 {
		classMap, err := cl.classRepo.GetClassesByIDs(ctx, classIDs)
		if err != nil {
			return nil, nil, err
		}

		for _, cs := range csArr {
			if class, exists := classMap[cs.ClassID]; exists {
				cs.Class = class
			}
		}
	}

	return csArr, p, nil
}

func (cl *ClassSession) BatchMarkAttendance(ctx context.Context, input *inputs.BatchMarkClassSessionAttendanceInput) error {
	if input == nil {
		return errors.New("input is required")
	}
	if input.ClassSessionID == "" {
		return errors.New("class session id is required")
	}
	if len(input.Attendances) == 0 {
		return errors.New("attendances list cannot be empty")
	}

	session, err := cl.classSessionRepo.GetByID(ctx, input.ClassSessionID)
	if err != nil {
		return err
	}
	if session == nil {
		return errors.New("class session not found")
	}

	attendances := make([]*models.ClassSessionAttendance, 0, len(input.Attendances))
	supplyRequests := make([]supplyUsageRequest, 0)

	for _, att := range input.Attendances {
		if att.CourseStudentID == "" {
			return errors.New("course_student_id is required for all attendance items")
		}

		attendances = append(attendances, &models.ClassSessionAttendance{
			ClassSessionID:  input.ClassSessionID,
			CourseStudentID: att.CourseStudentID,
			IsAttended:      att.IsAttended,
		})

		for _, sup := range att.Supplies {
			if sup.SupplyID == "" {
				return errors.New("supply_id is required for all supply usage items")
			}
			if sup.Quantity <= 0 {
				return errors.New("supply quantity must be greater than zero")
			}

			req := supplyUsageRequest{
				courseStudentID: att.CourseStudentID,
				supplyID:        sup.SupplyID,
				quantity:        sup.Quantity,
			}
			if sup.UnitPrice != nil {
				req.hasUnitPrice = true
				req.unitPriceOverride = decimal.NewFromFloat(*sup.UnitPrice)
			}
			supplyRequests = append(supplyRequests, req)
		}
	}

	var (
		createdUsages    []*models.SupplyUsage
		batchAdjustments map[string]int
	)

	if len(supplyRequests) > 0 {
		var err error
		createdUsages, batchAdjustments, err = cl.buildSupplyUsages(ctx, input.ClassSessionID, supplyRequests)
		if err != nil {
			return err
		}

		if err := cl.supplyUsageRepo.CreateUsagesAndDecreaseStock(ctx, createdUsages, batchAdjustments); err != nil {
			return err
		}
	}

	if err := cl.classSessionAttendances.BatchMarkAttendance(ctx, input.ClassSessionID, attendances); err != nil {
		if len(createdUsages) > 0 && batchAdjustments != nil {
			if rollbackErr := cl.supplyUsageRepo.RollbackUsages(ctx, createdUsages, batchAdjustments); rollbackErr != nil {
				return fmt.Errorf("attendance failed: %v; rollback error: %v", err, rollbackErr)
			}
		}
		return err
	}

	return nil
}

func (cl *ClassSession) buildSupplyUsages(ctx context.Context, classSessionID string, requests []supplyUsageRequest) ([]*models.SupplyUsage, map[string]int, error) {
	courseStudentIDs := make(map[string]struct{})
	supplyIDs := make(map[string]struct{})

	for _, req := range requests {
		courseStudentIDs[req.courseStudentID] = struct{}{}
		supplyIDs[req.supplyID] = struct{}{}
	}

	csIDList := make([]string, 0, len(courseStudentIDs))
	for id := range courseStudentIDs {
		csIDList = append(csIDList, id)
	}

	csMap, err := cl.courseStudentRepo.GetByIDs(ctx, csIDList)
	if err != nil {
		return nil, nil, err
	}

	supplyIDList := make([]string, 0, len(supplyIDs))
	for id := range supplyIDs {
		supplyIDList = append(supplyIDList, id)
	}

	batchesBySupply, err := cl.supplyBatchRepo.ListAvailableBySupplyIDs(ctx, supplyIDList)
	if err != nil {
		return nil, nil, err
	}

	batchAdjustments := make(map[string]int)
	usages := make([]*models.SupplyUsage, 0)

	for _, req := range requests {
		cs := csMap[req.courseStudentID]
		if cs == nil {
			return nil, nil, fmt.Errorf("course_student not found: %s", req.courseStudentID)
		}

		batches := batchesBySupply[req.supplyID]
		if len(batches) == 0 {
			return nil, nil, fmt.Errorf("no stock available for supply %s", req.supplyID)
		}

		qtyLeft := req.quantity
		for _, batch := range batches {
			available := batch.RemainingQuantity - batchAdjustments[batch.ID]
			if available <= 0 {
				continue
			}

			useQty := available
			if qtyLeft < useQty {
				useQty = qtyLeft
			}

			unitPrice := batch.PurchasePrice
			if req.hasUnitPrice {
				unitPrice = req.unitPriceOverride
			}
			totalPrice := unitPrice.Mul(decimal.NewFromInt(int64(useQty)))

			usages = append(usages, &models.SupplyUsage{
				BatchID:        batch.ID,
				StudentID:      cs.StudentID,
				ClassSessionID: classSessionID,
				Quantity:       useQty,
				UnitPrice:      unitPrice,
				TotalPrice:     totalPrice,
			})

			batchAdjustments[batch.ID] += useQty
			qtyLeft -= useQty

			if qtyLeft == 0 {
				break
			}
		}

		if qtyLeft > 0 {
			return nil, nil, fmt.Errorf("insufficient stock for supply %s", req.supplyID)
		}
	}

	return usages, batchAdjustments, nil
}
