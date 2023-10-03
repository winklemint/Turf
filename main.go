package main

import (
	"fmt"
	"turf/config"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
	config.SyncDB()

}

func main() {
	fmt.Println("Helloworld")
}
