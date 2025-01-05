package model

import "time"

type Domain struct {
	ID          string `gorm:"column:id;type:varchar(64);not null;primaryKey;"`
	Name        string `gorm:"column:name;type:varchar(64);not null;index:idx_name,unique;"`
	Description string `gorm:"column:description;type:text;not null;"`
	Extra       string `gorm:"column:extra;type:text;not null;"`
}

type User struct {
	ID          string    `gorm:"column:id;type:varchar(64);not null;primaryKey;"`
	Name        string    `gorm:"column:name;type:varchar(64);not null;index:idx_name_domain,unique;"`
	LastLoginAt time.Time `gorm:"column:last_login_at;not null;"`

	DomainID string `gorm:"column:domain_id;type:varchar(64);not null;index:idx_name_domain,unique;"`
	Domain   Domain `gorm:"foreignkey:DomainID;association_foreignkey:Refer;"`
}

type Team struct {
	ID          string `gorm:"column:id;type:varchar(64);not null;primaryKey;"`
	Name        string `gorm:"column:name;type:varchar(64);not null;index:idx_name_domain,unique;"`
	Description string `gorm:"column:description;type:text;not null;"`
	Extra       string `gorm:"column:extra;type:text;not null;"`

	DomainID string `gorm:"column:domain_id;type:varchar(64);not null;index:idx_name_domain,unique;"`
	Domain   Domain `gorm:"foreignkey:DomainID;association_foreignkey:Refer;"`
}

type Organization struct {
	ID          string `gorm:"column:id;type:varchar(64);not null;primaryKey;"`
	Name        string `gorm:"column:name;type:varchar(64);not null;index:idx_name_domain,unique;"`
	Description string `gorm:"column:description;type:text;not null;"`
	Extra       string `gorm:"column:extra;type:text;not null;"`

	DomainID string `gorm:"column:domain_id;type:varchar(64);not null;index:idx_name_domain,unique;"`
	Domain   Domain `gorm:"foreignkey:DomainID;association_foreignkey:Refer;"`
}

type TagOrganization struct {
	ID string `gorm:"column:id;type:varchar(64);not null;primaryKey;"`
}

type Project struct {
	ID          string `gorm:"column:id;type:varchar(64);not null;primaryKey;"`
	Name        string `gorm:"column:name;type:varchar(64);not null;index:idx_name_domain,unique;"`
	Description string `gorm:"column:description;type:text;not null;"`
	Extra       string `gorm:"column:extra;type:text;not null;"`

	OrganizationID string       `gorm:"column:organization_id;type:varchar(64);"`
	Organization   Organization `gorm:"foreignkey:OrganizationID;association_foreignkey:Refer;"`

	DomainID string `gorm:"column:domain_id;type:varchar(64);not null;index:idx_name_domain,unique;"`
	Domain   Domain `gorm:"foreignkey:DomainID;association_foreignkey:Refer;"`
}

type Role struct {
	ID          string `gorm:"column:id;type:varchar(64);not null;primaryKey;"`
	Name        string `gorm:"column:name;type:varchar(64);not null;index:idx_name,unique;"`
	Description string `gorm:"column:description;type:text;not null;"`
	Extra       string `gorm:"column:extra;type:text;not null;"`
}

type ProjectRoleAssignment struct {
	RoleID string `gorm:"column:role_id;type:varchar(64);not null;uniqueIndex:idx_project_role;"`
	Role   Role   `gorm:"foreignkey:RoleID;association_foreignkey:Refer;"`

	UserID *string `gorm:"column:user_id;type:varchar(64);uniqueIndex:idx_project_role;"`
	User   User    `gorm:"foreignkey:UserID;association_foreignkey:Refer;"`

	TeamID *string `gorm:"column:team_id;type:varchar(64);uniqueIndex:idx_project_role;"`
	Team   Team    `gorm:"foreignkey:TeamID;association_foreignkey:Refer;"`

	ProjectID string  `gorm:"column:project_id;type:varchar(64);not null;uniqueIndex:idx_project_role;"`
	Project   Project `gorm:"foreignkey:ProjectID;association_foreignkey:Refer;"`
}

