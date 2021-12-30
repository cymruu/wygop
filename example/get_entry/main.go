package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/cymruu/wygop"
	wygops_service "github.com/cymruu/wygop/service"
)

var (
	appkey = flag.String("name", "", "Name of repo to create in authenticated user's GitHub account.")
	secret = flag.String("description", "", "Description of created repo.")
)

func main() {
	flag.Parse()
	ctx := context.Background()

	client := wygop.CreateClient(*appkey, *secret, http.DefaultClient)
	service := wygops_service.CreateWykopService(client)

	entry, err := service.Entries.Entry(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", entry.Body)
}
