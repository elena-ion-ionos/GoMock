package main

import (
	"log"
	"net/http"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("failed to start server")
	}
}
