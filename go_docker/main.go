package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
	//fmt.Fprintf(w, "Hello World: " + r.RequestURI)
	resp, err := http.Get("https://www.nrk.no" +  r.RequestURI)
    if err != nil {
        print(err)
    }
   
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        print(err)
    }
	
	fmt.Fprintf(w, string(body))
    //fmt.Print(string(body))
}

func main() {
	fmt.Printf("Listening on :8080")
    http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
	fmt.Printf("Listening on :8080 ended")
}