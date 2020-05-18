// +build !mock

package test

import "fmt"

func (sts *storeTestStruct)iMustGetBackAStoreResponseWithKeyAndValue(key, value string) error {
	if sts.Sr.Key != key || sts.Sr.Value != value {
		return fmt.Errorf("Expected Key = " + key + " Got " + sts.Sr.Key + "Expected value = " +
			value + " Got " + sts.Sr.Value)
	}
	return nil
}