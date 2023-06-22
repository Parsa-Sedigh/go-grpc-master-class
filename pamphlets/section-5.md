# Section 5. [Hands-On] gRPC Unary

## 17-1. What's an Unary API
- unary RPC calls are the basic request/response that everyone is familiar with
- the client will send one message to the server and will receive one response from the server
- unary RPC calls will be the most common for your APIs
    - unary calls are very well suited when your data is small
    - start with unary when writing APIs and use streaming API if performance is an issue

For each RPC call we have to define a "Request" message and a "Response" message.

## 18-2. Greet API Definition
After creating a rpc function named for example `greet`, we will get `GreetServiceClient` interface, `NewGreetServiceClient` func, `GreetServiceServer` interface
and `RegisterGreetServiceServer` func. The interfaces have a function named `greet`.

## 19-3. Unary API Server Implementation
We now need to have our server struct to implement the `GreetServiceServer` interface(yeah server interface because we're writing the server side
code not client side).

## 20-4. Unary API Client Implementation
We'll implement a client call for our unary RPC.

## 21-5. [Solution] Sum API