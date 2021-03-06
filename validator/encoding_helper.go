package validator

import "github.com/herdius/herdius-core/libs/common"

// cdcEncode returns nil if the input is nil, otherwise returns
// cdc.MustMarshalBinaryBare(item)
func cdcEncode(item interface{}) []byte {
	if item != nil && !common.IsTypedNil(item) && !common.IsEmpty(item) {
		return cdc.MustMarshalBinaryBare(item)
	}
	return nil
}
