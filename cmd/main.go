package main

import (
	"proyecto/cli"
	"proyecto/db"
)

func main() {
	db.StartDB()
	cli.Run()
}
