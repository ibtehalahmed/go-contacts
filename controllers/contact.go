package controllers
import (
    "fmt"
    "gopkg.in/mgo.v2"
    "github.com/gin-gonic/gin"
    "gopkg.in/mgo.v2/bson"
    "github.com/ibtehal/go-contacts/models"
)

type Contact struct { Name string `json:"name" bson:"name"`
                      Phone string `json:"phone" bson:"phone"`
                      Id     bson.ObjectId `json:"id" bson:"_id"`
}
type (
    ContactController struct{session *mgo.Session}
)

func NewContactController(s *mgo.Session) *ContactController {
    return &ContactController{s}
}

func (con ContactController) Add(c *gin.Context) {

    a := con.session.DB("contacts").C("contacts")
    contact := models.Contact{}
    contact.Id = bson.NewObjectId()
    c.Bind(&contact)
    a.Insert(contact)

}

func (con ContactController) GetOne(c *gin.Context) {
    id := bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))}
    m := models.Contact{}
    a := con.session.DB("contacts").C("contacts")
    a.Find(id).One(&m)
    fmt.Println(m.Name)
}

func (con ContactController) GetAll(c *gin.Context) {

    contacts := []models.Contact{}
    a := con.session.DB("contacts").C("contacts")
    a.Find(nil).Sort("-updated_on").All(&contacts)

}

func (con ContactController) Update(c *gin.Context) {

    a := con.session.DB("contacts").C("contacts")
    contact := models.Contact{}
    c.Bind(&contact)

    query := bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))}
    doc := bson.M{
        "name": contact.Name,
        "phone": contact.Phone,
    }
    a.Update(query, doc)
}
func (con ContactController) Delete(c *gin.Context) {
    a := con.session.DB("contacts").C("contacts")
    query := bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))}
    a.Remove(query)

}