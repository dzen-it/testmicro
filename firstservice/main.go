package main

import (
	"context"
	"log"
	"net/http"

	"testmicro/secondservice/proto/request"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	web "github.com/micro/go-web"
)

var Client request.SecondClient

func main() {
	router := gin.Default()
	router.GET("/", GetRequest)
	service := web.NewService(
		web.Name("go.micro.srv.firstservice"),
		web.Address(":8088"),
		web.Handler(router),
	)

	service.Init()
	Client = request.NewSecondClient("go.micro.srv.secondservice", client.DefaultClient)

	if err := service.Run(); err != nil {
		log.Fatalln("ERROR", err)
	}
}

func GetRequest(c *gin.Context) {
	rsp, _ := Client.Get(context.TODO(), &request.GetRequest{
		Name: c.Query("name"),
	})
	c.JSON(http.StatusOK, gin.H{"Name": rsp.Resp})
	return
}
