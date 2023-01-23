package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"add/models"

	"github.com/gin-gonic/gin"
)


// GetListCar godoc
// @ID CalcInvesPrimeryKey
// @Router /investor-share [GET]
// @Summary Get List Calc
// @Description Get List Calc
// @Tags CalcInves
// @Accept json
// @Produce json
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Param search query string false "search"
// @Success 200 {object} models.GetListInves "GetCalcListBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetListInves(c *gin.Context) {
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

	res, err := h.storage.Calc().GetListInves(context.Background(),&models.GetListInvesRequest{
		Offset: int64(offset),
		Limit:  int64(limit),
		Search: search,
	})

	if err != nil {
		log.Println("error whiling get list GetListInves:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetListCar godoc
// @ID CalcBebtPrimeryKey
// @Router /debtor-client [GET]
// @Summary Get List Calc
// @Description Get List Calc
// @Tags CalcBebt
// @Accept json
// @Produce json
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Param search query string false "search"
// @Success 200 {object} models.GetListBebt "GetBabtListBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetListBebtor(c *gin.Context) {
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

	res, err := h.storage.Calc().GetListBebt(context.Background(),&models.GetListBebtRequest{
		Offset: int64(offset),
		Limit:  int64(limit),
		Search: search,
	})

	if err != nil {
		log.Println("error whiling get list GetListBabt:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}