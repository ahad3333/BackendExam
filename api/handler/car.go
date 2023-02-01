package handler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"app/models"

	"github.com/gin-gonic/gin"
)

// CreateCar godoc
// @ID create_car
// @Router /car [POST]
// @Summary Create Car
// @Description Create Car
// @Tags Car
// @Accept json
// @Produce json
// @Param Car body models.CreateCar true "CreateCarRequestBody"
// @Success 201 {object} models.Car "GetCarBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) CreateCar(c *gin.Context) {

	var car models.CreateCar

	err := c.ShouldBindJSON(&car)
	if err != nil {
		log.Println("error whiling marshal json:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Car().Insert(context.Background(), &car)
	if err != nil {
		log.Println("error whiling create car:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res, err := h.storage.Car().GetByID(context.Background(), &models.CarPrimeryKey{
		Id: id,
	})
	if err != nil {
		log.Println("error whiling get by id Car:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = h.cache.CarCache().Delete()
	if err != nil {
		log.Println("error whiling delete cache Car:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetByIDCar godoc
// @ID get_by_id_car
// @Router /car/{id} [GET]
// @Summary Get By ID Car
// @Description Get By ID Car
// @Tags Car
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Car "GetCarBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetByIdCar(c *gin.Context) {

	id := c.Param("id")

	res, err := h.storage.Car().GetByID(context.Background(), &models.CarPrimeryKey{
		Id: id,
	})
	if err != nil {
		log.Println("error whiling get by id Car:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetListCar godoc
// @ID get_list_car
// @Router /car [GET]
// @Summary Get List Car
// @Description Get List Car
// @Tags Car
// @Accept json
// @Produce json
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Success 200 {object} models.GetListCarResponse "GetCarListBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetListCar(c *gin.Context) {
	var (
		err       error
		offset    int
		limit     int
		offsetStr = c.Query("offset")
		limitStr  = c.Query("limit")
	)

	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			log.Println("error whiling offset:", err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			log.Println("error whiling limit:", err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	exists, err := h.cache.CarCache().Exists()
	if err != nil {
		log.Println("error whiling cache exists Car:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var res *models.GetListCarResponse
	if !exists {
		fmt.Println("Postgres")
		res, err = h.storage.Car().GetList(context.Background(), &models.GetListCarRequest{
			Offset: int64(offset),
			Limit:  int64(limit),
		})
		if err != nil {
			log.Println("error whiling get list Car:", err.Error())
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		err = h.cache.CarCache().Create(res)
		if err != nil {
			log.Println("error whiling create cache Car:", err.Error())
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		fmt.Println("Redis")
		res, err = h.cache.CarCache().GetList()
		if err != nil {
			log.Println("error whiling get list cache Car:", err.Error())
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusCreated, res)
}

// UpdateCar godoc
// @ID update_car
// @Router /car/{id} [PUT]
// @Summary Update Car
// @Description Update Car
// @Tags Car
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param Car body models.UpdateCarSwag true "UpdateCarRequestBody"
// @Success 202 {object} models.Car "UpdateCarBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) UpdateCar(c *gin.Context) {

	var (
		Car models.UpdateCar
	)

	Car.Id = c.Param("id")

	err := c.ShouldBindJSON(&Car)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := h.storage.Car().Update(context.Background(), &Car)

	if err != nil {
		log.Printf("error whiling update: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update").Error())
		return
	}

	if rowsAffected <= 0 {
		log.Printf("error whiling update: %v", sql.ErrNoRows)
		c.JSON(http.StatusInternalServerError, sql.ErrNoRows.Error())
		return
	}

	resp, err := h.storage.Car().GetByID(context.Background(), &models.CarPrimeryKey{
		Id: Car.Id,
	})
	if err != nil {
		log.Printf("error whiling get by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}
	err = h.cache.CarCache().Delete()
	if err != nil {
		log.Println("error whiling delete cache Car:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, resp)
}

// DeleteCar godoc
// @ID delete_car
// @Router /car/{id} [DELETE]
// @Summary Delete Car
// @Description Delete Car
// @Tags Car
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 204 {object} models.Empty "DeleteCarBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) DeleteCar(c *gin.Context) {
	id := c.Param("id")

	err := h.storage.Car().Delete(context.Background(), &models.CarPrimeryKey{Id: id})
	if err != nil {
		log.Println("error whiling delete car:", err.Error())
		c.JSON(http.StatusNoContent, err.Error())
		return
	}
	err = h.cache.CarCache().Delete()
	if err != nil {
		log.Println("error whiling delete cache Car:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, "Car deleted")
}
