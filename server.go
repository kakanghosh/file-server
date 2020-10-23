package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kakanghosh/fileserver/app/constants/appconstants"
	"github.com/kakanghosh/fileserver/app/constants/routes"
	"github.com/kakanghosh/fileserver/app/controller/uploader"
	"log"
)

func main() {

	router := gin.Default()

	router.MaxMultipartMemory = 8 << 25

	router.LoadHTMLGlob(appconstants.TemplateFolder)
	router.Static(routes.Assets, appconstants.AssetsFolder)

	accounts := gin.Accounts{
		appconstants.UserName: appconstants.Password,
	}
	authorized := router.Group("", gin.BasicAuth(accounts))

	authorized.GET(routes.Home, uploader.RedirectToIndexPage)
	authorized.GET(routes.Index, uploader.IndexPage)
	authorized.POST(routes.UploadFiles, uploader.UploadFile)
	authorized.GET(routes.DownloadFiles, uploader.DownloadFile)
	authorized.POST(routes.DeleteFiles, uploader.DeleteFile)

	if err := router.Run(); err != nil {
		log.Println(err)
	}
}
