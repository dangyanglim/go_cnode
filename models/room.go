package models

import (
	"errors"
	"log"
	"strconv"

	db "github.com/dangyanglim/go_cnode/database"
)

//"log"

type Room struct {
	Id         int    `json:"id" form:"id"`
	Floorid    string `json:"floorid" form:"floorid"`
	Room       string `json:"room" form:"room"`
	Propertyid string `json:"propertyid" form:"propertyid"`
}
type Floor struct {
	Id         int    `json:"id" form:"id"`
	Propertyid string `json:"Propertyid" form:"Propertyid"`
	Floor      string `json:"floor" form:"floor"`
	Buildingid string `json:"buildingid" form:"buildingid"`
}
type Building struct {
	Id         int    `json:"id" form:"id"`
	Propertyid string `json:"Propertyid" form:"Propertyid"`
	Building   string `json:"building" form:"building"`
	Gardenid   string `json:"gardenid" form:"gardenid"`
}
type Garden struct {
	Id         int    `json:"id" form:"id"`
	Propertyid string `json:"Propertyid" form:"Propertyid"`
	Garden     string `json:"garden" form:"garden"`
	Area       string `json:"area" form:"area"`
	City       string `json:"city" form:"city"`
	Province   string `json:"province" form:"province"`
}
type PropertyRoom struct {
	Floor      string `json:"floor" form:"floor"`
	Room       string `json:"room" form:"room"`
	Propertyid string `json:"propertyid" form:"propertyid"`
	Building   string `json:"building" form:"building"`
	Garden     string `json:"garden" form:"garden"`
}

func (p *Room) AddRoom() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO room(propertyid, floorid,room) VALUES (?, ?,?)", p.Propertyid, p.Floorid, p.Room)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}
func (p *Floor) AddFloor() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO floor(propertyid, buildingid,floor) VALUES (?, ?,?)", p.Propertyid, p.Buildingid, p.Floor)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}
func (p *Building) AddBuilding() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO building(propertyid, gardenid,building) VALUES (?, ?,?)", p.Propertyid, p.Gardenid, p.Building)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}
func (p *Garden) AddGarden() (id int64, err error) {
	// rs, err := db.SqlDB.Exec("INSERT INTO garden(propertyid, garden,area,city,province) VALUES (?, ?,?,?)", p.Propertyid, p.Garden, p.Area, p.City, p.Province)
	rs, err := db.SqlDB.Exec("INSERT INTO garden(propertyid, garden) VALUES (?, ?)", p.Propertyid, p.Garden)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}
