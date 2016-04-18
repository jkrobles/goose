# Goose
A IRC client library written in Golang 

Build Status: [![Build Status](https://travis-ci.org/jkrobles/goose.svg?branch=master)](https://travis-ci.org/jkrobles/goose)

## Examples ##
Import goose
```go
import (
    "github.com/jkrboles/goose"
)
```

Anywhere in your code you can create a Goose object. You can then set the parameters, most importantly being Server, Nick and Port. If Channels is empty it will not attempt to join on Connect(). You can issue the Join(channel) function later which will update the Channels array.
```go
    goose := new(goose.Goose)
    goose.Server = "irc.freenode.com"
    goose.Port = "6667"
    goose.Nick = "GooseBot" 
    goose.User = "Goose IRC Client in Go"
    goose.Channels = []string{"#goose"}
    goose.Password = "test" 
    goose.Connect()
```

Below is an example of a goose loop and how to get the line from the Reader. goose.Pings(line) is required to be ran during every loop iteration. This checks if the current line is a PING and responds with a PONG. Instead of printing the line you can do anything with it like logging. 
```go
    for ; goose.Reading() == true;  { 
        line, err := goose.Reader.ReadLine()
        if err != nil { 
           fmt.Println(err)
            break // break loop on errors
         } 
    
         goose.Pings(line)
         fmt.Printf("%s\n", line)

    }
```



## Contributing ##
If you are interested in contributing to this project whether it be providing test servers or debugging issues you may contact me at justice@justicerobles.com. I look forward to any help.
