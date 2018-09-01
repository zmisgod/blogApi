main: main.go
	rm ./main && go build -o main main.go && nohup ./main &