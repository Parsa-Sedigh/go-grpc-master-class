# Section 8. [Hands-On] gRPC Bi-Directional Streaming

## 32-1. What's a Bi-Directional Streaming API
This is possible thanks to http/2.

- The client can send many messages to the server and the server can send many responses to the client. The number of requests and responses
does not have to match
- bi directional streaming RPC is well suited for
    - when the client and the server needs to send a lot of data asynchronously(we don't want to wait for the response to the client to send the
    next message)
    - "Chat" protocol
    - long running connections where you wanna stream info between the client and server

## 33-2. GreetEveryone API Definition

## 34-3. Bi-Directional Streaming API Server Implementation
## 35-4. Bi-Directional Streaming API Client Implementation
## 36-5. [Solution] FindMaximum API
Now it is your turn to write code!

In this exercise, your goal is to implement a FindMaximum RPC Bi-Directional Streaming API in a CalculatorService:

The function takes a stream of Request message that has one integer, and returns a stream of Responses that represent the current maximum between all these integers
Remember to first implement the service definition in a .proto file, alongside the RPC messages
Implement the Server code first
Test the server code by implementing the Client
Example:

The client will send a stream of number (1,5,3,6,2,20) and the server will respond with a stream of (1,5,6,20)

Good luck!