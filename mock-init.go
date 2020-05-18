// +build mock

package storeservice

import "github.com/agorago/storeservice/internal/service"

func initializeRepo() service.Repo{
	return service.MakeHardcodedRepo()
}
