# Golang Protobuf (protocol Buffer)
![WorkerPool](https://user-images.githubusercontent.com/71340325/186423274-41d6c626-12bc-473e-a848-3e39c8b79bd3.jpg)

*  Google's language-neutral, platform-neutral, extensible mechanism for serializing structured data.
*  We define how we want our data to be structured once, and use special generated source code to write and read structured data to and from data streams.
* Currently supports Java, Python, C++, Kotlin, Dart, Go, Ruby, and C#.
* Creates a class that implements automatic encoding and parsing of the protocol buffer data with binary format.

## Instructions:
1. Starts with a .proto file. (Add a message for each data structure, and specify a name and type for each field)
2. .proto file starts with a package definition. (Prevents naming conflicts between different projects)
```
syntax = "proto3";
option go_package = "myProject/packageTest";
package packageTest;

```
3. Start your message definitions.
```
message Person{
    string name = 1;
    string gender = 2;
    int32 id = 3;
}...etc
```
* The "=1" marker identifies the unique "tag" that field uses in the binary encoding.
4. Install Go protocol buffers plugin. The plugin must be in your $PATH for the compiler protoc to find it.
5. Run the compiler. Mack sure to specify the source directory(where your applications's source code lives, if not provided, the current directory will be used.),the destination directory, and the path to your .proto.
```
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/test.proto
```
```
protoc SRCDIR --go_out=DESDIR
```
* Compile all proto files at once
```
./build/compile_by_protobuf.sh
```


<hr>

## Encoding / Serialize message

### Encoding
* We serialize messages to bytes Slice.
```
pbObj := PersonPB.PersonResponse{
    Error: ............
}
ByteSlice,err := proto.Marshal(pbObj)
```
* After serializing it, we can send this bytes and parse somewhere else.

### Decoding
1. Getting data from body and decode
```
binary,err := ioutil.ReadAll(r.Body) //get http request data
pbObj := PersonPB.PersonRequest{}
protojson.Unmarshal(binary,&pbObj)
```
2. DiscardUnknown
```
binary,err := ioutil.ReadAll(r.Body)
option := protojson.UnmarshalOptions(DiscardUnknown:true)
option.Unmarshal(binary,&pbObj)
```
