package models

//import "fmt"
import "labix.org/v2/mgo"
import "labix.org/v2/mgo/bson"
import "github.com/cfrco/mon"

type Message struct {
    // Always use "_id", and give it a new ObjectId when we create a record
    Id bson.ObjectId "_id" 
    mon.Meta "-" // Always use "-" for mon.Meta to avoid store `meta` to DB.
    AuthorName string
    AuthorMail string
    Content string
}

// Uesr-defined New function, like constructor
func NewMessage(collection *mgo.Collection, name string , mail string,
                content string) *Message {
    p := new(Message)
    p.Meta.Bind(collection, p)
    p.Id = bson.NewObjectId()

    p.AuthorName = name
    p.AuthorMail = mail
    p.Content = content

    return p
}
