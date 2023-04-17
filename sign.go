package signer

// docs : https://github.com/tsawler/signer

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	goalone "github.com/bwmarrin/go-alone"
)

const defaultSecret = "6nKcIGq19kxXJYOR"

var SignInstance *Signature

func InitSign(secret string) {

	s := defaultSecret
	if len(secret) > 0 {
		s = secret
	}

	SignInstance = &Signature{
		Secret: s,
	}
}

// Signature is the type for the package. Secret is the signer secret, a lengthy
// and hard to guess string we use to sign things. The secret must not exceed 64 characters.
type Signature struct {
	Secret string
}

// SignURL generates a signed url and returns it, stripping off http:// and https://
func (s *Signature) SignURL(data string) (string, error) {
	// verify this is a url
	_, err := url.ParseRequestURI(data)
	if err != nil {
		return "", err
	}

	var urlToSign string

	ex := strings.Split(data, "//")
	exploded := strings.Split(ex[1], "/")
	domain := exploded[0]
	exploded[0] = ""
	stringToSign := strings.Join(exploded, "/")

	pen := goalone.New([]byte(s.Secret), goalone.Timestamp)

	if strings.Contains(stringToSign, "?") {
		// handle case where URL contains query parameters
		urlToSign = fmt.Sprintf("%s&hash=", stringToSign)
	} else {
		// no query parameters
		urlToSign = fmt.Sprintf("%s?hash=", stringToSign)
	}

	tokenBytes := pen.Sign([]byte(urlToSign))
	token := string(tokenBytes)

	return fmt.Sprintf("%s//%s%s", ex[0], domain, token), nil
}

// VerifyURL verifies a signed url and returns true if it is valid,
// false if it is not. Note that http:// and https:// are stripped off
// before verification
func (s *Signature) VerifyURL(data string) (bool, error) {
	_, err := url.ParseRequestURI(data)
	if err != nil {
		return false, err
	}
	ex := strings.Split(data, "//")
	exploded := strings.Split(ex[1], "/")
	exploded[0] = ""
	stringToVerify := strings.Join(exploded, "/")

	pen := goalone.New([]byte(s.Secret), goalone.Timestamp)

	_, err = pen.Unsign([]byte(stringToVerify))
	if err != nil {
		// signature is not valid. Token was tampered with, forged, or maybe it's
		// not even a token at all! Either way, it's not safe to use it.
		return false, err
	}

	// valid hash
	return true, nil

}

// Expired checks to see if a token has expired. It returns true if
// the token was created within minutesUntilExpire, and false otherwise.
func (s *Signature) Expired(data string, minutesUntilExpire int) bool {
	exploded := strings.Split(data, "//")

	pen := goalone.New([]byte(s.Secret), goalone.Timestamp)
	ts := pen.Parse([]byte(exploded[1]))

	return time.Since(ts.Timestamp) > time.Duration(minutesUntilExpire)*time.Minute
}
