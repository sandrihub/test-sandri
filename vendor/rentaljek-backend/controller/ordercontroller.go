package controller

import (
	"encoding/json"
	"fmt"
	"rentaljek-backend/util"
	"gopkg.in/kataras/iris.v6"
	"io/ioutil"
	"rentaljek-backend/config"
	"rentaljek-backend/model"
	"log"
	"rentaljek-backend/view"
)

type Order		model.Order
type OrderView		view.OrderView

func OrderAdd(ctx *iris.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Println("error:", err)
	}

	db := config.ConnectDB()
	model.SetDatabase(db)

	var t Order
	err = json.Unmarshal(body, &t)
	if err != nil {
		fmt.Println("error:", err)
	}

	var statusView StatusView

	log.Println(t)
	err = db.Save(&t).Error
	if err != nil {
		statusView.Status = util.StatusFailed
		statusView.Message = err.Error()
		ctx.JSON(iris.StatusConflict, statusView)

		return
	}


	statusView.Status = util.StatusSuccess
	statusView.Message = util.MessageSuccessInsertTable

	ctx.JSON(iris.StatusCreated, statusView)

}

func OrderList(ctx *iris.Context) {
	db := config.ConnectDB()
	model.SetDatabase(db)
	customerid, err := ctx.ParamInt("customerId")

	var orders []*Order
	orderView := OrderView{}

	err = db.Where("customer_id = ?", customerid).Find(&orders).Error

	orderMar, err := json.Marshal(orders)
	if err != nil {
		fmt.Println("error:", err)
	}

	json.Unmarshal(orderMar, &orderView.HistoryOrder.Data)

	orderView.CutomerId = customerid
	orderView.HistoryOrder.Count = len(orders)
	if(len(orders)==0){
		orderView.HistoryOrder.Data = nil
	}


	for i, val := range orders {
		orderView.HistoryOrder.Data[i].StartOrder = val.CreatedAt.Format("2006-01-02")
		orderView.HistoryOrder.Data[i].OrderId = int(val.ID)
		orderView.HistoryOrder.Data[i].Rate = val.Rate
		orderView.HistoryOrder.Data[i].Message = val.Message
	}

	ctx.JSON(iris.StatusOK, orderView)
}
