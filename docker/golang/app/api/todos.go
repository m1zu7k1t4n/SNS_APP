package api

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/arasan01/app/model"
)

// createTodo add a new todo
func createTodo(c *gin.Context) {
  todo := model.TodoModel{Title: c.PostForm("title"), Body: c.PostForm("body")}
  db.Create(&todo)
  c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": todo.ID})
}

// fetchAllTodo fetch all todos
func fetchAllTodo(c *gin.Context) {
  var todos []model.TodoModel
  var _todos []model.TransformedTodo

  db.Find(&todos)
  if len(todos) <= 0 {
    c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
    return
  }

  //transforms the todos for building a good response
  for _, item := range todos {
    _todos = append(_todos, model.TransformedTodo{ID: item.ID, Title: item.Title, Body: item.Body})
  }
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})
}

// fetchSingleTodo fetch a single todo
func fetchSingleTodo(c *gin.Context) {
  var todo model.TodoModel
  todoID := c.Param("id")
  db.First(&todo, todoID)
  if todo.ID == 0 {
    c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
    return
  }

  _todo := model.TransformedTodo{ID: todo.ID, Title: todo.Title, Body: todo.Body}
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todo})
}

// updateTodo update a todo
func updateTodo(c *gin.Context) {
  var todo model.TodoModel
  todoID := c.Param("id")

  db.First(&todo, todoID)

  if todo.ID == 0 {
    c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
    return
  }

  db.Model(&todo).Update("title", c.PostForm("title"))
  db.Model(&todo).Update("body", c.PostForm("body"))
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
}

// deleteTodo remove a todo
func deleteTodo(c *gin.Context) {
  var todo model.TodoModel
  todoID := c.Param("id")

  db.First(&todo, todoID)

  if todo.ID == 0 {
    c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
    return
  }

  db.Delete(&todo)
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
}
