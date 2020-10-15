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
	multipartForm, _ := context.MultipartForm()
	files := multipartForm.File["files[]"]
	for _, file := range files {
		context.SaveUploadedFile(file, STATIC_FILES_PATH + "/" + file.Filename)
	}
	context.Redirect(http.StatusMovedPermanently, "/index")
}

func indexPage(context *gin.Context) {
	files, err := ioutil.ReadDir(STATIC_FILES_PATH)
	if err != nil {
		log.Fatal(err)
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
	fileName := context.PostForm("file")
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
	router.Static("/assets", "./assets")

	router.GET("/", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/index")
	})
	router.GET("/index", indexPage)
	router.POST("/upload-file", uploadFile)
	router.GET("/static-files/:fileName", downloadFile)
	router.POST("/static-files/delete", deleteFile)

	router.Run()
}