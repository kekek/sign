package sign

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"
)

type Url struct {
	Path  string
	Query url.Values
}

func NewUrl(rawUrl string, options ...func(*Url)) (*Url, error) {
	s := &Url{}

	for _, opt := range options {
		opt(s)
	}

	vals := url.Values{}
	var err error
	u, err := url.QueryUnescape(rawUrl)
	if err != nil {
		return s, err
	}

	ex := strings.Split(u, "?")
	if len(ex) <= 1 {
		return s, errors.New("参数错误:无鉴权参数")
	}
	vals, err = url.ParseQuery(ex[1])
	if err != nil {
		return s, err
	}

	if !vals.Has("sign") { // 有签名字段，就不添加时间戳
		vals.Add("ts", fmt.Sprintf("%d", time.Now().Unix()))
	}

	s.Path = ex[0]
	s.Query = vals

	return s, nil
}

func (u *Url) GetToSignStr(signKeyList []string) string {
	params := u.getSubResource(signKeyList)
	return fmt.Sprintf(u.Path+"?%s", params)
}

func (u *Url) GetSignedUrl(keyName, val string, signKeyList []string) string {
	params := u.getSubResource(signKeyList)
	return fmt.Sprintf(u.Path+"?%s&%s=%s", params, keyName, val)
}

func (u *Url) isParamSign(paramKey string, signKeyList []string) bool {
	for _, k := range signKeyList {
		if paramKey == k {
			return true
		}
	}
	return false
}

func (u *Url) getSubResource(signKeyList []string) string {
	// Sort
	params := u.Query
	res := url.Values{}
	for k := range params {
		if u.isParamSign(k, signKeyList) {
			if params[k] != nil {
				res[k] = params[k]
			}
		}
	}
	return res.Encode()
}

func (u *Url) GetParams(key string) string {
	// Sort
	return u.Query.Get(key)
}
