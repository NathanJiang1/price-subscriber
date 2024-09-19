package main

import (
	"context"
	"fmt"
	"price-subscriber/config"
	"price-subscriber/internal/dao"
)

func main() {
	config :=config.InitConfig()
	fmt.Print(config)
	result, err :=dao.NewListPriceDao().GetListPrice(context.Background(), "0010-925")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
