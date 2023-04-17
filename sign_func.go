package sign

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"hash"
	"io"
)

func SignStrSha256(toSignUrl, secret string) string {
	h := hmac.New(func() hash.Hash { return sha256.New() }, []byte(secret))
	io.WriteString(h, toSignUrl)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signedStr
}
func SignStrSha1(toSignUrl, secret string) string {
	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(secret))
	io.WriteString(h, toSignUrl)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signedStr
}

func SignStrMd5(toSignUrl, secret string) string {
	h := hmac.New(func() hash.Hash { return md5.New() }, []byte(secret))
	io.WriteString(h, toSignUrl)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signedStr
}

func Md5(toSignUrl, secret string) string {
	h := hmac.New(func() hash.Hash { return md5.New() }, []byte(secret))
	io.WriteString(h, toSignUrl)
	return fmt.Sprintf("%x", h.Sum(nil))
}
