# Application-Level Protocols

**Building any application requires design decisions before you start writing code.**

# Links

* [Application-Level Protocols | Network Programming with Go (golang)](https://ipfs.io/ipfs/QmfYeDhGH9bZzihBUDEQbCbTc5k5FZKURMUoUvfmc27BwL/applevelprotocols/index.html)

# The content of a protocol

* Version Control
* Message Format
* Data Format
* State Information

# Message Format

In the last chapter we discussed some possibilities for representing data to be sent across the wire. Now we look one level up, to the messages that may contain such data.

* The client and server will exchange messages with different meanings:
	* Login request
	* Login reply
	* Get record request
	* Record data reply
* The client will prepare a request, which must be understood by the server.
* The server will prepare a reply, which must be understood by the client.

Commonly, the first part of the message will be a message type.

* Client to server:

          LOGIN <name> <passwd>
          GET <subject> grade
* Server to client: LOGIN succeeded

          GRADE <subject> <grade>

# Data Format

There are two main format choices for messages: byte encoded or character encoded.

## Byte Format 

In the byte format:

* The first part of the message is typically a byte to distinguish between message types. 
* The message handler examines this first byte to distinguish the message type and then performs a switch to select the appropriate handler for that type.
* Further bytes in the message contain message content according to a predefined format (as discussed in the previous chapter).

Pseudocode for a byte-format server is as follows:


```go
handleConn(conn net.Conn) {
        for {
                b := conn.readByte()
                switch b {
                case msg1: ...
                case msg2: ...
		...
                }
        }
}
```

## Character Format

In this mode, everything is sent as characters if possible. For example, an integer 234 would be sent as, say, the three characters 2, 3, and 4 instead of as the one byte 234. Data that is inherently binary may be Base64 encoded to change it into a 7-bit format and then sent as ASCII characters, as discussed in the previous chapter.

In character format:

* A message is a sequence of one or more lines. The start of the first line of the message is typically a word that represents the message type.
* String-handling functions may be used to decode the message type and data.
* The rest of the first line and successive lines contain the data.
* Line-oriented functions and line-oriented conventions are used to manage this.

The pseudocode is as follows:

```go
handleConn(conn net.Conn) {
        for {
                line := conn.readLine()
                switch line.startWith() {
                case msg1: ...
                case msg2: ...
		...
                }
        }
}
```

