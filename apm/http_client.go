package apm

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chassis/go-chassis/pkg/util/httputil"
	"github.com/go-chassis/huawei-apm/common"
	"github.com/go-chassis/huawei-apm/utils"
	"github.com/go-mesh/openlogging"
)

const (
	DefaultPort string = "8923"
)

// httpDo use https protocol sent message to collector
func httpDo(client *http.Client, message *common.TAgentMessage, url, projectID string) error {
	// data to json
	data, err := utils.Serialize(message)
	if err != nil {
		openlogging.GetLogger().Errorf("use marshal to serialization TAgentMessage failed err : %s", err.Error())
		return err
	}

	// data into body
	var body io.Reader = bytes.NewReader(data)
	url = getURL(url, projectID)

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		openlogging.GetLogger().Errorf("new request of apm collector failed error : %v", err)
		return err
	}

	// set header
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	resp, err := client.Do(req)

	if err != nil {
		openlogging.GetLogger().Errorf("http call url[%s] collector  failed error : %v", url, err)
		return err
	}
	defer resp.Body.Close()
	bodyData := httputil.ReadBody(resp)
	fmt.Printf("call url is :%+v\n", url)
	fmt.Printf("respense data is :%+v\n", string(bodyData))
	// non 2xx code
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		openlogging.GetLogger().Errorf("http call url[%s]  collector  failed error : %v , code is %d", url, err, resp.StatusCode)
		return errors.New("call collector failed")
	}

	return nil
}

// getURL get access apm url
func getURL(u, projectID string) string {
	// get elb ip
	elbIP := utils.GetElbIP()

	u = utils.GetStringWithDefaultName(u, DefaultInventoryUrl)
	url := fmt.Sprintf("%s:%s%s", elbIP, DefaultPort, fmt.Sprintf(u, projectID))
	return url
}
