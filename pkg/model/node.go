package model

import (
	"bytes"
	"fmt"
	"time"
)

const (
	_DefaultPersistentKeepalive = 25
)

const (
	NodeTypeGateway = 1 + iota
	NodeTypeClient
)

type Node struct {
	ID                  int64
	NetworkID           int64
	Hostname            string
	VpnIP               string
	Type                int8
	ListenIP            string
	ListenPort          string
	DNS                 string
	PrivateKey          string
	PublicKey           string
	AllowedIPs          string
	PersistentKeepalive int32
	CreatedTime         time.Time
	UpdatedTime         time.Time
	IsDisabled          int8
}

func (node Node) Head() string {
	buff := &bytes.Buffer{}
	buff.WriteString("[Interface]\n")
	buff.WriteString(fmt.Sprintf("PrivateKey = %s\n", node.PrivateKey))
	buff.WriteString(fmt.Sprintf("Address = %s\n", node.VpnIP))
	// dns
	if len(node.DNS) > 0 {
		buff.WriteString(fmt.Sprintf("DNS = %s\n", node.DNS))
	}
	// listen port
	if len(node.ListenPort) > 0 {
		buff.WriteString(fmt.Sprintf("ListenPort = %s\n", node.ListenPort))
	}
	return buff.String()
}

func (node Node) Peer() string {
	buff := &bytes.Buffer{}
	if node.Type != NodeTypeGateway {
		buff.WriteString(fmt.Sprintf("# %s\n", node.Hostname))
	}
	buff.WriteString("[Peer]\n")
	buff.WriteString(fmt.Sprintf("PublicKey = %s\n", node.PublicKey))
	// endpoint
	if len(node.ListenIP) > 0 && len(node.ListenPort) > 0 {
		endpoint := fmt.Sprintf("%s:%s", node.ListenIP, node.ListenPort)
		buff.WriteString(fmt.Sprintf("Endpoint = %s\n", endpoint))
	}
	// allowedIPs
	allowedIPs := node.AllowedIPs
	if len(allowedIPs) == 0 {
		allowedIPs = node.VpnIP
	}
	buff.WriteString(fmt.Sprintf("AllowedIPs = %s\n", allowedIPs))
	// persistentKeepalive
	persistentKeepalive := node.PersistentKeepalive
	if persistentKeepalive == 0 {
		persistentKeepalive = _DefaultPersistentKeepalive
	}
	buff.WriteString(fmt.Sprintf("PersistentKeepalive = %d\n", persistentKeepalive))
	return buff.String()
}
