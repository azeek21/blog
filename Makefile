build:
	find cmd/*.go | xargs -n 1 go build -o ./bin/

generate: build
	./bin/generate_logo

migrate: build
	./bin/migrate

#  || xargs go build -ldflags "-s -w" -o ./* 

dev: build 
	./bin/start

dev_prepare: build generate migrate
	./bin/start

start: build generate
	GIN_MODE=release ./bin/start

clean:
	rm -rf ./bin/
	rm ./public/logo.png

.PHONY: build migrate start dev dev_prepare clean
