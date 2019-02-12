include .env

PROJECTNAME=$(shell basename "$(PWD)")

## exec: Run given command, wrapped with custom GOPATH. e.g; make exec run="go test ./..."
exec:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) $(run)

## gen: Generates API code from swagger tool 
gen: ( cd api ; swagger generate server -t gen -f ./swagger/swagger.yml -A spotter )
	# docker pull quay.io/goswagger/swagger
	# alias swagger="docker run --rm -it -e GOPATH=$HOME/go:/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger"
	# swagger version

help: Makefile
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
