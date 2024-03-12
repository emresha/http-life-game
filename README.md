# A simple http-server of life game (made with Ya.Lyceum)

This program is a game of life, but a server version.

## How to run
### First method
When you type `go run cmd/life/main.go` it runs a server on `localhost:8081`.

After that you need to visit `localhost:8081`, where you will see the magic happen:)
### Second method
Type `go build cmd/life/main.go` in root directory. You will get the `main` executable.

After you run it, you can now visit `localhost:8081` to see the game in action.

## Other
If you refresh the page, the game starts over again, where 7% of the field is filled with alive cells.<br>
This is just practice in go
