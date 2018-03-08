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

type Provider		model.Provider
type ProviderView	view.ProviderView
type DataProvider	view.DataProvider

func ProviderAdd(ctx *iris.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Println("error:", err)
	}

	db := config.ConnectDB()
	model.SetDatabase(db)

	var t Provider
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

func ProviderList(ctx *iris.Context) {
	db := config.ConnectDB()
	model.SetDatabase(db)

	var providers []*Provider
	providerView := []*ProviderView{}

	var err = db.Find(&providers).Error

	providerMar, err := json.Marshal(providers)
	if err != nil {
		fmt.Println("error:", err)
	}

	json.Unmarshal(providerMar, &providerView)


	for i, val := range providers {
		providerView[i].ProviderName = val.Name
		providerView[i].ProviderAddress = val.Address
		providerView[i].Id = int(val.ID)
		providerView[i].ProviderCode = val.ProviderCode
		providerView[i].ProviderPhone = val.Phone
		providerView[i].ProviderEmail = val.Email
		providerView[i].ProviderOwner = val.Owner
		providerView[i].Latitude = val.Latitude
		providerView[i].Longitude = val.Longitude
	}

	ctx.JSON(iris.StatusOK, providerView)
}
