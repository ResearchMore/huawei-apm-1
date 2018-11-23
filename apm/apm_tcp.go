package apm

import (
	"net"

	"fmt"

	"strings"

	"github.com/go-mesh/openlogging"
)

// NewConnection return tcp connect
func NewConnection(url, projectID string) (*net.TCPConn, error) {
	if projectID != "" && strings.Contains(url, "/%s") {
		url = fmt.Sprintf(url, projectID)
	}
	tcpAddr, err := net.ResolveTCPAddr("tcp", url)
	if err != nil {
		openlogging.GetLogger().Errorf("resolve collector  tcp addr failed , error : %v , Url is : %s", err, url)
		return nil, err
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		openlogging.GetLogger().Errorf("dial collector  tcp failed , error : %v , Url is : %s", err, url)
		return nil, err
	}
	return conn, nil
}
