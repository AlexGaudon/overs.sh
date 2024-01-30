build:
	bunx tailwindcss -i public/input.css -o public/main.css
	go build -o ./bin/overssh cmd/main.go

dev:
	DEV=true go run cmd/main.go

clean:
	rm public/main.css
	rm -rf bin
	rm -rf tmp

prod: build
	./bin/overssh

deploy: build
	sudo systemctl restart overssh.service