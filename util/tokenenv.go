package util

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Jehanv60/helper"
)

func TokenEnv(r *http.Request) string {
	helper.GoDoEnv()
	var Token = os.Getenv("Token")
	tokenn, _ := r.Cookie(Token)
	tokenstring, err := Decodetoken(tokenn.Value)
	helper.PanicError(err)
	return fmt.Sprintf("%v", tokenstring["pengguna"])
}
