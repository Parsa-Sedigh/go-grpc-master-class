# Section 6. [Hands-On] gRPC Server Streaming
## 22-1. What's a Server Streaming API
- it will take **one** GreetManyTimesRequest that contains a Greeting
- it will return **many** GreetManyTimesResponse that contains a result string

## 23-2. GreetManyTimes API Definition
**Everytime** you have a new rpc, you wanna create a new request and a new response message type, even if the fields with another rpc function is
the same. So always create a new message for each rpc req and res.

## 24-3. Server Streaming API Server Implementation

## 25-4. Server Streaming API Client Implementation
We'll implement a client call for our streaming server RPC.

When we have server streaming, the `Recv()` function can get called as many times as possible and when the stream is closed, we're gonna
receive an `io.EOF` error. To deal with stream, we use an infinite loop and inside it, we call Recv() and if the error was `io.EOF` means we reached
the end of the stream and we can break from the loop.

## 26-5. [Solution] PrimeNumberDecomposition API
Now it is your turn to write code!

In this exercise, your goal is to implement a PrimeNumberDecomposition RPC Server Streaming API in a CalculatorService:

The function takes a Request message that has one integer, and returns a stream of Responses that represent the prime number
decomposition of that number (see below for the algorithm).
Remember to first implement the service definition in a .proto file, alongside the RPC messages
Implement the Server code first
Test the server code by implementing the Client
Example:

The client will send one number (120) and the server will respond with a stream of (2,2,2,3,5), because 120=2*2*2*3*5

Algorithm (pseudo code):

k = 2
N = 210
while N > 1:
if N % k == 0:   // if k evenly divides into N
print k      // this is a factor
N = N / k    // divide N by k so that we have the rest of the number left.
else:
k = k + 1

Good luck!