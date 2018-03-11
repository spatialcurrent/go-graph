package backends

import (
  "bytes"
  "errors"
  "encoding/json"
  //"fmt"
  "io/ioutil"
  "net/http"
  "strconv"
  "time"
  //"gopkg.in/yaml.v2"
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

  if resp.StatusCode != 200 {
    return resp_body, errors.New("Backend returned status code "+strconv.Itoa(resp.StatusCode)+".\n"+string(resp_body))
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

  if resp.StatusCode != 200 {
    return resp_body, errors.New("Backend returned status code "+strconv.Itoa(resp.StatusCode)+".\n"+string(resp_body))
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

  //fmt.Println("Request Body")
  //req_yml, err := yaml.Marshal(data)
	//if err != nil {
	//	return err
	//}
  //fmt.Println(string(req_yml))

	resp_body, err := b.Post(u, "application/json", req_body, b.DefaultTimeout)
	if err != nil {
		return err
	}

  //fmt.Println("Response Body")
  //fmt.Println(len(resp_body))
  //fmt.Println(string(resp_body))

  if len(resp_body) > 0 {
    err = json.Unmarshal(resp_body, target)
  	if err != nil {
  		return err
  	}
  }

	return nil
}
