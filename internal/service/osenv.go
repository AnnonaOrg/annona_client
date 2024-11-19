package service

import (
	"strconv"
	"strings"

	"github.com/AnnonaOrg/osenv"
)

func IsEnableBlockLongText() bool {
	return strings.EqualFold(osenv.Getenv("BLOCK_LONG_TEXT_ENABLE"))
}

func GetBlockLongTextMaxCount() int {
	item := osenv.Getenv("BLOCK_LONG_TEXT_MAX_COUNT")
	count, err := strconv.Atoi(item)
	if err != nil || count == 0 {
		return 100
	}
	return count
}
