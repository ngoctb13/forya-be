package handler

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/forya-be/handler/models/request"
	"github.com/ngoctb13/forya-be/handler/models/response"
	"github.com/ngoctb13/forya-be/internal/domains/inputs"
	"github.com/ngoctb13/forya-be/pkg/csv"
)

func (h *Handler) CreateStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &request.CreateStudentRequest{}
		if err := c.ShouldBind(req); err != nil {
			log.Printf("parse request error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := req.Validate()
		if err != nil {
			log.Printf("validate request error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = h.student.CreateStudent(c, &inputs.CreateStudentInput{
			FullName:          req.FullName,
			Age:               req.Age,
			PhoneNumber:       req.PhoneNumber,
			ParentPhoneNumber: req.ParentPhoneNumber,
			Note:              req.Note,
		})

		if err != nil {
			log.Printf("CreateStudentUsescase fail with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "create student successfully"})
	}
}

func (h *Handler) ImportStudentsCSVFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing file"})
			return
		}

		f, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot open file"})
			return
		}
		defer f.Close()

		var reqArr []*request.CreateStudentRequest
		errMses, total, err := csv.ReadCSV(f, 5, true, func(record []string, line int) error {
			age, innerErr := strconv.Atoi(record[1])
			if innerErr != nil {
				return innerErr
			}

			req := &request.CreateStudentRequest{
				FullName:          record[0],
				Age:               age,
				PhoneNumber:       record[2],
				ParentPhoneNumber: record[3],
				Note:              record[4],
			}
			innerErr = req.Validate()
			if innerErr != nil {
				return innerErr
			}

			reqArr = append(reqArr, req)

			return nil
		})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if len(errMses) > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "failed validation",
				"total":   total,
				"errors":  errMses,
			})
			return
		}

		var inputArr []*inputs.CreateStudentInput

		for _, req := range reqArr {
			input := &inputs.CreateStudentInput{
				FullName:          req.FullName,
				Age:               req.Age,
				PhoneNumber:       req.PhoneNumber,
				ParentPhoneNumber: req.ParentPhoneNumber,
				Note:              req.Note,
			}

			inputArr = append(inputArr, input)
		}

		err = h.student.CreateStudents(c, inputArr)
		if err != nil {
			log.Printf("CreateStudentsUsescase fail with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "import students successfully"})
	}
}

func (h *Handler) ListClassStudents() gin.HandlerFunc {
	return func(c *gin.Context) {
		var queryOpts struct {
			joinedAt *time.Time
			leftAt   *time.Time
		}
		classID := c.Param("classId")
		if classID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid classId"})
			return
		}

		joinedAfterStr := c.Query("joined_after")
		if joinedAfterStr != "" {
			t, err := time.Parse(time.RFC3339, joinedAfterStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			queryOpts.joinedAt = &t
		}

		leftAfterStr := c.Query("left_after")
		if leftAfterStr != "" {
			t, err := time.Parse(time.RFC3339, leftAfterStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			queryOpts.leftAt = &t
		}

		studentArr, err := h.student.ListClassStudents(c, &inputs.ListClassStudentsInput{
			ClassID:  classID,
			JoinedAt: queryOpts.joinedAt,
			LeftAt:   queryOpts.leftAt,
		})
		if err != nil {
			log.Printf("ListClassStudentsUsecase fail with error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, studentArr)
	}
}

func (h *Handler) UpdateStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		studentId := c.Param("studentId")
		if studentId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid classId"})
			return
		}

		req := &request.UpdateStudentRequest{}
		if err := c.ShouldBind(req); err != nil {
			log.Printf("parse request error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := req.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		s, err := h.student.UpdateStudent(c, &inputs.UpdateStudentInput{
			StudentID: studentId,
			Fields:    req.Fields,
		})

		if err != nil {
			log.Printf("UpdateStudentUsecase fail with error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, s)
	}
}

func (h *Handler) ListStudents() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &request.ListStudentsRequest{}
		if err := c.ShouldBindQuery(req); err != nil {
			log.Printf("parse request error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		students, pagination, err := h.student.ListStudents(c, &inputs.ListStudentsInput{
			FullName:          req.FullName,
			AgeMin:            req.AgeMin,
			AgeMax:            req.AgeMax,
			PhoneNumber:       req.PhoneNumber,
			ParentPhoneNumber: req.ParentPhoneNumber,
			Page:              req.Page,
			Limit:             req.Limit,
		})
		if err != nil {
			log.Printf("ListStudentsUsecase fail with error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, response.ToListStudentsResponse(students, pagination))
	}
}
