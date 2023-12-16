package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	desc "ubulllka.com/rschir10-client/pkg/product/proto"
)

func GetAll(ctx *gin.Context, c desc.ProductV1Client) {
	res, err := c.GetAll(context.Background(), &desc.EmptyProduct{})
	if err != nil{
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(http.StatusOK, &res)
}

func Get(ctx *gin.Context, c desc.ProductV1Client) {
	id := ctx.Param("id")
	res, err := c.Get(context.Background(), &desc.ProductId{Id: id})
	if err != nil{
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(http.StatusOK, &res)
}

func Insert(ctx *gin.Context, c desc.ProductV1Client) {
	product := desc.Product{}
	if err := ctx.BindJSON(&product); err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	res, err := c.Insert(context.Background(), &desc.Product{Id: product.Id, Name: product.Name, Price: product.Price})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(http.StatusCreated, &res)
}

func Update(ctx *gin.Context, c desc.ProductV1Client) {
	product := desc.Product{}
	if err := ctx.BindJSON(&product); err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	res, err := c.Update(context.Background(), &desc.Product{Id: product.Id, Name: product.Name, Price: product.Price})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(http.StatusCreated, &res)
}

func Remove(ctx *gin.Context, c desc.ProductV1Client) {
	id := ctx.Param("id")
	res, err := c.Remove(context.Background(), &desc.ProductId{Id: id})
	if err != nil{
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(http.StatusOK, &res)
}