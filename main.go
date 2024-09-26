package main

import (
	"context"
	"fmt"
	"price-subscriber/config"
	"price-subscriber/internal/dao"
	datetable "price-subscriber/internal/datatable"
)

func main() {
	config := config.InitConfig()
	fmt.Print(config)
	result, err := dao.NewListPriceDao().GetListPrice(context.Background(), "0010-925")
	if err != nil {
		fmt.Println(err)
	}
	listprice := datetable.ListPriceItem{
		Id:            "3",
		CreatedAt:     "2024-09-19 16:52:50",
		UpdatedAt:     "2024-09-19 16:52:50",
		CreateBy:      "huan",
		UpdateBy:      "jiang",
		Product:       "test",
		ProdOpt:       "",
		PriceGeo:      "US",
		Currency:      "USD",
		PriceListType: "CIFE",
		Price:         52.36,
		StartEffDate:  "08-12-2015",
		EndEffDate:    "08-12-2018",
		Crttimestamp:  "20240808-150405",
	}

	bool, err := dao.NewListPriceDao().CreateListPrice(context.Background(), listprice)
	if err != nil && bool {
		fmt.Println(err)
	}
	fmt.Println(result)
}
