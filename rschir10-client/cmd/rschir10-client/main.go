package main

import (
	"github.com/gin-gonic/gin"
	"ubulllka.com/rschir10-client/pkg/book"
	"ubulllka.com/rschir10-client/pkg/product"
)


func main() {
	r := gin.Default()
	book.RegisterRoutes(r)
	product.RegisterRoutes(r)
	r.Run(":50050")
}