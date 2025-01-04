package model

type Domain struct {
	ID          string `gorm:"colomn:id;type:varchar(64);not null;primaryKey;"`
	Name        string `gorm:"colomn:name;type:varchar(64);not null;index:idx_name,unique;"`
	Description string `gorm:"colomn:extra;type:text;not null;"`
	Extra       string `gorm:"colomn:extra;type:text;not null;"`
}

// type User struct {
// 	ID          string `gorm:"colomn:id;type:varchar(64);not null;primaryKey;"`
// 	Name        string `gorm:"colomn:name;type:varchar(64);not null;index:idx_name_domain,unique;"`
// 	LastLoginAt time.Time
//
// 	DomainID string `gorm:"colomn:domain_id;type:varchar(64);not null;index:idx_name_domain,unique;"`
// 	Domain   Domain `gorm:"foreignkey:DomainID;association_foreignkey:Refer;"`
// }

type Team struct {
	ID          string `gorm:"colomn:id;type:varchar(64);not null;primaryKey;"`
	Name        string `gorm:"colomn:name;type:varchar(64);not null;index:idx_name_domain,unique;"`
	Description string `gorm:"colomn:extra;type:text;not null;"`
	Extra       string `gorm:"colomn:extra;type:text;not null;"`

	DomainID string `gorm:"colomn:domain_id;type:varchar(64);not null;index:idx_name_domain,unique;"`
	Domain   Domain `gorm:"foreignkey:DomainID;association_foreignkey:Refer;"`
}

type Organization struct {
	ID          string `gorm:"colomn:id;type:varchar(64);not null;primaryKey;"`
	Name        string `gorm:"colomn:name;type:varchar(64);not null;index:idx_name_domain,unique;"`
	Description string `gorm:"colomn:extra;type:text;not null;"`
	Extra       string `gorm:"colomn:extra;type:text;not null;"`

	DomainID string `gorm:"colomn:domain_id;type:varchar(64);not null;index:idx_name_domain,unique;"`
	Domain   Domain `gorm:"foreignkey:DomainID;association_foreignkey:Refer;"`
}

type TagOrganization struct {
	ID string `gorm:"colomn:id;type:varchar(64);not null;primaryKey;"`
}

type Project struct {
	ID          string `gorm:"colomn:id;type:varchar(64);not null;primaryKey;"`
	Name        string `gorm:"colomn:name;type:varchar(64);not null;index:idx_name_domain,unique;"`
	Description string `gorm:"colomn:extra;type:text;not null;"`
	Extra       string `gorm:"colomn:extra;type:text;not null;"`

	OrganizationID string       `gorm:"colomn:organization_id;type:varchar(64);"`
	Organization   Organization `gorm:"foreignkey:OrganizationID;association_foreignkey:Refer;"`

	DomainID string `gorm:"colomn:domain_id;type:varchar(64);not null;index:idx_name_domain,unique;"`
	Domain   Domain `gorm:"foreignkey:DomainID;association_foreignkey:Refer;"`
}

type Role struct {
	ID          string `gorm:"colomn:id;type:varchar(64);not null;primaryKey;"`
	Name        string `gorm:"colomn:name;type:varchar(64);not null;index:idx_name,unique;"`
	Description string `gorm:"colomn:extra;type:text;not null;"`
	Extra       string `gorm:"colomn:extra;type:text;not null;"`
}

type ProjectRoleAssignment struct {
	RoleID string `gorm:"colomn:role_id;type:varchar(64);not null;"`
	Role   Role   `gorm:"foreignkey:RoleID;association_foreignkey:Refer;"`

	UserID string `gorm:"colomn:user_id;type:varchar(64);not null;"`
	TeamID string `gorm:"colomn:team_id;type:varchar(64);not null;"`

	ProjectID string  `gorm:"colomn:project_id;type:varchar(64);not null;"`
	Project   Project `gorm:"foreignkey:ProjectID;association_foreignkey:Refer;"`
}

type DomainRoleAssignment struct {
	RoleID string `gorm:"colomn:role_id;type:varchar(64);not null;"`
	Role   Role   `gorm:"foreignkey:RoleID;association_foreignkey:Refer;"`

	UserID string `gorm:"colomn:user_id;type:varchar(64);not null;"`
	TeamID string `gorm:"colomn:team_id;type:varchar(64);not null;"`

	DomainID string `gorm:"colomn:domain_id;type:varchar(64);not null;"`
	Domain   Domain `gorm:"foreignkey:DomainID;association_foreignkey:Refer;"`
}

type OrganizationRoleAssignment struct {
	RoleID string `gorm:"colomn:role_id;type:varchar(64);not null;"`
	Role   Role   `gorm:"foreignkey:RoleID;association_foreignkey:Refer;"`

	UserID string `gorm:"colomn:user_id;type:varchar(64);not null;"`
	TeamID string `gorm:"colomn:team_id;type:varchar(64);not null;"`

	OrganizationID string       `gorm:"colomn:organization_id;type:varchar(64);not null;"`
	Organization   Organization `gorm:"foreignkey:OrganizationID;association_foreignkey:Refer;"`
}

type TeamRoleAssignment struct {
	RoleID string `gorm:"colomn:role_id;type:varchar(64);not null;"`
	Role   Role   `gorm:"foreignkey:RoleID;association_foreignkey:Refer;"`

	UserID string `gorm:"colomn:user_id;type:varchar(64);not null;"`

	TeamID string `gorm:"colomn:team_id;type:varchar(64);not null;"`
	Team   Team   `gorm:"foreignkey:TeamID;association_foreignkey:Refer;"`
}
