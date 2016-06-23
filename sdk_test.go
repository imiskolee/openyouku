package openyouku

import "testing"

func TestSDKUploader(t *testing.T) {

	sdk := &SDK{}

	sdk.ClientID = "6124c55d3315c23d"
	sdk.ClientSecret = "6e9f10f7035d38ccdd9564490d75c858"

	sdk.Get("youku.api.vod.upload.video", map[string]string{"title": "123"})

}

func TestUploader(t *testing.T) {

	sdk := &SDK{}

	sdk.ClientID = "6124c55d3315c23d"
	sdk.ClientSecret = "6e9f10f7035d38ccdd9564490d75c858"

	content := make([]byte, 100)

	uploader := sdk.GetUploader("test.mp4", content)

	uploader.Start()

}
