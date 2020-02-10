package service

import (
	"bytes"
	"github.com/LeeWaiHo/guardian/pkg/model"
)

type NetworkNodes struct {
	NetworkID   int64
	Nodes       []*model.Node
	GatewayNode *model.Node
	ClientNodes []*model.Node
}

func (table NetworkNodes) GatewayConfiguration() string {
	gatewayNode := table.GatewayNode
	if gatewayNode == nil {
		return ""
	}
	configuration := &bytes.Buffer{}
	configuration.WriteString(gatewayNode.Head() + "\n")
	for _, clientNode := range table.ClientNodes {
		configuration.WriteString(clientNode.Peer() + "\n")
	}
	return configuration.String()
}

func (table NetworkNodes) ClientConfiguration(clientNode *model.Node) string {
	if clientNode == nil {
		return ""
	}
	gatewayNode := table.GatewayNode
	if gatewayNode == nil {
		return ""
	}
	var found bool
	for _, v := range table.ClientNodes {
		if v == clientNode {
			found = true
			break
		}
	}
	if !found {
		return ""
	}
	configuration := &bytes.Buffer{}
	configuration.WriteString(clientNode.Head() + "\n")
	configuration.WriteString(gatewayNode.Peer() + "\n")
	return configuration.String()
}

func (table NetworkNodes) Node(nodeID int64) *model.Node {
	for _, node := range table.Nodes {
		if node.ID == nodeID {
			return node
		}
	}
	return nil
}
