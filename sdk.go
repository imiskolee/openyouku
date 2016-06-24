package openyouku

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

type SDK struct {
	ClientID     string
	ClientSecret string
	AccessToken  string
	User         string
	Password     string
}

func (sdk *SDK) init() {
}

func (sdk *SDK) Get(action string, params map[string]string) (*Response, error) {

	sysParam := &SysParams{}
	sysParam.Action = action
	sysParam.ClientID = sdk.ClientID
	sysParam.Format = "json"
	sysParam.Timestamp = fmt.Sprint(time.Now().Unix())
	sysParam.Version = "3.0"

	signParam := sysParam.SignParm(sdk.ClientSecret, params)

	jsonData, _ := json.Marshal(signParam)

	values := make(url.Values)

	for k, v := range params {
		values.Set(k, v)
	}

	uri := fmt.Sprintf("opensysparams=%s&%s", jsonData, values.Encode())

	url := "https://openapi.youku.com/router/rest.json?" + uri
	fmt.Fprintln(os.Stderr, url)

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	allBody, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(allBody))

	response := &Response{}

	err = json.Unmarshal(allBody, response)

	return response, err

}

func (sdk *SDK) Post(action string, params map[string]interface{}) *Response {

	return nil
}

func (sdk *SDK) GetUploader(name string, content []byte) *Uploader {

	return NewUploader(sdk, name, content)

}
