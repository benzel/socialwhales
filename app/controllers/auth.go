package controllers

import (
    "github.com/robfig/revel"
    "log"
)

type Auth struct {
	*revel.Controller
}


func (c Auth) Signup() revel.Result {
    log.Print(c.Params)
    return c.RenderJson(map[string] string {
        "message": "logout works",
    })
}


func (c Auth) Login() revel.Result {
    log.Print(c.Params)
    return c.RenderJson(map[string] string {
        "message": "login works",
    })
}


func (c Auth) Logout() revel.Result {
    log.Print(c.Params)
    return c.RenderJson(map[string] string {
        "message": "logout works",
    })
}
