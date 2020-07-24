## Install go lang
sudo apt-get update <br>
sudo apt-get upgrade <br>
cd /tmp <br>
wget https://dl.google.com/go/go1.11.linux-amd64.tar.gz <br>
sudo tar -xvf go1.11.linux-amd64.tar.gz <br>
sudo mv go /usr/local <br>
export GOROOT=/usr/local/go <br>
export GOPATH=$HOME/Projects/Proj1 <br>
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH <br>
go version <br>
go env <br>

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
go 0.15.1 (go team at google) <br>
live share (real-time collaborative) <br>

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

## Extension
install Go <br>
after that ctrl+shift+p <br>
type go: install and update tools <br>
checklist all and press ok button. <br>

## Error while install and update Go tools
if we got errors while trying to install these tools, make sure we already installed git and if we do, try and reinstall it. when we initially did this, the installs kept on failing, even though we had git installed. We uninstall git and reinstall and then the tools installed successfully. if most of the tools install, but some fail it could be because a newer version of golang is required, in this case updating to the newest version may fix the issue.

## NOTE!!!
## Note Golang
sudo rm -rf go <br>
sudo rm -rf /usr/local/go/ <br>
ls <br>
history | grep tar <br>
sudo tar -C /usr/local -xzf go1.14.6.linux-amd64.tar.gz <br>
go version <br>
cd go/src <br>
cd nano ~/.bashrc scroll until bottom and add 
export PATH=$PATH:/usr/local/go/bin <br>
export PROJECT=$GOPATH/src/github.com/juliaveronica02 below export nvm<br>
cd nano ~/.profile <br>
cd source ~/.bashrc <br>
cd go/scr/github.com/juliaveronica02/whatsapp-go <br>

## Note Struct
ID is for code and id is for json (postman).
```javascript
ID          string `json:"id"`
```

## vs code shortcut
alt+ arrow up and arrow down : pindah baris code. <br>
ctrl+/ : comment one line code. <br>
ctrl+g+number: go to line code. <br>
return: debugger (if we have error, next code enable to execute). <br>

## Extension
GO and live share.

## Postman
go to collections -> create newFolder (whatsapp-go) -> save -> add request (ctrl+N) -> requestName (get events) -> add -> no environment -> manage environment -> add -> environment name (whatsapp-go) -> key (host), value (localhost:8080) -> update.