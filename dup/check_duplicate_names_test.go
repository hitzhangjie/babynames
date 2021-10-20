package dup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckDuplicateNames(t *testing.T) {
	_, err := CheckDuplicateNames("张杰")
	assert.Nil(t, err)
}
