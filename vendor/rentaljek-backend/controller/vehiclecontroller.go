package controller

import (
	"encoding/json"
	"fmt"
	"gopkg.in/kataras/iris.v6"
	"io/ioutil"
	"rentaljek-backend/config"
	"rentaljek-backend/model"
	"log"
	"rentaljek-backend/util"
	"rentaljek-backend/view"
)

type Vehicle		model.Vehicle
type Photo		model.Photo
type PaymentType	model.PaymentType
type StatusView		view.StatusView
type VehicleView	view.VehicleView

func VehicleAPIMiddleware(ctx *iris.Context) {
	println("Request: " + ctx.Path())
	ctx.Next()
}

func VehicleNotFoundHandler(ctx *iris.Context) {
	ctx.HTML(iris.StatusNotFound, "<h1> Page not found </h1>")
}

func Index(ctx *iris.Context) {
	ctx.HTML(iris.StatusOK, "<h1> Welcome to RentalJek!</h1>")
}

func VehicleAdd(ctx *iris.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Println("error:", err)
	}

	db := config.ConnectDB()
	model.SetDatabase(db)

	var t Vehicle
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

	if t.ID != 0 {
		for _, v := range t.Photos {
			b, err := json.Marshal(v)
			if err != nil {
				fmt.Println("error:", err)
			}

			var pt Photo
			err = json.Unmarshal(b, &pt)
			if err != nil {
				fmt.Println("error:", err)
			}

			pt.VehicleCode= t.VehicleCode
			db.Save(&pt)

		}

		for _, v :=  range t.PaymentTypes {
			b, err := json.Marshal(v)
			if err != nil {
				fmt.Println("error:", err)
			}

			var pt PaymentType
			err = json.Unmarshal(b, &pt)
			if err != nil {
				fmt.Println("error:", err)
			}

			pt.VehicleCode= t.VehicleCode
			db.Save(&pt)
		}
	}

	statusView.Status = util.StatusSuccess
	statusView.Message = util.MessageSuccessInsertTable

	ctx.JSON(iris.StatusCreated, statusView)

}

func VehicleUpdate(ctx *iris.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Println("error:", err)
	}

	db := config.ConnectDB()
	model.SetDatabase(db)

	var t Vehicle
	err = json.Unmarshal(body, &t)
	if err != nil {
		fmt.Println("error:", err)
	}

	log.Println(t)
	db.Model(&t).Where("vehicle_code = ?", t.VehicleCode).Updates(t)

	if t.ID != 0 {
		for _, v := range t.Photos {
			b, err := json.Marshal(v)
			if err != nil {
				fmt.Println("error:", err)
			}

			var pt Photo
			err = json.Unmarshal(b, &pt)
			if err != nil {
				fmt.Println("error:", err)
			}

			pt.VehicleCode= t.VehicleCode
			db.Save(&pt)

		}

		for _, v :=  range t.PaymentTypes {
			b, err := json.Marshal(v)
			if err != nil {
				fmt.Println("error:", err)
			}

			var pt PaymentType
			err = json.Unmarshal(b, &pt)
			if err != nil {
				fmt.Println("error:", err)
			}

			pt.VehicleCode= t.VehicleCode
			db.Save(&pt)
		}
	}

}

func VehicleDelete(ctx *iris.Context) {
	db := config.ConnectDB()
	model.SetDatabase(db)
	vehiclecode := ctx.Param("code")
	log.Println(vehiclecode)

	var vehicles []*Vehicle

	db.Where("vehicle_code = ?", vehiclecode).Delete(&vehicles)
}

