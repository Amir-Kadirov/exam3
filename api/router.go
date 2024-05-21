package api

import (
	"exam/api/handler"
	"exam/pkg/logger"
	"exam/service"

	// "net/http"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New ...
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(service service.IServiceManager, log logger.ILogger) *gin.Engine {
	h := handler.NewStrg(service, log)

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/Customers", h.CreateCustomers)
	r.GET("/Customers", h.GetAllCustomerss)
	r.GET("/Customers/:id", h.GetById)
	r.PUT("/Customers/updateCustomers/:id", h.UpdateCustomers)
	r.DELETE("/Customers/deleteCustomers/:id", h.DeleteCustomers)

	return r
}
