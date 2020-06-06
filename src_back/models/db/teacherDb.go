package db

import (
    "fmt"
    "github.com/jinzhu/gorm"
    entity "../entity"
)

func open() *gorm.DB {
    DBMS := "mysql"
    USER := "root"
    PASS := "root"
    PROTOCOL := "tcp(localhost:3306)"
    DBNAME := "schoolmanagement"
    CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
    db, err := gorm.Open(DBMS, CONNECT)

    if err != nil {
        panic(err.Error())
    }

    db.Set("gorm:table_options", "ENGINE=InnoDB")
    db.LogMode(true)
    db.SingularTable(true)
    db.AutoMigrate(&entity.Teacher{})
    fmt.Println("db connected: ", &db)
    return db
}

func FindAllTeachers() []entity.Teacher {
    teachers := []entity.Teacher{}

    db := open()
    // select
    db.Order("ID asc").Find(&teachers)
    defer db.Close()
    return teachers
}

func FindTeacher(teacherID int) []entity.Teacher {
    teacher := []entity.Teacher{}

    db := open()
    // select
    db.First(&teacher, teacherID)
    defer db.Close()

    return teacher
}

func InsertTeacher(registerTeacher *entity.Teacher) {
    db := open()
    // insert
    db.Create(&registerTeacher)
    defer db.Close()
}

func UpdateStateTeacher(teacherID int, teacherState int) {
    teacher := []entity.Teacher{}

    db := open()
    // update
    db.Model(&teacher).Where("ID = ?", teacherID).Update("State", teacherState)
    defer db.Close()
}

func DeleteTeacher(teacherID int) {
    teacher := []entity.Teacher{}

    db := open()
    // delete
    db.Delete(&teacher, teacherID)
    defer db.Close()
}