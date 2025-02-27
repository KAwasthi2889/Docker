package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	secretMessage := os.Getenv("SECRET_MESSAGE")
	if secretMessage == "" {
		fmt.Println("Error getting message")
		return
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	message := "Listening the secret message: " + secretMessage + " in port " + port + "\n"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.OpenFile("/data/message.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("File Error!!")
		}
		file.WriteString(message)
		file.Seek(0, 0)
		io.Copy(w, file)
		file.Close()
	})
	fmt.Printf("Listening on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
