package sugar

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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

	fmt.Print("\n======http.Get======\n", string(respData))

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

	fmt.Print("\n======http.Post======\n", string(respData))

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

	fmt.Print("\n======http.Post======\n", string(respData))

	err = json.Unmarshal(respData, responseData)
	if err != nil {
		return err
	}
	return nil
}
