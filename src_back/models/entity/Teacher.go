package entity

type Teacher struct {
    ID      int    `gorm:"primary_key;not null"       json:"id"`
    Name    string `gorm:"type:varchar(200);not null" json:"name"`
    Salary  int    `gorm:"not null"                   json:"salary"`
    State   int    `gorm:"not null"                   json:"state"`
}