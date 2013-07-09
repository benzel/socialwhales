package controllers

import "io/ioutil"
import "encoding/json"
import "github.com/robfig/revel"
import "github.com/benzel/socialwhales/app/models"
import "github.com/benzel/socialwhales/app/services/storage"

var Profiles storage.ProfilesAPI = storage.GetStorage("MEM").GetProfiles()
var Accounts storage.AccountsAPI = storage.GetStorage("MEM").GetAccounts()

type Auth struct {
	*revel.Controller
}

type AuthErrorResponse struct {
	ErrorFlag    int    `json:"error"`
	ErrorCode    int    `json:"error_code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
	ErrorDetails string `json:"error_details,omitempty"`
}

type SignInResponse struct {
	ErrorFlag    int    `json:"error"`
	AccountId    int64  `json:"user_id"`
	AccountToken string `json:"user_token"`
}

func (c Auth) Signup() revel.Result {
	var credentials models.Credentials

	// Read request parameters and deserialize an instance of credentials
	json_data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return c.RenderJson(AuthErrorResponse{
			ErrorFlag:    1,
			ErrorMessage: "Error while reading request body.",
			ErrorDetails: err.Error(),
		})
	}

	if err := json.Unmarshal([]byte(json_data), &credentials); err != nil {
		return c.RenderJson(AuthErrorResponse{
			ErrorFlag:    1,
			ErrorMessage: "Invalid json format in the request body.",
			ErrorDetails: err.Error(),
		})
	}

	// perform validation
	credentials.Validate(c.Validation)
	if c.Validation.HasErrors() {
		return c.RenderJson(AuthErrorResponse{
			ErrorFlag:    1,
			ErrorMessage: "Invalid user credentials.",
			ErrorDetails: err.Error(),
		})
	}

	accountId, err := Accounts.Create(credentials)

	if err != nil {
		return c.RenderJson(AuthErrorResponse{
			ErrorFlag:    1,
			ErrorMessage: "API error while creating new account.",
			ErrorDetails: err.Error(),
		})

	}
	return c.RenderJson(SignInResponse{
		ErrorFlag:    0,
		AccountId:    accountId,
		AccountToken: "<token>",
	})
}

func (c Auth) Login() revel.Result {
	return c.RenderJson(AuthErrorResponse{
		ErrorFlag:    1,
		ErrorMessage: "Not implemented.",
	})
}

func (c Auth) Logout() revel.Result {
	return c.RenderJson(AuthErrorResponse{
		ErrorFlag:    1,
		ErrorMessage: "Not implemented.",
	})
}
