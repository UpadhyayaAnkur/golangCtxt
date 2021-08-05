package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctxt := context.Background()
	cctxt, cfunc := context.WithTimeout(ctxt, time.Second*4)
	req, err := http.NewRequestWithContext(cctxt, http.MethodGet, "http://127.0.0.1:2000/mycontext", nil)

	defer cfunc()

	rsp, err := http.DefaultClient.Do(req)
	//rsp, err := http.Get("http://localhost:2000/mycontext")
	if err != nil {
		log.Fatal("Get Error: ", err.Error())
	}

	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		log.Fatal("Rsp status code: ", rsp.StatusCode)
	}

	io.Copy(os.Stdout, rsp.Body)
}
