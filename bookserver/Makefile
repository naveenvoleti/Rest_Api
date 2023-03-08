URL = http://localhost:8080
input := main.go
output_dir := ./build

build:
	echo "Building target"
	GOOS=freebsd GOARCH=386 go build -o $(output_dir)/main-freebsd-386 $(input)
	GOOS=linux GOARCH=386 go build -o $(output_dir)/main-linux-386 $(input)
	GOOS=windows GOARCH=386 go build -o $(output_dir)/main-windows-386 $(input)

run:
	echo "Running target"
	go run $(input)

test:
	go test

clean:
	rm $(output_dir)
	echo "Deleted"

hello:
	curl $(URL)/hello?name=Naveen

#all: run hello

books:
	echo "Getting Books"
	curl -i $(URL)/books

return:
	echo "Updating Books"
	curl -i -X PUT -H 'Content-Type: application/json' -d '{"id":1}' $(URL)/books/return

insert:
	echo "Inserting Books"
	curl -i -X POST -H 'Content-Type: application/json' -d '{"id":4,"title":"This is test Book","author":"Test Author","quantity":10}' $(URL)/books/add

delete:
	echo "Deleting Books"
	curl -i -X DELETE -H 'Content-Type: application/json' -d '{"id": 1}' $(URL)/books/delete

borrow:
	echo "Borrowing Books"
	curl -i -X PUT -H 'Content-Type: application/json' -d '{"id": 1}' $(URL)/books/borrow