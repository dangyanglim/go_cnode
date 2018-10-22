package models

import (
	"log"

	db "github.com/dangyanglim/go_cnode/database"
)

//"log"

type Renter struct {
	Id         int    `json:"id" form:"id"`
	Renter     string `json:"renter" form:"renter"`
	Name       string `json:"name" form:"name"`
	Propertyid string `json:"propertyid" form:"propertyid"`
	TimeBegin  string `json:"timeBegin" form:"timeBegin"`
	TimeEnd    string `json:"timeEnd" form:"timeEnd"`
	Roomid     string `json:"roomid" form:"roomid"`
	Status     string `json:"status" form:"status"`
}

func (p *Renter) AddRenter() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO renter(propertyid, renter,name,timeBegin,timeEnd,roomid,status) VALUES (?, ?,?,?,?,?,?)", p.Propertyid, p.Renter, p.Name, p.TimeBegin, p.TimeEnd, p.Roomid, p.Status)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}
func (p *Renter) SetRenterStatus() (ra int64, err error) {
	rs, err := db.SqlDB.Exec("update renter set status=? WHERE id=? and propertyid=? and roomid=?", p.Status, p.Id, p.Propertyid, p.Roomid)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err = rs.RowsAffected()
	return
}
func (p *Renter) GetRenterByPropertyidByRoomid() (*[]map[string]string, error) {
	rows, err := db.SqlDB.Query("SELECT * FROM renter WHERE propertyid=? and roomid=? and status=?", p.Propertyid, p.Roomid, p.Status)
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

func (p *Renter) GetRenterByMobile() (*[]map[string]string, error) {
	// err = db.SqlDB.QueryRow("SELECT * FROM renter WHERE renter=?  and status=? order by timeBegin desc LIMIT 1", p.Renter, p.Status).Scan(
	// 	&renter.Renter, &renter.Status, &renter.Name, &renter.Propertyid, &renter.TimeBegin, &renter.TimeEnd, &renter.Roomid, &renter.Id,
	// )
	rows, err := db.SqlDB.Query("SELECT * FROM renter WHERE renter=?  and status=? order by timeBegin desc LIMIT 1", p.Renter, p.Status)
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
