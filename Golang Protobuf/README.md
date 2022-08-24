# Golang Protobuf (protocol Buffer)

*  Google's language-neutral, platform-neutral, extensible mechanism for serializing structured data.
*  We define how we want our data to be structured once, and use special generated source code to write and read structured data to and from data streams.
* Currently supports Java, Python, C++, Kotlin, Dart, Go, Ruby, and C#.
* Creates a class that implements automatic encoding and parsing of the protocol buffer data with binary format.

## Instructions:
1. Starts with a .proto file. (Add a message for each data structure, and specify a name and type for each field)
2. .proto file starts with a package definition. (Prevents naming conflicts between different projects)
```
syntax = "proto3";
package test;
import "google/protobuf/timestamp.proto";

```
3. Start your message definitions.
```
message Person{
    string name = 1;
    string gender = 2;
    int32 id = 3;
}...etc
```
* The "=1" marker identify the unique "tag" that field uses in the binary encoding.
4. Install Go protocol buffers plugin. The plugin must be in your $PATH for the compiler protoc to find it.
5. Run the compiler. Mack sure to specify the source directory(where your applications's source code lives, if not provided, the current directory will be used.),the destination directory, and the path to your .proto.
```
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/test.proto
```
