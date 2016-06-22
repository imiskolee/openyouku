package openyouku

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
	Param    UploaderParam
	content  []byte
	fileName string
}

func NewUploader(sdk *SDK, name string, content []byte) *Uploader {
	uploader := new(Uploader)
	uploader.content = content
	uploader.fileName = name
	return uploader
}

//获取上传权限
func (uploader *Uploader) getUploadToken() {}

//分片
func (uploader *Uploader) makeSlice() {}

//开始上传
func (uploader *Uploader) doUpload() {}
