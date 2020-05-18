package service

import (
	"context"
	api "github.com/agorago/storeapi/api"
)

type storeservice struct{
	repo Repo
}

func MakeStoreService(repo Repo) storeservice {
	return storeservice{
		repo: repo,
	}
}

func (sd storeservice) Store(_ context.Context, key,value string)(api.StoreResponse,error) {
	_,err := sd.repo.Store(key, value)
	if err != nil{
		return api.StoreResponse{},nil
	}
	return api.StoreResponse{
		Key:   key,
		Value: value,
	},nil
}

func (sd storeservice) Retrieve(_ context.Context, key string)(api.StoreResponse,error) {
	value, err := sd.repo.Retrieve(key)
	if err != nil {
		return api.StoreResponse{}, nil
	}
	return api.StoreResponse{
		Key:   key,
		Value: value,
	}, nil
}