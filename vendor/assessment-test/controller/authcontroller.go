package controller

import (
	"gopkg.in/kataras/iris.v6"
	"fmt"
	"assessment-test/config"
	"assessment-test/model"
	"assessment-test/util"
	"crypto/rand"
	"crypto/sha1"
	"io"
	"os"
	"log"
	"encoding/hex"

)

type Signup		model.Signup

func AuthAPIMiddleware(ctx *iris.Context) {
	println("Request: " + ctx.Path())
	ctx.Next()
}

func AuthNotFoundHandler(ctx *iris.Context) {
	ctx.HTML(iris.StatusNotFound, "<h1> Page not found </h1>")
}


func AuthSignup(ctx *iris.Context) {
	//body, err := ioutil.ReadAll(ctx.Request.MultipartForm)
	name := ctx.FormValue("name")
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")
	var statusView StatusView


	db := config.ConnectDB()
	model.SetDatabase(db)

	var t Signup

	pwd := []byte(password)
	salt := GenerateSalt(pwd)
	combination := string(salt) + string(pwd)
	passwordHash := sha1.New()
	io.WriteString(passwordHash, combination)

	t.Password = hex.EncodeToString(passwordHash.Sum(nil))
	t.Email = email
	t.Name = name

	log.Println(t)

	err := db.Save(&t).Error
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
