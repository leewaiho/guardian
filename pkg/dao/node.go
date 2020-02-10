package dao

import (
	"context"
	"github.com/LeeWaiHo/guardian/pkg/model"
)

func (dao *Dao) ListNetworkNodes(networkID int64) ([]*model.Node, error) {
	rows, err := dao.DB.QueryContext(
		context.Background(),
		"SELECT id, network_id, hostname, vpn_ip, type, listen_ip, listen_port, dns, private_key, public_key, allowed_ips, persistent_keepalive, created_time, updated_time, is_disabled from node WHERE network_id = ? AND is_disabled = 0",
		networkID,
	)
	if err != nil {
		return nil, err
	}
	var nodes []*model.Node
	for rows.Next() {
		node := &model.Node{}
		_ = rows.Scan(&node.ID, &node.NetworkID, &node.Hostname, &node.VpnIP, &node.Type, &node.ListenIP, &node.ListenPort, &node.DNS, &node.PrivateKey, &node.PublicKey, &node.AllowedIPs, &node.PersistentKeepalive, &node.CreatedTime, &node.UpdatedTime, &node.IsDisabled)
		nodes = append(nodes, node)
	}
	return nodes, nil
}
