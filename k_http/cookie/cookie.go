package cookie

import (
	"bufio"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport"
	"net/http"
	"strings"
)

// GetCookieToken  获取cookie中key为ut的值
func GetCookieToken(header transport.Transporter, key string) string {
	rawCookies := header.RequestHeader().Get("Cookie")
	rawRequest := fmt.Sprintf("GET / HTTP/1.0\r\nCookie: %s\r\n\r\n", rawCookies)
	req, err := http.ReadRequest(bufio.NewReader(strings.NewReader(rawRequest)))
	var token string
	if err == nil {
		cookies := req.Cookies()
		for _, cookie := range cookies {
			if cookie.Name == key {
				token = cookie.Value
			}
		}
	}
	return token
}
