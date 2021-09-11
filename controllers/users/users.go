package users

import (
	// "encoding/json"
	// "io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lavinas-science/learn-users-api/domain/users"
	"github.com/lavinas-science/learn-users-api/services"
	"github.com/lavinas-science/learn-users-api/utils/errors"
)

func GetUser(c *gin.Context) {
	uid, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Status, err)
		return
	}

	u, err2 := services.GetUser(uid)
	if err2 != nil {
		c.JSON(err2.Status, err2)
		return
	}
	c.JSON(http.StatusOK, u)
}

func CreateUser(c *gin.Context) {
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
		err := errors.NewBadRequestError("invalid json body")
		c.JSON(err.Status, err)
		return
	}

	cr, err := services.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, cr)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement search user !\n")
}


func UpdateUser(c *gin.Context) {
	var user users.User

	uid, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Status, err)
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.NewBadRequestError("invalid json body")
		c.JSON(err.Status, err)
		return
	}

	user.Id = uid

	us, errUp := services.UpdateUser(user)
	if err != nil {
		c.JSON(errUp.Status, err)
		return
	}
	c.JSON(http.StatusOK, us)

}