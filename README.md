# A simple http-server of life game (made with Ya.Lyceum)

This program is a game of life, but a server version.

When you type `go run cmd/life/main.go` it runs a server on `localhost:8081`.

When you visit `http://localhost:8081/nextstate`, it sends you a json file of cells in game, where<br>
true - the cell is alive
false - the cell is dead

this is just practice in go
