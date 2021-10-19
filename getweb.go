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

func GetAPI(url string, user string, pass string) ([]byte, error) {
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return []byte(""), err
	}

	request.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString(getUserPass()))

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return []byte(""), err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return []byte(""), err
	}

	return responseData, err
}

//This function is not meant to stay the way it is. Please encode your credentials !
func getUserPass() string {
	var userStr := "user"
	var passStr := "pass"
	
	return user+":"+pass
}

func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(getUserPass()))
	
	return nil
}
