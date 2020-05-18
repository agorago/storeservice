package service

// This illustrates the mocking concept in WeGO
// This is a Repo with two implementations.
// For mocking HardcodedRepo is always returned
// HardcodedRepo always returns the same thing
// MapRepo stores the name, value pair in a map

// Implements a simple string repo
type Repo interface{
	Store(key,value string)(string,error)
	Retrieve(key string)(string,error)
}


// MapRepo is the actual repo and HardcodedRepo is used for mocking
type MapRepo struct {
	pairs map[string]string
}

func MakeMapRepo() Repo{
	return MapRepo{
		pairs: make(map[string]string),
	}
}

func (mr MapRepo)Store(key,value string)(string,error){
	mr.pairs[key]  = value
	return value,nil
}

func (mr MapRepo)Retrieve(key string)(string,error){
	return mr.pairs[key],nil
}

// A different implementation using a hard coded repo
type HardcodedRepo struct{}
func MakeHardcodedRepo()Repo{
	return HardcodedRepo{}
}
func (mr HardcodedRepo)Store(key,value string)(string,error){
	return value,nil // does nothing
}
func (mr HardcodedRepo)Retrieve(key string)(string,error){
	return "SAMETHING",nil
}