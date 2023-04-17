
## Sign

Sign 用于对url进行签名和验证。

这对于发送带有可验证链接的地址以及防篡改的电子邮件等事情非常有用。

### Installation

`go get github.com/kekek/sign@latest`

### Features

- 可自定签名参数key
- 可配置参与签名的参数返回
- 可自定签名函数

### Usage 

```golang
package main

import (
	"fmt"
	"github.com/kekek/sign"
)

const secret = "somelongsecuresecret"

func main() {
	// Create a variable of type Signature, and pass it a secret.
	sign := sign.NewSignature{Secret: secret}

	// Call the SignURL to get a signed version. Note that only the part after 
	// https://somesite.com or http://somesite.com is actually signed, but you 
	// must pass the full url. This way, we can use the package in development 
	// without worrying about the domain name of a particular site.
	signed, _ := sign.SignURL("https://example.com/test?id=1")
	fmt.Println("Signed url:", signed)

	// Output is something like:
	// https://example.com/test?id=1&ts=1681720708&sign=a0786345aee299a095554419061030ab
	// Verify that a signed URL is valid, and was  issued by this application. Here, 
	// valid is true if the URL has a valid signature, and false if it is not.
	valid, _ := sign.VerifyURL(signed)
	fmt.Println("Valid url:", valid)

	// You can also check for expiry. Here, the signed url expires after 30 minutes.
	expired, _ := sign.Expired(signed, 30)
	fmt.Println("Expired:", expired)
}

 ```

- 配置参与签名参数列表

``` golang  
var list = []string{"id", "name", "key"}
var s := NewSignature("111111", WithSignParamKeyList(list))

url := "https://example.com/test?id=1&age=1"

```

在上面的例子中，list中的参数，会参与签名，其它的url参数不会参与签名；id会参与签名，age不会


### 其它的库

- https://github.com/tsawler/signer

上面这个库，有坑，签名后，path中的反斜线会被删除；本人修改了一版，在本项目的分支：signer