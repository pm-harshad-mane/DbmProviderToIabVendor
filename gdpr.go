package main

import(
	"strconv"
)

func getPurposeIdInteger(allowedPurposes []int) (int64) {
	origBitArray := []byte("000000000000000000000000")
	for _, purposeId := range allowedPurposes {
		purposeId = purposeId - 1
		origBitArray[ purposeId ] = origBitArray[ purposeId ] + 1
	}
	if i, err := strconv.ParseInt(string(origBitArray), 2, 64); err != nil {
        	return 0
 	} else {
    	return i
	}
}
