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

// CreateClient godoc
// @ID CreateClient
// @Router /client [POST]
// @Summary CreateClient
// @Description CreateClient
// @Tags Client
// @Accept json
// @Produce json
// @Param client body models.UpdateClientSwag true "CreateClientRequestBody"
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
		log.Println("error whiling create Client:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storage.Client().GetByID(context.Background(), &models.ClientPrimeryKey{
		Id: id,
	})
	if err != nil {
		log.Println("error whiling get by id Client in CreatClient:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetByIDClient godoc
// @ID Get_By_IDClient
// @Router /client/{id} [GET]
// @Summary GetByID Client
// @Description GetByID Client
// @Tags Client
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 201 {object} models.Client "GetByIDClientBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetByIDClient(c *gin.Context) {

	id := c.Param("id")

	res, err := h.storage.Client().GetByID(context.Background(), &models.ClientPrimeryKey{
		Id: id,
	})

	if err != nil {
		log.Println("error whiling get by id Client:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetListClient godoc
// @ID ClientPrimeryKey
// @Router /client [GET]
// @Summary Get List Client
// @Description Get List Client
// @Tags Client
// @Accept json
// @Produce json
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Param search query string false "search"
// @Success 200 {object} models.GetListClientResponse "GetClientkListBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetListClient(c *gin.Context) {
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

	res, err := h.storage.Client().GetList(context.Background(),&models.GetListClientRequest{
		Offset: int64(offset),
		Limit:  int64(limit),
		Search: search,
	})

	if err != nil {
		log.Println("error whiling get list Client:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// UpdateClient godoc
// @ID UpdateClient
// @Router /client/{id} [PUT]
// @Summary Update Client
// @Description Update Client
// @Tags Client
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param client body models.UpdateClientSwag true "UpdateClientRequestBody"
// @Success 202 {object} models.Client "UpdateClientBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) UpdateClient(c *gin.Context) {

	var (
		client models.UpdateClient
	)


	err := c.ShouldBindJSON(&client)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	client.Id = c.Param("id")
	 err = h.storage.Client().Update(context.Background(),&models.UpdateClient{
		Id: client.Id,
		First_name: client.First_name,
		Last_name: client.Last_name,
		Address: client.Address,
		Phone_number: client.Phone_number,
	})
	if err != nil {
		log.Printf("error whiling update Client: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update").Error())
		return
	}


	resp, err := h.storage.Client().GetByID(context.Background(),&models.ClientPrimeryKey{Id: client.Id})
	if err != nil {
		log.Printf("error whiling get by id Up: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

// DeleteClient godoc
// @ID DeleteClient
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

	err := h.storage.Client().Delete(context.Background(),&models.ClientPrimeryKey{Id: id})
	if err != nil {
		log.Println("error whiling delete  Client:", err.Error())
		c.JSON(http.StatusNoContent, err.Error())
		return
	}
	c.JSON(http.StatusAccepted, "delete Client")
}
