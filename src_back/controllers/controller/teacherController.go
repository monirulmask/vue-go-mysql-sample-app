package controller

import (
    strconv "strconv"
    "github.com/gin-gonic/gin"
    entity "../../models/entity"
    db "../../models/db"
)

const (
    InActive = 0
    Active = 1
)

func FetchAllTeachers(c *gin.Context) {
    resultTeachers := db.FindAllTeachers()
    c.JSON(200, resultTeachers)
}

func FindTeacher(c *gin.Context) {
    teacherIDStr := c.Query("teacherID")
    teacherID, _ := strconv.Atoi(teacherIDStr)
    resultTeacher := db.FindTeacher(teacherID)
    c.JSON(200, resultTeacher)
}

func AddTeacher(c *gin.Context) {
    teacherName := c.PostForm("teacherName")
    teacherSalary, _ := strconv.Atoi(c.PostForm("teacherSalary"))

    var teacher = entity.Teacher{
        Name:  teacherName,
        Salary:  teacherSalary,
        State: Active,
    }

    db.InsertTeacher(&teacher)
}

func ChangeStateTeacher(c *gin.Context) {
    reqTeacherID := c.PostForm("teacherID")
    reqTeacherState := c.PostForm("teacherState")

    teacherID, _ := strconv.Atoi(reqTeacherID)
    teacherState, _ := strconv.Atoi(reqTeacherState)
    changeState := InActive

    if teacherState == InActive {
        changeState = Active
    } else {
        changeState = InActive
    }

    db.UpdateStateTeacher(teacherID, changeState)
}

func DeleteTeacher(c *gin.Context) {
    teacherIDStr := c.PostForm("teacherID")
    teacherID, _ := strconv.Atoi(teacherIDStr)
    db.DeleteTeacher(teacherID)
}