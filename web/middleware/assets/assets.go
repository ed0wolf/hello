package assets

import (
	"net/http"
	"strings"
)

const prefix string = "/assets"

type AssetsHandler struct {
	retriever AssetsRetriever
}

func (handler *AssetsHandler) HandleAssets(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if isAssetRequest(r.RequestURI) {
			handler.retriever.Retrieve(getAssetPath(r.RequestURI))
		} else {
			h.ServeHTTP(w, r)
		}
	}

	return http.HandlerFunc(fn)
}

func isAssetRequest(requestUri string) bool {
	return strings.HasPrefix(requestUri, prefix)
}

func getAssetPath(requestUri string) string {
	return strings.TrimPrefix(requestUri, prefix)
}
