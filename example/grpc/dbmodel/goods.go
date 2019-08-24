package dbmodel

type Goods struct {
	ID         uint64  `gorm:"NOT NULL;column:id;PRIMARY_KEY;bigint;" json:"ID,omitempty"`
	Name       string  `gorm:"NOT NULL;column:name;type:varchar(500)" json:"Name,omitempty"`
	Bn         string  `gorm:"NOT NULL;column:bn;type:varchar(200)" json:"Bn,omitempty"`
	Price      float64 `gorm:"NOT NULL;column:price;type:double" json:"Price,omitempty"`
	Pic        string  `gorm:"NOT NULL;column:pic;type:varchar(500)" json:"Pic,omitempty"`
	Content    string  `gorm:"NOT NULL;column:content;type:mediumtext" json:"Content,omitempty"`
	CreateTime uint64  `gorm:"NOT NULL;column:create_time;bigint" json:"CreateTime,omitempty"`
	UpdateTime uint64  `gorm:"NOT NULL;column:update_time;bigint" json:"UpdateTime,omitempty"`
}
