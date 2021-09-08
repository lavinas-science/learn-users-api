package users

import (
	// "encoding/json"
	// "io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lavinas-science/learn-users-api/domain/users"
	"github.com/lavinas-science/learn-users-api/services"
)

func GetUser (c *gin.Context) {	
	c.String(http.StatusNotImplemented, "Implement get user !\n")
}

func CreateUser (c *gin.Context) {
	var user users.User
	/* see refact below
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// TODO: Handle error
		return
	}
	if err = json.Unmarshal(b, &user); err != nil {
		// TODO: Handle error
		return
	}
	*/
	if err := c.ShouldBindJSON(&user); err != nil {
		// TODO: Handle error
		return
	}

	cr, err := services.CreateUser(user)
	if err != nil {
		// TODO: Handle error
		return
	}
	c.JSON(http.StatusCreated, cr)
}


func SearchUser (c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement search user !\n")
}