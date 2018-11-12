// util pkg provide all the project use publish func
package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
	"zipkin-test-project/common"
)

// EncryptionMD5 return md5ed string
func EncryptionMD5(str string) string {
	hash := md5.New()
	io.WriteString(hash, str)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

//GetAPMKey return apm will use key for report data
func GetAPMKey(keys ...string) string {
	return strings.Join(keys, common.APMKwySeparator)
}

// GetTracingID return tracing ID
func GetTracingID()string  {
	return ""
}