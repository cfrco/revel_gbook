package controllers

import "github.com/robfig/revel"
import "labix.org/v2/mgo"
import "labix.org/v2/mgo/bson"
import "gbook/app/models"
import "github.com/cfrco/mon"
import "fmt"

type App struct {
    *revel.Controller
}

func (c App) Index() revel.Result {
    session, err := mgo.Dial("localhost")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    // Optional. Switch the session to a monotonic behavior.
    session.SetMode(mgo.Monotonic, true)
    
    co := session.DB("SE_test").C("message")
    if err != nil {
        panic(err)
    }

    var messages []models.Message
    mon.QueryAll(co,co.Find(bson.M{}),&messages)

    return c.Render(messages)
}

func (c App) New(name string) revel.Result {
    session, err := mgo.Dial("localhost")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    // Optional. Switch the session to a monotonic behavior.
    session.SetMode(mgo.Monotonic, true)
    
    co := session.DB("SE_test").C("message")
    if err != nil {
        panic(err)
    }
    
    method := c.Controller.Params.Values["METHOD"][0]
    if method == "GET" {
    } else if method == "POST" {
        //fmt.Println(c.Controller.Params.Values)
        message := models.NewMessage(co,
            c.Controller.Params.Values["name"][0],
            c.Controller.Params.Values["mail"][0],
            c.Controller.Params.Values["content"][0])
        message.Insert()
        return c.Redirect("/")
    }
    return c.Render()
}
