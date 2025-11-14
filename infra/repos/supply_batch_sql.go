package repos

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	"gorm.io/gorm"
)

type supplyBatchSQLRepo struct {
	db *gorm.DB
}

func NewSupplyBatchSQLRepo(db *gorm.DB) *supplyBatchSQLRepo {
	return &supplyBatchSQLRepo{db: db}
}

func (r *supplyBatchSQLRepo) Create(ctx context.Context, sb *models.SupplyBatch) error {
	return r.db.WithContext(ctx).Create(sb).Error
}

func (r *supplyBatchSQLRepo) GetByID(ctx context.Context, id string) (*models.SupplyBatch, error) {
	var sb models.SupplyBatch
	if err := r.db.WithContext(ctx).First(&sb, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &sb, nil
}

func (r *supplyBatchSQLRepo) GetByIDs(ctx context.Context, ids []string) (map[string]*models.SupplyBatch, error) {
	result := make(map[string]*models.SupplyBatch)
	if len(ids) == 0 {
		return result, nil
	}

	var batches []*models.SupplyBatch
	if err := r.db.WithContext(ctx).
		Where("id IN ?", ids).
		Find(&batches).Error; err != nil {
		return nil, err
	}

	for _, batch := range batches {
		result[batch.ID] = batch
	}

	return result, nil
}

func (r *supplyBatchSQLRepo) ListAvailableBySupplyIDs(ctx context.Context, supplyIDs []string) (map[string][]*models.SupplyBatch, error) {
	result := make(map[string][]*models.SupplyBatch)
	if len(supplyIDs) == 0 {
		return result, nil
	}

	var batches []*models.SupplyBatch
	if err := r.db.WithContext(ctx).
		Where("supply_id IN ? AND remaining_quantity > 0", supplyIDs).
		Order("supply_id ASC, purchase_date ASC, id ASC").
		Find(&batches).Error; err != nil {
		return nil, err
	}

	for _, batch := range batches {
		result[batch.SupplyID] = append(result[batch.SupplyID], batch)
	}

	return result, nil
}

func (r *supplyBatchSQLRepo) List(ctx context.Context, queries map[string]interface{}, pagination *models.Pagination) ([]*models.SupplyBatch, *models.Pagination, error) {
	query := r.db.WithContext(ctx).Model(&models.SupplyBatch{})

	for k, v := range queries {
		switch k {
		case "supply_id":
			query = query.Where("supply_id = ?", v)
		case "purchase_date_start":
			query = query.Where("purchase_date >= ?", v)
		case "purchase_date_end":
			query = query.Where("purchase_date <= ?", v)
		}
	}

	var (
		total int64
		res   []*models.SupplyBatch
	)

	if err := query.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	pagination.SetTotal(total)

	// Stable default ordering
	query = query.Order("purchase_date DESC")
	query = pagination.ApplyToQuery(query)

	if err := query.Find(&res).Error; err != nil {
		return nil, nil, err
	}

	return res, pagination, nil
}

func (r *supplyBatchSQLRepo) UpdateWithFields(ctx context.Context, sb *models.SupplyBatch, fields map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(sb).Updates(fields).Error
}
