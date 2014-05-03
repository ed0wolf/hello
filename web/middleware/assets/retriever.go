package assets

import (
	"io/ioutil"
	"net/http"
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

	w.Write(content)
}
