package main

import (
  "github.com/jinzhu/gorm"
)

type (
  userModel struct {
    Email			string `json:"email", gorm:"unique;not null"`
    UserName	string `json:"username", gorm:"not null"`
    Password	string `json:"password", gorm:"not null"`
  }

  // todoModel describes a todoModel type
  todoModel struct {
    gorm.Model
    Title     string `json:"title"`
    Body 			string `json:"body"`
  }

  // transformedTodo represents a formatted todo
  transformedTodo struct {
    ID        uint   `json:"id"`
    Title     string `json:"title"`
    Body 			string `json:"body"`
  }
)
