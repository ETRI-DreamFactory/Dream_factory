package models

import "time"

type Developer struct {
	Developer_id string `gorm:"primary_key;AUTO_INCREMENT" json:"d_id"`
	Member       Member `gorm:"foreignkey:MemberRefer"`
	MemberRefer  uint
	Bio          string   `gorm:"type:text" json:"d_bio"`
	Past_project Projects `gorm:"foreignKey:Project_id"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Developers []Developer
