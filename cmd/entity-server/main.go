package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"sme-education-backend/cmd/entity-server/config"
	"sme-education-backend/cmd/entity-server/models"
	"sme-education-backend/cmd/entity-server/router"
)

func main() {
	//get env param
	env := os.Getenv("ENV")
	if env == "" {
		os.Setenv("ENV", "local")
		env = "local"
	}
	log.Printf("Env set to ** %s **", env)
	//get port param
	port := os.Getenv("PORT")
	if port == "" {
		os.Setenv("PORT", "7171")
		port = "7171"
	}
	log.Printf("Port set to ** %s **", port)

	basePath, _ := os.Getwd()
	config.InitFromFile("config/config.toml", basePath)
	closeFunc, _ := models.InitFromSQLLite(config.GetConfig().DbConnection)
	routersInit := router.InitRouter(env)
	readTimeout := time.Minute
	writeTimeout := time.Minute
	endPoint := fmt.Sprintf(":%s", port) //config.GetConfig().ServerPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	_ = server.ListenAndServe()
	defer closeFunc()
}
