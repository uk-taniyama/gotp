build: bin/main

bin/main: main.go controllers/*.go models/*.go services/*.go utils/*.go
	go build -o bin/main main.go

run: build
	rm main.log
	bin/main

try:
	rm cookiejar
	curl -c cookiejar http://localhost:3000/login -X POST -d "username=test&password=test"
	curl -b cookiejar -D - http://localhost:3000/hello -X POST -d "message=Hello"
	curl -b cookiejar http://localhost:3000/logout
	curl -b cookiejar -D - http://localhost:3000/hello -X POST -d "message=Hello"
	curl -c cookiejar http://localhost:3000/login -X POST -d "username=test&password=test2"
	curl -b cookiejar http://localhost:3000/logout
	# curl -v -b cookiejar http://localhost:3000/logout
	# curl -v -b cookiejar http://localhost:3000/doSomething

pre:
	go install github.com/volatiletech/sqlboiler/v4@latest
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest

gen:
	sqlboiler psql
