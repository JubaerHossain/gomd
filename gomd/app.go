package gomd

import (
	"fmt"
	"net/http"

	"github.com/JubaerHossain/gomd/config"
	"github.com/gin-gonic/gin"
)

var Res ResponseBuilder

func LoadRoute() {

	gin.ForceConsoleColor()

	Router = gin.Default()

	Router.GET("/health-check", func(c *gin.Context) {
		Res.Code(http.StatusOK).Message("Up and Running").Data(gin.H{"app": "OK"}).Json(c)
	})
}

func LoadConfig() {
	config.Config = config.NewConfig()

	config.Config.Load()
}

func InitializeLogger() LoggerBuilder {
	return NewLogger()
}

func NoSqlConnection() {
	Mongo = NewNoSqlDB()
}

func New() {
	LoadConfig()
	LoadRoute()
}

func Start() {
	NoSqlConnection()
	InitializeLogger()
}

func Run() {
	if Mongo != nil {
		defer Mongo.Client.Disconnect(Mongo.Ctx)
	}

	port, _ := config.Config.Int("App.Port")

	if port == 0 {
		port = 8080
	}

	fmt.Sprintf("Server is running on port :%d", port)

	Router.Run(fmt.Sprintf(":%d", port))
}
