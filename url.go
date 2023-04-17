package sign

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"
)

type Url struct {
	Path                 string
	Query                url.Values
	signKeyName          string // 签名key
	timestampName        string
	toSignParamsNameList []string // 待签名的参数名称
}

func NewUrl(rawUrl string, opts ...Option) (*Url, error) {
	s := &Url{
		signKeyName:          "sign",
		timestampName:        "ts",
		toSignParamsNameList: []string{"id", "name"},
	}

	for _, opt := range opts {
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

	if !vals.Has(s.signKeyName) { // 有签名字段，就不添加时间戳
		vals.Add(s.timestampName, fmt.Sprintf("%d", time.Now().Unix()))
	}

	s.Path = ex[0]
	s.Query = vals

	return s, nil
}

func (u *Url) GetToSignStr() string {
	params := u.getSubResource()
	return fmt.Sprintf(u.Path+"?%s", params)
}

func (u *Url) GetSignedUrl(val string) string {
	params := u.getSubResource()
	return fmt.Sprintf(u.Path+"?%s&%s=%s", params, u.signKeyName, val)
}

func (u *Url) isParamSign(paramKey string, signKeyList []string) bool {
	for _, k := range signKeyList {
		if paramKey == k {
			return true
		}
	}
	return false
}

func (u *Url) getSubResource() string {
	// Sort
	params := u.Query
	res := url.Values{}
	for k := range params {
		if k == u.signKeyName {
			continue
		}
		if u.isParamSign(k, u.toSignParamsNameList) {
			if params[k] != nil {
				res[k] = params[k]
			}
		}
	}
	return res.Encode()
}

func (u *Url) GetTimestamp() string {
	return u.getParams(u.timestampName)
}

func (u *Url) GetSign() string {
	return u.getParams(u.signKeyName)
}

func (u *Url) getParams(key string) string {
	// Sort
	return u.Query.Get(key)
}
