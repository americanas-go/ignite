package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/db/nosql/go.mongodb.org/mongo-driver.v1"
	ilog "github.com/americanas-go/ignite/log/americanas-go/log.v1"
	"github.com/americanas-go/log"
	"github.com/hashicorp/vault/api"
	vault "github.com/mittwald/vaultgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var conn *mongo.Conn
var username, password string

func init() {
	os.Setenv("IGNITE_MONGO_URI", "mongodb://localhost:27004/database")
}

func main() {

	config.Load()
	ilog.New()

	mongoConn()

	go renew(mongoConn)

	for {

		mongoQuery(conn)

		time.Sleep(2 * time.Second)
	}
}

func mongoQuery(conn *mongo.Conn) {

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

func mongoConn() {
	var err error
	options, _ := mongo.NewOptions()
	options.Auth.Username = username
	options.Auth.Password = password
	conn, err = mongo.NewConnWithOptions(context.Background(), options)
	if err != nil {
		log.Error(err)
	}
}

func renew(f func()) {

	client, err := vault.NewClient(
		"http://vault.example.com",
		vault.WithCaPath(""),
		vault.WithAuthToken("s.O12jxPL9bbJhS7laeVP2h6fK"),
	)

	if err != nil {
		panic(err)
	}

	var response api.Secret
	err = client.Read([]string{"v1/database/creds/recency_vault_hml"}, &response, &vault.RequestOptions{
		Parameters:  nil,
		SkipRenewal: false,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("lease_id: ", response.LeaseID)
	fmt.Println("data: ", response.Data)
	fmt.Println("lease_duration: ", response.LeaseDuration)

	username = response.Data["username"].(string)
	password = response.Data["password"].(string)

	f()

	secretesTokenWatcher := api.LifetimeWatcherInput{
		Secret:    &response,
		Increment: 120,
	}
	watcher, err := client.NewLifetimeWatcher(&secretesTokenWatcher)
	if err != nil {
		panic(err)
	}
	go watcher.Start()

	for {
		select {
		case rawData := <-watcher.RenewCh():
			fmt.Printf("received renewal at: %+v \n", rawData.RenewedAt)
			fmt.Printf("received renewal Secret: %+v \n", rawData.Secret)
		case er := <-watcher.DoneCh():
			fmt.Println("Got watcher error: ", er)
			watcher.Stop()
			go renew(f)
			return
		}
	}
}
