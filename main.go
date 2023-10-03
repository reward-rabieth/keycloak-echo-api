package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/reward-rabieth/Authclk/api/middlewares"
	"github.com/reward-rabieth/Authclk/api/routes"
	"github.com/spf13/viper"
	"log"
	"strings"
)

// todo:Register user no Auth Requires
// todo:Get all products (those with role "viewer" can use)
// todo:create product (those with role "admin ' can use
func main() {
	initViper()

	e := echo.New()
	e.Use(ServerHeader)
	middlewares.InitEchoMiddleware(e, routes.InitPublicRoute, routes.InitProtectedRoute)
	var listenIp = viper.GetString("ListenIP")
	var listenPort = viper.GetString("ListenPort")
	log.Printf("will be  listening to %v:%v", listenIp, listenPort)
	err := e.Start(fmt.Sprintf("%v:%v", listenIp, listenPort))
	if err != nil {
		log.Fatal(err)
	}

}

func initViper() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("demo")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("unable to initialize viper: %w", err))
	}
	log.Println("viper config initialized")
}

// ServerHeader middleware adds a `Server` header to the response.
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//set the Server header
		c.Response().Header().Set(echo.HeaderServer, "Echo/4.0")
		// Set the AppName header
		c.Response().Header().Set("AppName", "AuthClk")
		return next(c)
	}
}
