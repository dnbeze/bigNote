package db

import (
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func InitDB(c echo.Context, uri string) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(c.Request().Context(), opts) //echo.Context isnt context.context but you can get one from Request().Context() seen here :)
	if err != nil {
		log.Fatalf("failed to connect to database")
	}
	return client, nil
}

///
//thanks stackoverflow
//An echo.Context can't be used as a context.Context, but it does refer to one; if c is the echo context then c.Request().Context() is the context.Context for the incoming request, which will be canceled if the user closes the connection, for instance.

//You can copy other useful values from the echo context to the stdlib context, if needed, by using the Get method on the one hand and the WithValue method on the other. If there are some values that you always want copied, then you can write a helper function to do so.
//Share
//Improve this answer
//Follow
//answered Dec 16, 2017 at 9:07
//hobbs's user avatar
//hobbs
//230k1919 gold badges217217 silver badges294
