package function

import (
	"os"

	"github.com/alexellis/hmac"
)

// Handle a serverless request
func Handle(req []byte) string {
	xHubSignature := os.Getenv("Http_X_Hub_Signature")

	res := hmac.Validate(req, xHubSignature, os.Getenv("hmac_secret"))

	if res != nil {
		return res.Error()
	}

	return "You got in - thanks for using the secret to sign your payload"
}
