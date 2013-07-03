package controllers

import (
	"io/ioutil"
	"encoding/json"
	"github.com/robfig/revel"
	"github.com/benzel/socialwhales/app/models"
	"github.com/benzel/socialwhales/app/services/storage"
)

type Auth struct {
	*revel.Controller
}

var users storage.Users = storage.GetStorage("memory").GetUsers()

func (c Auth) Signup() revel.Result {
	var credentials models.Credentials
	json_data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return c.RenderJson(map[string]string{
			"message": "error",
		})
	}
	
	if err := json.Unmarshal([]byte(json_data), &credentials); err != nil {
		return c.RenderJson(map[string]string{
			"message": "error",
		})
	}
	
	// user := models.User{}
	// user.Get("test", "test")
	//
	// revel.INFO.Println("Email: " + credentials.Email)
	// revel.INFO.Println("Password: " + credentials.Password)
	return c.RenderJson(map[string]string{
		"message": "error",
	})
}

func (c Auth) Login() revel.Result {
	// log.Print(c.Params)
	return c.RenderJson(map[string]string{
		"message": "login works",
	})
}

func (c Auth) Logout() revel.Result {
	// log.Print(c.Params)
	return c.RenderJson(map[string]string{
		"message": "logout works",
	})
}
