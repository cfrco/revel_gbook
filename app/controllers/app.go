package controllers

import "github.com/robfig/revel"
import "labix.org/v2/mgo"
import "labix.org/v2/mgo/bson"
import "gbook/app/models"
import "github.com/cfrco/mon"
//import "fmt"

type App struct {
    *revel.Controller
    DBsession *mgo.Session
}

func (c *App) connectDB() {
    session, err := mgo.Dial("localhost")
    if err != nil {
        panic(err)
    }
    c.DBsession = session

    // Optional. Switch the session to a monotonic behavior.
    c.DBsession.SetMode(mgo.Monotonic, true)
}

func (c App) Index() revel.Result {
    c.connectDB()
    defer c.DBsession.Close()
    co := c.DBsession.DB("SE_test").C("message")

    var messages []models.Message
    mon.QueryAll(co,co.Find(bson.M{}),&messages)

    return c.Render(messages)
}

func (c App) New(name string, mail string, content string) revel.Result {
    c.connectDB()
    defer c.DBsession.Close()
    co := c.DBsession.DB("SE_test").C("message")
    
    method := c.Controller.Params.Values["METHOD"][0]
    if method == "GET" {
    } else if method == "POST" {
        //fmt.Println(c.Controller.Params.Values)
        message := models.NewMessage(co, name, mail, content)
        message.Insert()
        return c.Redirect("/")
    }
    return c.Render()
}

func (c App) Edit(mid string) revel.Result {
    c.connectDB()
    defer c.DBsession.Close()
    co := c.DBsession.DB("SE_test").C("message")

    message := models.NewMessage(co, "", "", "")
    message.Find(bson.M{"_id":bson.ObjectIdHex(mid)})

    method := c.Controller.Params.Values["METHOD"][0]
    if method == "GET" {
    } else if method == "POST" {
        message.AuthorName = c.Controller.Params.Values["name"][0]
        message.AuthorMail = c.Controller.Params.Values["mail"][0]
        message.Content = c.Controller.Params.Values["content"][0]
        message.Update()
        return c.Redirect("/")
    }

    return c.Render(message)
}

func (c App) Remove(mid string) revel.Result {
    c.connectDB()
    defer c.DBsession.Close()
    co := c.DBsession.DB("SE_test").C("message")

    message := models.NewMessage(co, "", "", "")
    message.Find(bson.M{"_id":bson.ObjectIdHex(mid)})

    message.Remove()

    return c.Redirect("/")
}
