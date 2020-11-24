package main

import (
	"gmopg-card/client"
	"gmopg-card/client/card"
	"gmopg-card/config"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func getClient() client.Client {
	shopID := os.Getenv("GMO_PG_SHOP_ID")
	shopPass := os.Getenv("GMO_PG_SHOP_PASS")
	isProdMode := os.Getenv("GMO_PG_PRODUCTION_MODE") != ""
	useDebugOption := os.Getenv("GMO_PG_DEBUG_OPTION") != ""

	conf, err := config.New(shopID, shopPass)
	if err != nil {
		panic(err)
	}

	if isProdMode {
		conf.SetAsProduction()
	}

	cli := client.New(conf)
	cli.Option.Retry = true
	cli.Option.Debug = useDebugOption
	return cli
}

func main() {
	cli := getClient()

	// orderID 001_abcdefg のテスト
	entryAPI := card.EntryTranCard{
		OrderID: "001-abcdefg",
		JobCd:   "AUTH",
		Amount:  2000,
	}

	// 実行
	entryResp, err := entryAPI.Do(cli)
	switch {
	case err != nil:
		panic(err)
	case !entryResp.IsSuccess():
		// error process...
		return
	}
}
