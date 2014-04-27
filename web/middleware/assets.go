package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

const keyword string = "assets"

type AssetsHandler struct {
	rootFolder string
}

func (hello *AssetsHandler) HandleAssets(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if isAssetRequest(r.RequestURI) {
			fmt.Fprint(w, r.RequestURI)
		} else {
			h.ServeHTTP(w, r)
		}
	}

	return http.HandlerFunc(fn)
}

func isAssetRequest(requestUri string) bool {
	return strings.HasPrefix(requestUri, "/"+keyword)
}
