set windows-shell := ["powershell.exe", "-c"]

default:
	just --list

documentation:
	go doc -all -u -http

build:
	go build ./...

[group('test')]
test:
	go test -v ./... --cover -coverprofile=reports/coverage.out --covermode set --coverpkg=./...

[group('test')]
show-coverage-report:
	go tool cover -html=reports/coverage.out

[group('test')]
coverage-report: test show-coverage-report

[group('benchmark')]
benchmark:
	go test -benchmem -run=^$$ -benchtime 2s -bench . ./tests/benchmarks/...

[group('benchmark')]
benchmark-package package:
	go test -benchmem -run=^$$ -benchtime 10s -cpuprofile ./cpu-{{package}}.pprof -bench . ./tests/benchmarks/{{package}}

[group('checks')]
lint:
	go tool golangci-lint run -v --fix

alias fmt := format

[group('checks')]
format:
	go fmt ./...

[group('checks')]
pprof package:
	go tool pprof --http=:8080 ./cpu-{{package}}.pprof

[group('test')]
fuzz:
	go test -fuzz=Fuzz -fuzztime=30s ./...

[group('pgo')]
pgo:
    go tool pprof -proto ./cpu.pprof > default.pgo

[group('pgo')]
build-pgo:
	just benchmark-package bloomfilters
	just benchmark-package bloomhashes
	go tool pprof -proto ./cpu-bloomfilters.pprof ./cpu-bloomhashes.pprof > default.pgo