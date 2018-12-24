// util pkg provide all the project use publish func
package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"

	"github.com/go-chassis/huawei-apm/common"
	"github.com/go-mesh/openlogging"
)

const defaultCACrtFileName string = "ca.crt"

const defaultK8sCrtFileName string = "kubecfg.crt"

const defaultK8sKeyFileName string = "kubecfg_crypto.key"

// EncryptionMD5 return md5ed string
func EncryptionMD5(str string) string {
	hash := md5.New()
	io.WriteString(hash, str)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

//GetAPMKey return collector will use key for report data
func GetAPMKey(keys ...string) string {
	return strings.Join(keys, common.APMKeySeparator)
}

// UUID16 return 16 uuid
func UUID16() string {
	result := make([]byte, 16)

	rand.Seed(time.Now().UTC().UnixNano())
	tmp := rand.Int63()

	rand.Seed(tmp)

	for i := 0; i < 16; i++ {
		result[i] = byte(rand.Intn(16))
	}

	return fmt.Sprintf("%x-%x-%x-%x-%x", result[0:4], result[4:6], result[6:8], result[8:10], result[10:])
}

// GetTimeMillisecond get time millisecond num
func GetTimeMillisecond() int64 {
	return time.Now().UnixNano() / 1e6
}

// GetHostname get hostname
func GetHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		openlogging.GetLogger().Errorf("get hostname failed err:[%v]", err)
		return ""
	}
	return hostname
}

//GetLocalIP get host ip
func GetLocalIP() string {

	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addresses {
		// Parse IP
		var ip net.IP
		if ip, _, err = net.ParseCIDR(address.String()); err != nil {
			return ""
		}
		// Check if Isloopback and IPV4
		if ip != nil && !ip.IsLoopback() && ip.To4() != nil {
			return ip.String()
		}
	}
	return ""
}

// getFilePath
func getFilePath(path, fileName, defaultFileName string) string {
	fileName = GetStringWithDefaultName(fileName, defaultFileName)
	path = GetStringWithDefaultName(path, common.DefaultCAPath)
	path = strings.Replace(path, "\\", "/", -1)
	if path[len(path)-1] == 47 {
		return fmt.Sprintf("%s%s", path, fileName)
	} else if path[len(path)-4:] == ".crt" || path[len(path)-4:] == ".key" {
		return path
	}
	return fmt.Sprintf("%s/%s", path, fileName)
}

// GetStringWithDefaultName
func GetStringWithDefaultName(filename, defaultName string) string {
	if filename == "" {
		return defaultName
	}
	return filename
}
