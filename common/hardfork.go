package common

import "github.com/coschain/contentos-go/common/constants"

var hardForks = []uint64{
	constants.Original,
	constants.HardFork1,
}

func GetHardFork(blockNum uint64) uint64 {
	r := constants.Original
	for _, hf := range hardForks {
		if blockNum >= hf {
			r = hf
		}
	}
	return r
}
