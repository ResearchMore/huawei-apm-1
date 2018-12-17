package apm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"strings"

	"github.com/go-chassis/huawei-apm/common"
	"github.com/go-chassis/huawei-apm/utils"
	"github.com/go-mesh/openlogging"
)

const (
	DefaultPort      string = "8923"
	DefaultPrefixUrl string = "https://100.125.5.235"
	//DefaultPrefixUrl string = "https://117.78.44.160"
)

// httpDo use https protocol sent message to collector
func httpDo(client *http.Client, message *common.TAgentMessage, url, projectID string) error {
	openlogging.GetLogger().Warn("send data to apm test")

	data, err := json.Marshal(message)
	if err != nil {
		openlogging.GetLogger().Errorf("use marshal to serialization TAgentMessage failed err : %s", err.Error())
		return err
	}

	var body io.Reader = bytes.NewReader(data)

	openlogging.GetLogger().Warn("send data to apm test")
	url = getURL(url, projectID)
	fmt.Println("===>", url)

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		openlogging.GetLogger().Errorf("new request for collector failed error : %v", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	resp, err := client.Do(req)

	fmt.Printf("resp is %+v \n err is %+v \n", resp, err)

	if err != nil {
		openlogging.GetLogger().Errorf("http call url[%s] collector  failed error : %v", url, err)
		return err
	}
	defer resp.Body.Close()

	// non 2xx code
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		openlogging.GetLogger().Errorf("http call url[%s]  collector  failed error : %v , code is %d", url, err, resp.StatusCode)
		return errors.New("call collector failed")
	}

	return nil
}

// getURL get access apm url
func getURL(u, projectID string) string {
	//g := config.GlobalDefinition
	//
	//c := g.Cse
	//s := c.Service
	//r := s.Registry
	//
	//urlPre := r.Address
	//if urlPre == "" {
	urlPre := DefaultPrefixUrl
	//}
	us := strings.Split(urlPre, ":")
	if len(us) > 1 {
		urlPre = strings.Join([]string{us[0], us[1]}, ":")
	} else {
		urlPre = DefaultPrefixUrl
	}
	u = utils.GetStringWithDefaultName(u, DefaultInventoryUrl)
	url := fmt.Sprintf("%s:%s%s", urlPre, DefaultPort, fmt.Sprintf(u, "d801598753ce4aa6a611bc2815a2eed2"))
	openlogging.GetLogger().Warnf("url is [%v]", url)
	return url
}
