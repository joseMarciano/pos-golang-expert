module pos-golang-expert

// to create go mod we need run go mod init pos-golang-expert

// go mod tidy. This command is to update our dependencies, remove unused dependencies, and update the go.sum file.
// the go.sum file is a file that contains a hash of the dependencies and their versions.
go 1.21.5

require (
	github.com/google/uuid v1.5.0 // indirect
	golang.org/x/exp v0.0.0-20240112132812-db7319d0e0e3 // indirect
)
