package main

import (
	"context"
	"fmt"
	"github.com/linqcod/student-testing-app/app/cmd/api"
	"github.com/linqcod/student-testing-app/app/pkg/config"
	"github.com/linqcod/student-testing-app/app/pkg/database"
	"github.com/spf13/viper"
	"log"
)

func init() {
	config.InitConfig()
}

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	port := fmt.Sprintf(":%d", viper.GetInt("app.port"))

	app := api.InitRouter(context.Background(), db)
	app.Run(port)
}
