package decorator

import (
	"log"
	"net/http"
	"testing"
)

func TestRouter(t *testing.T) {
	http.HandleFunc("/v1/hello", WithServerHeader(WithAuthCookie(hello)))
	http.HandleFunc("/v2/hello", WithServerHeader(WithBasicAuth(hello)))
	http.HandleFunc("/v3/hello", WithServerHeader(WithBasicAuth(WithDebugLog(hello))))
	var opts []HttpHandlerDecorator
	opts = append(opts, WithServerHeader)
	opts = append(opts, WithBasicAuth)
	opts = append(opts, WithDebugLog)
	http.HandleFunc("/v4/hello", Handler(hello, opts...))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
