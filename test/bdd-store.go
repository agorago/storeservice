package test

import (
	"context"
	"github.com/agorago/storeapi"
	"github.com/agorago/wego/fw"
	"log"

	"github.com/DATA-DOG/godog"
	"github.com/agorago/storeapi/api"
)

type storeTestStruct struct {
	Sr  api.StoreResponse
	Err error
}

var storeServiceProxy api.StoreService
// FeatureContext - the GODOG feature context that defines the step definitions for the BDD tests
// that need to be executed for hello service
func FeatureContext(commandCatalog fw.CommandCatalog,s *godog.Suite) {
	var err error
	var sts = &storeTestStruct{}
	storeServiceProxy,err = storeapi.GetStoreProxy(commandCatalog)
	if err != nil {
		log.Fatalf("Test application has an error. Error = %v\n",err)
	}
	s.Step(`^that I invoke Store with key "([^"]*)" and value "([^"]*)"$`, sts.iInvokeStore)
	s.Step(`^I invoke Retrieve with key "([^"]*)"$`, sts.iInvokeRetrieveWithKey)
	s.Step(`^I must get back a StoreResponse with key "([^"]*)" and value "([^"]*)"$`, sts.iMustGetBackAStoreResponseWithKeyAndValue)

}

func (sts *storeTestStruct) iInvokeStore(key string, value string) error {
	ctx := context.TODO()

	sr, err := storeServiceProxy.Store(ctx, key,value)
	if err != nil {
		return err
	}
	sts.Sr = sr
	return nil
}

func (sts *storeTestStruct)iInvokeRetrieveWithKey(key string) error {
	ctx := context.TODO()

	sr, err := storeServiceProxy.Retrieve(ctx, key)
	if err != nil {
		return err
	}
	sts.Sr = sr
	return nil
}


