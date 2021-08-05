package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctxt := context.Background()
	//cctxt, cfunc := context.WithCancel(ctxt)
	cctxt, cfunc := context.WithTimeout(ctxt, time.Second)
	go myfunc(cfunc)

	select {
	case <-cctxt.Done():
		fmt.Println("Main got child ctxt cancel notification")
	case <-ctxt.Done():
		fmt.Println("Main got parent ctxt cancel notification")
	}
}

func myfunc(cfunc context.CancelFunc) {
	time.Sleep(2 * time.Second)
	fmt.Println("waking after sleep and cancel ctxt")
	cfunc()
}
