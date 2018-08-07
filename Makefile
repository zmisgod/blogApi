main: main.go
	go build -o main main.go && nohup ./main &
