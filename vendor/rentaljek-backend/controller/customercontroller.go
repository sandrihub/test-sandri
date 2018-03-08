package controller

import (
	"encoding/json"
	"fmt"
	"rentaljek-backend/util"
	"gopkg.in/kataras/iris.v6"
	"io/ioutil"
	"rentaljek-backend/config"
	"rentaljek-backend/model"
	"rentaljek-backend/view"
	"crypto/rand"
	"crypto/sha1"
	"io"
	"os"
	"encoding/hex"
)

type Customer		model.Customer
type CustomerView	view.Customer

func CustomerAdd(ctx *iris.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Println("error:", err)
	}

	db := config.ConnectDB()
	model.SetDatabase(db)

	var t Customer
	err = json.Unmarshal(body, &t)
	if err != nil {
		fmt.Println("error:", err)
	}

	password := []byte(t.Password)
	fmt.Println("Password : ", string(password))

	// generate salt from given password
	salt := GenerateSalt(password)
	fmt.Printf("Salt : %x \n", salt)

	// generate password + salt hash to store into database
	combination := string(salt) + string(password)
	passwordHash := sha1.New()
	io.WriteString(passwordHash, combination)
	fmt.Printf("Password Hash : %x \n", passwordHash.Sum(nil))
	t.Password = hex.EncodeToString(passwordHash.Sum(nil))

	var statusView StatusView

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

func CustomerDetail(ctx *iris.Context) {
	db := config.ConnectDB()
	model.SetDatabase(db)
	customerid, err := ctx.ParamInt("customerId")

	var customers []*Customer
	customerView := CustomerView{}

	err = db.First(&customers, customerid).Error

	customerMar, err := json.Marshal(customers)
	if err != nil {
		fmt.Println("error:", err)
	}

	json.Unmarshal(customerMar, &customerView)

	customerView.Id = int(customers[0].ID)
	customerView.Fullname = customers[0].Fullname
	customerView.Email = customers[0].Email
	customerView.CreatedAt = customers[0].CreatedAt
	customerView.Username = customers[0].Username
	customerView.Rating = customers[0].Rating

	ctx.JSON(iris.StatusOK, customerView)
}

const saltSize = 16

func GenerateSalt(secret []byte) []byte {
	buf := make([]byte, saltSize, saltSize+sha1.Size)
	_, err := io.ReadFull(rand.Reader, buf)

	if err != nil {
		fmt.Printf("random read failed: %v", err)
		os.Exit(1)
	}

	hash := sha1.New()
	hash.Write(buf)
	hash.Write(secret)
	return hash.Sum(buf)
}