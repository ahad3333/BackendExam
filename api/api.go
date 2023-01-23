package api

import (
	_ "add/api/docs"
	"add/api/handler"
	"add/storage"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


func NewApi(r *gin.Engine, storage storage.StorageI) {

	handlerV1 := handler.NewHandler(storage)

	r.POST("/client", handlerV1.CreateClient)
	r.GET("/client/:id", handlerV1.GetByIDClient)
	r.GET("/client", handlerV1.GetListClient)
	r.PUT("/client/:id", handlerV1.UpdateClient)
	r.DELETE("/client/:id", handlerV1.DeleteClient)

	r.POST("/investor", handlerV1.CreateInvestor)
	r.GET("/investor/:id", handlerV1.GetByIdInvestor)
	r.GET("/investor", handlerV1.GetListInvestor)
	r.DELETE("/investor/:id", handlerV1.DeleteInvestor)
	r.PUT("/investor/:id", handlerV1.UpdateInvestor)

	r.POST("/car", handlerV1.CreateCar)
	r.GET("/car/:id", handlerV1.GetByIDCar)
	r.GET("/car", handlerV1.GetListCar)
	r.PUT("/car/:id", handlerV1.UpdateCar)
	r.DELETE("/car/:id", handlerV1.DeleteCar)

	r.POST("/order", handlerV1.CreateOrder)
	r.GET("/order/:id", handlerV1.GetByIDOrder)
	r.GET("/order", handlerV1.GetListOrder)
	r.PUT("/order/:id", handlerV1.UpdateOrder)
	r.PUT("/return/:id", handlerV1.UpdateReturn)
	r.DELETE("/order/:id", handlerV1.DeleteOrder)

	r.GET("/investor-share", handlerV1.GetListInves)
	r.GET("/debtor-client", handlerV1.GetListBebtor)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
