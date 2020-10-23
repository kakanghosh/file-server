package uploader

import (
	"github.com/gin-gonic/gin"
	"github.com/kakanghosh/fileserver/app/constants/appconstants"
	"github.com/kakanghosh/fileserver/app/constants/routes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func UploadFile(context *gin.Context) {
	multipartForm, _ := context.MultipartForm()
	files := multipartForm.File["files[]"]
	for _, file := range files {
		err := context.SaveUploadedFile(file, appconstants.StaticFilesPath+"/"+file.Filename)
		if err != nil {
			log.Println(err)
		}
	}
	context.Redirect(http.StatusMovedPermanently, routes.Index)
}

func IndexPage(context *gin.Context) {
	files, err := ioutil.ReadDir(appconstants.StaticFilesPath)
	if err != nil {
		log.Fatal(err)
	}
	context.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Upload content",
		"files": files,
	})
}

func DownloadFile(context *gin.Context) {
	fileName := context.Param("fileName")
	context.File(appconstants.StaticFilesPath + "/" + fileName)
}

func DeleteFile(context *gin.Context) {
	fileName := context.PostForm("file")
	err := os.Remove(appconstants.StaticFilesPath + "/" + fileName)
	if err != nil {
		log.Println("Failed to delete" + err.Error())
	}
	context.Redirect(http.StatusMovedPermanently, routes.Index)
}

func RedirectToIndexPage(context *gin.Context) {
	context.Redirect(http.StatusMovedPermanently, routes.Index)
}
