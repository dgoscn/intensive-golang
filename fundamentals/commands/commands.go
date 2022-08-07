package main

import "fmt"

func main() {
	fmt.Printf("Another program using %s!\n", "Go")
}

/*
Line command with go

go
go help list
godoc -http=:6060 Documentation offline
go env // GOPATH -> dependency, GOROOT
go doc cmd/vet -> used to detect death code or suspect code <go vet commands.go>
go build commands.go
go run commands.go
go get -u guthub.com/go-sql-driver/mysql -> Instaling dependencies

*/
