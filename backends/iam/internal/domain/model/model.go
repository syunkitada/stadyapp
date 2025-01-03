package model

type Project struct {
	ID    string `gorm:"colomn:id;type:varchar(64);not null;primaryKey;"`
	Name  string `gorm:"colomn:name;type:varchar(64);not null;index:idx_name_deleted,unique;"`
	Extra string `gorm:"colomn:extra;type:text;not null;"`
}

type Role struct {
	ID    string `gorm:"colomn:id;type:varchar(64);not null;primaryKey;"`
	Name  string `gorm:"colomn:name;type:varchar(64);not null;index:idx_name_deleted,unique;"`
	Extra string `gorm:"colomn:extra;type:text;not null;"`
}

type UserRoleAssignment struct {
	UserID    string  `gorm:"not null;"`
	Project   Project `gorm:"foreignkey:ProjectID;association_foreignkey:Refer;"`
	ProjectID string  `gorm:"not null;"`
	Role      Role    `gorm:"foreignkey:RoleID;association_foreignkey:Refer;"`
	RoleID    string  `gorm:"not null;"`
}
