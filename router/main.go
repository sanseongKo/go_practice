package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"router/model"
	context "router/util"
	"strconv"
)

var r context.Router

func main() {
	r := gin.Default()

	r.POST("/test", TestPostFunction)
	r.PUT("/test", testPutFunction)
	r.GET("/test", testGetFunction)
	r.DELETE("/test", testDeleteFunction)

	//r.Use(gin.LoggerWithFormatter(testLogger))

	_ = r.Run(":8080")
}

func TestPostFunction(c *gin.Context) {
	fmt.Println(reflect.TypeOf(c).String())

	con := context.GinContext{
		Context: c,
	}

	r = &con

	//Request Body로 받았을 때 데이터 바인딩
	req := model.RequestPostStructure{}

	if err := r.RequestBody(req); err != nil {
		//return nil
	}

	_ = model.ResponsePostStructure{
		Id:      req.Id,
		Name:    req.Name,
		Message: req.Message,
	}

	//return con.Response(response)
}

func testPutFunction(c *gin.Context) {
	con := context.GinContext{
		Context: c,
	}

	r = &con

	id, _ := strconv.Atoi(r.RequestQueryParam("id"))

	req := &model.RequestPutStructure{}

	if err := r.RequestBody(req); err != nil {
		//return nil
	}

	_ = model.ResponsePutStructure{
		Id:      id,
		Name:    req.Name,
		Message: req.Message,
	}

	//return r.Response(response)
}

func testGetFunction(c *gin.Context) {
	con := context.GinContext{
		Context: c,
	}

	r = &con

	id, _ := strconv.Atoi(r.RequestQueryParam("id"))

	_ = model.ResponseGetStructure{
		id, "테스트 성공",
	}

	//return r.Response(res)
}

func testDeleteFunction(c *gin.Context) {
	_ = *new(error)

	//return err
}
