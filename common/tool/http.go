package tool

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"time"
)

func HttpPost(url string, req any, resp any) (err error) {
	var request *http.Request
	if req == nil || reflect.ValueOf(req).IsNil() {
		request, err = http.NewRequest(http.MethodPost, url, nil)
		if err != nil {
			return
		}
	} else {
		var body []byte
		if body, err = json.Marshal(req); err != nil {
			return
		}

		request, err = http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
		if err != nil {
			return
		}

		request.Header.Add("Content-Type", "application/json")
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if resp != nil && !reflect.ValueOf(resp).IsNil() {
		if err = json.NewDecoder(response.Body).Decode(resp); err != nil {
			return err
		}
	}

	return nil
}

func GetWithHeader(path string, m map[string]string, proxy string) ([]byte, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancelFunc()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	if m != nil {
		for k, v := range m {
			httpReq.Header.Set(k, v)
		}
	}
	httpReq.Header.Add("Content-Type", "application/json")

	client := http.DefaultClient
	if proxy != "" {
		proxyAddress, _ := url.Parse(proxy)
		client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyAddress),
			},
		}
	}

	httpRsp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	rspBody, err := io.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}
	return rspBody, nil
}
