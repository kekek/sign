package my

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"hash"
	"io"
	"net/url"
	"sort"
	"strings"
)

var signKeyList = []string{"acl", "uploads", "location", "cors"}

type Signer struct {
	secret string
}

func (s Signer) isParamSign(paramKey string) bool {
	for _, k := range signKeyList {
		if paramKey == k {
			return true
		}
	}
	return false
}

func (s Signer) getSubResource(params map[string]interface{}) string {
	// Sort
	keys := make([]string, 0, len(params))
	signParams := make(map[string]string)
	for k := range params {
		if s.isParamSign(k) {
			keys = append(keys, k)
			if params[k] != nil {
				signParams[k] = params[k].(string)
			}
		}
	}
	sort.Strings(keys)

	// Serialize
	var buf bytes.Buffer
	for _, k := range keys {
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(k)
		if _, ok := signParams[k]; ok {
			if signParams[k] != "" {
				buf.WriteString("=" + signParams[k])
			}
		}
	}
	return buf.String()
}

func (s Signer) getURLParams(params map[string]interface{}) string {
	// Sort
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Serialize
	var buf bytes.Buffer
	for _, k := range keys {
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(url.QueryEscape(k))
		if params[k] != nil && params[k].(string) != "" {
			buf.WriteString("=" + strings.Replace(url.QueryEscape(params[k].(string)), "+", "%20", -1))
		}
	}

	return buf.String()
}

func (s Signer) getSignURL(bucket, object, params string) string {
	return fmt.Sprintf("%s://%s%s?%s", "", "", "", params)
}

func (s Signer) signStr(toSignUrl string) string {
	h := hmac.New(func() hash.Hash { return sha256.New() }, []byte(s.secret))
	io.WriteString(h, toSignUrl)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signedStr
}
func (s Signer) signStrSha1(toSignUrl string) string {
	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(s.secret))
	io.WriteString(h, toSignUrl)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signedStr
}

func (s Signer) signStrMd5(toSignUrl string) string {
	h := hmac.New(func() hash.Hash { return md5.New() }, []byte(s.secret))
	io.WriteString(h, toSignUrl)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signedStr
}

// ------------- 签名
// 首先获取url
// 确定签名参数
// 签名并返回

// ------------- 验签
// 首先获取url
// 确定签名参数 + 签名
// 验证
