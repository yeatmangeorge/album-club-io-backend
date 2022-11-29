package util

import (
	"testing"

	"github.com/yeatmangeorge/album-club-io-backend/domain/src/util"

	"github.com/stretchr/testify/assert"
)

func TestUintToString(t *testing.T){
	var value uint= 5
	result := util.UintToString(value)
	assert.Equal(t, "5", result)
}