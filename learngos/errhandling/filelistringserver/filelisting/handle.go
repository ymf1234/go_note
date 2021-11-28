package filelisting

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	index := strings.Index(request.URL.Path, prefix)

	if index != 0 {
		return errors.New("path must start with " + prefix)
	}

	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		//http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)
	return nil
}
