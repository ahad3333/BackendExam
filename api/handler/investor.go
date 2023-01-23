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

// CreateInvestor godoc
// @ID create_Investor
// @Router /investor [POST]
// @Summary Create Investor
// @Description Create Investor
// @Tags Investor
// @Accept json
// @Produce json
// @Param investor body models.UpdateInvestorSwag true "CreateInvestorRequestBody"
// @Success 201 {object} models.Investor "GetInvestorBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) CreateInvestor(c *gin.Context) {

	var investor models.CreateInvestor

	err := c.ShouldBindJSON(&investor)
	if err != nil {
		log.Println("error whiling marshal json:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Investor().Insert(context.Background(), &investor)
	if err != nil {
		log.Println("error whiling create Investor:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res, err := h.storage.Investor().GetByID(context.Background(), &models.InvestorPrimeryKey{
		Id: id,
	})
	if err != nil {
		log.Println("error whiling get by id Investor:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetByIDInvestor godoc
// @ID get_by_id_investor
// @Router /investor/{id} [GET]
// @Summary Get By ID Investor
// @Description Get By ID Investor
// @Tags Investor
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Investor "GetInvestorBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetByIdInvestor(c *gin.Context) {

	id := c.Param("id")

	res, err := h.storage.Investor().GetByID(context.Background(), &models.InvestorPrimeryKey{
		Id: id,
	})
	if err != nil {
		log.Println("error whiling get by id Investor:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetListInvestor godoc
// @ID get_list_Investor
// @Router /investor [GET]
// @Summary Get List Investor
// @Description Get List Investor
// @Tags Investor
// @Accept json
// @Produce json
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Param search query string false "search"
// @Success 200 {object} models.GetListInvestorResponse "GetInvestorListBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetListInvestor(c *gin.Context) {
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

	res, err := h.storage.Investor().GetList(context.Background(), &models.GetListInvestorRequest{
		Offset: int64(offset),
		Limit:  int64(limit),
		Search: search,
	})

	if err != nil {
		log.Println("error whiling get list Investor:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// UpdateInvestor godoc
// @ID update_investor
// @Router /investor/{id} [PUT]
// @Summary Update Investor
// @Description Update Investor
// @Tags Investor
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param investor body models.UpdateInvestorSwag true "UpdateInvestorRequestBody"
// @Success 202 {object} models.Investor "UpdateInvestorBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) UpdateInvestor(c *gin.Context) {

	var (
		investor models.UpdateInvestor
	)

	id := c.Param("id")

	err := c.ShouldBindJSON(&investor)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	 err = h.storage.Investor().Update(context.Background(),&models.UpdateInvestor{
		Id: id,
		Name: investor.Name,
	})
	if err != nil {
		log.Printf("error whiling update Investor: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update Investor").Error())
		return
	}

	

	resp, err := h.storage.Investor().GetByID(context.Background(),&models.InvestorPrimeryKey{
		Id: id,
	})
	if err != nil {
		log.Printf("error whiling get by id Investor: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}

	c.JSON(http.StatusAccepted, resp)
}

// DeleteInvestor godoc
// @ID delete_Investor
// @Router /investor/{id} [DELETE]
// @Summary Delete Investor
// @Description Delete Investor
// @Tags Investor
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 204 {object} models.Empty "DeleteInvestorBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) DeleteInvestor(c *gin.Context) {

	id := c.Param("id")
	

	err := h.storage.Investor().Delete(context.Background(),&models.InvestorPrimeryKey{Id: id})
	if err != nil {
		log.Println("error whiling delete  Investor:", err.Error())
		c.JSON(http.StatusNoContent, err.Error())
		return
	}
	
	c.JSON(http.StatusNoContent, "Deletet Investor")
}