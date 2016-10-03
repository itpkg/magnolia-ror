# magnolia

A web application by go-lang.

## Install go

### For ubuntu

```
add-apt-repository ppa:ubuntu-lxc/lxd-stable
sudo apt-get update
sudo apt-get install golang
```

### For archlinux

```
sudo pacman -S go go-tools
```

### Add to your .bashrc or .zshrc

```
GOPATH=$HOME/go
PATH=$GOPATH/bin:$PATH
export GOPATH PATH
```

### Some other packages

```
go get -u github.com/nsf/gocode
go get -u github.com/derekparker/delve/cmd/dlv
go get -u github.com/alecthomas/gometalinter
go get -u github.com/golang/lint/golint

go get -u github.com/kardianos/govendor
```

## Database creation

### postgresql

```
psql -U postgres
CREATE DATABASE db-name WITH ENCODING = 'UTF8';
CREATE USER user-name WITH PASSWORD 'change-me';
GRANT ALL PRIVILEGES ON DATABASE db-name TO user-name;
```

## Build

```
go get github.com/itpkg/magnolia
cd $GOPATH/src/github.com/itpkg/magnolia
make
ls -lh dist
```

### Development
```
cd $GOPATH/src/github.com/itpkg/magnolia
# start backend server
cd demo
go run main.go s
# start front server
cd front
npm run dev
```

## Documents

- [go-plus](https://atom.io/packages/go-plus)
- [gorm](http://jinzhu.me/gorm/)
- [gin](https://github.com/gin-gonic/gin)
- [cli](https://github.com/urfave/cli)
- [locale](https://blog.golang.org/matchlang)
- [govendor](https://github.com/kardianos/govendor)
