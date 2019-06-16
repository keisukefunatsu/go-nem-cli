package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type AccountData struct {
	Account AccountInfo
}

type AccountInfo struct {
	Address string
	Balance int64
}

func getAccount(address string) (AccountData, error) {
	url := "http://go.nem.ninja:7890/account/get?address=" + address
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else if res.StatusCode != 200 {
		log.Fatal("Unable to get this url : http status ", res.StatusCode)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var accountData AccountData
	if err := json.Unmarshal(body, &accountData); err != nil {
		log.Fatal(err)
	}
	return accountData, nil
}

func main() {
	app := cli.NewApp()
	app.Name = "nem_cli"
	app.Usage = "nem cli tools"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "address, a",
			Value: "",
			Usage: "nem address",
		},
	}
	app.Action = func(c *cli.Context) error {
		var address = c.String("address")
		account, err := getAccount(address)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		fmt.Println(account)
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
