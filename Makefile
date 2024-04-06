server:
	go run apps/main.go server
mod:
	cd apps && go mod tidy && cd ..
	