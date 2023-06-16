# 1. gRPC Course Overview
The microservices must exchange information and need to agree on:
- the api to exchange data
- the data format
- the error patterns
- load balancing
- many other

One of the popular choices for building API is REST(HTTP-JSON)

### Building an API is hard
- need to think about data model
    - JSON
    - XML
    - something binary?
- need to think about endpoints 
    - GET /api/v1/user/123/post/456
    - POST
- need to think about how to invoke it and handle errors
- efficiency of api
    - too much data
    - too little data -> many API calls?
- latency?
- scalability to 1000s of clients?
- load balancing?
- inter operability between many languages?
- authentication, monitoring, logging?

An API is a contract, saying:
- send me this request(client)
- I'll send you this response(server)
so it's all about data. We'll leave the rest(hard stuff) to gRPC framework.

### What's gRPC?
- part of CNCF
- at a high level, it allows you to define request and response for RPC(remote procedure calls) and handles all the rest for you
- it's built on top of HTTP2, low latency, supports streaming, language independent and makes it easy to plug in authentication, load balancing, logging
and monitoring

### What's an RPC?
- in your client code, it looks like you're just calling a function directly on the server
- RPC is not a new concept(CORBA had this before)
- but with gRPC, it's implemented very cleanly and solves a lof of problems

- you need to define messages and services using protocol buffers
- the rest of the gRPC code will be generated for us and we'll just have to provide an implementation for services
- one .proto file works for 12 programming languages(server and client) and allows you to use a framework that scales to millions of RPC per seconds

### Why protocol buffers?
- protocol buffers are language agnostic
- the code can be generated for pretty much any language
- data is binary and efficiently serialized(small payloads)
- very convenient for transporting a lot of data
- protocol buffers allows for easy API evolution using rules

## 2-2. Course Objective

## 3-3. About your instructor

## 4-4. Important Message