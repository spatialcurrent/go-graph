package backends

import (
  "bytes"
  "encoding/json"
  "io/ioutil"
  "net/http"
  "time"
)

type HttpBackend struct {
  DefaultTimeout int `json:"default_timeout" bson:"default_timeout" yaml:"default_timeout" hcl:"default_timeout"`
}

func (b *HttpBackend) Get(url string, contentType string, timeout int) ([]byte, error) {

	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, err
	}

	req.Header.Set("content-type", contentType)

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return resp_body, nil
}

func (b *HttpBackend) Post(url string, contentType string, data []byte, timeout int) ([]byte, error) {

	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return []byte{}, err
	}

	req.Header.Set("content-type", contentType)

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return resp_body, nil
}

func (b *HttpBackend) GetJson(u string, target interface{}) error {

	resp_body, err := b.Get(u, "application/json", b.DefaultTimeout)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp_body, target)
	if err != nil {
		return err
	}

	return nil
}

func (b *HttpBackend) PostJson(u string, data map[string]interface{}, target interface{}) error {

	req_body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp_body, err := b.Post(u, "application/json", req_body, b.DefaultTimeout)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp_body, target)
	if err != nil {
		return err
	}

	return nil
}
