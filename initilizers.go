package storeservice

import (
	"github.com/agorago/storeapi"
	"github.com/agorago/wego"
	"github.com/agorago/wego/fw"
)

var Initializers = []fw.Initializer{
	wego.MakeWegoInitializer(),
	storeapi.MakeStoreApiInitializer(),
	MakeStoreServiceInitializer(),
}
