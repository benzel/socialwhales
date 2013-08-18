package controllers

import "fmt"
import "net/http"
import "io/ioutil"
import "encoding/json"
import "github.com/robfig/revel"
import "github.com/benzel/socialwhales/app/models"
import "github.com/benzel/socialwhales/app/services/storage"

var Profiles storage.ProfilesAPI
var Accounts storage.AccountsAPI

func InitAuth() {
	revel.INFO.Println("Initializing AUTH controller")
	storageAlias, found := revel.Config.String("app.storage")
	if !found {
		revel.ERROR.Fatal("storage: Storage service is not specified in configs.")
	}
	revel.INFO.Println("Using " + storageAlias + " storage")
	Profiles = storage.GetStorage(storageAlias).GetProfiles()
	Accounts = storage.GetStorage(storageAlias).GetAccounts()
	if Profiles == nil {
		revel.ERROR.Panic("Profiles is nil")
	}
	if Accounts == nil {
		revel.ERROR.Panic("Accounts is nil")
	}
	revel.INFO.Println(storage.GetStorage(storageAlias))
	revel.INFO.Println(Accounts)
}

type Auth struct {
	*revel.Controller
}

type AuthError struct {
	Error        int    `json:"error"`
	ErrorDetails string `json:"error_details"`
	Flash        string `json:"flash,omitempty"`
}

type SignUpResponse struct {
	Error int    `json:"error"`
	Flash string `json:"flash,omitempty"`
}

type LoginResponse struct {
	Error        int    `json:"error"`
	AccountId    int64  `json:"id"`
	AccountEmail string `json:"email"`
	AccountToken string `json:"token"`
	Flash        string `json:"flash,omitempty"`
}

func (c Auth) Signup() revel.Result {
	var credentials models.Credentials

	json_data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Response.Status = http.StatusUnauthorized
		return c.RenderJson(AuthError{
			Error:        1,
			Flash:        "Error while reading request body.",
			ErrorDetails: err.Error(),
		})
	}

	if err := json.Unmarshal([]byte(json_data), &credentials); err != nil {
		c.Response.Status = http.StatusUnauthorized
		return c.RenderJson(AuthError{
			Error:        1,
			Flash:        "Invalid JSON format in the request body.",
			ErrorDetails: err.Error(),
		})
	}

	credentials.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Response.Status = http.StatusUnauthorized
		return c.RenderJson(AuthError{
			Error:        1,
			Flash:        "Invalid credentials.",
			ErrorDetails: err.Error(),
		})
	}

	revel.INFO.Println(credentials)
	_, err = Accounts.Create(&credentials)

	if err != nil {
		c.Response.Status = http.StatusUnauthorized
		return c.RenderJson(AuthError{
			Error:        1,
			Flash:        "API error while creating new account. " + err.Error(),
			ErrorDetails: err.Error(),
		})
	}

	// TODO:(zaytsev@usc.edu): send activation email
	return c.RenderJson(SignUpResponse{
		Error: 0,
		Flash: fmt.Sprintf("New account was successfully created. Activation email sent to %s.", credentials.Email),
	})
}

func (c Auth) Login() revel.Result {
	var credentials models.Credentials

	json_data, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.Response.Status = http.StatusUnauthorized
		return c.RenderJson(AuthError{
			Error:        1,
			Flash:        "Error while reading request body.",
			ErrorDetails: err.Error(),
		})
	}

	if err := json.Unmarshal([]byte(json_data), &credentials); err != nil {
		c.Response.Status = http.StatusUnauthorized
		return c.RenderJson(AuthError{
			Error:        1,
			Flash:        "Invalid json format in the request body.",
			ErrorDetails: err.Error(),
		})
	}

	account, err := Accounts.Read(&credentials)

	if err != nil {
		c.Response.Status = http.StatusUnauthorized
		return c.RenderJson(AuthError{
			Error:        1,
			Flash:        "Invalid credentials.",
			ErrorDetails: err.Error(),
		})
	}

	return c.RenderJson(LoginResponse{
		Error:        0,
		AccountId:    account.Id,
		AccountEmail: account.Email,
		AccountToken: account.Token,
		Flash:        "You have signed in.",
	})
}

func (c Auth) Logout() revel.Result {
	c.Response.Status = http.StatusNotImplemented
	return c.RenderJson(AuthError{
		Error:        1,
		Flash:        "Not implemented.",
		ErrorDetails: "Not implemented.",
	})
}
