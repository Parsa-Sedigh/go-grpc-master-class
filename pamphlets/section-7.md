# Section 7. [Hands-On] gRPC Client Streaming

## 27-1. What's a Client Streaming API
- client streaming RPC API are a new kind of API enabled thanks to HTTP/2
- the client will send **many** messages to the server and will receive **one** response from the server(at any time). The response
sent from the server can be received at anytime, there's no guarantee that it will be sent after it has received all the messages from
the client, it's really up to the server to decide when it wants to send the response to the client. Obviously if you design the server
to only send the response once client has sent everything, that's perfect.
- streaming client is well suited for
    - when the client needs to send a lot of data(big data)
    - when the server processing(for each message) is expensive and should happen as the client sends data. For example if you have to
    process 10,000 messages, it's better to start processing them one by one as soon as we can VS 10,000 at a time processing them all at once.
    - when the client needs to **push** data to the server without really expecting a response

## 28-2. LongGreet API Definition

## 29-3. Client Streaming API Server Implementation
The server will only respond to the client once the client is done sending requests. But in theory, the server can respond whenever it wants.

Create LongGreet func on the server.

## 30-4. Client Streaming API Client Implementation

## 31-5. [Solution] ComputeAverage API
Now it is your turn to write code!

In this exercise, your goal is to implement a ComputeAverage RPC Client Streaming API in a CalculatorService:

The function takes a stream of Request message that has one integer, and returns a Response with a double that represents the computed average
Remember to first implement the service definition in a .proto file, alongside the RPC messages
Implement the Server code first
Test the server code by implementing the Client
Example:

The client will send a stream of number (1,2,3,4) and the server will respond with (2.5), because (1+2+3+4)/4 = 2.5

Good luck!