type ProjectRoleAssignmentDetail struct {
	RoleID   string `gorm:"column:role_id;"`
	RoleName string `gorm:"column:role_name;"`

	UserID   *string `gorm:"column:user_id;"`
	UserName string  `gorm:"column:user_name;"`

	TeamID   *string `gorm:"column:team_id;"`
	TeamName string  `gorm:"column:team_name;"`

	ProjectID   string `gorm:"column:project_id;"`
	ProjectName string `gorm:"column:project_name;"`

	DomainID   string `gorm:"column:domain_id;"`
	DomainName string `gorm:"column:domain_name;"`
}

type DomainRoleAssignment struct {
	RoleID string `gorm:"column:role_id;type:varchar(64);not null;uniqueIndex:idx_domain_role;"`
	Role   Role   `gorm:"foreignkey:RoleID;association_foreignkey:Refer;"`

	UserID *string `gorm:"column:user_id;type:varchar(64);uniqueIndex:idx_domain_role;"`
	User   User    `gorm:"foreignkey:UserID;association_foreignkey:Refer;"`

	TeamID *string `gorm:"column:team_id;type:varchar(64);uniqueIndex:idx_domain_role;"`
	Team   Team    `gorm:"foreignkey:TeamID;association_foreignkey:Refer;"`

	DomainID string `gorm:"column:domain_id;type:varchar(64);not null;uniqueIndex:idx_domain_role;"`
	Domain   Domain `gorm:"foreignkey:DomainID;association_foreignkey:Refer;"`
}

type DomainRoleAssignmentDetail struct {
	RoleID   string `gorm:"column:role_id;"`
	RoleName string `gorm:"column:role_name;"`

	UserID   *string `gorm:"column:user_id;"`
	UserName string  `gorm:"column:user_name;"`

	TeamID   *string `gorm:"column:team_id;"`
	TeamName string  `gorm:"column:team_name;"`

	DomainID   string `gorm:"column:domain_id;"`
	DomainName string `gorm:"column:domain_name;"`
}

type OrganizationRoleAssignment struct {
	RoleID string `gorm:"column:role_id;type:varchar(64);not null;uniqueIndex:idx_organization_role;"`
	Role   Role   `gorm:"foreignkey:RoleID;association_foreignkey:Refer;"`

	UserID *string `gorm:"column:user_id;type:varchar(64);uniqueIndex:idx_organization_role;"`
	User   User    `gorm:"foreignkey:UserID;association_foreignkey:Refer;"`

	TeamID *string `gorm:"column:team_id;type:varchar(64);uniqueIndex:idx_organization_role;"`
	Team   Team    `gorm:"foreignkey:TeamID;association_foreignkey:Refer;"`

	OrganizationID string       `gorm:"column:organization_id;type:varchar(64);not null;uniqueIndex:idx_organization_role;"`
	Organization   Organization `gorm:"foreignkey:OrganizationID;association_foreignkey:Refer;"`
}

type OrganizationRoleAssignmentDetail struct {
	RoleID   string `gorm:"column:role_id;"`
	RoleName string `gorm:"column:role_name;"`

	UserID   *string `gorm:"column:user_id;"`
	UserName string  `gorm:"column:user_name;"`

	TeamID   *string `gorm:"column:team_id;"`
	TeamName string  `gorm:"column:team_name;"`

	OrganizationID   string `gorm:"column:organization_id;"`
	OrganizationName string `gorm:"column:organization_name;"`

	DomainID   string `gorm:"column:domain_id;"`
	DomainName string `gorm:"column:domain_name;"`
}

type TeamRoleAssignment struct {
	RoleID string `gorm:"column:role_id;type:varchar(64);not null;uniqueIndex:idx_team_role;"`
	Role   Role   `gorm:"foreignkey:RoleID;association_foreignkey:Refer;"`

	UserID string `gorm:"column:user_id;type:varchar(64);not null;uniqueIndex:idx_team_role;"`
	User   User   `gorm:"foreignkey:UserID;association_foreignkey:Refer;"`

	TeamID string `gorm:"column:team_id;type:varchar(64);not null;uniqueIndex:idx_team_role;"`
	Team   Team   `gorm:"foreignkey:TeamID;association_foreignkey:Refer;"`
}

type TeamRoleAssignmentDetail struct {
	RoleID   string `gorm:"column:role_id;"`
	RoleName string `gorm:"column:role_name;"`

	UserID   string `gorm:"column:user_id;"`
	UserName string `gorm:"column:user_name;"`

	TeamID   string `gorm:"column:team_id;"`
	TeamName string `gorm:"column:team_name;"`

	DomainID   string `gorm:"column:domain_id;"`
	DomainName string `gorm:"column:domain_name;"`
}
