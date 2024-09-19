package datetable

type ListPriceItem struct {
	ID               string
	CreatedAt        string
	UpdatedAt        string
	CreatedBy        string
	UpdatedBy        string
	IsDeleted        bool
	Product          string
	ProdOpt          string
	PriceGeo         string
	Currency         string
	PriceListType    string
	Price            float64
	StartEffDate     string
	EndEffDate       string
	CurrentTimeStamp string
}

func (*ListPriceItem) TableName() string {
	return "listprice"
}
