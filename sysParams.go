package openyouku

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/url"
	"sort"
)

//SysParams API定义的系统级别参数
type SysParams struct {
	Action      string `json:"action"`
	ClientID    string `json:"client_id"`
	AccessToken string `json:"access_token"`
	Format      string `json:"format"`
	Timestamp   string `json:"timestamp"`
	Version     string `json:"version"`
	Sign        string `json:"sign"`
}

func (p *SysParams) String() (string, error) {

	data, err := json.Marshal(p)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

//SignParm is a method to sign this request
func (p *SysParams) SignParm(clientSecret string, param map[string]string) map[string]string {

	signParam := map[string]string{
		"action":    p.Action,
		"timestamp": p.Timestamp,
		"client_id": p.ClientID,
		"version":   p.Version,
	}

	var keys []string
	for key := range signParam {
		keys = append(keys, key)
	}

	for key := range param {
		keys = append(keys, key)
	}
	//对系统参数与业务参数的key进行排序
	sort.Strings(keys)
	buffer := bytes.NewBufferString("")
	for _, k := range keys {
		//如果冲突，则系统参数在前，所以，这里的顺序很重要
		if s, ok := signParam[k]; ok {
			buffer.WriteString(k)
			buffer.WriteString(s)
		}
		if s, ok := param[k]; ok {
			buffer.WriteString(k)
			buffer.WriteString(s)
		}
	}
	buffer.WriteString(clientSecret)
	signStr := buffer.String()
	signStr = url.QueryEscape(signStr)
	md5Hahser := md5.New()
	md5Hahser.Write([]byte(signStr))
	md5Sum := md5Hahser.Sum(nil)
	signStr = hex.EncodeToString(md5Sum[:])
	signParam["sign"] = signStr
	return signParam
}
