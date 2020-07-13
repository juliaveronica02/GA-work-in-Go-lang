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