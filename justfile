default:
	just --list

documentation:
    go doc -all -u -http

build:
    go build ./...

test:
    go test -v -x ./... --cover -coverprofile=reports/coverage.out --covermode atomic --coverpkg=./...

show-coverage-report:
    go tool cover -html=reports/coverage.out

coverage-report: test show-coverage-report

lint:
    go tool golangci-lint run -v --fix

format:
    go fmt ./...

pprof:
    go tool pprof --http=:8080 ./cpu.out

fuzz:
    go test -fuzz=Fuzz -fuzztime=30s ./...