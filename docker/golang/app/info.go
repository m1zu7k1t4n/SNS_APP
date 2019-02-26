package main

import (
  "net/http"
  "os"
  "github.com/gin-gonic/gin"
)

func versionGET(c *gin.Context) {
  version := os.Getenv("VERSION")
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "version": version})
}
