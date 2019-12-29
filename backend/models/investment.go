package models

import "time"

type Investment struct {
	Investment_id uint    `gorm:"primary_key;AUTO_INCREMENT"`
	Investor      Member  `gorm:"foreignkey:Member_id"`
	Value         uint    `gorm:"NOT NULL" json:"i_value"`
	Project       Project `gorm:"foreignkey:Project_id"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Investments []Investment
