dst=dist

build:
	go build -ldflags "-s -X main.version=`git rev-parse --short HEAD`" -o $(dst)/magnolia demo/main.go
	-cp -rv demo/locales $(dst)/
	cd front && npm run build
	-cp -rv front/dist $(dst)/public

clean:
	-rm -rv $(dst)
