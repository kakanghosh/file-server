# file-server
This is an application for uploading and downloading files.

# To run this project locally (without docker)

# Install go version >= 1.15 

$ go get

$ go mod tidy

# Export environment variables 
1. USER_NAME 
2. PASSWORD
3. STATIC_FILE_PATH (Your prefered folder location)
        

$ go run main.go

Your application should be running on localhost:8080

# To run this project locally (with docker)

You can build the docker image from the Dockerfile
# Use env variables
1. USER_NAME 
2. PASSWORD
3. STATIC_FILE_PATH=/app/static-files

# Or, you can simply go with the docker-compose setup, just run

$ docker-compose up -d