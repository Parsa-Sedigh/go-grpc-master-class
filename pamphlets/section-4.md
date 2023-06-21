# Section 4. [Hands-On] gRPC Project Overview & Setup

## 13-1. Go Dependencies Setup
Install grpc-go:
```shell
go get -u google.golang.org/grpc
```

The second thing to install is: protoc-gen-go:
```shell
go get -u github.com/golang/protobuf/protoc-gen-go
```

## 14-2. Code Generation Test
In `go` directory, we have `bin`, `package` and `src` directories.

All the code you write, goes into `src/github.com/<your github username>/<project name>`.

Create `greet/greetpb` directory then greet.proto .

To generate code, run(**this doesn't work, look at the next command**):
```shell
protoc greet/greetpb/greet.proto --go_out=plugins=grpc:. # Not working
```
Use this instead:
```shell
protoc ./greet/greetpb/greet.proto --go-grpc_out=.
```

To not repeat these kind of commands, create a `generate.sh` file and put this command in there.

## 15-3. Server Setup Boilerplate Code
Create greet_server directory.

`50051` is the default port for gRPC.

## 16-4. Client Setup Boilerplate Code
Create greet_client folder and client.go inside there.

By default, gRPC has SSL in it but for now because we don't have SSL certificate, we use grpc.WithInsecure passed to grpc.Dial().
Note: Remove this option when going to production and use SSL.

---
This course assumes you already know Golang. But in case you need to have the same setup as me, perform step 1, 2 & 3. Otherwise, just perform step 3.

1) Go Setup
   Instructions to install Golang are here: https://golang.org/doc/install

Make sure that the go binaries are in your PATH  (see installation instructions), so that you can run the following command in a terminal:

go version
(returns: go version go1.10 darwin/amd64 or different version)


2) VSCode Setup
   You are free to use a text editor of your choice, but in this course I will use VSCode.  It is not necessary to have the same text editor as 
me to complete the course.

=========== Instructions for VSCode ==========

1. Install VSCode: https://code.visualstudio.com/

2. Install VSCode extensions (https://code.visualstudio.com/docs/editor/extension-gallery#_browse-and-install-extensions):vscode-proto3  & Clang-Format 
(if you want to format your files automatically)

3. Install Clang-Format (if you want to format your files automatically):

MacOSX: brew install clang-format
Windows / Ubuntu: See http://www.codepool.biz/vscode-format-c-code-windows-linux.html
4. Install protoc (see below)

3) Protoc Setup
   In order to perform code generation, you will need to install protoc  on your computer.

============ MacOSX =============

It is actually very easy, open a command line interface and type brew install protobuf

============ Ubuntu (Linux) ============

Find the correct protocol buffers version based on your Linux Distro: https://github.com/google/protobuf/releases

Example with x64:

### Make sure you grab the latest version
curl -OL https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip
### Unzip
unzip protoc-3.5.1-linux-x86_64.zip -d protoc3
### Move protoc to /usr/local/bin/
sudo mv protoc3/bin/* /usr/local/bin/
### Move protoc3/include to /usr/local/include/
sudo mv protoc3/include/* /usr/local/include/
### Optional: change owner
sudo chown [user] /usr/local/bin/protoc
sudo chown -R [user] /usr/local/include/google
============ Windows ============

Download the windows archive: https://github.com/google/protobuf/releases

Example: https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-win32.zip

Extract all to C:\proto3

Your directory structure should now be

C:\proto3\bin

C:\proto3\include

Finally, add C:\proto3\bin to your PATH:

From the desktop, right-click My Computer and click Properties.
In the System Properties window, click on the Advanced tab
In the Advanced section, click the Environment Variables button.
Finally, in the Environment Variables window (as shown below), highlight the Path variable in the Systems Variable section 
and click the Edit button. Add or modify the path lines with the paths you wish the computer to access. Each different directory is
separated with a semicolon as shown below.

C:\Program Files; C:\Winnt; ...... ; C:\proto3\bin
(you need to add ; C:\proto3\bin  at the end)