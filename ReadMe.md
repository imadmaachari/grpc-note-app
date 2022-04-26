# gRPC Note App

This is an example of using gRPC with Golang ,You can check the story on medium using this [MEDIUM](www.medium.com) link .

## Table of content

- [Installation](#installation)
    - [Go binaries](#typo3-extension-repository)
    - [Protoc](#composer)
    - [gRPC dependencies](#gRPC-Dependencies)
- [Project setup](#Project-Setup)
    - [Note Module](#Create-a-new-go-module)
    - [Screenshots](#Screenshots)
- [Links](#links)


## Installation

There is some prerequisties to have before diving in the code 

- Go 
- Protoc
- gRPC dependencies

### Go
Dawnload the last version of golang from [here](https://go.dev/dl/)
### ProtoC
Install the latest release of the protocol compiler from pre-compiled binaries, download from [official site](github.com/google/protobuf/releases) the zip file corresponding to your operating system . then , update your environments to include the path to the protoc executable(ex. C:\Tools\protoc-3.18.0-win64\bin).
### gRPC Dependencies
Install Go specific gRPC dependencies:

`go install http://google.golang.org/protobuf/cmd/protoc-gen-go@v1.26`

`go install http://google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1`


## Project Setup
### Create a new go module 
```bash
go mod init github.com/imadmaachari/grpc-note-app
```
inside the project , create 3 folders :

```bash
mkdir "proto" "client/cli" "server"
```
### Project Tree
```
|github.com/imadmaachari/grpc-note-app'
├── protos
├── client
├   └── cli
├       └── main.go
└── server
    └── main.go

```
### Screenshots
- Proto file :
- Server implementation :
- Client implementation :
