# Section 11. Next Steps

## 56-1. gRPC Services in the Real Word

Google Pub/Sub .proto : https://github.com/googleapis/googleapis/blob/master/google/pubsub/v1/pubsub.proto

Google Spanner .proto : https://github.com/googleapis/googleapis/blob/master/google/spanner/v1/spanner.proto

## 57-2. Congrats & Next Steps

## 58-3. THANK YOU!

---
### What about Gogo
In the open source world, it is not unusual to see active community members fork a popular project and add features they find the community needs.

One notable project is "gogo", created by Walter Schulze, which aims to provide an alternative at how Protocol Buffers and gRPC are implemented 
in Go. This different implementation from the one provided by Google has many benefits, mainly around performance and usability, as well as 
providing more code generation options, which makes this project one of the most important in the gRPC ecosystem today. You can find
the project alongside its 1800+ stars on GitHub here: https://github.com/gogo/protobuf

This course does not go over how to use this project as the code implementation will be different than Google's, but it is worth mentioning it and
having the students investigate it on their own. It is not a guarantee that the two will eventually merge, but many people in the community hope
to see a convergence of both the implementation by Google and Gogo some day. Time will tell!

If you'd like to read about it, I strongly recommend this blog: https://jbrandhorst.com/post/gogoproto/

Regardless of the implementation you will use for your projects, the gRPC concepts are the exact same. Happy learning!