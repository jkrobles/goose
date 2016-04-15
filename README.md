# Goose
A IRC client library written in Golang 

Build Status: [![Build Status](https://travis-ci.org/jkrobles/goose.svg?branch=master)](https://travis-ci.org/jkrobles/goose)

## Examples ##

```go
import (
    "goose"
    "fmt"
    "strings"
)

func main(){ 
    goose := new(goose.Goose)
    goose.Server = "irc.freenode.com"
    goose.Port = "6667"
    goose.Nick = "GooseBot" 
    goose.User = "Goose IRC Client in Go"
    goose.Channels = []string{"#goose"}
    goose.Password = "test" 
    goose.Connect()

    for ; goose.Reading() == true;  { 
        line, err := goose.Reader.ReadLine()
        if err != nil { 
           fmt.Println(err)
            break // break loop on errors
         } 
    
         goose.Pings(line)
         l := strings.Split(line," ")
        if l[1] == "MODE" {
            fmt.Println(goose.Channels[0])
            goose.Join("#goose")
        }
         fmt.Printf("%s\n", line)

    }



```



## Contributing ##
If you are interested in contributing to this project whether it be providing test servers or debugging issues you may contact me at justice@justicerobles.com. I look forward to any help.
