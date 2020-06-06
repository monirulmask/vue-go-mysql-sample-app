package main

import (

    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    controller "./src_back/controllers/controller"
)

func main() {
    serve()
}

func serve() {
    router := gin.Default()
    router.Static("/src_front", "./src_front")
    router.StaticFS("/schoolapp", http.Dir("./src_front/pages/teacher"))
    router.GET("/fetchAllTeachers", controller.FetchAllTeachers)
    router.GET("/fetchTeacher", controller.FindTeacher)
    router.POST("/addTeacher", controller.AddTeacher)
    router.POST("/changeStateTeacher", controller.ChangeStateTeacher)
    router.POST("/deleteTeacher", controller.DeleteTeacher)
    
    if err := router.Run(":8080"); err != nil {
        log.Fatal("Server Run Failed.: ", err)
    }
}
