[grpc](https://grpc.io/docs/languages/)
[blog post](https://medium.com/@rohitk12r/design-and-develop-grpc-service-using-golang-d0b0ba212af0)
[blog post](https://medium.com/@tj.ruesch/getting-started-with-grpc-and-golang-45407211f04d)
-
[blog post](https://medium.com/@muradu.iurie.1986/grpc-in-go-unlocking-high-performance-microservices-661164d06434)

```sh
$ go get google.golang.org/protobuf/cmd/protoc-gen-go
go: downloading google.golang.org/protobuf v1.34.2
go: added google.golang.org/protobuf v1.34.2
```

```sh
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go: downloading google.golang.org/grpc v1.64.0
go: downloading google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.4.0
go: downloading google.golang.org/protobuf v1.34.1
```

generate protc files
```sh
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    routeguide/route_guide.proto
```

Analyzing my command

```sh
protoc --go_out=src/grpc \
    --go-grpc_out=src/grpc \
    src/proto/todo.proto

....
option go_package = "src/grpc";
```
This placed the file in the src/grpc/src/grpc
the first path from the protoc, the second from the options

```sh
option go_package = "./todo";
```
src/grpc/todo

```bash
[AI Prompt](https://www.phind.com/search?cache=jgei95zcty0b7uqm49ik3yaf)

The go get, go install, and go get -u commands in Go serve different purposes related to managing and installing packages and binaries within a Go project. Here's a detailed explanation of each:

go get
Purpose: Used to download and install packages along with their dependencies from remote repositories. It updates the go.mod and go.sum files with the versions of the downloaded packages. If the package is already present, go get will update it to the latest version available.
Behavior: Downloads the package and its dependencies, updating the go.mod and go.sum files accordingly. If the package is already installed, it fetches the latest version and updates the files.
go install
Purpose: Compiles and installs the package in your project's GOPATH or GOBIN. Unlike go get, it does not update the go.mod or go.sum files. Instead, it uses the information in those files to determine the correct versions of the dependencies to use.
Behavior: Compiles and installs the package, placing the resulting binary in the $GOPATH/bin directory (or the directory specified by GOBIN if set). It does not modify the go.mod or go.sum files.
go get -u
Purpose: Updates all the dependencies of the current module to their latest minor or patch upgrades. It checks the go.mod file for the current versions of the dependencies and compares them against the latest versions available in the remote repository.
Behavior: Updates the go.mod and go.sum files to reflect the latest versions of the dependencies, fetching and downloading the new versions if necessary.
Summary of Differences
go get is primarily used for adding new packages or updating them to their latest versions, affecting the go.mod and go.sum files.
go install compiles and installs packages without modifying the go.mod or go.sum files, relying on their current state to resolve dependencies.
go get -u specifically targets updating dependencies to newer versions, reflecting changes in the go.mod and go.sum files.
These commands are essential tools in managing dependencies and binaries in Go projects, allowing developers to easily add, update, and install packages and their dependencies as needed
```


```bash
In gRPC with Go, the choice between embedding UnimplementedTranslationServer and TranslationServer in your service implementation struct depends on the intended behavior of your server regarding the implementation of the service methods.

Using UnimplementedTranslationServer
Purpose: Embedding UnimplementedTranslationServer in your service implementation struct is a way to indicate that you are implementing a gRPC service but have not yet implemented any of the service methods. This is useful during development or when you are planning to implement the service methods later.
Behavior: Any attempt to call a method on the embedded UnimplementedTranslationServer will result in a runtime error because the methods are not implemented. This serves as a clear indication that the service is incomplete.
Use Case: You might use this approach when starting a new service implementation or when you want to ensure that all service methods are explicitly implemented before the server starts accepting calls.
Using TranslationServer
Purpose: Embedding TranslationServer (assuming TranslationServer is a type that implements the service interface) in your service implementation struct indicates that you are providing concrete implementations for all the service methods defined in the TranslationServer interface.
Behavior: With TranslationServer embedded, the service methods are expected to be implemented, and the server will respond correctly to client calls for these methods.
Use Case: You would use this approach when you have completed the implementation of all required service methods and are ready to start serving clients.
Why Choose One Over the Other?
Development Stage: If you are still developing the service and have not yet implemented all methods, using UnimplementedTranslationServer signals that the service is not yet ready. As you progress and implement methods, you can replace UnimplementedTranslationServer with TranslationServer to indicate readiness.
Explicit Implementation: Using TranslationServer requires you to explicitly implement all service methods, which can be beneficial for ensuring completeness and adherence to the service contract.
```

cross platform class from node client to golang server
edited the node clinet to include `package proto` which was contained in here
```
https://www.phind.com/search?cache=jgei95zcty0b7uqm49ik3yaf
```
