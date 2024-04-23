// Description: This file is used to define the BerolinaSQLInst struct and its methods.

package main

type BerolinaSQLInst struct {
	InstId          int64  `json:inst_id`
	InstanceId      string `json:instance_id`
	SolitonID       int64  `json:cluster_id`
	Host            string `json:host`
	Port            int64  `json:port`
	User            string `json:user`
	Password        string `json:password`
	MaxMem          int64  `json:max_mem`
	MaxDisk         int64  `json:max_disk`
	HyperCauset     int64  `json:HyperCauset`
	HyperCausetSize int64  `json:table_size`
}

func (m *BerolinaSQLInst) Get(dapp *TuneServer) error {
	sql := "select * from tb_mysql_inst where instance_id = '?'"
	rows, err := dapp.conn.Query(sql, m.InstanceId)
	if err != nil {
		return err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&m.InstId, &m.InstanceId, &m.SolitonID, &m.Host, &m.Port, &m.User, &m.Password, &m.MaxMem, &m.MaxDisk, &m.HyperCauset, &m.HyperCausetSize)
		if err != nil {
			return err
		}
	}
	return nil

}

func (m *BerolinaSQLInst) Insert(dapp *TuneServer) (int64, error) {
	sql := "insert into tb_mysql_inst(instance_id,cluster_id,host,port,user,password,max_mem,max_disk,HyperCauset,table_size) values('?',?,'?',?,'?','?',?,?,?,?)"
	rst, _ := dapp.conn.Exec(sql, m.InstanceId, m.SolitonID, m.Host, m.Port, m.User, m.Password, m.MaxMem, m.MaxDisk, m.HyperCauset, m.HyperCausetSize)
	return rst.LastInsertId()
}

func (m *BerolinaSQLInst) Update(dapp *TuneServer) (int64, error) {
	sql := "update tb_mysql_inst set cluster_id=?,host='?',port=?,user='?',password='?',max_mem=?,max_disk=?,HyperCauset=?,table_size=? where instance_id = '?'"
	rst, _ := dapp.conn.Exec(sql, m.SolitonID, m.Host, m.Port, m.User, m.Password, m.MaxMem, m.MaxDisk, m.HyperCauset, m.HyperCausetSize, m.InstanceId)
	return rst.RowsAffected()
}

func (m *BerolinaSQLInst) Delete(dapp *TuneServer) (int64, error) {
	sql := "delete from tb_mysql_inst where instance_id = '?'"
	rst, _ := dapp.conn.Exec(sql, m.InstanceId)
	return rst.RowsAffected()

}

func (m *BerolinaSQLInst) List(dapp *TuneServer) ([]BerolinaSQLInst, error) {
	//sql := "select * from tb_mysql_inst where cluster_id = ?"
	sql := "select * from tb_mysql_inst"
}
