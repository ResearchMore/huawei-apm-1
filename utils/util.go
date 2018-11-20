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

	"crypto/x509"

	"io/ioutil"

	"crypto/tls"

	"github.com/go-chassis/huawei-apm/common"
	"github.com/go-mesh/openlogging"
)

const defaultServerCrtFileName string = "ca.crt"

const defaultClientCrtFileName string = "kubecfg.crt"
const defaultClientKeyFileName string = "kubecfg_crypto.key"

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

// GetCertPool load server
func GetCertPool(path, fileName string) *x509.CertPool {

	fileName = getFileName(fileName, defaultServerCrtFileName)
	filePath := getFilePath(path, fileName)

	caCert, err := ioutil.ReadFile(filePath)
	if err != nil {
		openlogging.GetLogger().Errorf("read server.crt failed please check this it exist error : [%v]", err)
		return nil
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caCert)

	return certPool
}

// GetCertificate load client crt file and key file
func GetCertificate(path, crt, key string) tls.Certificate {

	crt = getFileName(crt, defaultClientCrtFileName)
	key = getFileName(key, defaultClientKeyFileName)

	crtFilePath := getFilePath(path, crt)
	keyFilePath := getFilePath(path, key)

	certificate, err := tls.LoadX509KeyPair(crtFilePath, keyFilePath)
	if err != nil {
		return tls.Certificate{}
	}
	return certificate
}

// getFilePath
func getFilePath(path, fileName string) string {
	path = strings.Replace(path, "\\", "/", -1)
	if path[len(path)-1] == 47 {
		return fmt.Sprintf("%s%s", path, fileName)
	}
	return fmt.Sprintf("%s/%s", path, fileName)
}

// getFileName
func getFileName(filename, defaultName string) string {
	if filename == "" {
		return defaultName
	}
	return filename
}
