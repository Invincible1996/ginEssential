package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goEssential/common"
	"github.com/goEssential/routes"
	"github.com/spf13/viper"
	"os"
)

func main() {
	InitConfig()
	common.GetDB()
	r := gin.Default()
	r = routes.CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("init config err:" + err.Error())
	}
}
