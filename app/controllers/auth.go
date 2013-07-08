package controllers

import "io/ioutil"
import "encoding/json"
import "github.com/robfig/revel"
import "github.com/benzel/socialwhales/app/models"
import "github.com/benzel/socialwhales/app/services/storage"

var Profiles storage.ProfilesAPI = storage.GetStorage("MEM").GetProfiles()

type Auth struct {
	*revel.Controller
}

func (c Auth) Signup() revel.Result {
	var credentials models.Credentials

	// Read request parameters and deserialize an instance of credentials
	json_data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return c.RenderJson(map[string]string{
			"error":         "1",
			"error_message": "error while reading request body",
		})
	}
	if err := json.Unmarshal([]byte(json_data), &credentials); err != nil {
		return c.RenderJson(map[string]string{
			"error":         "1",
			"error_message": "invalid json format in the request body",
		})
	}

	// perform validation
	credentials.Validate(c.Validation)
	if c.Validation.HasErrors() {
		return c.RenderJson(map[string]string{
			"error":         "1",
			"error_message": "invalid user credentials",
		})
	}

	return c.RenderJson(map[string]string{
		"error":           "0",
		"success_message": "successfully performed account user signup",
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
