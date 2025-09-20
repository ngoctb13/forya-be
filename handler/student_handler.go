package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/forya-be/handler/models"
	dm "github.com/ngoctb13/forya-be/internal/domain/models"
	"github.com/ngoctb13/forya-be/pkg/csv"
)

func (h *Handler) CreateStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &models.CreateStudentRequest{}
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

		err = h.student.CreateStudent(c, &dm.CreateStudentInput{
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

		var reqArr []*models.CreateStudentRequest
		errMses, total, err := csv.ReadCSV(f, 5, true, func(record []string, line int) error {
			age, innerErr := strconv.Atoi(record[1])
			if innerErr != nil {
				return innerErr
			}

			req := &models.CreateStudentRequest{
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

		var inputArr []*dm.CreateStudentInput

		for _, req := range reqArr {
			input := &dm.CreateStudentInput{
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
