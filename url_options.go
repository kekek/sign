package sign

type Option func(*Url)

// SetSignKeyName 设置签名的名字
func SetSignKeyName(name string) Option {
	return func(o *Url) {
		o.signKeyName = name
	}
}

// SetTimestampName 设置签名的名字
func SetTimestampName(name string) Option {
	return func(o *Url) {
		o.timestampName = name
	}
}

// SetSignParamsNameList 待签名参数名称列表：在此列表之内的参数，参与签名
func SetSignParamsNameList(keyList []string) Option {
	return func(o *Url) {
		o.toSignParamsNameList = keyList
	}
}
