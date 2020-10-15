package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"io/ioutil"
	"os"
)

//"/home/ghosh/Desktop/static-files"
var STATIC_FILES_PATH 	= os.Getenv("STATIC_FILE_PATH")
var USER_NAME			= os.Getenv("USER_NAME")
var PASSWORD			= os.Getenv("PASSWORD")

const HOME 			 	= "/"
const INDEX			 	= "/index"
const UPLOAD_FILES   	= "/upload-file"
const DOWNLOAD_FILES 	= "/static-files/:fileName"
const DELETE_FILES   	= "/static-files/delete"

func uploadFile(context *gin.Context) {
	multipartForm, _ := context.MultipartForm()
	files := multipartForm.File["files[]"]
	for _, file := range files {
		context.SaveUploadedFile(file, STATIC_FILES_PATH + "/" + file.Filename)
	}
	context.Redirect(http.StatusMovedPermanently, INDEX)
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
	context.Redirect(http.StatusMovedPermanently, INDEX)
}

func redirectToIndexPage(context *gin.Context) {
	context.Redirect(http.StatusMovedPermanently, INDEX)
}

func main() {
	
	router := gin.Default()

	router.MaxMultipartMemory = 8 << 25

	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")

	authorized := router.Group("", gin.BasicAuth(gin.Accounts{USER_NAME: PASSWORD}))

	authorized.GET(HOME, redirectToIndexPage)
	authorized.GET(INDEX, indexPage)
	authorized.POST(UPLOAD_FILES, uploadFile)
	authorized.GET(DOWNLOAD_FILES, downloadFile)
	authorized.POST(DELETE_FILES, deleteFile)

	router.Run()
}