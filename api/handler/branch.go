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

// CreateBranch godoc
// @ID create_Branch
// @Router /branch [POST]
// @Summary Create Branch
// @Description Create Branch
// @Tags Branch
// @Accept json
// @Produce json
// @Param branch body models.CreateBranch true "CreateBranchRequestBody"
// @Success 201 {object} models.Branch "GetBranchBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) CreateBranch(c *gin.Context) {

	var branch models.CreateBranch

	err := c.ShouldBindJSON(&branch)
	if err != nil {
		log.Println("error whiling marshal json:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Branch().Insert(context.Background(), &branch)
	if err != nil {
		log.Println("error whiling create Branc:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res, err := h.storage.Branch().GetByID(context.Background(), &models.BranchPrimeryKey{
		Id: id,
	})
	if err != nil {
		log.Println("error whiling get by id Branch:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetByIDBranch godoc
// @ID get_by_id_Branch
// @Router /branch/{id} [GET]
// @Summary Get By ID Branch
// @Description Get By ID Branch
// @Tags Branch
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Branch "GetBranchBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetByIdBranch(c *gin.Context) {

	id := c.Param("id")

	res, err := h.storage.Branch().GetByID(context.Background(), &models.BranchPrimeryKey{
		Id: id,
	})
	if err != nil {
		log.Println("error whiling get by id Branch:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetListBranch godoc
// @ID get_list_Branch
// @Router /branch [GET]
// @Summary Get List Branch
// @Description Get List Branch
// @Tags Branch
// @Accept json
// @Produce json
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Param search query string false "search"
// @Success 200 {object} models.GetListBranchResponse "GetBranchListBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetListBranch(c *gin.Context) {
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

	res, err := h.storage.Branch().GetList(context.Background(), &models.GetListBranchRequest{
		Offset: int64(offset),
		Limit:  int64(limit),
	})

	if err != nil {
		log.Println("error whiling get list Branch:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// UpdateBranch godoc
// @ID update_Branch
// @Router /branch/{id} [PUT]
// @Summary Update Branch
// @Description Update Branch
// @Tags Branch
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param Branch body models.UpdateBranchSwag true "UpdateBranchRequestBody"
// @Success 202 {object} models.Branch "UpdateBranchBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) UpdateBranch(c *gin.Context) {

	var (
		Branch models.UpdateBranch
	)

	Branch.Id = c.Param("id")

	err := c.ShouldBindJSON(&Branch)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.storage.Branch().Update(context.Background(), &models.UpdateBranch{
		Id:   Branch.Id,
		Name: Branch.Name,
	})

	if err != nil {
		log.Printf("error whiling update: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update").Error())
		return
	}

	resp, err := h.storage.Branch().GetByID(context.Background(), &models.BranchPrimeryKey{
		Id: Branch.Id,
	})
	if err != nil {
		log.Printf("error whiling get by id Branch: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}

	c.JSON(http.StatusAccepted, resp)
}

// DeleteBranch godoc
// @ID delete_Branch
// @Router /branch/{id} [DELETE]
// @Summary Delete Branch
// @Description Delete Branch
// @Tags Branch
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 204 {object} models.Empty "DeleteBranchBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) DeleteBranch(c *gin.Context) {
	id := c.Param("id")

	err := h.storage.Branch().Delete(context.Background(), &models.BranchPrimeryKey{Id: id})
	if err != nil {
		log.Println("error whiling delete  Branch:", err.Error())
		c.JSON(http.StatusNoContent, err.Error())
		return
	}
	c.JSON(http.StatusCreated, "Branch deleted")
}