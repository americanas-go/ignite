package main

import (
	"context"
	"os"
	"time"

	"github.com/americanas-go/config"
	ilog "github.com/americanas-go/ignite/americanas-go/log.v1"
	mgo "github.com/americanas-go/ignite/go.mongodb.org/mongo-driver.v1"
	vault "github.com/americanas-go/ignite/mittwald/vaultgo.v0"
	"github.com/americanas-go/ignite/mittwald/vaultgo.v0/managers/contrib/go.mongodb.org/mongo-driver.v1"
	"github.com/americanas-go/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	os.Setenv("IGNITE_MONGO_URI", "mongodb://localhost:27004/database")
	os.Setenv("IGNITE_VAULT_MANAGERS_MONGO_SECRET__PATH", "v1/database/creds/recency_vault_hml")
	os.Setenv("IGNITE_VAULT_ADDR", "http://vault.example.com")
	os.Setenv("IGNITE_VAULT_TOKEN", "s.O12jxPL9bbJhS7laeVP2h6fK")
	os.Setenv("IGNITE_VAULT_TYPE", "TOKEN")
	os.Setenv("IGNITE_LOGRUS_CONSOLE_LEVEL", "TRACE")
}

var conn *mgo.Conn

func main() {

	config.Load()
	ilog.New()

	ctx := context.Background()

	var err error

	conn, err = mgo.NewConn(ctx)
	if err != nil {
		log.Error(err)
	}

	mgoManager := mongo.NewManager(conn)
	vault.ManageAll(ctx, mgoManager)

	for {
		mongoQuery()
		time.Sleep(2 * time.Second)
	}
}

func mongoQuery() {

	if conn == nil {
		return
	}

	col := conn.Database.Collection("teste")
	objID, _ := primitive.ObjectIDFromHex("5f19a4416e6274c01d474089")
	result := col.FindOne(context.Background(), bson.M{"_id": objID})
	if result.Err() != nil {
		log.Error(result.Err())
	} else {
		log.Infof("success")
	}

}
