package util

import (
	"bytes"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

const TIMEOUT = 4

func Get(url string, headerData map[string]string) (string, error) {
	client := &http.Client{Timeout: time.Duration(TIMEOUT * time.Second)}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	if len(headerData) > 0 {
		for k, v := range headerData {
			request.Header.Set(k, v)
		}
	}

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	if response == nil {
		return "", nil
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func Post(url string, params []byte, headerData map[string]string) ([]byte, error) {
	client := &http.Client{Timeout: time.Duration(TIMEOUT * time.Second)} //4秒超时
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(params))
	if err != nil {
		return nil, err
	}

	if len(headerData) > 0 {
		for k, v := range headerData {
			req.Header.Set(k, v)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
func PostV2(url string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Println(err)
		return "", err
	}
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(body), nil
}

func GetV2(url string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logx.Info(err)
		return "", err
	}
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		logx.Info(err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logx.Info(err)
		return "", err
	}
	return string(body), nil
}

func SetHeader(w http.ResponseWriter, filename string) {
	// 获取文件扩展名
	ext := filepath.Ext(filename)
	// 根据扩展名设置ContentType
	switch ext {
	case ".jpg", ".jpeg":
		w.Header().Set("Content-Type", "image/jpeg")
	case ".gif":
		w.Header().Set("Content-Type", "image/gif")
	case ".png":
		w.Header().Set("Content-Type", "image/png")
	default:
		w.Header().Set("Content-Type", "application/octet-stream")
	}
}
