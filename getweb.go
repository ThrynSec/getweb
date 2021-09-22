package getweb

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func Download(filepath string, url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func GetAPI(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return []byte(""), err
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte(""), err
	}
	return responseData, err
}
