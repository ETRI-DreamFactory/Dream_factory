package models

import "time"

type Member struct {
	Member_id string `gorm:"primary_key"`
	Id        string `gorm:"type:varchar(255);NOT NULL;UNIQUE" json:"m_id" binding:"required"`
	Passwd    string `gorm:"type:varchar(255);NOT NULL" json:"m_passwd" binding:"required"`
	Email     string `gorm:"type:varchar(255);NOT NULL;UNIQUE" json:"m_email"`
	Name      string `gorm:"type:varchar(255);NOT NULL;" json:"m_name"`
	Phone     string `gorm:"type:varchar(100);NOT NULL;UNIQUE;UNIQUE_INDEX" json:"m_phone" binding:"required"`
	Nickname  string `gorm:"type:varchar(255);NOT NULL;UNIQUE" json:"m_nickname"`
	Wallet    string `gorm:"type:varchar(255)" json:"m_wallet"`
	Birth     string `gorm:"type:varchar(100)" json:"m_birth"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Members []Member
