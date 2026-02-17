package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

	log.Println("core-go listening on :8080")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}