package models

import (
	"log"

	db "github.com/dangyanglim/go_cnode/database"
)

//"log"

type Ammeter struct {
	Id           int    `json:"id" form:"id"`
	Ammeter_addr string `json:"ammeter_addr" form:"ammeter_addr"`
	Voltage      string `json:"voltage" form:"voltage"`
	Current      string `json:"current" form:"current"`
	Energy       string `json:"energy" form:"energy"`
	Spaceid      string `json:"spaceid" form:"spaceid"`
	Spacetype    string `json:"spacetype" form:"spacetype"`
	Propertyid   string `json:"propertyid" form:"propertyid"`
	GateWay      string `json:"gateway" form:"gateway"`
}

func (p *Ammeter) AddAmmeter() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO ammeter(voltage,current,energy,ammeter_addr,spaceid,spacetype,propertyid,gateway) VALUES (?,?,?,?,?,?,?,?)", p.Voltage, p.Current, p.Energy, p.Ammeter_addr, p.Spaceid, p.Spacetype, p.Propertyid, p.GateWay)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

func (p *Ammeter) UpdateAmmeter() (ra int64, err error) {

	stmt, err := db.SqlDB.Prepare("UPDATE ammeter SET voltage=? , current=? ,energy=? WHERE Ammeter_addr=?")
	defer stmt.Close()
	if err != nil {
		return
	}
	rs, err := stmt.Exec(p.Voltage, p.Current, p.Energy, p.Ammeter_addr)
	if err != nil {
		return
	}
	ra, err = rs.RowsAffected()
	return
}
func (p *Ammeter) GetAmmeterStates() (ammeter Ammeter, err error) {
	err = db.SqlDB.QueryRow("SELECT voltage,current,energy FROM ammeter WHERE spaceid=? and spacetype=?", p.Spaceid, p.Spacetype).Scan(
		&ammeter.Voltage, &ammeter.Current, &ammeter.Energy,
	)
	return
}
func (p *Ammeter) GetAmmeter() (ammeter Ammeter, err error) {
	err = db.SqlDB.QueryRow("SELECT id,ammeter_addr, gateway FROM ammeter WHERE spaceid=? and spacetype=?", p.Spaceid, p.Spacetype).Scan(
		&ammeter.Id, &ammeter.GateWay, &ammeter.Ammeter_addr,
	)
	return
}
func (p *Ammeter) GetAmmeters() (*[]map[string]string, error) {
	// err = db.SqlDB.Query("SELECT * FROM room WHERE propertyid=?", p.Propertyid).Scan(
	// 	&roomtemp.Id, &roomtemp.Propertyid, &roomtemp.Room, &roomtemp.Floorid,
	// )
	// return
	rows, err := db.SqlDB.Query("SELECT id,ammeter_addr, gateway FROM ammeter WHERE spaceid=? and spacetype=? and propertyid=?", p.Spaceid, p.Spacetype, p.Propertyid)
	if err != nil {
		log.Print(err.Error())
	}
	defer rows.Close()
	//建立一个列数组
	// cols, err := rows.Columns()
	// //var roomtemp2 = make([]interface{}, len(cols))
	// for i := 0; i < len(cols); i++ {
	// 	roomtemp[i] = Room{}

	// }
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
func (p *Ammeter) DelAmmeter() (ra int64, err error) {
	rs, err := db.SqlDB.Exec("DELETE FROM ammeter WHERE  id=?", p.Id)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err = rs.RowsAffected()
	return
}
func (p *Ammeter) GetAmmeterCount() (int, error) {
	var count int
	row := db.SqlDB.QueryRow("SELECT count(*) FROM ammeter WHERE spaceid=? and spacetype=? and propertyid=?", p.Spaceid, p.Spacetype, p.Propertyid)
	err := row.Scan(&count)
	return count, err
}
