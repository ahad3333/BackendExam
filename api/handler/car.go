package handler

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"add/models"

	"github.com/gin-gonic/gin"
)

// CreateCar godoc
// @ID CreateCar
// @Router /car [POST]
// @Summary CreateCar
// @Description CreateCar
// @Tags Car
// @Accept json
// @Produce json
// @Param car body models.UpdateCarSwag true "CreateCarRequestBody"
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
		log.Println("error whiling create Car:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storage.Car().GetByID(context.Background(), &models.CarPrimeryKey{
		Id: id,
	})
	if err != nil {
		log.Println("error whiling get by id Client in CreatCar:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetByIDCar godoc
// @ID Get_By_IDCar
// @Router /car/{id} [GET]
// @Summary GetByID Car
// @Description GetByID Car
// @Tags Car
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 201 {object} models.Car "GetByIDCarBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetByIDCar(c *gin.Context) {

	id := c.Param("id")

	res, err := h.storage.Car().GetByID(context.Background(), &models.CarPrimeryKey{
		Id: id,
	})

	if err != nil {
		log.Println("error whiling get by id Car:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetListCar godoc
// @ID CarPrimeryKey
// @Router /car [GET]
// @Summary Get List Car
// @Description Get List Car
// @Tags Car
// @Accept json
// @Produce json
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Param search query string false "search"
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
		search    = c.Query("search")
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

	res, err := h.storage.Car().GetList(context.Background(),&models.GetListCarRequest{
		Offset: int64(offset),
		Limit:  int64(limit),
		Search: search,
	})

	if err != nil {
		log.Println("error whiling get list Car:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// UpdateCar godoc
// @ID UpdateCar
// @Router /car/{id} [PUT]
// @Summary Update Car
// @Description Update Car
// @Tags Car
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param car body models.UpdateCarSwag true "UpdateCarRequestBody"
// @Success 202 {object} models.Car "UpdateCarBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) UpdateCar(c *gin.Context) {

	var (
		car models.UpdateCar
	)


	err := c.ShouldBindJSON(&car)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	car.Id = c.Param("id")
	 err = h.storage.Car().Update(context.Background(),&models.UpdateCar{
	 	Id:           car.Id,
	 	State_number: car.State_number,
	 	Model:        car.Model,
	 	Price:        car.Price,
	 	Daily_limit:  car.Daily_limit,
	 	Over_limit:   car.Over_limit,
	 	Investor_id:  car.Investor_id,
	 	Km:           car.Km,
	 	UpdatedAt:    car.UpdatedAt,
	 })
	if err != nil {
		log.Printf("error whiling update Car: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update").Error())
		return
	}


	resp, err := h.storage.Car().GetByID(context.Background(),&models.CarPrimeryKey{Id: car.Id})
	if err != nil {
		log.Printf("error whiling get by id Car: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

// DeleteCar godoc
// @ID DeleteCar
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

	err := h.storage.Car().Delete(context.Background(),&models.CarPrimeryKey{Id: id})
	if err != nil {
		log.Println("error whiling delete  Car:", err.Error())
		c.JSON(http.StatusNoContent, err.Error())
		return
	}
	c.JSON(http.StatusAccepted, "delete Car")
}
