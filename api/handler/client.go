package handler

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strconv"

	"app/models"

	"github.com/gin-gonic/gin"
)

// CreateClient godoc
// @ID create_client
// @Router /client [POST]
// @Summary Create Client
// @Description Create Client
// @Tags Client
// @Accept json
// @Produce json
// @Param Client body models.CreateClient true "CreateClientRequestBody"
// @Success 201 {object} models.Client "GetClientBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) CreateClient(c *gin.Context) {

	var client models.CreateClient

	err := c.ShouldBindJSON(&client)
	if err != nil {
		log.Println("error whiling marshal json:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Client().Insert(context.Background(), &client)
	if err != nil {
		log.Println("error whiling create client:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res, err := h.storage.Client().GetByID(context.Background(), &models.ClientPrimeryKey{
		Id: id,
	})
	if err != nil {
		log.Println("error whiling get by id client:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetByIDClient godoc
// @ID get_by_id_client
// @Router /client/{id} [GET]
// @Summary Get By ID Client
// @Description Get By ID Client
// @Tags Client
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Client "GetClientBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetByIdClient(c *gin.Context) {

	id := c.Param("id")

	res, err := h.storage.Client().GetByID(context.Background(), &models.ClientPrimeryKey{
		Id: id,
	})
	if err != nil {
		log.Println("error whiling get by id Client:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetListClient godoc
// @ID get_list_client
// @Router /client [GET]
// @Summary Get List Client
// @Description Get List Client
// @Tags Client
// @Accept json
// @Produce json
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Success 200 {object} models.GetListClientResponse "GetClientListBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetListClient(c *gin.Context) {
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

	res, err := h.storage.Client().GetList(context.Background(), &models.GetListClientRequest{
		Offset: int64(offset),
		Limit:  int64(limit),
	})

	if err != nil {
		log.Println("error whiling get list client:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// UpdateClient godoc
// @ID update_client
// @Router /client/{id} [PUT]
// @Summary Update Client
// @Description Update Client
// @Tags Client
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param Client body models.UpdateClientSwag true "UpdateClientRequestBody"
// @Success 202 {object} models.Client "UpdateClientBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) UpdateClient(c *gin.Context) {

	var (
		client models.UpdateClient
	)

	client.Id = c.Param("id")

	err := c.ShouldBindJSON(&client)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := h.storage.Client().Update(context.Background(), &client)

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

	resp, err := h.storage.Client().GetByID(context.Background(), &models.ClientPrimeryKey{
		Id: client.Id,
	})
	if err != nil {
		log.Printf("error whiling get by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}

	c.JSON(http.StatusAccepted, resp)
}

// DeleteClient godoc
// @ID delete_client
// @Router /client/{id} [DELETE]
// @Summary Delete Client
// @Description Delete Client
// @Tags Client
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 204 {object} models.Empty "DeleteClientBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) DeleteClient(c *gin.Context) {
	id := c.Param("id")

	err := h.storage.Client().Delete(context.Background(), &models.ClientPrimeryKey{Id: id})
	if err != nil {
		log.Println("error whiling delete  client:", err.Error())
		c.JSON(http.StatusNoContent, err.Error())
		return
	}
	c.JSON(http.StatusCreated, "Client deleted")
}
