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

	"github.com/go-chassis/go-chassis/security"
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

// GetX509CACertPool load server
func GetX509CACertPool(path, crt string) *x509.CertPool {

	//crt = GetStringWithDefaultName(crt, defaultCACrtFileName)

	filePath := getFilePath(path, crt, defaultK8sCrtFileName)

	caCert, err := ioutil.ReadFile(filePath)
	if err != nil {
		openlogging.GetLogger().Errorf("read server.crt failed please check this it exist error : [%v]", err)
		return nil
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caCert)

	return certPool
}

// getX509CACertPool use X509 to get CA Cert pool
func getX509CACertPool(path, crt string) (*x509.CertPool, error) {

	//crt = GetStringWithDefaultName(crt, defaultCACrtFileName)

	filePath := getFilePath(path, crt, defaultK8sCrtFileName)

	caCert, err := ioutil.ReadFile(filePath)
	if err != nil {
		openlogging.GetLogger().Errorf("read server.crt failed please check this it exist error : [%v]", err)
		return nil, err
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caCert)

	return certPool, nil
}

// GetCertificate load client crt file and key file
func GetCertificate(path, crt, key string) tls.Certificate {
	getCertificate(path, crt, "", key)
	crtFilePath := getFilePath(path, crt, defaultK8sCrtFileName)
	keyFilePath := getFilePath(path, key, defaultK8sKeyFileName)

	certificate, err := tls.LoadX509KeyPair(crtFilePath, keyFilePath)
	if err != nil {
		openlogging.GetLogger().Errorf("get client failed err is : %v\n", err)
		return tls.Certificate{}
	}

	return certificate
}

// getCertificate load client crt file and key file
func getCertificate(path, crt, ca, key string) ([]tls.Certificate, error) {
	// get ca ,kub ctr file ,k8s key file
	//	caFilePath := getFilePath(path, ca, defaultCACrtFileName)
	crtFilePath := getFilePath(path, crt, defaultK8sCrtFileName)
	keyFilePath := getFilePath(path, key, defaultK8sKeyFileName)
	// 证书读取
	crtContent, err := ioutil.ReadFile(crtFilePath)
	if err != nil {
		return nil, fmt.Errorf("read cert file %s failed", crtFilePath)
	}
	keyContent, err := ioutil.ReadFile(keyFilePath)
	if err != nil {
		return nil, fmt.Errorf("read key file %s failed", keyFilePath)
	}
	//caCrtContent, err := ioutil.ReadFile(caFilePath)
	//if err != nil {
	//	return nil, fmt.Errorf("read key file %s failed", keyFilePath)
	//}
	// test
	reply, err := DecryptKey(keyContent, crtContent)
	fmt.Println("reply  ==>", reply)
	fmt.Println("err ===>", err)
	fmt.Println("")
	return nil, nil
	//test
	fmt.Println(crtContent, keyContent)
	return []tls.Certificate{}, nil
}

// 解密
func DecryptKey(ciphertext, key []byte) ([]byte, error) {
	cipher, err := security.GetCipherNewFunc("aes")
	if err != nil {
		return nil, err
	}
	aes := cipher()
	s, err := aes.Decrypt(string(ciphertext))
	fmt.Println("====>", s, "\n====>", err)

	return []byte(s), err
}

func GetTLSConfig(path, CAFile string) (*tls.Config, error) {
	pool, err := getX509CACertPool(path, CAFile)
	if err != nil {
		return nil, err
	}
	return &tls.Config{
		ClientCAs: pool,
	}, nil
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
