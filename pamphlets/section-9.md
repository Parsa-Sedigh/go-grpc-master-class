# Section 9. [Hands-On] gRPC Advanced Features Deep Dive

## 37-1. [Theory] Errors in gRPC
### Error codes
- in http(REST APIs are built on top of them), there are many error codes
- while http codes are standardized they're not usually clear and people misuse them

- gRPC approach is instead of so many error codes, there are just a few error codes.
- there is also a complete reference to implementation of error codes that close a lot of gaps with documentation: `https://avi.im/grpc-errors`
- if an app needs to return extra information on top of an error code, it can use the metadata context

## 38-2. [Hands-On] Errors implementation
- Let's implement an error message for a new `SquareRoot` unary API
- create `SquareRoot` RPC
- implement the server with the error handling
- implement the client with the error handling

One thing you wanna do usually when you can have errors in an rpc, is document the kind of errors that will be in that rpc.
You can add the docs in proto file's rpc definition.

We use `google.golang.org/grpc/status` to send errors and `google.golang.org/grpc/codes` to specify the actual error codes.

## 39-3. [Theory] Deadlines

Deadlines Doc : https://grpc.io/blog/deadlines

## 40-4. [Hands-On] Deadlines
## 41-5. [Theory] SSL Security
## 42-6. [Hands-On] SSL Security
Code Samples for Security : https://grpc.io/docs/guides/auth.html

## 43-7. [Demo] Language Interoperability
## 44-8. gRPC Reflection & Evans CLI