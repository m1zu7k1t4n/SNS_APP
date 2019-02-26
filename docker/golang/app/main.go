package main

import (
  "os"
  "fmt"
  "log"
  "golang.org/x/crypto/bcrypt"

  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"

  "github.com/arasan01/app/env"
)

var db *gorm.DB

func init() {
  env.Env_load()
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
  db.AutoMigrate(&userModel{})
  db.AutoMigrate(&todoModel{})
}

func main() {

  router := gin.Default()
  router.Use(cors.Default())

  router.GET("/", versionGET)

  v1 := router.Group("/api/v1/todos")
  {
    v1.POST("/", createTodo)
    v1.GET("/", fetchAllTodo)
    v1.GET("/:id", fetchSingleTodo)
    v1.PUT("/:id", updateTodo)
    v1.DELETE("/:id", deleteTodo)
  }
  router.Run(":8080")
}

func registerUser(c *gin.Context) {
  hash, err := bcrypt.GenerateFromPassword([]byte(c.PostForm("password")), bcrypt.DefaultCost)
  if err != nil {
    log.Fatal("Error Generating Password")
  }

  user := userModel{UserName: c.PostForm("username"), Password: string(hash)}
  db.Create(&user)
}

