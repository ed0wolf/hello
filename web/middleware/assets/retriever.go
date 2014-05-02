package assets

type AssetsRetriever interface {
	Retrieve(assetPath string)
}

type FileAssetsRetriever struct {
}

func (assetsRetriever *FileAssetsRetriever) Retrieve(assetPath string) {

}
