package openyouku

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

type UploaderParam struct {
	Title      string        `json:"title"`
	Tags       string        `json:"tags"`
	Desc       string        `json:"desc"`
	IP         string        `json:"ip"`
	CategoryID int64         `json:"category_id"`
	UserID     string        `json:"user_id"`
	AttrValues []interface{} `json:"attr_values"`
	MD5        string        `json:"md5"`
	FileSize   int64         `json:"file_size"`
	FileFormat string        `json:"file_format"`
}

type Uploader struct {
	content  []byte
	fileName string
	apiParam map[string]string
	sdk      *SDK
}

func NewUploader(sdk *SDK, name string, content []byte) *Uploader {
	uploader := new(Uploader)
	uploader.content = content
	uploader.fileName = name
	uploader.apiParam = make(map[string]string, 0)
	uploader.Set("title", name)
	uploader.Set("file_size", fmt.Sprint(len(uploader.content)))
	uploader.sdk = sdk
	return uploader
}

func (uploader *Uploader) Set(k, v string) {
	uploader.apiParam[k] = v
}

//获取上传权限
func (uploader *Uploader) getUploadToken() {}

//开始上传
func (uploader *Uploader) doUpload() {}

func (uploader *Uploader) Start() (map[string]interface{}, error) {

	resp, err := uploader.sdk.Get("youku.api.vod.upload.video", uploader.apiParam)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	data, ok := resp.Data.(map[string]interface{})

	if !ok {
		return nil, errors.New("DATA EERROR")
	}
	uploadURL := data["upload_url"].(string)
	//token := data["token"].(string)
	uploadURL = uploadURL + "?id=" + data["fid"].(string) + "&sign=" + data["token"].(string)

	postBuffer := bytes.NewBufferString("")

	form := multipart.NewWriter(postBuffer)
	file, _ := form.CreateFormFile("file", "file.mp4")
	file.Write(uploader.content)
	client := &http.Client{}

	form.Close()

	request, err := http.NewRequest("POST", uploadURL, postBuffer)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", form.FormDataContentType())

	rsp, _ := client.Do(request)

	allBody, _ := ioutil.ReadAll(rsp.Body)

	log.Println("[Uploader]", string(allBody))

	apiResp := &Response{}

	err = json.Unmarshal(allBody, apiResp)

	if err != nil {
		return nil, err
	}
	return data, nil
}
