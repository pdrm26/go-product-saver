package server

import (
	"fmt"
	"net/http"
)

const PORT = ":8080"

func Connect() {
	if err := http.ListenAndServe(PORT, nil); err != nil {
		fmt.Println("Server error: ", err)
	}

	fmt.Printf("Server connect successfully on %s", PORT)
}
