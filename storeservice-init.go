package storeservice

import (
	"github.com/agorago/storeapi"
	"github.com/agorago/storeservice/internal/register"
	"github.com/agorago/storeservice/internal/service"
	"github.com/agorago/wego"
	"github.com/agorago/wego/fw"
)

const (
	StoreService = "StoreService"
)

func MakeStoreServiceInitializer()fw.Initializer{
	return storeServiceInitializer{}
}

type storeServiceInitializer struct{}

func (storeServiceInitializer)ModuleName() string{
	return StoreService
}
// The storeservice initializer
func (storeServiceInitializer)Initialize(commandCatalog fw.CommandCatalog)(fw.CommandCatalog,error){
	rs,err := wego.GetWego(commandCatalog)
	if err != nil {
		return commandCatalog,err
	}
	repo := initializeRepo()
	register.RegisterStoreService(rs,service.MakeStoreService(repo),
		[]fw.Middleware {},
		storeapi.GetProxyMiddlewares())
	return commandCatalog,nil
}
