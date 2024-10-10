package models

type Person struct {
	IDPerson   				int    			`gorm:"primaryKey;autoIncrement;column:id_person" json:"id_person"`
	Name       				string 			`gorm:"size:255;not null" json:"name"`
	Birthday 				int64		    `json:"birthday"`
	Email                   string          `gorm:"size:255;not null" json:"email"`
	Insta                   string          `gorm:"size:255;not null" json:"insta"`
	Telephone               string          `gorm:"size:20;not null"  json:"telephone"`
	Info                    string          `gorm:"size:511;not null" json:"info"`
	Photo                   *string         `gorm:"size:2000"         json:"photo"`
}