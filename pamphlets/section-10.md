# Section 10. [Hands-On] CRUD API with MongoDB

## 45-1. Install MongoDB
If you run the mongodb for the first time, it will throw an error saying:
`exception in initAndListen: NonExistentPath: Data directory /data/db not found.`

## 46-2. Install MongoDB UI - Robo 3T
Robo 3T(previous name was robo mongo).

## 47-3. Blog Service Golang Setup

## 48-4. MongoDB Driver Golang Setup
```shell
# old version:
# go get github.com/mongodb/mongo-go-driver/mongo

# for newer versions:
go get go.mongodb.org/mongo-driver/mongo
```

Note: We won't use `go-mgo/mgo` package.

## 49-5. CreateBlog Server

## 50-6. CreateBlog Client

## 51-7. ReadBlog Server
In `ReadBlog` rpc, if we don't find the blog, then we should return an errors that says `NotFound`.

Note: After changes to .proto files, run the related protoc command.

To run things:
```shell
go run blog/blog_server/server.go

go run blog/blog_client/client.go
```

## 52-8. ReadBlog Client

## 53-9. UpdateBlog Server
For finding out the signature of the CRUD functions that you need to implement, look at the `BlogServiceServer` interface in the
generated `blog_grpc.pb.go` .

## 54-10. UpdateBlog Client

## 55-11. DeleteBlog Server

## 56-12. DeleteBlog Client

## 57-13. ListBlog Server
This endpoint is gonna be server streaming.

## 58-14. ListBlog Client

## 55-15. Evans CLI test with CRUD
After registering the reflection on server, run:
```shell
evans -p 50051 -r
show package
show service
service BlogService

# the prompt would be: default.BlogService@127.0.0.1:50051
call CreateBlog

# will prompt for message body
# note: if the values have space, enter them with double quotes around the string

call DeleteBlog

call ListBlog
```

---

### Code Changes
Important: Code Changes
Some MongoDB methods (just a few) I'm using in the next videos have been deprecated.

The changes are minor and summarized here: https://github.com/simplesteph/grpc-go-course/pull/1/files

If you download the code from https://github.com/simplesteph/grpc-go-course you will have the latest and working code

Overall, this doesn't change any understanding for this section, but should avoid a little bit of frustration :) Happy learning!