## BUILDING COMMAND LINE TOOL for Youtube Video

Basically I am performing CRUD functionality on videos data and then run this application by using Command Line Tool and I am running this application within a docker container.

## we can run go in a small docker container without installing go in our Syste

## creating docker image
docker build --target dev . -t go

## run docker image
docker run -it -v ${PWD}:/work go sh

## in the docker container we can check our go version
go version

## creating go module
go mod init module_name

## go build used for creating binary file
go build

## and run this static binary by using this command
./videos 

##  creating new file with powershell command
type nul > main.go
