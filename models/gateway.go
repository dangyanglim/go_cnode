package models

import (
	"log"

	db "github.com/dangyanglim/go_cnode/database"
)

//"log"

type Gateway struct {
	Id         int    `json:"id" form:"id"`
	Spaceid    string `json:"spaceid" form:"spaceid"`
	Spacetype  string `json:"spacetype" form:"spacetype"`
	Propertyid string `json:"propertyid" form:"propertyid"`
	GateWay    string `json:"gateway" form:"gateway"`
	Name       string `json:"name" form:"name"`
}

func (p *Gateway) AddGateway() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO gateway(spaceid,spacetype,propertyid,gateway,name) VALUES (?,?,?,?,?)", p.Spaceid, p.Spacetype, p.Propertyid, p.GateWay, p.Name)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}
func (p *Gateway) GetGateways() (*[]map[string]string, error) {
	rows, err := db.SqlDB.Query("SELECT id,gateway,name FROM gateway WHERE spaceid=? and spacetype=? and propertyid=?", p.Spaceid, p.Spacetype, p.Propertyid)
	if err != nil {
		log.Print(err.Error())
	}
	defer rows.Close()
	columns, _ := rows.Columns()
	length := len(columns)

	values := make([]string, length)
	columnPointers := make([]interface{}, length)
	for i := 0; i < length; i++ {
		columnPointers[i] = &values[i]
	}
	ret := make([]map[string]string, 0)
	for rows.Next() {

		err := rows.Scan(columnPointers...)
		//rows.Scan(roomtemp...)
		if err != nil {
			log.Print(err.Error())
		}
		data := make(map[string]string)
		for i := 0; i < length; i++ {
			columnName := columns[i]
			columnValue := columnPointers[i].(*string)
			data[columnName] = *columnValue
		}
		ret = append(ret, data)
	}

	err = rows.Err()
	if err != nil {
		log.Print(err.Error())
	}

	return &ret, err
}
func (p *Gateway) DelGateway() (ra int64, err error) {
	rs, err := db.SqlDB.Exec("DELETE FROM gateway WHERE  id=?", p.Id)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err = rs.RowsAffected()
	return
}
func (p *Gateway) GetGateway() (gateway Gateway, err error) {
	err = db.SqlDB.QueryRow("SELECT id,gateway,propertyid from gateway WHERE gateway=? and propertyid=?", p.GateWay, p.Propertyid).Scan(
		&gateway.Id, &gateway.GateWay, &gateway.Propertyid,
	)
	return
}
func (p *Gateway) GetGatewayCount() (int, error) {
	var count int
	row := db.SqlDB.QueryRow("SELECT count(*) FROM gateway WHERE spaceid=? and spacetype=? and propertyid=?", p.Spaceid, p.Spacetype, p.Propertyid)
	err := row.Scan(&count)
	return count, err
}
