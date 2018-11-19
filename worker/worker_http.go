package worker

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"crypto/x509"

	"github.com/go-chassis/huawei-apm/common"
	"github.com/go-mesh/openlogging"
)

// httpDo use https protocol sent message to collector
func httpDo(client *http.Client, message *common.TAgentMessage, url, projectID string) error {
	data, err := json.Marshal(message)
	if err != nil {
		openlogging.GetLogger().Errorf("use marshal to serialization TAgentMessage failed err : %s", err.Error())
		return err
	}

	var body io.Reader = bytes.NewReader(data)

	url = fmt.Sprintf(url, projectID)
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		openlogging.GetLogger().Errorf("new request for collector failed error : %v", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	resp, err := client.Do(req)
	if err != nil {
		openlogging.GetLogger().Errorf("http call  collector  failed error : %v", err)
		return err
	}
	defer resp.Body.Close()

	// non 2xx code
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		openlogging.GetLogger().Errorf("http call  collector  failed error : %v , code is %d", err, resp.StatusCode)
		return errors.New("call collector failed")
	}

	return nil
}

func getCAs(path string) *x509.CertPool {
	return nil
}
