### usage

```golang 
const secret = "somelongsecuresecret"

func main() {
	// Create a variable of type Signature, and pass it a secret, <= 64 characters.
	sign := signer.Signature{Secret: secret}

	// Call the SignURL to get a signed version. Note that only the part after 
	// https://somesite.com or http://somesite.com is actually signed, but you 
	// must pass the full url. This way, we can use the package in development 
	// without worrying about the domain name of a particular site.
	signed, _ := sign.SignURL("https://example.com/test?id=1")
	fmt.Println("Signed url:", signed)
	
	// Output is something like:
	// https://example.com/test?id=1&hash=.3w4TgJ.pAJWBPAO5k1cimZJ-nrRKnlvosOY1Krrp3ALf1rOAds
	
	// Verify that a signed URL is valid, and was  issued by this application. Here, 
	// valid is true if the URL has a valid signature, and false if it is not.
	valid, _ := sign.VerifyURL(signed)
	fmt.Println("Valid url:", valid)

	// You can also check for expiry. Here, the signed url expires after 30 minutes.
	expired := sign.Expired(signed, 30)
	fmt.Println("Expired:", expired)
}


```