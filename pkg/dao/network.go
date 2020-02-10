package dao

import (
	"database/sql"
	"github.com/LeeWaiHo/guardian/pkg/model"
	"log"
)

func (dao *Dao) ListNetwork() (data []*model.Network, err error) {
	var rows *sql.Rows
	rows, err = dao.DB.Query("SELECT id, name, created_time, updated_time from network")
	if err != nil {
		log.Printf("查询VPN网络数据异常 error(%v)", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		network := &model.Network{}
		_ = rows.Scan(&network.ID, &network.Name, &network.CreatedTime, &network.UpdatedTime)
		data = append(data, network)
	}
	return
}
