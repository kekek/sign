package sign

import (
	"fmt"
	"strconv"
	"time"
)

var defaultSignKeyList = []string{"id", "timestamp", "ts"}

type Signature struct {
	Secret       string                      // 签名密码
	SignParamKey string                      // 签名key
	SignKeyList  []string                    // 参与签名的参数名称
	SignFunc     func(string, string) string // 签名函数
}

var defaultSignature = &Signature{
	Secret:       "11111",
	SignFunc:     Md5,
	SignParamKey: "sign",
	SignKeyList:  defaultSignKeyList,
}

// 加签
func (s *Signature) SignURL(data string) (string, error) {

	u, err := NewUrl(data)
	if err != nil {
		return "", err
	}

	toSignStr := u.GetToSignStr(s.SignKeyList)

	sign := s.SignFunc(toSignStr, s.Secret)

	return u.GetSignedUrl(s.SignParamKey, sign, s.SignKeyList), nil
}

// 验证签名
func (s *Signature) VerifyURL(data string) (bool, error) {
	u, err := NewUrl(data)
	if err != nil {
		return false, err
	}

	toSignStr := u.GetToSignStr(s.SignKeyList)

	sign := s.SignFunc(toSignStr, s.Secret)

	oldSign := u.GetParams(s.SignParamKey)

	fmt.Println("sign", sign, oldSign)
	return sign == oldSign, nil
}

func (s *Signature) Expired(data string, minutesUntilExpire int) (bool, error) {

	u, err := NewUrl(data)
	if err != nil {
		return true, err
	}

	ts := u.GetParams("ts")

	tsInt64, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return true, nil
	}

	start := time.Unix(tsInt64, 0)

	return time.Since(start) > time.Duration(minutesUntilExpire)*time.Minute, nil
}
