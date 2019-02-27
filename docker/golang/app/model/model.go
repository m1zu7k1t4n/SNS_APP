package model

import (
  "github.com/jinzhu/gorm"
  "time"
)

type (
  UserModel struct {
    Email			string `json:"email" gorm:"unique;not null"`
    UserName	string `json:"username" gorm:"not null"`
    Password	string `json:"password" gorm:"not null"`
    Token     string `json:"token" gorm:"unique;not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time
  }

  // todoModel describes a todoModel type
  TodoModel struct {
    gorm.Model
    Title     string `json:"title"`
    Body 			string `json:"body"`
  }

  // transformedTodo represents a formatted todo
  TransformedTodo struct {
    ID        uint   `json:"id"`
    Title     string `json:"title"`
    Body 			string `json:"body"`
  }
)
