# gRPC tutorial

# [gRPC Documentation](https://grpc.io/docs)

---

## Introduction to gRPC

![gRPC theory](https://grpc.io/img/landing-2.svg)

### Working with Protocol Buffers

By default, gRPC uses Protocol Buffers, Protocol buffer data is structured as **message**.  
Containing a series of name-value pairs called fields.
```Proto
message Person{
  string name=1;
  int32 id =2;
  bool has_ponycopter=3;
}
```

Then,once you've ispecified your data structures, you use the protocol buffer compiler `protoc` togennerate data access classes in your preferred languages from your proto definition. These provide simple accessors for each field, `like name()` and `set_name()`,as well as methods to serialize/parse the whole structure to/from raw bytes.  
Running the compiler on the example above will generate a class called `Person`. You can then usethis class in your application to populate, serialize, and retrieve `Person` protocol buffer messages.

```
// The greeter service definition.
service Greeter{
  //Sends a greeting
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

// The request message containing the user's name.
message HelloRequest{
  string name=1;
}

// The response message containing the greetings
message HelloReply{
string message=1;
}
```


