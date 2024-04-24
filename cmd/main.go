package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/vsafonkin/involta/internal/api"
	"github.com/vsafonkin/involta/internal/config"
	"github.com/vsafonkin/involta/internal/db"
	"github.com/vsafonkin/involta/internal/model"

	"github.com/vsafonkin/involta/internal/gendata"
)

var configPath = flag.String("config", "", "main -config <path>")

func main() {
	flag.Parse()
	if *configPath == "" {
		fmt.Println("You need to set path to config file: ./main -config <path>")
		os.Exit(1)
	}

	err := config.LoadConfig(*configPath)
	if err != nil {
		fmt.Println(err)
	}
	conn, err := db.NewConnect()
	if err != nil {
		panic(err)
	}

	if err := model.NewDocModel(conn); err != nil {
		panic(err)
	}

	result, err := conn.List(config.Namespace())
	if err != nil {
		panic(err)
	}
	if len(result) == 0 {
		fmt.Println("Generate data...")
		if err := gendata.GenRandomData(100); err != nil {
			panic(err)
		}
	}

	if err := api.RunAPIServer(); err != nil {
		panic(err)
	}
}
