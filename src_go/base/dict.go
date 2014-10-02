package base

// 字典数据
type Dict struct {
	BaseObject
	// 类型
	Type string
	// 代码
	Code string
	// 标题
	Title string
}

// 字典版本
type DictVer struct {
	BaseObject
	// 类型
	Type string
	// 版本
	Ver int
}

// 保存字典，修改字典版本
func (u *Dict) BeforeSave() (err error) {
	var dv DictVer
	db.Where("type = ?", u.Type).First(&dv)
	if dv.Id == 0 {
		dv.Type = u.Type
		dv.Ver = 1
	} else {
		dv.Ver++
	}
	db.Save(&dv)
	return
}

// 初始化
func init() {
	db.AutoMigrate(&Dict{}, &DictVer{})
	db.Model(&Dict{}).AddUniqueIndex("idx_dict_code", "type", "code")
	db.Model(&DictVer{}).AddUniqueIndex("idx_dict_ver", "type")
	db.Create(&Dict{
		Type:  "type",
		Code:  "gender",
		Title: "性别",
	})
	db.Create(&Dict{
		Type:  "type",
		Code:  "province",
		Title: "省份",
	})
	db.Create(&Dict{
		Type:  "gender",
		Code:  "M",
		Title: "男",
	})
	db.Create(&Dict{
		Type:  "gender",
		Code:  "F",
		Title: "女",
	})
	db.Create(&Dict{
		Type:  "province",
		Code:  "SD",
		Title: "山东省",
	})
}
