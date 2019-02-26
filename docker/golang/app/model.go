package main

import (
  "github.com/jinzhu/gorm"
)

type (
  userModel struct {
    gorm.Model
    Email			string `json:"email", db:"email", gorm:"size:255;not null;unique"`
    UserName	string `json:"username", db:"username", gorm:"size:255;not null"`
    Password	string `json:"password", db:"password", gorm:"size:255"`
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
