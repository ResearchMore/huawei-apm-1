package apm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-chassis/go-chassis/core/config"
	"github.com/go-chassis/go-chassis/pkg/util/httputil"
	"github.com/go-chassis/huawei-apm/common"
	"github.com/go-chassis/huawei-apm/utils"
	"github.com/go-mesh/openlogging"
)

const (
	DefaultPort      string = "8923"
	DefaultPrefixUrl string = "https://100.125.5.235"
)

// httpDo use https protocol sent message to collector
func httpDo(client *http.Client, message *common.TAgentMessage, url, projectID string) error {
	// data to json
	data, err := json.Marshal(message)
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
	urlPre := config.GlobalDefinition.Cse.APM.Address
	if urlPre == "" {
		urlPre = DefaultPrefixUrl
	}
	us := strings.Split(urlPre, ":")
	if len(us) > 1 {
		urlPre = strings.Join([]string{us[0], us[1]}, ":")
	} else {
		urlPre = DefaultPrefixUrl
	}
	u = utils.GetStringWithDefaultName(u, DefaultInventoryUrl)
	// 此处的projectID需要修改
	url := fmt.Sprintf("%s:%s%s", urlPre, DefaultPort, fmt.Sprintf(u, config.GlobalDefinition.Cse.APM.Project))
	openlogging.GetLogger().Warnf("url is [%v]", url)
	return url
}
