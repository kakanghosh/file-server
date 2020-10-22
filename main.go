package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//"/home/ghosh/Desktop/static-files"
var (
	staticFilesPath = os.Getenv("STATIC_FILE_PATH")
	userName        = os.Getenv("USER_NAME")
	password        = os.Getenv("PASSWORD")
)

const (
	home          = "/"
	index         = "/index"
	uploadFiles   = "/upload-file"
	downloadFiles = "/static-files/:fileName"
	deleteFiles   = "/static-files/delete"
	assets        = "/assets"
)

const (
	templateFolder = "templates/*"
	assetsFolder   = "assets/"
)

func uploadFile(context *gin.Context) {
	multipartForm, _ := context.MultipartForm()
	files := multipartForm.File["files[]"]
	for _, file := range files {
		err := context.SaveUploadedFile(file, staticFilesPath+"/"+file.Filename)
		if err != nil {
			log.Println(err)
		}
	}
	context.Redirect(http.StatusMovedPermanently, index)
}

func indexPage(context *gin.Context) {
	files, err := ioutil.ReadDir(staticFilesPath)
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
	context.File(staticFilesPath + "/" + fileName)
}

func deleteFile(context *gin.Context) {
	fileName := context.PostForm("file")
	err := os.Remove(staticFilesPath + "/" + fileName)
	if err != nil {
		log.Println("Failed to delete" + err.Error())
	}
	context.Redirect(http.StatusMovedPermanently, index)
}

func redirectToIndexPage(context *gin.Context) {
	context.Redirect(http.StatusMovedPermanently, index)
}

func main() {

	router := gin.Default()

	router.MaxMultipartMemory = 8 << 25

	router.LoadHTMLGlob(templateFolder)
	router.Static(assets, assetsFolder)

	authorized := router.Group("", gin.BasicAuth(gin.Accounts{userName: password}))

	authorized.GET(home, redirectToIndexPage)
	authorized.GET(index, indexPage)
	authorized.POST(uploadFiles, uploadFile)
	authorized.GET(downloadFiles, downloadFile)
	authorized.POST(deleteFiles, deleteFile)

	if err := router.Run(); err != nil {
		log.Println(err)
	}
}
