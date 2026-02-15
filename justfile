default:
	just --list

documentation:
	go doc -all -u -http

build:
	go build ./...

test:
	go test -v ./... --cover -coverprofile=reports/coverage.out --covermode set --coverpkg=./...

show-coverage-report:
	go tool cover -html=reports/coverage.out

coverage-report: test show-coverage-report

benchmark:
	go test -benchmem -run=^$$ -bench . ./tests/benchmarks/...

benchmark-package package:
	go test -benchmem -run=^$$ -benchtime 10s -cpuprofile ./cpu-{{package}}.pprof -bench . ./tests/benchmarks/{{package}}

lint:
	go tool golangci-lint run -v --fix

format:
	go fmt ./...

pprof package:
	go tool pprof --http=:8080 ./cpu-{{package}}.pprof

fuzz:
	go test -fuzz=Fuzz -fuzztime=30s ./...

pgo:
    go tool pprof -proto ./cpu.pprof > default.pgo

build-pgo:
	just benchmark-package bloomfilters
	just benchmark-package bloomhashes
	go tool pprof -proto ./cpu-bloomfilters.pprof ./cpu-bloomhashes.pprof > default.pgo