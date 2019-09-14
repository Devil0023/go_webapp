package ghttp

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//Get
func Get(url string, params map[string]string) (resBody string, err error) {

	url += "?"

	for key, value := range params {
		url += key + "=" + value + "&"
	}

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

//Post
func Post(url string, params map[string]string) (resBody string, err error) {

	paramString := ""

	for key, value := range params {
		paramString += key + "=" + value + "&"
	}

	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(paramString))

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

//Download
func Download(url, savePath string) (written int, err error) {

	resp, err := http.Get(url)

	fmt.Println(url)

	if err != nil {
		return 0, err
	}

	file, err := os.Create(savePath)

	defer file.Close()

	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	buf := make([]byte, 4096)

	for {

		n, err := resp.Body.Read(buf)
		w, _ := file.Write(buf[:n])

		written += w

		if err == io.EOF {
			break
		} else if err != nil {
			return 0, err
		}
	}

	return written, nil
}

//Do TODO 需要测试
func Do(method, url string, params, headers map[string]string) (resBody string, err error) {

	paramString := ""

	for key, value := range params {
		paramString += key + "=" + value + "&"
	}

	reqBody := strings.NewReader(paramString)

	req, err := http.NewRequest(method, url, reqBody)

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil

}
