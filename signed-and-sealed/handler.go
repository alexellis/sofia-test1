package function

import (
	"os"

	"github.com/alexellis/hmac"
)

// Handle a serverless request
func Handle(req []byte) string {
	res := hmac.Validate(req, os.Getenv("Http_X_Github_Event"), os.Getenv("hmac_secret"))

	if res != nil {
		return res.Error()
	}

	return "You got in - thanks for using the secret to sign your payload"
}
