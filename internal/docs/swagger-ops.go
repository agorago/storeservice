
package docs

import (
"github.com/agorago/storeapi/api"
)

// swagger:route GET /store Store-tag StoreRequestWrapper
//  Store - Stores the key, value pair
// responses:
//   200: StoreResponseWrapper

//  StoreResponse - contains the stored name,value pair
// swagger:response StoreResponseWrapper
type StoreResponseWrapper struct {
// in:body
Body api.StoreResponse
}

// swagger:parameters StoreRequestWrapper
type StoreRequestWrapper struct{
// 
// in:body

}

// swagger:route GET /retrieve Retrieve-tag RetrieveRequestWrapper
//  Retrieve - retrieves the stored key and value
// responses:
//   200: RetrieveResponseWrapper

//  StoreResponse - contains the stored name,value pair
// swagger:response RetrieveResponseWrapper
type RetrieveResponseWrapper struct {
// in:body
Body api.StoreResponse
}

// swagger:parameters RetrieveRequestWrapper
type RetrieveRequestWrapper struct{
// 

}


