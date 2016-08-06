package controllers
import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	elastigo "github.com/mattbaird/elastigo/lib"
)

var (
	host *string = flag.String("host", "localhost", "Elasticsearch Host")
)

type Contact struct {
	Name string
	Phone string
}

type (
	ContactController struct{}
)

func NewContactController() *ContactController {
	return &ContactController{}
}

func (con ContactController) Add(g *gin.Context)  {
	id:=g.Param("id")

	contact := Contact{}
	g.Bind(&contact)

	c := elastigo.NewConn()
	c.Domain = *host
	// Index a doc using Structs
	r, _ := c.Index("contacts", "contact", id, nil, Contact{contact.Name, contact.Phone})
	log.Printf("post: %v", r.Exists)
}

func (con ContactController) Update(g *gin.Context){
	id:=g.Param("id")
	contact := Contact{}
	g.Bind(&contact)
	c := elastigo.NewConn()
	c.Domain = *host
	// Index a doc using Structs
	r, _ := c.Index("contacts", "contact", id, nil, Contact{contact.Name, contact.Phone})
	log.Printf("put: %v", r.Exists)
}

func (con ContactController) Delete(g *gin.Context)  {
	id:=g.Param("id")
	c := elastigo.NewConn()
	c.Domain = *host
	r, _ := c.Delete("contacts","contact",id, nil)
	log.Printf("Delete: %v", r.Exists)
}

func (con ContactController) Get(g *gin.Context){
	id:=g.Param("id")
	c := elastigo.NewConn()
	c.Domain = *host
	response, _:= c.Get("contacts","contact", id, nil)
	fmt.Println(response)
}