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

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	uid, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		err := errors.NewBadRequestError("invalid user id")
		return 0, err
	}
	return uid, nil
}

func Get(c *gin.Context) {
	uid, err := getUserId(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	u, err2 := services.UserService.GetUser(uid)
	if err2 != nil {
		c.JSON(err2.Status, err2)
		return
	}
	c.JSON(http.StatusOK, u.Marshall(c.GetHeader("X-Public") == "true"))
}

func Create(c *gin.Context) {
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
	cr, err := services.UserService.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, cr.Marshall(c.GetHeader("X-Public") == "true"))
}

func Update(c *gin.Context) {
	var user users.User
	uid, err := getUserId(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.NewBadRequestError("invalid json body")
		c.JSON(err.Status, err)
		return
	}
	user.Id = uid
	isPart := c.Request.Method == http.MethodPatch
	us, errUp := services.UserService.UpdateUser(isPart, user)
	if errUp != nil {
		c.JSON(errUp.Status, errUp)
		return
	}
	c.JSON(http.StatusOK, us.Marshall(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {
	uid, err := getUserId(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	if err := services.UserService.DeleteUser(uid); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": strconv.Itoa(http.StatusOK), "message": "Deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UserService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}

func Login(c *gin.Context) {
	var lg users.Login
	if err := c.ShouldBindJSON(&lg); err != nil {
		err := errors.NewBadRequestError("invalid login json body")
		c.JSON(err.Status, err)
		return
	}
	u, err := services.UserService.LoginUser(lg)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, u.Marshall(c.GetHeader("X-Public") == "true"))
	
}