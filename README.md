# Go Intensivo 3

> FullCycle / CodeEdu (2024-05-13 a 15)

## Golang : Oportunidades e Criando nosso primeiro programa

### Sobre o Go

![slide-01](/files/slide-01.png)

![slide-02](/files/slide-02.png)

![slide-03](/files/slide-03.png)

![slide-04](/files/slide-04.png)

![slide-05](/files/slide-05.png)

![slide-06](/files/slide-06.png)

![slide-07](/files/slide-07.png)

![slide-08](/files/slide-08.png)

![slide-09](/files/slide-09.png)

![slide-10](/files/slide-10.png)

![slide-11](/files/slide-11.png)

![slide-12](/files/slide-12.png)

![slide-13](/files/slide-13.png)

![slide-14](/files/slide-14.png)

![slide-15](/files/slide-15.png)

![slide-16](/files/slide-16.png)

![slide-17](/files/slide-17.png)

![slide-18](/files/slide-18.png)

![slide-19](/files/slide-19.png)

![slide-20](/files/slide-20.png)

![slide-21](/files/slide-21.png)

![slide-22](/files/slide-22.png)

![slide-23](/files/slide-23.png)

![slide-24](/files/slide-24.png)

![slide-25](/files/slide-25.png)

![slide-26](/files/slide-26.png)

![slide-27](/files/slide-27.png)

![slide-28](/files/slide-28.png)

### Código fonte

[Código fonte](day1/)

### Comandos

- go mod init github.com/rodolfoHOk/fullcycle.go-intensivo-3/day1

- go run main.go

- go build -o "hello"

- ./hello

- GOOS=windows go build -o "hello.exe"

## Go: Alta Performance, Multithreading e Profiling na prática

### Sobre o Go

![slide-29](/files/slide-29.png)

![slide-30](/files/slide-30.png)

![slide-31](/files/slide-31.png)

![slide-32](/files/slide-32.png)

![slide-33](/files/slide-33.png)

![slide-34](/files/slide-34.png)

![slide-35](/files/slide-35.png)

![slide-36](/files/slide-36.png)

![slide-37](/files/slide-37.png)

![slide-38](/files/slide-38.png)

![slide-39](/files/slide-39.png)

![slide-40](/files/slide-40.png)

![slide-41](/files/slide-41.png)

### Código fonte

[Código fonte](day2/)

## Go: API sem Framework, gRPC e Profiling

### Código fonte

[Código fonte](day3/)

### Comandos

- go mod tidy

### Profiling

- endpoint: http://localhost:6060/debug/pprof

- go tool pprof -seconds 5 http://localhost:6060/debug/pprof/profile

  - top
  - list CPUIntensiveEndpoint
  - web

- go test -bench=. -benchmem -memprofile mem1.prof -cpuprofile cpu1.prof -count 10 > 1.bench

- go test -bench=. -benchmem -memprofile mem2.prof -cpuprofile cpu2.prof -count 10 > 2.bench

- benchstat 1.bench 2.bench

- go tool pprof cpu1.prof

  - top

- go tool pprof cpu2.prof

  - top

- go tool pprof mem1.prof

  - top

- go tool pprof mem2.prof

  - top
  - list GenerateLargeString
