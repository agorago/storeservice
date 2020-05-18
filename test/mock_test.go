// +build mock

package test
import "fmt"

// Since mock uses a hard coded response ignore the value
func (sts *storeTestStruct)iMustGetBackAStoreResponseWithKeyAndValue(key, value string) error {
	value = "SAMETHING"
	if sts.Sr.Key != key || sts.Sr.Value != value {
		return fmt.Errorf("Expected Key = " + key + " Got " + sts.Sr.Key + "Expected value = " +
			value + " Got " + sts.Sr.Value)
	}
	return nil
}