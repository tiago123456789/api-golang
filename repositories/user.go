package repositories

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo"
)


type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserRepository struct{}

var number = 1
var users = map[int]*User{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) CreateUser(c echo.Context) error {
	u := &user{
		Id: number,
	}
  
	err := c.Bind(u);
	if err != nil {
	  return err
	}
  
	users[u.Id] = u
	number++
	  
	return c.JSON(http.StatusOK, u)
}

func (ur *UserRepository) UpdateUser(c echo.Context) error {
	u := new(user);
	if err := c.Bind(u); err != nil {
		return err;
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err;
	}
	
	users[id].Name = u.Name
	
	return c.JSON(http.StatusOK, users[id])
}

func (ur *UserRepository) getUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err;
	}

	return c.JSON(http.StatusOK, users[id]);
}

func (ur *UserRepository) getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}
