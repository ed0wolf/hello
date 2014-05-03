package assets

import (
	"net/http/httptest"
	"os"
	"testing"
)

const expectedFileBody string = "BODY"

var rootDir string
var writer *httptest.ResponseRecorder
var retriever *FileAssetsRetriever

func initRetrieverTest() {
	rootDir = os.TempDir()
	writer = httptest.NewRecorder()
	retriever = &FileAssetsRetriever{rootDir}
}

func TestWhenTheAssetDoesNotExist(t *testing.T) {
	initRetrieverTest()

	retriever.Retrieve("this_really_shouldnt_exist.js", writer)

	if writer.Code != 404 {
		t.Errorf("should have returned 404 Not Found but instead returned: %v", writer.Code)
	}
}

func TestWhenTheReqeustedAssetExists(t *testing.T) {
	initRetrieverTest()
	assetPath := "/should_exist_for_test.js"
	file, err := os.Create(rootDir + assetPath)
	if err != nil {
		t.Fatal(err)
	} else {
		file.WriteString(expectedFileBody)
	}

	retriever.Retrieve(assetPath, writer)

	os.Remove(file.Name())

	writer.Flush()
	body := writer.Body.String()
	if writer.Code != 200 {
		t.Errorf("should have returned 200 OK but instead returned: %v", writer.Code)
	}
	if body != expectedFileBody {
		t.Errorf("should have written %0v to responseWriter but wrote: %1v", expectedFileBody, body)
	}
}
