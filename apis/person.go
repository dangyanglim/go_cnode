package apis

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	. "github.com/dangyanglim/go_cnode/models"
	"github.com/gin-gonic/gin"
)

func IndexApi(c *gin.Context) {
	//c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	c.String(http.StatusOK, "It works")
}

//var person=new(models.Person)
func AddPersonApi(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")

	p := Person{FirstName: firstName, LastName: lastName}
	//var p models.InsetPerson

	//err:=c.BindJSON(&p)
	ra, err := p.AddPerson()
	if err != nil {
		//log.Fatalln(err)
		log.Fatal(err.Error())
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
func GetPersonsApi(c *gin.Context) {

	var p Person
	persons, err := p.GetPersons()
	if err != nil {
		//log.Fatalln(err)
		//log.Fatal(err.Error())
		log.Print(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"data": persons,
	})
}
func GetPersonApi(c *gin.Context) {
	cid := c.Param("id")
	fmt.Println("getpersonapi")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	p := Person{Id: id}

	person, err := p.GetPerson()
	if err != nil {
		//log.Fatalln(err)
		log.Print(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"person": person,
	})
}
func ModPersonApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	p := Person{Id: id}
	err = c.Bind(&p)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err := p.ModPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("Update person %d successful %d", p.Id, ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
func DelPersonApi(c *gin.Context) {
	cid := c.Param("id")
	fmt.Println("delpersonapi")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	p := Person{Id: id}
	ra, err := p.DelPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("Delete person %d successful %d", id, ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
