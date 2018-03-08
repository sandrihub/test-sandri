# ASSESSMENT-TEST #

API Using Golang (Iris Framework) as Programming Language

## Directory structure
Your project directory structure should look like this
```
  + your_gopath/
  |
  +--+ src/
  |  |
  |  +--+ assessment/
  |     |
  |     +--+ assessment-test/
  |        |
  |        +--+ main.go
  |           + model/
  |           + controller/
  |           + view/
  |           + ... any other source code
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
  mkdir -p src/assessment/
  cd src/assessment/
  git clone <url>
  rename directory to api-server
```

## Running Application
  ```
  cd to $GOPATH/src/assessment/assessment-test
  go build
  nohup ./assessment-test &
  if you want to kill the process, type lsof -i :8080
  kill PID_NUMBER
  ```
