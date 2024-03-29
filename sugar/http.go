package sugar

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// HttpGet get请求 responseData必须是指针类型
func HttpGet(urlStr string, requestData map[string]string, responseData interface{}, headerData map[string]string) error {
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return err
	}
	if len(headerData) > 0 {
		for k, v := range headerData {
			req.Header.Set(k, v)
		}
	}
	// 设置参数
	params := req.URL.Query()
	for k, v := range requestData {
		params.Add(k, v)
	}
	req.URL.RawQuery = params.Encode()
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return err
	}
	defer resp.Body.Close()
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("response StatusCode is not StatusOK; code:%+v", resp.StatusCode))
	}

	fmt.Println("\n======http.Get======\n", req.URL, "\n", string(respData))

	err = json.Unmarshal(respData, responseData)
	if err != nil {
		return err
	}
	return nil
}

// HttpPost post请求 responseData必须是指针类型
func HttpPostForm(urlStr string, requestData map[string]string, responseData interface{}, headerData map[string]string) error {
	formData := url.Values{}
	for k, v := range requestData {
		formData.Set(k, v)
	}
	data := formData.Encode()
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(data))
	if err != nil {
		fmt.Printf("%s", err.Error())
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if len(headerData) > 0 {
		for k, v := range headerData {
			req.Header.Set(k, v)
		}
	}
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return err
	}
	defer resp.Body.Close()
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("response StatusCode is not StatusOK; code:%+v", resp.StatusCode))
	}

	fmt.Println("\n======http.Post======\n", formData, string(respData))

	err = json.Unmarshal(respData, responseData)
	if err != nil {
		return err
	}
	return nil
}

// HttpPost post请求 header json responseData必须是指针类型
func HttpPost(urlStr string, requestData interface{}, responseData interface{}, headerData map[string]string) error {
	data, _ := json.Marshal(requestData)
	req, err := http.NewRequest("POST", urlStr, bytes.NewReader(data))
	if err != nil {
		fmt.Printf("%s", err.Error())
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	if len(headerData) > 0 {
		for k, v := range headerData {
			req.Header.Set(k, v)
		}
	}
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return err
	}
	defer resp.Body.Close()
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("response StatusCode is not StatusOK; code:%+v", resp.StatusCode))
	}

	fmt.Println("\n======http.Post======\n", string(data), string(respData))

	err = json.Unmarshal(respData, responseData)
	if err != nil {
		return err
	}
	return nil
}

func PHPPostRequest(urlStr string, requestData interface{}, responseData interface{}, headerData map[string]string) error {
	data, _ := json.Marshal(requestData)
	respData, err := BaseRequest(urlStr, "POST", data, headerData, map[string]string{})
	extra.RegisterFuzzyDecoders()
	err = jsoniter.Unmarshal(respData, responseData)
	if err != nil {
		return err
	}
	return nil
}

func PHPGetRequest(urlStr string, requestData map[string]string, responseData interface{}, headerData map[string]string) error {
	respData, err := BaseRequest(urlStr, "GET", nil, headerData, requestData)
	extra.RegisterFuzzyDecoders()
	err = jsoniter.Unmarshal(respData, responseData)
	if err != nil {
		return err
	}
	return nil
}

func BaseRequest(url string, method string, data []byte, headerData map[string]string, queryParams map[string]string) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(data))
	params := req.URL.Query()
	for k, v := range queryParams {
		params.Add(k, v)
	}
	req.URL.RawQuery = params.Encode()
	if err != nil {
		fmt.Printf("%s", err.Error())
		return []byte{}, err
	}
	if len(headerData) > 0 {
		for k, v := range headerData {
			req.Header.Set(k, v)
		}
	}
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return []byte{}, err
	}
	defer resp.Body.Close()
	respData, err := io.ReadAll(resp.Body)

	fmt.Println("\n======http."+method+"======\n", url, string(data), string(respData))

	if err != nil {
		fmt.Printf("%s", err.Error())
		return []byte{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return []byte{}, errors.New(fmt.Sprintf("response StatusCode is not StatusOK; code:%+v", resp.StatusCode))
	}
	return respData, nil
}
