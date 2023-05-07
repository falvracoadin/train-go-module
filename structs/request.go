package structs

type Filter struct {
	ID     int64  `json:"id" gorm:"column:id"`
	Filter string `param:"filter" query:"filter" form:"filter" json:"filter" xml:"filter"`
	Offset int32  `param:"offset" query:"offset" form:"offset" json:"offset" xml:"offset"`
	Limit  int32  `param:"limit" query:"limit" form:"limit" json:"limit" xml:"limit"`
}
