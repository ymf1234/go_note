package main

import (
	"io/ioutil"
	"learngos/errhandling/filelistringserver/filelisting"
	"net/http"
	"os"
)



func main() {
	http.HandleFunc("/list/",
		filelisting.HandleFileList())

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
