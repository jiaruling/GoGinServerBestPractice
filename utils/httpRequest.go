package utils

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/tidwall/gjson"
)

/*
   功能说明: http请求
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1.sql/19 10:44
*/

func RequestHasBody(method, url string, reqBody *[]byte) (respBody *[]byte, err error) {
	return request(method, url, reqBody)
}

func RequestHasNoBody(method, url string) (respBody *[]byte, err error) {
	return request(method, url, nil)
}

// 解析请求体数据
func ParseRespBody(respBody *[]byte) (data []byte, err error) {
	/**
	情况一:
		 {
			"code": 200
	        "data": {
						"id": 1.sql,
						"name": "123"
					}
			"msg": "success"
		 }
	情况二:
	     {
			"code": 200
	        "data": {
						"data": [{"id": 1.sql, "name": "A"}, {"id": 2, "name": "B"}],
	   					"page": 1.sql,
						"page_size": 10,
						"total": 2
					}
			"msg": "success"
		 }
	*/
	bodyStr := string(*respBody)
	code := gjson.Get(bodyStr, "code").Int()
	if code == 0 || (code >= 200 && code < 300) {
		// 请求数据成功
		if gjson.Get(bodyStr, "data.data").String() == "" {
			// data 格式为单个
		} else {
			// data 格式为
		}
	} else {
		// 请求数据失败
		return nil, errors.New(gjson.Get(bodyStr, "msg").String())
	}
	return
}

func request(method, url string, reqBody *[]byte) (respBody *[]byte, err error) {
	var req *http.Request
	url = spliceUrl(url)
	client := &http.Client{Timeout: 1 * time.Second}
	if reqBody == nil {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, bytes.NewReader(*reqBody))
	}
	if err != nil {
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}()
	code := resp.StatusCode
	if !(code >= 200 && code < 300) {
		return nil, errors.New(strconv.Itoa(code))
	}
	// 获取响应数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	respBody = &body
	return respBody, nil
}

// 拼接 ulr
func spliceUrl(url string) string {
	matched, _ := regexp.MatchString("^http", url)
	if matched {
		return url
	}
	return "http://" + url
}
