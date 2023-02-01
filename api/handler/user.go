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

// CreateUser godoc
// @ID create_User
// @Router /user [POST]
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept json
// @Produce json
// @Param User body models.CreateUser true "CreateUserRequestBody"
// @Success 201 {object} models.User "GetUserBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *Handler) CreateUser(c *gin.Context) {
	var User models.CreateUser

	err := c.ShouldBindJSON(&User)
	if err != nil {
		log.Printf("error whiling create: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.User().Create(context.Background(), &User)
	if err != nil {
		log.Printf("error whiling Create: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling Create").Error())
		return
	}

	resp, err := h.storage.User().GetByPKey(
		context.Background(),
		&models.UserPrimarKey{Id: id},
	)

	if err != nil {
		log.Printf("error whiling GetByPKey: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling GetByPKey").Error())
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetByIdUser godoc
// @ID get_by_id_User
// @Router /user/{id} [GET]
// @Summary Get By Id User
// @Description Get By Id User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.User "GetUserBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *Handler) GetUserById(c *gin.Context) {

	id := c.Param("id")
	resp, err := h.storage.User().GetByPKey(
		context.Background(),
		&models.UserPrimarKey{Id: id},
	)

	if err != nil {
		log.Printf("error whiling GetByPKey: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling GetByPKey").Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetListUser godoc
// @Security ApiKeyAuth
// @ID get_list_User
// @Router /v1/user [GET]
// @Summary Get List User
// @Description Get List User
// @Tags User
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Success 200 {object} models.GetListUserResponse "GetUserBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *Handler) GetUserList(c *gin.Context) {
	var (
		limit  int
		offset int
		err    error
	)

	limitStr := c.Query("limit")
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			log.Printf("error whiling limit: %v\n", err)
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	offsetStr := c.Query("offset")
	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			log.Printf("error whiling limit: %v\n", err)
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	resp, err := h.storage.User().GetList(
		context.Background(),
		&models.GetListUserRequest{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	)

	if err != nil {
		log.Printf("error whiling get list: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get list").Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// // UpdateUser godoc
// // @ID update_User
// // @Router /user/{id} [PUT]
// // @Summary Update User
// // @Description Update User
// // @Tags User
// // @Accept json
// // @Produce json
// // @Param id path string true "id"
// // @Param User body models.UpdateUserSwagger true "CreateUserRequestBody"
// // @Success 200 {object} models.User "GetUsersBody"
// // @Response 400 {object} string "Invalid Argument"
// // @Failure 500 {object} string "Server Error"
// func (h *HandlerV1) UpdateUser(c *gin.Context) {

// 	var (
// 		User models.UpdateUser
// 	)

// 	id := c.Param("id")

// 	if id == "" {
// 		log.Printf("error whiling update: %v\n", errors.New("required User id").Error())
// 		c.JSON(http.StatusBadRequest, errors.New("required User id").Error())
// 		return
// 	}

// 	err := c.ShouldBindJSON(&User)
// 	if err != nil {
// 		log.Printf("error whiling update: %v\n", err)
// 		c.JSON(http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	User.Id = id

// 	rowsAffected, err := h.storage.User().Update(
// 		context.Background(),
// 		&User,
// 	)

// 	if err != nil {
// 		log.Printf("error whiling update: %v", err)
// 		c.JSON(http.StatusInternalServerError, errors.New("error whiling update").Error())
// 		return
// 	}

// 	if rowsAffected == 0 {
// 		log.Printf("error whiling update rows affected: %v", err)
// 		c.JSON(http.StatusInternalServerError, errors.New("error whiling update rows affected").Error())
// 		return
// 	}

// 	resp, err := h.storage.User().GetByPKey(
// 		context.Background(),
// 		&models.UserPrimarKey{Id: id},
// 	)

// 	if err != nil {
// 		log.Printf("error whiling GetByPKey: %v\n", err)
// 		c.JSON(http.StatusInternalServerError, errors.New("error whiling GetByPKey").Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, resp)
// }

// // DeleteByIdUser godoc
// // @ID delete_by_id_User
// // @Router /user/{id} [DELETE]
// // @Summary Delete By Id User
// // @Description Delete By Id User
// // @Tags User
// // @Accept json
// // @Produce json
// // @Param id path string true "id"
// // @Success 200 {object} models.User "GetUserBody"
// // @Response 400 {object} string "Invalid Argument"
// // @Failure 500 {object} string "Server Error"
// func (h *HandlerV1) DeleteUser(c *gin.Context) {

// 	id := c.Param("id")
// 	if id == "" {
// 		log.Printf("error whiling update: %v\n", errors.New("required User id").Error())
// 		c.JSON(http.StatusBadRequest, errors.New("required User id").Error())
// 		return
// 	}

// 	err := h.storage.User().Delete(
// 		context.Background(),
// 		&models.UserPrimarKey{
// 			Id: id,
// 		},
// 	)

// 	if err != nil {
// 		log.Printf("error whiling delete: %v", err)
// 		c.JSON(http.StatusInternalServerError, errors.New("error whiling delete").Error())
// 		return
// 	}

// 	c.JSON(http.StatusNoContent, nil)
// }
