package main

import (
	"learngos/errhandling/filelistringserver/filelisting"
	"net/http"
	"os"
)

type addHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler addHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				//http.Error(writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/list/", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
