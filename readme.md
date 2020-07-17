## Install go lang
sudo apt-get update <br>
sudo apt-get upgrade <br>
cd /tmp <br>
wget https://dl.google.com/go/go1.11.linux-amd64.tar.gz <br>
sudo tar -xvf go1.11.linux-amd64.tar.gz <br>
sudo mv go /usr/local <br>
export GOROOT=/usr/local/go <br>

OR We can use: <br>
sudo apt-get update <br>
sudo apt-get install golang <br>

## Set up go environment
export GOROOT=/usr/local/go <br>
export GOPATH=$HOME/go <br>
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH <br>

## Update current shell session
source ~/.profile <br>

## Verify installation
go version <br>

## Extension GO
go 0.15.1 (go team at google)

## Create main.go
create folderName <br>
create fileName.go<br>

## Run
cd folderName <br>
go run main.go

## fmt
https://golang.org/pkg/fmt/

## Golang
package main : will be executable program. <br>
import "fmt" : this is from go lang. <br>
fmt is mean format package for provide input and output . <br>
func main(){} : main function is going to be called for execution whatever application is executed and function is going to open. <br>

listener have 3 methods: <br>
1. Accept() (Conn, error) 
Accept waits for and returns the next connection the listener. at here we can accept when we are listening also we can accept and get back a connection and an air.
2. Close() error
Close closes the listener when we finish our programs. Any blocked Accept operations will be unblocked and returns errors.
3. Addr() Addr
Addr returns the listener's network address.