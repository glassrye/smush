package main

import (
	_ "github.com/go-sql-driver/mysql"
)

/*
	This main.go file is doing way too much logic and heavy liftings
	We need to make a commangs.go file that has the functions from the cli
	Main really only needs to initiate the CLI code
*/


func main() {
	// Set some variables for application user
	cli()
}
