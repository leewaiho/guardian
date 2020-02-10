package service

import (
	"github.com/LeeWaiHo/guardian/pkg/dao"
	"github.com/LeeWaiHo/guardian/pkg/model"
)

type Service struct {
	dao *dao.Dao
}

func New() *Service {
	return &Service{dao: dao.New()}
}

func (s *Service) ScanNetworkNodes(networkID int64) (*NetworkNodes, error) {
	networkNodes := &NetworkNodes{
		NetworkID: networkID,
	}
	nodes, err := s.dao.ListNetworkNodes(networkID)
	if err != nil {
		return nil, err
	}
	networkNodes.Nodes = nodes
	for _, v := range nodes {
		switch v.Type {
		case model.NodeTypeGateway:
			networkNodes.GatewayNode = v
		case model.NodeTypeClient:
			networkNodes.ClientNodes = append(networkNodes.ClientNodes, v)
		}
	}
	return networkNodes, nil
}
