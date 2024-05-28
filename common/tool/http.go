package tool

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
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
