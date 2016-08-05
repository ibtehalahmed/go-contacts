package main
import (
    "github.com/gin-gonic/gin"
    "gopkg.in/mgo.v2"
    "github.com/ibtehal/go-contacts/controllers"
)

func main() {
	r := gin.Default()
	con := controllers.NewContactController(getSession())

	r.POST("/contact", con.Add)
	r.GET("/contact/:_id", con.GetOne)
	r.GET("/contact", con.GetAll)
	r.PUT("/contact/:_id", con.Update)
	r.DELETE("/contact/:_id", con.Delete)

	r.Run()
}

func getSession() *mgo.Session {
    // Connect to our local mongo
    s, err := mgo.Dial("mongodb://localhost")

    // Check if connection error, is mongo running?
    if err != nil {
        panic(err)
    }
    return s
}