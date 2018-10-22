package models

import (
	"log"

	db "github.com/dangyanglim/go_cnode/database"
)

type Property struct {
	Id     int    `json:"id" form:"id"`
	Pwd    string `json:"pwd" form:"pwd"`
	Mobile string `json:"mobile" form:"mobile"`
	Name   string `json:"name" form:"name"`
}
type PropertyToken struct {
	Id           int    `json:"id" form:"id"`
	Token        string `json:"token" form:"token"`
	Mobile       string `json:"mobile" form:"mobile"`
	TokenExptime string `json:"tokenExptime" form:"tokenExptime"`
}
type PropertyCheckCode struct {
	Id           int    `json:"id" form:"id"`
	CheckCode    string `json:"checkcode" form:"checkcode"`
	Mobile       string `json:"mobile" form:"mobile"`
	TokenExptime string `json:"tokenExptime" form:"tokenExptime"`
}

func (p *Property) AddProperty() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO property(mobile, password) VALUES (?, ?)", p.Mobile, p.Pwd)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}
func (p *PropertyToken) AddPropertyToken() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO propertytoken(mobile, token,tokenExptime) VALUES (?, ?,?)", p.Mobile, p.Token, p.TokenExptime)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}
func (p *PropertyCheckCode) AddPropertyCheckCode() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO propertycheckcode(mobile, checkcode,tokenExptime) VALUES (?, ?,?)", p.Mobile, p.CheckCode, p.TokenExptime)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

// func (p *Person) GetPersons() (persons []Person, err error) {
// 	persons = make([]Person, 0)
// 	rows, err := db.SqlDB.Query("SELECT id, first_name, last_name FROM person")
// 	defer rows.Close()

// 	if err != nil {
// 		return
// 	}

// 	for rows.Next() {
// 		var person Person
// 		rows.Scan(&person.Id, &person.FirstName, &person.LastName)
// 		persons = append(persons, person)
// 	}
// 	if err = rows.Err(); err != nil {
// 		return
// 	}
// 	return
// }
func (p *Property) GetProperty() (propery Property, err error) {
	err = db.SqlDB.QueryRow("SELECT id,mobile,password,name FROM property WHERE mobile=?", p.Mobile).Scan(
		&propery.Id, &propery.Mobile, &propery.Pwd, &propery.Name,
	)
	return
}
func (p *PropertyToken) GetPropertyToken() (token PropertyToken, err error) {
	err = db.SqlDB.QueryRow("SELECT id,mobile,token,tokenExptime FROM propertytoken WHERE mobile=?", p.Mobile).Scan(
		&token.Id, &token.Mobile, &token.Token, &p.TokenExptime,
	)
	return
}
func (p *PropertyCheckCode) GetPropertyCheckCode() (checkcode PropertyCheckCode, err error) {
	err = db.SqlDB.QueryRow("SELECT id,mobile,checkcode,tokenExptime FROM propertycheckcode WHERE mobile=?", p.Mobile).Scan(
		&checkcode.Id, &checkcode.Mobile, &checkcode.CheckCode, &p.TokenExptime,
	)
	return
}

// func (p *Person) ModPerson() (ra int64, err error) {
// 	stmt, err := db.SqlDB.Prepare("UPDATE person SET first_name=?, last_name=? WHERE id=?")
// 	defer stmt.Close()
// 	if err != nil {
// 		return
// 	}
// 	rs, err := stmt.Exec(p.FirstName, p.LastName, p.Id)
// 	if err != nil {
// 		return
// 	}
// 	ra, err = rs.RowsAffected()
// 	return
// }
func (p *Property) DelProperty() (ra int64, err error) {
	rs, err := db.SqlDB.Exec("DELETE FROM property WHERE mobile=?", p.Id)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err = rs.RowsAffected()
	return
}
func (p *PropertyToken) DelPropertyToken() (ra int64, err error) {
	rs, err := db.SqlDB.Exec("DELETE FROM propertytoken WHERE mobile=?", p.Mobile)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err = rs.RowsAffected()
	return
}
func (p *PropertyCheckCode) DelPropertyCheckCode() (ra int64, err error) {
	rs, err := db.SqlDB.Exec("DELETE FROM propertycheckcode WHERE mobile=?", p.Mobile)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err = rs.RowsAffected()
	return
}

func (p *PropertyCheckCode) UpdatePropertyCode() (ra int64, err error) {

	stmt, err := db.SqlDB.Prepare("UPDATE propertycheckcode SET checkcode=? , tokenExptime =? WHERE mobile=?")
	defer stmt.Close()
	if err != nil {
		return
	}
	rs, err := stmt.Exec(p.CheckCode, p.TokenExptime, p.Mobile)
	if err != nil {
		return
	}
	ra, err = rs.RowsAffected()
	return
}
func (p *PropertyToken) UpdatePropertyToken() (ra int64, err error) {

	stmt, err := db.SqlDB.Prepare("UPDATE propertytoken SET token=? , tokenExptime =? WHERE mobile=?")
	defer stmt.Close()
	if err != nil {
		return
	}
	rs, err := stmt.Exec(p.Token, p.TokenExptime, p.Mobile)
	if err != nil {
		return
	}
	ra, err = rs.RowsAffected()
	return
}
func (p *Property) UpdatePropertyPwd() (ra int64, err error) {

	stmt, err := db.SqlDB.Prepare("UPDATE property SET password=? WHERE mobile=?")
	defer stmt.Close()
	if err != nil {
		return
	}
	rs, err := stmt.Exec(p.Pwd, p.Mobile)
	if err != nil {
		return
	}
	ra, err = rs.RowsAffected()
	return
}
func (p *Property) UpdatePropertyName() (ra int64, err error) {

	stmt, err := db.SqlDB.Prepare("UPDATE property SET name=? WHERE mobile=?")
	defer stmt.Close()
	if err != nil {
		return ra, err
	}
	rs, err := stmt.Exec(p.Name, p.Mobile)
	if err != nil {
		return ra, err
	}
	ra, err = rs.RowsAffected()
	return
}
