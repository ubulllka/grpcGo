package book

import (
	"github.com/gin-gonic/gin"
	"ubulllka.com/rschir10-client/pkg/book/routes"
)

func RegisterRoutes(r *gin.Engine) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(),
	}

	routes := r.Group("/book")
	routes.GET("/", svc.GetAll)
	routes.GET("/:id", svc.Get)
	routes.POST("/", svc.Insert)
	routes.PUT("/", svc.Update)
	routes.DELETE("/:id", svc.Remove)

	return svc
}

func (svc *ServiceClient) GetAll(ctx *gin.Context) {
	routes.GetAll(ctx, svc.Client)
}

func (svc *ServiceClient) Get(ctx *gin.Context) {
	routes.Get(ctx, svc.Client)
}

func (svc *ServiceClient) Insert(ctx *gin.Context) {
	routes.Insert(ctx, svc.Client)
}

func (svc *ServiceClient) Update(ctx *gin.Context) {
	routes.Update(ctx, svc.Client)
}

func (svc *ServiceClient) Remove(ctx *gin.Context) {
	routes.Remove(ctx, svc.Client)
}