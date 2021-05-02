## tcp-chatting

This is a project to teach some features of Golang to friends.

### How To Run
You should run the server first to enbale clients connecting to it. 
Run command `go run cmd/server/main.go` in the root of project. 
Note: Make sure you have Go language installed on your machine. [How to install Go](https://golang.org/doc/install)

After running the server, it should be available to accept new connections (ie. seeing the log like `server started listening on port...`).
To create a new client run the command `go run cmd/client/main.go`.

Run multiple clients and write a message in one of them, then hit enter and you should be able to
see the message logged with client id in all the other clients.


### Todos
- [ ] handle clients that are disconnected
- [ ] implement chat rooms
- [ ] implement private (1to1) chat


### License
 [Apache 2.0](http://www.apache.org/licenses/LICENSE-2.0)