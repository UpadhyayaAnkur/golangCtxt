package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/mycontext", myhandler)
	http.ListenAndServe("127.0.0.1:2000", nil)
}

func myhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler invoked")
	fmt.Fprintln(w, "My handler invoked")
	defer fmt.Fprintln(w, "My handler ended")
	ctxt := r.Context()

	select {
	case <-ctxt.Done():
		err := ctxt.Err()
		log.Println("Context Error: ", err.Error())
		//http.Error(w, err.Error(), http.StatusInternalServerError)

	case <-time.After(time.Second * 5):
		fmt.Println("Timer of 5 seconds expired")
		http.Error(w, "Session expired", http.StatusRequestTimeout)
	}
}
