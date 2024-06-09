/*
Stack is a small service that lets me save things for later.

I use it to put a pin on links, documents and ideas that I
want to rediscover at a later stage.
*/
package main

import (
	"github.com/peder2911/stack/pkg/stack"
	"github.com/peder2911/stack/internal/server"
	"github.com/peder2911/stack/internal/files"
	"net/http"
	"fmt"
)

func main() {
	files,err := files.DefaultFiles()
	stack := stack.Stack(files.Database)
	server,err := server.NewServer(&stack)
	port := "8000"
	fmt.Printf("Serving on %s\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s",port),server)
	if err != nil {
		panic(err)
	}
}
