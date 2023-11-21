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
	config.MustLoadConfig(".env")
}

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	port := fmt.Sprintf(":%s", viper.GetString("SERVER_PORT"))

	app := api.InitRouter(context.Background(), db)
	app.Run(port)
}
