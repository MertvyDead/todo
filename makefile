run:
	go run main.go storage.go task.go

build:
	go build -o todo_app main.go

clean:
	rm -f todo_app