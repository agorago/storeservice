package test

import (
	"github.com/agorago/storeservice"
	_ "github.com/agorago/wego" // to initialize the http module
	wegotest "github.com/agorago/wego/test"
	"testing"
)

func TestMain(m *testing.M) {
	wegotest.BDD(m, FeatureContext, storeservice.Initializers...)
}
