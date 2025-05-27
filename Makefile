# for testing individual modules
# go test -run TestMyFunction ./...
compile:
	go build -o ./bin/gamegorl -gcflags='all=-N -l' ./cmd/gamegorl   2> ./errors.err
	
c:
	go build -o ./bin/gamegorl -gcflags='all=-N -l' ./cmd/gamegorl
	
run:
	./bin/gamegorl

compile_test:
	go test -v ./cmd/gamegorl 2> ./errors.err

test:
	go test -v ./cmd/gamegorl
	
test_f:
	go test -v -run $(FN) ./cmd/gamegorl

debug:debug
	env --chdir=./cmd/gamegorl gdlv debug
	
debugt:debugt
	env --chdir=./cmd/gamegorl gdlv test
	
debug_f:
	env --chdir=./cmd/gamegorl gdlv test -run $(FN)
