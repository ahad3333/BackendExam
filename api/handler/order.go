package handler

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"app/models"

	"github.com/gin-gonic/gin"
)

// CreateOrder godoc
// @ID create_order
// @Router /order [POST]
// @Summary Create Order
// @Description Create Order
// @Tags Order
// @Accept json
// @Produce json
// @Param Order body models.CreateOrder true "CreateOrderRequestBody"
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

	c.JSON(http.StatusCreated, id)
}

// GetByIDOrder godoc
// @ID get_by_id_order
// @Router /order/{id} [GET]
// @Summary Get By ID Order
// @Description Get By ID Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Order "GetOrderBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetByIdOrder(c *gin.Context) {

	id := c.Param("id")

	res, err := h.storage.Order().GetByID(context.Background(), &models.OrderPrimeryKey{
		Id: id,
	})
	if err != nil {
		log.Println("error whiling get by id Order:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetListOrder godoc
// @Security ApiKeyAuth
// @ID get_list_order
// @Router /v1/order [GET]
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

	res, err := h.storage.Order().GetList(context.Background(), &models.GetListOrderRequest{
		Offset: int64(offset),
		Limit:  int64(limit),
	})

	if err != nil {
		log.Println("error whiling get list order:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// UpdateOrder godoc
// @ID update_order
// @Router /order/{id} [PUT]
// @Summary Update Order
// @Description Update Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param Order body models.UpdateOrderSwag true "UpdateOrderRequestBody"
// @Success 202 {object} models.Order "UpdateOrderBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) UpdateOrder(c *gin.Context) {

	var (
		order models.UpdateOrder
	)

	order.Id = c.Param("id")

	err := c.ShouldBindJSON(&order)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.storage.Order().Update(context.Background(), &order)

	if err != nil {
		log.Printf("error whiling update: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update").Error())
		return
	}

	resp, err := h.storage.Order().GetByID(context.Background(), &models.OrderPrimeryKey{
		Id: order.Id,
	})
	if err != nil {
		log.Printf("error whiling get by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}

	c.JSON(http.StatusAccepted, resp)
}

// UpdatePatchOrder godoc
// @ID update_patch_order
// @Router /order/{id} [PATCH]
// @Summary Update Patch Order
// @Description Update Patch Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param Order body models.UpdatePatch true "UpdatePatchRequestBody"
// @Success 202 {object} models.Order "OrderBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) UpdatePatchOrder(c *gin.Context) {

	var (
		order models.UpdatePatch
	)

	order.Id = c.Param("id")

	err := c.ShouldBindJSON(&order)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.storage.Order().UpdatePatch(context.Background(), &order)

	if err != nil {
		log.Printf("error whiling update: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update").Error())
		return
	}

	resp, err := h.storage.Order().GetByID(context.Background(), &models.OrderPrimeryKey{
		Id: order.Id,
	})
	if err != nil {
		log.Printf("error whiling get by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}

	c.JSON(http.StatusAccepted, resp)
}

// DeleteOrder godoc
// @ID delete_order
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

	err := h.storage.Order().Delete(context.Background(), &models.OrderPrimeryKey{Id: id})
	if err != nil {
		log.Println("error whiling delete  order:", err.Error())
		c.JSON(http.StatusNoContent, err.Error())
		return
	}
	c.JSON(http.StatusCreated, "Order deleted")
}
