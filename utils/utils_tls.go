package utils

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/go-chassis/go-chassis/security"
	"github.com/go-chassis/huawei-apm/common"
	"github.com/go-mesh/openlogging"
)

// GetTLSConfig  get https access  certificate config
func GetTLSConfig(path, kubecrt, kubekey string) (*tls.Config, error) {
	kubecrt = getFilePath(path, kubecrt, common.DefaultK8sCrtFileName)
	kubekey = getFilePath(path, kubekey, common.DefaultK8sKeyFileName)

	certificates, err := getCertificate(kubecrt, kubekey)
	if err != nil {
		return nil, err
	}
	return &tls.Config{
		Certificates: certificates,
	}, nil
}

// getCertificate load client crt file and key file
func getCertificate(kubeCrt, kubeKey string) ([]tls.Certificate, error) {
	// read crt and key file
	crtContent, err := ioutil.ReadFile(kubeCrt)
	if err != nil {
		return nil, fmt.Errorf("read kubeCrt file [%s] failed", kubeCrt)
	}
	keyContent, err := ioutil.ReadFile(kubeKey)
	if err != nil {
		return nil, fmt.Errorf("read kubeKey file [%s] failed", kubeKey)
	}
	// decrypt k8s key data
	decryptData, err := decryptKey(keyContent)

	cer, err := tls.X509KeyPair(crtContent, decryptData)
	//cer, err := tls.X509KeyPair(crtContent, keyContent)
	if err != nil {
		return nil, err
	}

	return []tls.Certificate{cer}, nil
}

// decryptKey decrypt kubecfg_crypto.key file
func decryptKey(ciphertext []byte) ([]byte, error) {
	// use chassis aes plugin
	cipher, err := security.GetCipherNewFunc("aes")

	if err != nil {
		return nil, err
	}
	aes := cipher()
	if aes == nil {
		err := errors.New("use plugin func to get aes failed")
		openlogging.GetLogger().Error(err.Error())
		return nil, err
	}
	s, err := aes.Decrypt(string(ciphertext))
	return []byte(s), err
}
