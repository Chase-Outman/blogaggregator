package main

import (
	"fmt"

	"github.com/Chase-Outman/blogaggregator/internal/config"
)

func main() {
	c, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)

	//add user info

	c.SetUser("Chase")
	newc, _ := config.Read()

	fmt.Println(newc)
}
