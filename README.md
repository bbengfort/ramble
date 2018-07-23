# Ramble

**A streaming gRPC point to point chat system.**

This app is a quick proof of concept for a bidirectional streaming service using gRPC. The app implements both a chat server and a chat client.

## Connecting to a Chat

To install ramble, first `go get` the project:

```
$ go get github.com/bbengfort/ramble/...
```

If you have placed `$GOPATH/bin` in your `$PATH` you should now be able to lookup `ramble` directly; otherwise, `cd $GOPATH/src/github.com/bbengfort/ramble` and run `make install`. Once ramble is connected to your path, you can enter the chat client as follows:

```
$ ramble chat -n username -a 192.168.1.1:3265
```

Replacing `username` with the name you wish to be identified as, and `192.168.1.1:3265` with the address of the chat server. You should now be connected to the chat and see a terminal UI.

In order to start chatting, press the `TAB` button to enter the chat window; you can then type your message and press `ENTER` to send it. If you hit tab again, you'll be taken to the message history window, which will allow you to scroll through messages. Note that the chat history only keeps 150 of the most recent messages in memory at a time. To quit, use `CTRL+C`.

## Running a Server

Running the server is as simple as running:

```
$ ramble serve -p 3265
```

Replacing 3265 with the port you'd like the server to listen for chat messages on (the default is 3265).

Of course, things get more complicated when deploying a server to run as a long running-service. Instructions for deploying the chat server as a systemd service on Ubuntu follow.

### Ubuntu systemd service

Clone the repository to your `$GOPATH` and change directories into the project directory. If you have already run `go get`, then make sure that the binary in `$GOPATH/bin` is symlinked to `/usr/local/bin/ramble`, otherwise run `make install` to build the binary in this location. 
