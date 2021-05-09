package main

import "fmt"

func main() {
	app := New()
	err := app.Run()
	if err != nil {
		fmt.Printf("%+v", err)
	}
}
