package database

import "fmt"

var connection string

func init() {
	fmt.Println("Init is called")
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
