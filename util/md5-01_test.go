package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAPMKey(t *testing.T) {
	key := GetAPMKey("a", "b", "c")
	assert.Equal(t, key, "a|b|c")
}
func TestEncryptionMD5(t *testing.T) {
	str01 := "abc123"
	md5Str01 := EncryptionMD5(str01)

	str02 := "abc123"
	md5Str02 := EncryptionMD5(str02)
	assert.Equal(t, md5Str01, md5Str02)

	str03 := "123abc"
	md5Str03 := EncryptionMD5(str03)
	assert.NotEqual(t, md5Str01, md5Str03)
}
