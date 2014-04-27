package middleware

import (
	"testing"
)

func WhenRequestUriIsN(t *testing.T) {
	var ah = AssetsHandler{}
	ah.HandleAssets()
	t.Error()
}
