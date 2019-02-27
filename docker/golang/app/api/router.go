package api

import (
  "os"
  "fmt"

  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"

  "github.com/arasan01/app/config"
  "github.com/arasan01/app/model"
)

var db *gorm.DB

func Router_init() {
  config.Env_load()
  //open a db connection
  var err error
  db, err = gorm.Open("mysql",
                      fmt.Sprintf(
                        "%s:%s@tcp(mysql_host:3306)/%s?parseTime=true",
                        os.Getenv("MYSQL_USER"),
                        os.Getenv("MYSQL_PASSWORD"),
                        os.Getenv("MYSQL_DATABASE"),
                      ),
  )
  if err != nil {
    panic("failed to connect database")
  }

  db.LogMode(true)

  //Migrate the schema
  db.AutoMigrate(&model.UserModel{})
  // db.AutoMigrate(&model.TodoModel{})
}

func Router() {

  router := gin.Default()
  router.Use(cors.Default())

  router.GET("/", versionGET)

  account := router.Group("/api/v1/account")
  {
    account.POST("/token", tokenGET)
    account.POST("/register", registerUser)
  }

  // v1 := router.Group("/api/v1/todos")
  // {
  //   v1.POST("/", createTodo)
  //   v1.GET("/", fetchAllTodo)
  //   v1.GET("/:id", fetchSingleTodo)
  //   v1.PUT("/:id", updateTodo)
  //   v1.DELETE("/:id", deleteTodo)
  // }
  router.Run(":8080")
  defer db.Close()
}

