package controllers

import "net/http"
import "io/ioutil"

import "github.com/robfig/revel"

type Html string

func (r Html) Apply(req *revel.Request, resp *revel.Response) {
	resp.WriteHeader(http.StatusOK, "text/html")
	resp.Out.Write([]byte(r))
}

var IndexCache Html

func CacheIndex() {
	content, err := ioutil.ReadFile(revel.AppPath + "/views/App/Index.html")
	if err != nil {
		revel.ERROR.Println(err.Error())
	}
	IndexCache = Html(content)

}

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	c.Response.ContentType = "text/html"
	// We will return just raw text because go templates
	// conflict with angular js templates.
	return IndexCache
}
