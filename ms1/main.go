package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request){
		log.Println("Hello world")
		d, err := io.ReadAll(r.Body)
		if(err != nil){
			http.Error(rw, "Oops", http.StatusBadRequest)
			return 
		}
		log.Printf("Data %s\n",d)

		fmt.Fprintf(rw, "Hello %s", d)
	})
	http.ListenAndServe(":9090", nil)
	
}