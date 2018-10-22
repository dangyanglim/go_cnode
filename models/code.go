package models

import db "github.com/dangyanglim/go_cnode/database"

//"log"

type Code struct {
	Id           int    `json:"id" form:"id"`
	Mobile       string `json:"mobile" form:"mobile"`
	Token        string `json:"token" form:"token"`
	TokenExptime string `json:"tokenExptime" form:"tokenExptime"`
}
type UserToken struct {
	Id           int    `json:"id" form:"id"`
	Mobile       string `json:"mobile" form:"mobile"`
	Token        string `json:"token" form:"token"`
	TokenExptime string `json:"tokenExptime" form:"tokenExptime"`
}
type User struct {
	Id     int    `json:"id" form:"id"`
	Mobile string `json:"mobile" form:"mobile"`
	Pwd    string `json:"pwd" form:"pwd"`
}

func (p *Code) AddCode() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO usercode(mobile, token,tokenExptime) VALUES (?, ?,?)", p.Mobile, p.Token, p.TokenExptime)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}
func (p *UserToken) AddToken() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO usertoken(mobile, token,tokenExptime) VALUES (?, ?,?)", p.Mobile, p.Token, p.TokenExptime)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}
func (p *User) AddUser() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO user(mobile, pwd) VALUES (?, ?)", p.Mobile, p.Pwd)
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
func (p *Code) GetUserCode() (usercode Code, err error) {
	err = db.SqlDB.QueryRow("SELECT id, mobile, token ,tokenExptime FROM usercode WHERE mobile=?", p.Mobile).Scan(
		&usercode.Id, &usercode.Mobile, &usercode.Token, &usercode.TokenExptime,
	)
	return
}
func (p *UserToken) GetUserToken() (usercode UserToken, err error) {
	err = db.SqlDB.QueryRow("SELECT id, mobile, token ,tokenExptime FROM usertoken WHERE mobile=?", p.Mobile).Scan(
		&usercode.Id, &usercode.Mobile, &usercode.Token, &usercode.TokenExptime,
	)
	return
}
func (p *User) GetUser() (user User, err error) {
	err = db.SqlDB.QueryRow("SELECT * from user WHERE mobile=?", p.Mobile).Scan(
		&user.Id, &user.Mobile, &user.Pwd,
	)
	return
}
func (p *Code) UpdateCode() (ra int64, err error) {

	stmt, err := db.SqlDB.Prepare("UPDATE usercode SET token=? , tokenExptime =? WHERE mobile=?")
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
func (p *User) UpdateUserPwd() (ra int64, err error) {

	stmt, err := db.SqlDB.Prepare("UPDATE user SET pwd=? WHERE mobile=?")
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
func (p *UserToken) UpdateToken() (ra int64, err error) {

	stmt, err := db.SqlDB.Prepare("UPDATE usertoken SET token=? , tokenExptime =? WHERE mobile=?")
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

// func (p *Person) DelPerson() (ra int64, err error) {
// 	rs, err := db.SqlDB.Exec("DELETE FROM person WHERE id=?", p.Id)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	ra, err = rs.RowsAffected()
// 	return
// }
