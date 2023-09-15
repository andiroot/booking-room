package handler

import (
	"booking-room/schedule"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
)

type scheduleHandler struct {
	scheduleService schedule.Service
}

func NewScheduleHandler(service schedule.Service) *scheduleHandler {
	return &scheduleHandler{service}
}

func (h *scheduleHandler) PostScheduleHandler(c *gin.Context) {
	var scheduleRequest schedule.ScheduleRequest

	err := c.ShouldBindJSON(&scheduleRequest)
	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errorMessages := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field: %s,condition: %s", e.Field(), e.ActualTag())

				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		case *json.UnmarshalTypeError:
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": err.Error(),
			})
			return
		}
	}
	jwtClaims, _ := c.Get("jwtClaims")
	claims, _ := jwtClaims.(jwt.MapClaims)
	userID, _ := claims["sub"].(float64)
	book, err := h.scheduleService.Create(scheduleRequest, uint(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (h *scheduleHandler) GetSchedules(c *gin.Context) {
	jwtClaims, _ := c.Get("jwtClaims")
	claims, _ := jwtClaims.(jwt.MapClaims)
	userID, _ := claims["sub"].(uint)
	schedules, err := h.scheduleService.FindAllByUser(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	var schedulesResponse []schedule.ScheduleResponse
	for _, b := range schedules {
		scheduleResponse := schedule.ConvertToScheduleResponse(b)
		schedulesResponse = append(schedulesResponse, scheduleResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": schedulesResponse,
	})
}

func (h *scheduleHandler) GetSchedule(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("schedule_id"))
	b, err := h.scheduleService.FindByID(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	scheduleResponse := schedule.ConvertToScheduleResponse(b)
	c.JSON(http.StatusBadRequest, gin.H{
		"data": scheduleResponse,
	})
}

func (h *scheduleHandler) UpdateScheduleHandler(c *gin.Context) {
	var scheduleRequest schedule.ScheduleRequest
	err := c.ShouldBindJSON(&scheduleRequest)
	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errorMessages := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field: %s,condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		case *json.UnmarshalTypeError:
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": err.Error(),
			})
			return
		}
	}

	ID, _ := strconv.Atoi(c.Param("schedule_id"))
	b, err := h.scheduleService.Update(ID, scheduleRequest)
	bookResponse := schedule.ConvertToScheduleResponse(b)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *scheduleHandler) DeleteSchedule(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	b, err := h.scheduleService.Delete(ID)
	scheduleResponse := schedule.ConvertToScheduleResponse(b)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": scheduleResponse,
	})
}
