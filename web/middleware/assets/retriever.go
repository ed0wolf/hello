package assets

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type AssetsRetriever interface {
	Retrieve(assetPath string, w http.ResponseWriter)
}

type FileAssetsRetriever struct {
	rootDir string
}

func (assetsRetriever *FileAssetsRetriever) Retrieve(assetPath string, w http.ResponseWriter) {
	content, err := ioutil.ReadFile(assetsRetriever.rootDir + "/" + assetPath)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	mime, mimeErr := getAssetMime(assetPath)
	if mimeErr == nil {
		w.Header().Set("Content-Type", mime)
	}

	w.Write(content)
}

func getAssetMime(assetPath string) (string, error) {
	if strings.HasSuffix(assetPath, ".js") {
		return "application/javascript", nil
	} else if strings.HasSuffix(assetPath, ".css") {
		return "text/css", nil
	}
	return "", errors.New("Couldn't find MIME for" + assetPath)
}
