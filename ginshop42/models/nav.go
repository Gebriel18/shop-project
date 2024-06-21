package models

type Nav struct {
	Id         int     `json:"id"`
	Title      string  `json:"title"`
	Link       string  `json:"link"`
	Position   int     `json:"position"`
	IsOpennew  int     `json:"isOpennew"`
	Relation   string  `json:"relation"`
	Sort       int     `json:"sort"`
	Status     int     `json:"status"`
	AddTime    int     `json:"addTime"`
	GoodsItems []Goods `gorm:"-" json:"goodsItems"` // 忽略本字段
}

func (Nav) TableName() string {
	return "nav"
}
