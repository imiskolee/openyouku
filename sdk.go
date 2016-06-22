package openyouku

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func (sdk *SDK) Get(action string, params map[string]interface{}) *Response {

	sysParam := &SysParams{}
	sysParam.Action = action
	sysParam.ClientID = sdk.ClientID
	sysParam.Format = "json"
	sysParam.Timestamp = fmt.Sprint(time.Now().Unix())
	sysParam.Version = "3.0"

	signParam := sysParam.SignParm(sdk.ClientSecret)

	jsonData, _ := json.Marshal(signParam)

	fmt.Fprintln(os.Stderr, string(jsonData))

	uri := fmt.Sprintf("opensysparams=%s&title=%s", jsonData, "test")

	//	uri = url.QueryEscape(uri)

	url := "https://openapi.youku.com/router/rest.json?" + uri

	fmt.Fprintln(os.Stderr, url)

	resp, _ := http.Get(url)

	allBody, _ := ioutil.ReadAll(resp.Body)

	fmt.Fprintln(os.Stderr, string(allBody))

	return nil
}

func (sdk *SDK) Post(action string, params map[string]interface{}) *Response {

	return nil
}

func (sdk *SDK) GetUploader(name string, content []byte) *Uploader {

	return NewUploader(sdk, name, content)

}
