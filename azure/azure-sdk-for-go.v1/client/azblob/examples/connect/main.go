package main

import (
	"context"
	"github.com/americanas-go/ignite"
	ilog "github.com/americanas-go/ignite/americanas-go/log.v1"
	"github.com/americanas-go/ignite/azure/azure-sdk-for-go.v1"
	"github.com/americanas-go/ignite/azure/azure-sdk-for-go.v1/client/azblob"
	"github.com/americanas-go/log"
	"os"
)

func init() {
	os.Setenv("IGNITE_AZURE_AZBLOB_ACCOUNT__NAME", "pricing")
}

func main() {
	ignite.Boot()
	ilog.New()

	ctx := context.Background()

	credential, err := azure.NewCredential(ctx)
	if err != nil {
		panic(err)
	}

	client, err := azblob.NewClient(ctx, credential)
	if err != nil {
		panic(err)
	}

	log.Info("connected")

	container, err := client.CreateContainer(ctx, "test", nil)
	if err != nil {
		panic(err)
	}

	log.Info(container.LastModified)

}
