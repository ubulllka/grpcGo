package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	desc "ubulllka.com/rschir10-client/pkg/book/proto"
)

func GetAll(ctx *gin.Context, c desc.BookV1Client) {
	res, err := c.GetAll(context.Background(), &desc.EmptyBook{})
	if err != nil{
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(http.StatusOK, &res)
}

func Get(ctx *gin.Context, c desc.BookV1Client) {
	id := ctx.Param("id")
	res, err := c.Get(context.Background(), &desc.BookId{Id: id})
	if err != nil{
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(http.StatusOK, &res)
}

func Insert(ctx *gin.Context, c desc.BookV1Client) {
	book := desc.Book{}
	if err := ctx.BindJSON(&book); err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	res, err := c.Insert(context.Background(), &desc.Book{Id: book.Id, Name: book.Name, Author: book.Author})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(http.StatusCreated, &res)
}

func Update(ctx *gin.Context, c desc.BookV1Client) {
	book := desc.Book{}
	if err := ctx.BindJSON(&book); err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	res, err := c.Update(context.Background(), &desc.Book{Id: book.Id, Name: book.Name, Author: book.Author})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(http.StatusCreated, &res)
}

func Remove(ctx *gin.Context, c desc.BookV1Client) {
	id := ctx.Param("id")
	res, err := c.Remove(context.Background(), &desc.BookId{Id: id})
	if err != nil{
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(http.StatusOK, &res)
}