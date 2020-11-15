build:
	@echo "Building for current OS"
	@go build -o passwdvault main.go

compile:
	@echo "Compiling for every OS and Platform"
	@GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
	@GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	@GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go

link:
	@echo "Symlinking $(shell pwd)/passwdvault to /usr/local/bin/passwdvault"
	@ln -s $(shell pwd)/passwdvault /usr/local/bin/passwdvault

unlink:
	@echo "Removing link"
	@rm /usr/local/bin/passwdvault

clean:
	@rm -rf bin
	@rm passwdvault

install: build link

uninstall: unlink clean