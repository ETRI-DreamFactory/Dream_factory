package models

import "time"

type Project struct {
	Project_id   uint       `gorm:"primary_key;AUTO_INCREMENT"`
	Team_wallet  string     `gorm:"type:varchar(255);NOT NULL" json:"p_wallet" binding:"required"`
	Project_name string     `gorm:"type:varchar(255);NOT NULL;UNIQUE" json:"p_name" binding:"required"`
	Proposer     Developer  `gorm:"foreignkey:Developer_id"`
	Developers   Developers `gorm:"foreignkey:Developer_id"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Projects []Project