func VehicleDetail(ctx *iris.Context) {
	db := config.ConnectDB()
	model.SetDatabase(db)
	vehiclecode := ctx.Param("code")

	var vehicles Vehicle
	var photo []view.PhotoView
	var category view.CategoryView
	var paymentType []view.PaymentTypeView
	var vehicleView VehicleView


	err := db.Where("vehicle_code = ?", vehiclecode).Find(&vehicles).Error
	if err != nil {
		log.Println(err)
	}


	vehicleMar, err := json.Marshal(vehicles)
	if err != nil {
		fmt.Println("error:", err)
	}

	json.Unmarshal(vehicleMar, &vehicleView)

	if vehicles.ID > 0 {
		vehicleView.Price.Daily.ActualPrice = vehicles.ActualPriceDaily
		vehicleView.Price.Daily.Discount = vehicles.DiscountDaily
		vehicleView.Price.Daily.InitialPrice = vehicles.InitialPriceDaily

		vehicleView.Price.Weekly.ActualPrice = vehicles.ActualPriceWeekly
		vehicleView.Price.Weekly.Discount = vehicles.DiscountWeekly
		vehicleView.Price.Weekly.InitialPrice = vehicles.InitialPriceWeekly

		vehicleView.Price.Monthly.ActualPrice = vehicles.ActualPriceMonthly
		vehicleView.Price.Monthly.Discount = vehicles.DiscountMonthly
		vehicleView.Price.Monthly.InitialPrice = vehicles.InitialPriceMonthly

		categorySel := db.Table("categories").Where("category_code = ?", vehicles.CategoryCode ).Find(&category)
		categoryMar, err := json.Marshal(categorySel.Value)
		if err != nil {
			fmt.Println("error:", err)
		}

		err = json.Unmarshal(categoryMar, &category)
		if err != nil {
			fmt.Println(err.Error())
		}

		vehicleView.CategoryId = int(category.ID)
		vehicleView.CategoryName = category.Name

		photoSel := db.Table("photos").Where("vehicle_code = ?", vehicles.VehicleCode).Find(&photo)
		photoMar, err := json.Marshal(photoSel.Value)
		if err != nil {
			fmt.Println("error:", err)
		}

		err = json.Unmarshal(photoMar, &photo)
		if err != nil {
			fmt.Println(err.Error())
		}

		vehicleView.Photos = photo

		paymentTypeSel := db.Table("payment_types").Where("vehicle_code = ?", vehicles.VehicleCode).Find(&paymentType)
		paymentTypeMar, err := json.Marshal(paymentTypeSel.Value)
		if err != nil {
			fmt.Println("error:", err)
		}

		err = json.Unmarshal(paymentTypeMar, &paymentType)
		if err != nil {
			fmt.Println(err.Error())
		}

		vehicleView.Price.PaymentTypes = paymentType
		vehicleView.Price.CancelationFee = vehicles.CancelationFee
		ctx.JSON(iris.StatusOK, vehicleView)

	}

}

func VehicleList(ctx *iris.Context) {
	db := config.ConnectDB()
	model.SetDatabase(db)

	var vehicles []*Vehicle
	var photo []view.PhotoView
	var category view.CategoryView
	var paymentType []view.PaymentTypeView
	vehicleView := []*VehicleView{}

	db.Find(&vehicles)

	vehicleMar, err := json.Marshal(vehicles)
	if err != nil {
		fmt.Println("error:", err)
	}

	json.Unmarshal(vehicleMar, &vehicleView)

	for k, val := range vehicles {

		vehicleView[k].Price.Daily.ActualPrice = val.ActualPriceDaily
		vehicleView[k].Price.Daily.Discount = val.DiscountDaily
		vehicleView[k].Price.Daily.InitialPrice = val.InitialPriceDaily

		vehicleView[k].Price.Weekly.ActualPrice = val.ActualPriceWeekly
		vehicleView[k].Price.Weekly.Discount = val.DiscountWeekly
		vehicleView[k].Price.Weekly.InitialPrice = val.InitialPriceWeekly

		vehicleView[k].Price.Monthly.ActualPrice = val.ActualPriceMonthly
		vehicleView[k].Price.Monthly.Discount = val.DiscountMonthly
		vehicleView[k].Price.Monthly.InitialPrice = val.InitialPriceMonthly

		categorySel := db.Table("categories").Where("category_code = ?", val.CategoryCode ).Find(&category)
		categoryMar, err := json.Marshal(categorySel.Value)
		if err != nil {
			fmt.Println("error:", err)
		}

		err = json.Unmarshal(categoryMar, &category)
		if err != nil {
			fmt.Println(err.Error())
		}

		vehicleView[k].CategoryId = int(category.ID)
		vehicleView[k].CategoryName = category.Name

		photoSel := db.Table("photos").Where("vehicle_code = ?", val.VehicleCode ).Find(&photo)
		photoMar, err := json.Marshal(photoSel.Value)
		if err != nil {
			fmt.Println("error:", err)
		}

		err = json.Unmarshal(photoMar, &photo)
		if err != nil {
			fmt.Println(err.Error())
		}

		vehicleView[k].Photos = photo

		paymentTypeSel := db.Table("payment_types").Where("vehicle_code = ?", val.VehicleCode ).Find(&paymentType)
		paymentTypeMar, err := json.Marshal(paymentTypeSel.Value)
		if err != nil {
			fmt.Println("error:", err)
		}

		err = json.Unmarshal(paymentTypeMar, &paymentType)
		if err != nil {
			fmt.Println(err.Error())
		}

		vehicleView[k].Price.PaymentTypes = paymentType
		vehicleView[k].Price.CancelationFee = val.CancelationFee
	}

	ctx.JSON(iris.StatusOK, vehicleView)
}
