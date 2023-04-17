package sign

import (
	"fmt"
	"strconv"
	"time"
)

var defaultSignKeyList = []string{"id", "timestamp", "ts", "age"}

type Signature struct {
	Secret           string   // 签名密码
	SignParamKeyName string   // 签名key
	SignParamKeyList []string // 参与签名的参数名称
	SignFunc         SignFunc // 签名函数
}

var defaultSignature = NewSignature("123456")

func NewSignature(secret string, opts ...SignatureOption) *Signature {
	s := &Signature{
		Secret:           secret,
		SignParamKeyName: "sign",
		SignParamKeyList: defaultSignKeyList,
		SignFunc:         Md5,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

// 加签
func (s *Signature) SignURL(data string) (string, error) {

	u, err := NewUrl(data, SetSignKeyName(s.SignParamKeyName), SetSignParamsNameList(s.SignParamKeyList))
	if err != nil {
		return "", err
	}

	toSignStr := u.GetToSignStr()

	sign := s.SignFunc(toSignStr, s.Secret)

	return u.GetSignedUrl(sign), nil
}

// 验证签名
func (s *Signature) VerifyURL(data string) (bool, error) {
	u, err := NewUrl(data, SetSignKeyName(s.SignParamKeyName), SetSignParamsNameList(s.SignParamKeyList))
	if err != nil {
		return false, err
	}

	toSignStr := u.GetToSignStr()

	sign := s.SignFunc(toSignStr, s.Secret)

	oldSign := u.GetSign()

	fmt.Println("sign", sign, oldSign)
	return sign == oldSign, nil
}

func (s *Signature) Expired(data string, minutesUntilExpire int) (bool, error) {

	u, err := NewUrl(data)
	if err != nil {
		return true, err
	}

	ts := u.GetTimestamp()

	tsInt64, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return true, nil
	}

	start := time.Unix(tsInt64, 0)

	return time.Since(start) > time.Duration(minutesUntilExpire)*time.Minute, nil
}
