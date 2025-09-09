First application in Go language

# To create module
go mod init "module name"

# To generate binarie file / executable
go build

# To execute
./modulo.exe

# To install external package
go get "url", example: go get github.com/badoux/checkmail

# Remove all dependencies which are not use
go mod tidy