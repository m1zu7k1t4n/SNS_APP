package api

import (
  "net/http"
  "log"

  "github.com/gin-gonic/gin"
  "golang.org/x/crypto/bcrypt"
  "github.com/arasan01/app/model"

  "github.com/arasan01/app/tool"
)


func registerUser(c *gin.Context) {
  var check model.UserModel
  email := tool.SaniAtSign(c.PostForm("email"))
  db.Where("email = ?", email).First(&check)
  if check.Email == email {
    c.JSON(http.StatusConflict, gin.H{"status": http.StatusConflict, "message": "Already Registered."})
    return
  }

  hash, err := bcrypt.GenerateFromPassword([]byte(c.PostForm("password")), bcrypt.DefaultCost)
  if err != nil {
    log.Fatal("Error Generating Password")
  }

  token, err := bcrypt.GenerateFromPassword([]byte(email+c.PostForm("username")+string(hash)), bcrypt.DefaultCost)
  if err != nil {
    log.Fatal("Error Generating Password")
  }

  user := model.UserModel{Email: email,
                          UserName: c.PostForm("username"),
                          Password: string(hash),
                          Token: string(token)}
  db.Create(&user)
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "username": c.PostForm("username"), "email": c.PostForm("email")})
}

func tokenGET(c *gin.Context) {
  var user model.UserModel
  email := tool.SaniAtSign(c.PostForm("email"))
  db.Where("email = ?", email).First(&user)
  if user.Email == "" {
    c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
    return
  }

  err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(c.PostForm("password")))
  if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Not Match Password!"})
        log.Println("Not match Password")
        return
  }

  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "token": user.Token})
}
