package openyouku

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
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
func (p *SysParams) SignParm(clientSecret string) map[string]string {

	signParam := map[string]string{
		"action":    p.Action,
		"timestamp": p.Timestamp,
		"client_id": p.ClientID,
		"version":   p.Version,
		"title":     "test",
	}

	var keys []string

	for key := range signParam {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	buffer := bytes.NewBufferString("")

	for _, k := range keys {
		fmt.Fprintln(os.Stderr, k)
		buffer.WriteString(k)
		val, ok := signParam[k]
		if ok {
			buffer.WriteString(val)
		}
	}
	buffer.WriteString(clientSecret)

	fmt.Fprintln(os.Stderr, buffer.String())

	signStr := buffer.String()
	signStr = url.QueryEscape(signStr)
	md5Hahser := md5.New()
	md5Hahser.Write([]byte(signStr))
	md5Sum := md5Hahser.Sum(nil)
	signStr = hex.EncodeToString(md5Sum[:])

	signParam["sign"] = signStr
	return signParam
}
