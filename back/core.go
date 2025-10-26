package main

import (
	"fmt"
	"log/slog"
	"net/http"

	// "github.com/IliyaG10/Download-Manager/common"
)


func main() {
	
	link := "https://dl.konkur.in/post/Book/Paye/Adams-Calculus-A-Complete-Course-8th-Edition-%5Bkonkur.in%5D.pdf"
	
	resp , err := http.Get(link)
	if err != nil {
		slog.Info(err.Error())
	}

	
	fmt.Println(resp.Status)


	fmt.Println("\n--- Headers ---")
    for key, values := range resp.Header {
        for _, value := range values {
            fmt.Printf("%s: %s\n", key, value)
        }
    }

	contentLengthStr := resp.Header.Get("Content-Length")

	fmt.Println(contentLengthStr)



	
}