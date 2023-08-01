package tools

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/*
GetUrl 发送 Get 请求
*/
func GetUrl(addr string) ([]byte, error) {
	resp, err := http.Get(addr)
	if err != nil {
		return nil, errors.Errorf(fmt.Sprintf("http get url(%v) failed, err: %v", addr, err))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("get url(%v) failed, statusCode: %d", addr, resp.StatusCode))
	}

	log.Debugf("get url:%s success, status:%d", addr, resp.StatusCode)
	return ioutil.ReadAll(resp.Body)
}

/*
PostUrl 发送 Post 请求
contentType: application/json  application/x-www-form-urlencoded
*/
func PostUrl(addr, contentType string, params []byte) ([]byte, error) {
	if contentType == "" {
		contentType = "application/json"
	}
	// 发起请求
	resp, err := http.Post(addr, contentType, bytes.NewReader(params))
	if err != nil {
		log.Errorf("post url:%s failed with error:%s", addr, err.Error())
		return nil, errors.WithMessage(err, "post url failed")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("post url(%v) failed, statusCode: %d", addr, resp.StatusCode))
	}

	log.Debugf("post url:%s success, status:%d", addr, resp.StatusCode)

	return ioutil.ReadAll(resp.Body)
}

/*
GetUrlHeader 携带请求头参数发送 Get 请求
*/
func GetUrlHeader(addr string, headers map[string]string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, addr, nil)
	if err != nil {
		return nil, err
	}
	for key, header := range headers {
		req.Header.Set(key, header)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

/*
PostHeader 携带请求头参数发送 Post 请求
*/
func PostHeader(url string, msg []byte, headers map[string]string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(msg)))
	if err != nil {
		return nil, err
	}
	for key, header := range headers {
		req.Header.Set(key, header)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

/*
PostUrlEncode 适用于 application/x-www-form-urlencoded
*/
func PostUrlEncode(postUrl string, msg, headers map[string]string) ([]byte, error) {
	client := &http.Client{}
	// post要提交的数据
	DataUrlVal := url.Values{}
	for key, val := range msg {
		DataUrlVal.Add(key, val)
	}

	req, err := http.NewRequest("POST", postUrl, strings.NewReader(DataUrlVal.Encode()))
	if err != nil {
		return nil, err
	}
	for key, header := range headers {
		req.Header.Set(key, header)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func PostForm(url string, data url.Values) ([]byte, error) {
	resp, err := http.PostForm(url, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	return body, nil
}
