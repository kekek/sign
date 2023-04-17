package sign

type SignatureOption func(*Signature)

type SignFunc func(string, string) string

func WithSignFunc(f SignFunc) SignatureOption {
	return func(o *Signature) {
		o.SignFunc = f
	}
}

func WithSignParamKeyName(s string) SignatureOption {
	return func(o *Signature) {
		o.SignParamKeyName = s
	}
}

func WithSignParamKeyList(paramNameList []string) SignatureOption {
	return func(o *Signature) {
		o.SignParamKeyList = paramNameList
	}
}
