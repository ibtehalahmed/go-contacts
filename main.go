package main
import (
	"github.com/gin-gonic/gin"
	"github.com/ibtehal/go-contacts/controllers"
)

func main() {
	r := gin.Default()
	con := controllers.NewContactController()
	r.POST("/contacts/contact/:id", con.Add)
	r.DELETE("/contacts/contact/:id",con.Delete)
	r.GET("/contacts/contact",con.Get)
	r.PUT("/contacts/contact/:id",con.Update)
	r.POST("/contacts/contact/:name",con.SearchByName)
	r.Run()
}