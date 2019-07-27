all: sensor-node.out

sensor-node.out: main.go count/count.go
	go build -o sensor-node.out

run:
	./sensor-node.out