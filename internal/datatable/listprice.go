package datetable

type ListPriceItem struct {
	Id            string
	CreatedAt     string
	UpdatedAt     string
	CreateBy      string
	UpdateBy      string
	IsDeleted     bool
	Product       string
	ProdOpt       string
	PriceGeo      string
	Currency      string
	PriceListType string
	Price         float64
	StartEffDate  string
	EndEffDate    string
	Crttimestamp  string
}

func (*ListPriceItem) TableName() string {
	return "listprice"
}