func (p *Room) DelRoom() (ra int64, err error) {
	rs, err := db.SqlDB.Exec("DELETE FROM room WHERE propertyid=? and id=?", p.Propertyid, p.Id)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err = rs.RowsAffected()
	return
}
func (p *PropertyRoom) AddPropertyRoom() (err error) {
	//var id int64 = 0
	garden := Garden{Propertyid: p.Propertyid, Garden: p.Garden}
	id, err := garden.AddGarden()
	if err != nil {
		return
	}
	s := strconv.FormatInt(id, 10)
	building := Building{Propertyid: p.Propertyid, Gardenid: s, Building: p.Building}
	id, err = building.AddBuilding()
	if err != nil {
		return
	}
	s = strconv.FormatInt(id, 10)
	floor := Floor{Propertyid: p.Propertyid, Buildingid: s, Floor: p.Floor}
	id, err = floor.AddFloor()
	if err != nil {
		return
	}
	s = strconv.FormatInt(id, 10)
	room := Room{Propertyid: p.Propertyid, Floorid: s, Room: p.Room}
	id, err = room.AddRoom()
	if err != nil {
		return
	}
	// log.Print(err.Error())
	//log.Print(id)
	// rs, err := db.SqlDB.Exec("INSERT INTO room(propertyid, garden,building,floor,room) VALUES (?, ?,?,?,?)", p.Propertyid, p.Garden, p.Building, p.Floor, p.Room)
	// if err != nil {
	// 	return
	// }
	// id, err = rs.LastInsertId()
	return
}
func (p *Room) GetRoomByPropertyid() (*[]map[string]string, error) {
	rows, err := db.SqlDB.Query("SELECT * FROM room WHERE propertyid=?", p.Propertyid)
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
func (p *Room) GetRoomByPropertyidPage(start int, end int) (*[]map[string]string, error) {

	rows, err := db.SqlDB.Query("SELECT * FROM room WHERE propertyid=? limit ?,?", p.Propertyid, start, end)
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
func (p *Room) GetRoomByid() (*[]map[string]string, error) {
	// err = db.SqlDB.Query("SELECT * FROM room WHERE propertyid=?", p.Propertyid).Scan(
	// 	&roomtemp.Id, &roomtemp.Propertyid, &roomtemp.Room, &roomtemp.Floorid,
	// )
	// return
	rows, err := db.SqlDB.Query("SELECT * FROM room WHERE propertyid=? and id=?", p.Propertyid, p.Id)
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
func GetRoomCount(propertyid string) (string, error) {
	var count string
	row := db.SqlDB.QueryRow("SELECT count(*) FROM room WHERE propertyid=?", propertyid)
	err := row.Scan(&count)
	return count, err
}
func (p *Room) GetRoom() (roomtemp Room, err error) {
	err = db.SqlDB.QueryRow("SELECT * FROM room WHERE id=?", p.Id).Scan(
		&roomtemp.Id, &roomtemp.Propertyid, &roomtemp.Room, &roomtemp.Floorid,
	)
	return
}

func (p *Floor) GetFloor() (floortemp Floor, err error) {
	err = db.SqlDB.QueryRow("SELECT * FROM floor WHERE id=?", p.Id).Scan(
		&floortemp.Id, &floortemp.Propertyid, &floortemp.Floor, &floortemp.Buildingid,
	)
	return
}
func (p *Building) GetBuilding() (buildingtemp Building, err error) {
	err = db.SqlDB.QueryRow("SELECT * FROM building WHERE id=?", p.Id).Scan(
		&buildingtemp.Id, &buildingtemp.Propertyid, &buildingtemp.Building, &buildingtemp.Gardenid,
	)
	return
}
func (p *Garden) GetGarden() (gardentemp Garden, err error) {
	err = db.SqlDB.QueryRow("SELECT * FROM garden WHERE id=?", p.Id).Scan(
		&gardentemp.Id, &gardentemp.Propertyid, &gardentemp.Garden, &gardentemp.Area, &gardentemp.City, &gardentemp.Province,
	)
	return
}

func GetPropertyRoom(propertyid string, start int, end int) (*[]map[string]string, error) {
	var err error
	var roomtemp *[]map[string]string
	var propertyroom []map[string]string
	room := Room{Propertyid: propertyid}
	log.Print(propertyid)
	roomtemp, err = room.GetRoomByPropertyidPage(start, end)
	if err != nil {
		log.Print(err.Error())
		return &propertyroom, errors.New("room不存在")
	}
	log.Print(*roomtemp)
	for col := range *roomtemp {
		log.Print((*roomtemp)[col])
		log.Print((*roomtemp)[col]["room"])
		//}
		//for i := 0; i < 3; i++ {
		log.Print("ks")
		//log.Print(i)
		var id int
		id, err = strconv.Atoi((*roomtemp)[col]["floorid"])

		floor := Floor{Id: id}
		floor, err = floor.GetFloor()
		if err != nil {
			log.Print(err.Error())
			return &propertyroom, errors.New("floor不存在")
		}
		id, err = strconv.Atoi(floor.Buildingid)
		building := Building{Id: id}
		building, err = building.GetBuilding()
		if err != nil {
			log.Print(err.Error())
			return &propertyroom, errors.New("building不存在")
		}
		log.Print(building.Gardenid)
		id, err = strconv.Atoi(building.Gardenid)
		garden := Garden{Id: id}
		garden, err = garden.GetGarden()
		if err != nil {
			log.Print(err.Error())
			return &propertyroom, errors.New("garden不存在")
		}
		//(*propertyroom)[col] = PropertyRoom{Room: room.Room, Floor: floor.Floor, Propertyid: room.Propertyid, Building: building.Building, Garden: garden.Garden}
		//(*propertyroom)[col]["room"] = room.Room
		data := make(map[string]string)
		data["id"] = (*roomtemp)[col]["id"]
		data["room"] = (*roomtemp)[col]["room"]
		data["propertyid"] = (*roomtemp)[col]["propertyid"]
		data["floor"] = floor.Floor
		data["building"] = building.Building
		data["garden"] = garden.Garden
		propertyroom = append(propertyroom, data)
	}
	return &propertyroom, err
}

func GetPropertyRoomById(propertyid string, id int) (*[]map[string]string, error) {
	var err error
	var roomtemp *[]map[string]string
	var propertyroom []map[string]string
	room := Room{Propertyid: propertyid, Id: id}
	log.Print(propertyid)
	roomtemp, err = room.GetRoomByid()
	if err != nil {
		log.Print(err.Error())
		return &propertyroom, errors.New("room不存在")
	}
	log.Print(*roomtemp)
	for col := range *roomtemp {
		log.Print((*roomtemp)[col])
		log.Print((*roomtemp)[col]["room"])
		//}
		//for i := 0; i < 3; i++ {
		log.Print("ks")
		//log.Print(i)
		var id int
		id, err = strconv.Atoi((*roomtemp)[col]["floorid"])

		floor := Floor{Id: id}
		floor, err = floor.GetFloor()
		if err != nil {
			log.Print(err.Error())
			return &propertyroom, errors.New("floor不存在")
		}
		id, err = strconv.Atoi(floor.Buildingid)
		building := Building{Id: id}
		building, err = building.GetBuilding()
		if err != nil {
			log.Print(err.Error())
			return &propertyroom, errors.New("building不存在")
		}
		log.Print(building.Gardenid)
		id, err = strconv.Atoi(building.Gardenid)
		garden := Garden{Id: id}
		garden, err = garden.GetGarden()
		if err != nil {
			log.Print(err.Error())
			return &propertyroom, errors.New("garden不存在")
		}
		//(*propertyroom)[col] = PropertyRoom{Room: room.Room, Floor: floor.Floor, Propertyid: room.Propertyid, Building: building.Building, Garden: garden.Garden}
		//(*propertyroom)[col]["room"] = room.Room
		data := make(map[string]string)
		data["id"] = (*roomtemp)[col]["id"]
		data["room"] = (*roomtemp)[col]["room"]
		data["propertyid"] = (*roomtemp)[col]["propertyid"]
		data["floor"] = floor.Floor
		data["building"] = building.Building
		data["garden"] = garden.Garden
		propertyroom = append(propertyroom, data)
	}
	return &propertyroom, err
}

// func (p *Room) GetRoom() (roomtemp Room, err error) {
// 	err = db.SqlDB.QueryRow("SELECT * FROM room WHERE mobile=?", p.Mobile).Scan(
// 		&roomtemp.Id, &roomtemp.Mobile, &roomtemp.Roomid, &roomtemp.Roomname,
// 	)
// 	return
// }

// func (p *Room) UpdateRoom() (ra int64, err error) {

// 	stmt, err := db.SqlDB.Prepare("UPDATE room SET roomid=? , roomname =? WHERE mobile=?")
// 	defer stmt.Close()
// 	if err != nil {
// 		return
// 	}
// 	rs, err := stmt.Exec(p.Roomid, p.Roomname, p.Mobile)
// 	if err != nil {
// 		return
// 	}
// 	ra, err = rs.RowsAffected()
// 	return
// }
