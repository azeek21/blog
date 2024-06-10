build:
	find cmd/*.go | xargs -n 1 go build -o ./bin/

generate_views:
	templ generate

generate: generate_views build
	./bin/generate_logo

migrate: build
	./bin/migrate

#  || xargs go build -ldflags "-s -w" -o ./* 

dev: generate 
	./bin/start

dev_prepare: build generate migrate
	./bin/start

start: generate
	GIN_MODE=release ./bin/start

clean:
	rm -rf ./bin/
	rm -f ./public/logo.png
	find ./views -name '*_templ.go' -delete

.PHONY: build migrate start dev dev_prepare clean generate_views
