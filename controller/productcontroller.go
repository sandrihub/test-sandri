package controller

import (
	"gopkg.in/kataras/iris.v6"
	"fmt"
	"assessment-test/config"
	"assessment-test/model"
	"log"
	"assessment-test/util"
	"assessment-test/view"
	"io/ioutil"
	"encoding/json"
)

type Product		model.Product
type StatusView		view.StatusView
type ProductView	view.ProductView
type MessageView	view.MessageView

func ProductAPIMiddleware(ctx *iris.Context) {
	println("Request: " + ctx.Path())
	ctx.Next()
}

func ProductNotFoundHandler(ctx *iris.Context) {
	ctx.HTML(iris.StatusNotFound, "<h1> Page not found </h1>")
}

func Index(ctx *iris.Context) {
	ctx.HTML(iris.StatusOK, "<h1> Welcome to Assessment Test!</h1>")
}

func ProductAdd(ctx *iris.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	var statusView StatusView

	if err != nil {
		fmt.Println("error:", err)
		statusView.Status = util.StatusFailed
		statusView.Error = err.Error()
		ctx.JSON(iris.StatusConflict, statusView)

		return
	}

	db := config.ConnectDB()
	model.SetDatabase(db)

	var t Product
	err = json.Unmarshal(body, &t)
	if err != nil {
		fmt.Println("error:", err)
		statusView.Status = util.StatusFailed
		statusView.Error = err.Error()
		ctx.JSON(iris.StatusConflict, statusView)

		return
	}

	err = db.Save(&t).Error
	if err != nil {
		statusView.Status = util.StatusFailed
		statusView.Error = err.Error()
		ctx.JSON(iris.StatusConflict, statusView)

		return
	}

	statusView.Status = util.StatusSuccess
	statusView.Result = t

	ctx.JSON(iris.StatusCreated, statusView)
}

func ProductList(ctx *iris.Context) {
	db := config.ConnectDB()
	model.SetDatabase(db)

	var products []*Product
	var statusView StatusView
	productView := []*ProductView{}

	db.Find(&products)

	productMar, err := json.Marshal(products)
	if err != nil {
		fmt.Println("error:", err)
		statusView.Status = util.StatusFailed
		statusView.Error = err.Error()
		ctx.JSON(iris.StatusConflict, statusView)

		return
	}

	json.Unmarshal(productMar, &productView)

	for k, val := range products {
		productView[k].Id = val.Id
		productView[k].Name = val.Name
		productView[k].Price = val.Price
		productView[k].Imageurl = val.Imageurl
		productView[k].CreatedAt = val.CreatedAt
		productView[k].UpdatedAt = val.UpdatedAt
	}

	statusView.Status = util.StatusSuccess
	statusView.Result = productView

	ctx.JSON(iris.StatusOK, statusView)
}

func ProductDetail(ctx *iris.Context) {
	db := config.ConnectDB()
	model.SetDatabase(db)
	productId := ctx.Param("id")

	var products Product
	var productView ProductView
	var statusView StatusView

	err := db.Where("id = ?", productId).Find(&products).Error
	if err != nil {
		log.Println(err)
		statusView.Status = util.StatusFailed
		statusView.Error = err.Error()
		ctx.JSON(iris.StatusConflict, statusView)

		return
	}


	productMar, err := json.Marshal(products)
	if err != nil {
		fmt.Println("error:", err)
		statusView.Status = util.StatusFailed
		statusView.Error = err.Error()
		ctx.JSON(iris.StatusConflict, statusView)

		return
	}

	json.Unmarshal(productMar, &productView)

	if products.Id > 0 {
		productView.Id = products.Id
		productView.Name = products.Name
		productView.Price = products.Price
		productView.Imageurl = products.Imageurl
		productView.CreatedAt = products.CreatedAt
		productView.UpdatedAt = products.UpdatedAt

		statusView.Status = util.StatusSuccess
		statusView.Result = productView

		ctx.JSON(iris.StatusOK, statusView)
	}

}

func ProductUpdate(ctx *iris.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	productId := ctx.Param("id")
	var statusView StatusView
	var products Product
	var productView ProductView

	if err != nil {
		fmt.Println("error:", err)
		statusView.Status = util.StatusFailed
		statusView.Error = err.Error()
		ctx.JSON(iris.StatusConflict, statusView)

		return
	}

	db := config.ConnectDB()
	model.SetDatabase(db)

	var t Product
	err = json.Unmarshal(body, &t)
	if err != nil {
		fmt.Println("error:", err)
		statusView.Status = util.StatusFailed
		statusView.Error = err.Error()
		ctx.JSON(iris.StatusConflict, statusView)

		return
	}

	log.Println(t)
	db.Model(&t).Where("id = ?", productId).Updates(t)

	err = db.Where("id = ?", productId).Find(&products).Error
	if err != nil {
		log.Println(err)
		statusView.Status = util.StatusFailed
		statusView.Error = err.Error()
		ctx.JSON(iris.StatusConflict, statusView)

		return
	}


	productMar, err := json.Marshal(products)
	if err != nil {
		fmt.Println("error:", err)
		statusView.Status = util.StatusFailed
		statusView.Error = err.Error()
		ctx.JSON(iris.StatusConflict, statusView)

		return
	}

	json.Unmarshal(productMar, &productView)

	if products.Id > 0 {
		productView.Id = products.Id
		productView.Name = products.Name
		productView.Price = products.Price
		productView.Imageurl = products.Imageurl
		productView.CreatedAt = products.CreatedAt
		productView.UpdatedAt = products.UpdatedAt

		statusView.Status = util.StatusSuccess
		statusView.Result = productView

		ctx.JSON(iris.StatusOK, statusView)
	}

}

func ProductDelete(ctx *iris.Context) {
	db := config.ConnectDB()
	model.SetDatabase(db)
	productId := ctx.Param("id")
	var statusView StatusView
	var messageView MessageView

	var products []*Product

	db.Where("id = ?", productId).Delete(&products)

	messageView.Message = productId + " deleted"
	statusView.Status = util.StatusSuccess
	statusView.Result = messageView

	ctx.JSON(iris.StatusOK, statusView)
}

func ProductListV2(ctx *iris.Context) {
	var messageView MessageView

	messageView.Message = "Hello there"

	ctx.JSON(iris.StatusOK, messageView)
}

