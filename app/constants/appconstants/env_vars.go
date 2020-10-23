package appconstants

import "os"

//"/home/ghosh/Desktop/static-files"
var (
	StaticFilesPath = os.Getenv("STATIC_FILE_PATH")
	UserName        = os.Getenv("USER_NAME")
	Password        = os.Getenv("PASSWORD")
)
