package models

type User struct {
	SiteID                uint   `gorm:"primaryKey"`
	FixletID              uint   `gorm:"primaryKey"`
	Name                  string `gorm:"size:255;not null"`
	Criticality           string `gorm:"size:50;not null"`
	RelevantComputerCount int    `gorm:"not null"`
}
