package main

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title    string `gorm:"size:255;not null;unique" json:"title"`
	Author   string `gorm:"size:255;not null;unique" json:"author"`
	Quantity int16  `gorm:"size:255;not null;unique" json:"quantity"`
}
