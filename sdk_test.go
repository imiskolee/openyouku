package openyouku

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestSDKUploader(t *testing.T) {

	sdk := &SDK{}

	sdk.ClientID = "6124c55d3315c23d"
	sdk.ClientSecret = "6e9f10f7035d38ccdd9564490d75c858"

	//sdk.Get("youku.api.vod.upload.video", map[string]string{"title": "123"})

}

func TestUploader(t *testing.T) {

	sdk := &SDK{}

	sdk.ClientID = "6124c55d3315c23d"
	sdk.ClientSecret = "6e9f10f7035d38ccdd9564490d75c858"

	b, e := ioutil.ReadFile("/Users/MiskoLee/developer/go_workspace/src/github.com/imiskolee/openyouku/test.mp4")
	fmt.Fprintln(os.Stderr, e, len(b))
	//b := []byte{1, 2, 3}
	uploader := sdk.GetUploader("test.mp4", b)
	r, e := uploader.Start()

	fmt.Fprintln(os.Stderr, r, e)

}
