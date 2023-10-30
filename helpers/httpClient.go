package helpers

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"log"
	"m2ps/config"
	"net/http"
)

func CallAPI(method string, body []byte, suffixUrl string) (*string, error) {
	url := config.BASEApiUrl + suffixUrl

	httpReq, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		log.Println("Err http.NewRequest ", url, " :", err.Error())
		return nil, err
	}
	defer httpReq.Body.Close()

	httpReq.Close = true
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Basic bWtwbW9iaWxlOk1LUG1vYmlsZTEyMzQ1NkA=")
	httpReq.Header.Set("Connection", "close")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	log.Println("URL:", url)
	log.Println("Request:", string(body))

	response, err := client.Do(httpReq)
	if err != nil {
		log.Println("Err - client.Do :", err.Error())
		return nil, err
	}

	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Err ReadAll :", err.Error())
		return nil, err
	}

	bodyString := string(body)
	log.Println("Response:", bodyString)

	//err = json.Unmarshal(body, &result)
	//if err != nil {
	//	log.Println("Err Unmarshall :", err.Error())
	//	return nil, err
	//}

	return &bodyString, nil
}

func CallPaymentAPI(method string, body []byte, suffixUrl string) (*string, error) {
	url := config.BASEPaymentUrl + suffixUrl

	httpReq, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		log.Println("Err http.NewRequest ", url, " :", err.Error())
		return nil, err
	}
	defer httpReq.Body.Close()

	httpReq.Close = true
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Basic bWtwbW9iaWxlOm1rcG1vYmlsZTEyMw==")
	httpReq.Header.Set("Connection", "close")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	log.Println("URL:", url)
	log.Println("Request:", string(body))

	response, err := client.Do(httpReq)
	if err != nil {
		log.Println("Err - client.Do :", err.Error())
		return nil, err
	}

	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Err ReadAll :", err.Error())
		return nil, err
	}

	bodyString := string(body)
	log.Println("Response:", bodyString)

	//err = json.Unmarshal(body, &result)
	//if err != nil {
	//	log.Println("Err Unmarshall :", err.Error())
	//	return nil, err
	//}

	return &bodyString, nil
}
