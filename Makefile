build:
	mkdir -p ./dist
	go build -o ./dist/ptxt

clean:
	rm -rf ./dist
