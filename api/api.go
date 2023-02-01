package api

import (
	"errors"

	"net/http"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"

	_ "app/api/docs"
	"app/api/handler"
	"app/config"
	"app/pkg/helper"
	"app/storage"
)

func NewApi(cfg *config.Config, r *gin.Engine, storage storage.StorageI, cache storage.CacheStorageI) {

	handlerV1 := handler.NewHandler(cfg, storage, cache)

	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization

	v1 := r.Group("v1")

	r.Use(customCORSMiddleware())
	v1.Use(SecurityMiddleware())

	v1.GET("/user", handlerV1.GetUserList)
	v1.GET("/order",handlerV1.GetListOrder)

	r.POST("/investor", handlerV1.CreateInvestor)
	r.GET("/investor/:id", handlerV1.GetByIdInvestor)
	r.GET("/investor", handlerV1.GetListInvestor)
	r.DELETE("/investor/:id", handlerV1.DeleteInvestor)
	r.PUT("/investor/:id", handlerV1.UpdateInvestor)

	r.POST("/car", handlerV1.CreateCar)
	r.GET("/car/:id", handlerV1.GetByIdCar)
	r.GET("/car", handlerV1.GetListCar)
	r.DELETE("/car/:id", handlerV1.DeleteCar)
	r.PUT("/car/:id", handlerV1.UpdateCar)

	r.POST("/client", handlerV1.CreateClient)
	r.GET("/client/:id", handlerV1.GetByIdClient)
	r.GET("/client", handlerV1.GetListClient)
	r.DELETE("/client/:id", handlerV1.DeleteClient)
	r.PUT("/client/:id", handlerV1.UpdateClient)

	r.POST("/order", handlerV1.CreateOrder)
	r.GET("/order/:id", handlerV1.GetByIdOrder)
	r.DELETE("/order/:id", handlerV1.DeleteOrder)
	r.PUT("/order/:id", handlerV1.UpdateOrder)
	r.PATCH("/order/:id", handlerV1.UpdatePatchOrder)

	r.POST("/branch", handlerV1.CreateBranch)
	r.GET("/branch/:id", handlerV1.GetByIdBranch)
	r.GET("/branch", handlerV1.GetListBranch)
	r.DELETE("/branch/:id", handlerV1.DeleteBranch)
	r.PUT("/branch/:id", handlerV1.UpdateBranch)

	// Report
	r.GET("/report/debtors", handlerV1.GetDebtors)
	r.GET("/report/investor-share", handlerV1.GetInvestorShare)
	r.GET("/report/company-share", handlerV1.GetBranchShare)

	r.POST("/login", handlerV1.Login)

	r.POST("/user", handlerV1.CreateUser)
	r.GET("/user/:id", handlerV1.GetUserById)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))
}

func SecurityMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		key := config.Load().AuthSecretKey

		if len(c.Request.Header["Authorization"]) > 0 {

			token := c.Request.Header["Authorization"][0]

			_, err := helper.ParseClaims(token, key)

			if err != nil {
				c.JSON(http.StatusUnauthorized, struct {
					Code int
					Err  string
				}{
					Code: http.StatusUnauthorized,
					Err:  errors.New("error access denied 2").Error(),
				})
				c.Abort()
				return
			}
		} else {
			c.JSON(http.StatusUnauthorized, struct {
				Code int
				Err  string
			}{
				Code: http.StatusUnauthorized,
				Err:  errors.New("error access denied 1").Error(),
			})
			c.Abort()
			return
		}
		c.Next()
		
		}
	}



func Parse(s string) {
	panic("unimplemented")
}

func customCORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
