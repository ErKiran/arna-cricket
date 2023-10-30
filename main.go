package main

import (
	"fmt"
	"os"

	controllers "arna-cricket/api"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("err", err)
		os.Exit(0)
	}
	if err := controllers.InitRouter().Run(); err != nil {
		fmt.Println("unable to init routes", err)
	}
}
