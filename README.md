# ASSESSMENT-TEST #

API Using Golang (Iris Framework) as Programming Language

## Directory structure
Your project directory structure should look like this
```
  + your_gopath/
  |
  +--+ src/
  |  |
  |  +--+ test-sandri/
  |     |
  |     +--+ main.go
  |        + model/
  |        + controller/
  |        + view/
  |        + config/
  |        + ... any other source code
  |
  |
  |
  +--+ bin/
  |  |
  |  +-- ... executable file
  |
  +--+ pkg/
     |
     +-- ... all dependency_library required

```

## Setup and Build

Setup Golang <https://golang.org/>

Under `$GOPATH`, do the following command :
```
  cd src/
  git clone <url>
```

## Running Application
  ```
  cd to $GOPATH/src/test-sandri
  go build
  nohup ./test-sandri &
  if you want to kill the process, type lsof -i :8089
  kill PID_NUMBER
  ```
