package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"assessment-test/controller"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func main() {
	app := iris.New()
	app.Adapt(httprouter.New())

	app.Get("/", controller.Index)

	productAPI := app.Party("/v1", controller.ProductAPIMiddleware)
	{
		productAPI.OnError(404, controller.ProductNotFoundHandler)
		productAPI.Post("/products", controller.ProductAdd)
		productAPI.Put("/products/:id", controller.ProductUpdate)
		productAPI.Get("/products", controller.ProductList)
		productAPI.Get("/products/:id", controller.ProductDetail)
		productAPI.Delete("/products/:id", controller.ProductDelete)
	}

	productV2API := app.Party("/v2", controller.ProductAPIMiddleware)
	{
		productV2API.OnError(404, controller.ProductAPIMiddleware)
		productV2API.Get("/products", controller.ProductListV2)
	}

	authAPI := app.Party("/auth", controller.AuthAPIMiddleware)
	{
		authAPI.OnError(404, controller.AuthNotFoundHandler)
		authAPI.Post("/signup", controller.AuthSignup)
	}




	app.Listen(":8089")
}


