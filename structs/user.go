package structs

type User struct {
	File         string `param:"file" query:"file" form:"file" json:"file" xml:"file" gorm:"-"`
	Name		string		`param:"name" query:"name" form:"name" json:"name" gorm:"column:name" xml:"name"`
	Email		string		`param:"email" query:"query" form:"email" json:"email" gorm:"column:email" xml:"email"`
	Path		string		`param:"path" query:"path" form:"path" json:"path" gorm:"column:path" xml:"path"`
	Foto		string		`param:"foto" query:"foto" form:"foto" json:"foto" gorm:"column:foto" xml:"foto"`
	Url			string		`param:"url" query:"url" form:"url" json:"url" gorm:"column:url" xml:"url"`
	Username	string		`param:"username" query:"username" form:"username" json:"username" gorm:"column:username" xml:"username"`
	BirthDate	string		`param:"birth_date" query:"birth_date" form:"birth_date" json:"birth_date" gorm:"column:birth_date" xml:"birth_date"`
	IsDeleted	int32		`param:"is_deleted" query:"is_deleted" form:"is_deleted" json:"is_deleted" gorm:"column:is_deleted" xml:"is_deleted"`
	IsActive	int32		`param:"is_active" query:"is_active" form:"is_active" json:"is_active" gorm:"column:is_active" xml:"is_active"`
	Password	string		`param:"password" query:"password" form:"password" json:"password" gorm:"column:password" xml:"password"`
	Address		string		`param:"address" query:"address" form:"address" json:"address" gorm:"column:address" xml:"address"`
}