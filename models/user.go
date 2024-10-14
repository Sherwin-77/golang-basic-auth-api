package models

type User struct {
	BaseModel
	Username string `json:"username" gorm:"type:varchar(255);not null"`
	Email    string `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Password string `json:"-" gorm:"type:text;not null"`

	Roles []*Role `json:"roles,omitempty" gorm:"many2many:user_roles;"`
	Todos []*Todo `json:"todos,omitempty" gorm:"foreignKey:UserID"`
}
