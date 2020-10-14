package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"io/ioutil"
	"os"
)

const STATIC_FILES_PATH = "/home/ghosh/Desktop/static-files"

func uploadFile(context *gin.Context) {
	file, _ := context.FormFile("file")
	log.Println(file.Filename)
	context.SaveUploadedFile(file, STATIC_FILES_PATH + "/" + file.Filename)
	context.Redirect(http.StatusMovedPermanently, "/index")
}

func indexPage(context *gin.Context) {
	files, err := ioutil.ReadDir(STATIC_FILES_PATH)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
        log.Println(file.ModTime())
    }
	context.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Upload content",
		"files": files,
	})
}

func downloadFile(context *gin.Context) {
	fileName := context.Param("fileName")
	context.File(STATIC_FILES_PATH+"/"+fileName)
}

func deleteFile(context *gin.Context) {
	fileName := context.Param("fileName")
	err := os.Remove(STATIC_FILES_PATH+"/"+fileName)
	if err != nil {
		log.Println("Failed to delete"+err.Error())
	}
	context.Redirect(http.StatusMovedPermanently, "/index")
}

func main() {
	
	router := gin.Default()

	router.MaxMultipartMemory = 8 << 25

	router.LoadHTMLGlob("templates/*")

	router.GET("/index", indexPage)
	router.POST("/upload-file", uploadFile)
	router.GET("/static-files/:fileName", downloadFile)
	router.GET("/static-files/:fileName/delete", deleteFile)

	router.Run()
}