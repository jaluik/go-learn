package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
)

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
		//http.Error(writer, err.Error(),http.StatusInternalServerError)
		//return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	_, err = writer.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}
