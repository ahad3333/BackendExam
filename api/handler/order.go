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

// CreateOrder godoc
// @ID CreateOrder
// @Router /order [POST]
// @Summary CreateOrder
// @Description CreateOrder
// @Tags Order
// @Accept json
// @Produce json
// @Param order body models.CreateOrderSwag true "CreateOrderRequestBody"
// @Success 201 {object} models.Order "GetOrderBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) CreateOrder(c *gin.Context) {

	var order models.CreateOrder

	err := c.ShouldBindJSON(&order)
	if err != nil {
		log.Println("error whiling marshal json:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Order().Insert(context.Background(), &order)
	if err != nil {
		log.Println("error whiling create order:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storage.Order().GetByID(context.Background(), &models.OrderPrimeryKey{
		Id: id,
	})
	if err != nil {
		log.Println("error whiling get by id Order", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetByIDOrder godoc
// @ID Get_By_IDOrder
// @Router /order/{id} [GET]
// @Summary GetByID Order
// @Description GetByID Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 201 {object} models.Order "GetByIDCarBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetByIDOrder(c *gin.Context) {

	id := c.Param("id")

	res, err := h.storage.Order().GetByID(context.Background(), &models.OrderPrimeryKey{
		Id: id,
	})

	if err != nil {
		log.Println("error whiling get by id Order:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetListOrder godoc
// @ID OrderPrimeryKey
// @Router /order [GET]
// @Summary Get List Order
// @Description Get List Order
// @Tags Order
// @Accept json
// @Produce json
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Success 200 {object} models.GetListOrderResponse "GetOrderListBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetListOrder(c *gin.Context) {
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

	res, err := h.storage.Order().GetList(context.Background(),&models.GetListOrderRequest{
		Offset: int64(offset),
		Limit:  int64(limit),
	})

	if err != nil {
		log.Println("error whiling get list Order:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// UpdateOrder godoc
// @ID UpdateOrder
// @Router /order/{id} [PUT]
// @Summary Update Order
// @Description Update Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string fl "id"
// @Param order body models.UpdateOrderSwag true "UpdateOrderRequestBody"
// @Success 202 {object} models.Order "UpdateCarBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) UpdateOrder(c *gin.Context) {

	var (
		order models.UpdateOrder
		// order1 models.InvestorBenefit

	)


	err := c.ShouldBindJSON(&order)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	order.Id = c.Param("id")
	 err = h.storage.Order().Update(context.Background(),&models.Order{
	 	Id:          order.Id,
	 	Car_id:      order.Car_id,
	 	Client_id:   order.Client_id,
	 	Total_price: order.Total_price,
	 	Paid_price:  order.Paid_price,
	 	Day_count:   order.Day_count,
	 	Give_km:     order.Give_km,
	 	Receive_km:  order.Receive_km,
	 	UpdatedAt:   order.UpdatedAt,
	 })
	//   err = h.storage.Order().Return(context.Background(), &order1)
	if err != nil {
		log.Println("error whiling create order:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err != nil {
		log.Printf("error whiling update Order: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update").Error())
		return
	}


	resp, err := h.storage.Order().GetByID(context.Background(),&models.OrderPrimeryKey{Id: order.Id})
	if err != nil {
		log.Printf("error whiling get by id Order: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

// DeleteOrder godoc
// @ID DeleteOrder
// @Router /order/{id} [DELETE]
// @Summary Delete Order
// @Description Delete Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 204 {object} models.Empty "DeleteOrderBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")

	err := h.storage.Order().Delete(context.Background(),&models.OrderPrimeryKey{Id: id})
	if err != nil {
		log.Println("error whiling delete  Order:", err.Error())
		c.JSON(http.StatusNoContent, err.Error())
		return
	}
	c.JSON(http.StatusAccepted, "delete Order")
}

// ReturnCar godoc
// @ID ReturnCar
// @Router /return/{id} [PUT]
// @Summary ReturnCar Order
// @Description ReturnCar Order
// @Tags ReturnCar
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param order body models.ReturnCar true "UpdateOrderRequestBody"
// @Success 202 {object} models.Order "UpdateCarBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) UpdateReturn(c *gin.Context) {

	var (
		order models.UpdateOrder
	)


	err := c.ShouldBindJSON(&order)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, errors.New("error whiling update").Error())
		return
	}
	order.Id = c.Param("id")
	 
	 err = h.storage.Order().Return(context.Background(),&models.Order{
		Id: order.Id,
		Car_id: order.Car_id,
		Client_id: order.Client_id,
		Receive_km: order.Receive_km,
	 })
	if err != nil {
		log.Println("error whiling create order:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err != nil {
		log.Printf("error whiling update ReturnCar: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update").Error())
		return
	}


	resp, err := h.storage.Order().GetByID(context.Background(),&models.OrderPrimeryKey{Id: order.Id})
	if err != nil {
		log.Printf("error whiling get by id ReturnCar: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}
	c.JSON(http.StatusAccepted, resp)
}