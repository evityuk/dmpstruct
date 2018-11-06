# dmpstruct
[![Build Status](https://travis-ci.com/evityuk/dmpstruct.svg?branch=master)](https://travis-ci.com/evityuk/dmpstruct)
[![GoDoc](https://godoc.org/github.com/evityuk/dmpstruct?status.svg)](https://godoc.org/github.com/evityuk/dmpstruct)

Go try by writing utility package and cmd for dumping structs to map[string]interface{}

```Go
import "github.com/evityuk/dmpstruct"
```

*Check next packages for more mature implementations:*
 - https://github.com/davecgh/go-spew(+++)
 - https://github.com/fatih/structtag(++)
 - https://github.com/fatih/gomodifytags(structtag modification - "offtopic one")

## Usage

See [cmd/main.go]() for details

```Go
package main                                     

import (                                                                           
  logrus "github.com/sirupsen/logrus"                                              
  "os"                                                                             
  "github.com/evityuk/dmpstruct"                                                              
)                                                                                  
                                                                                   
type Employee struct {                                                             
  department string                                                                
  position   string                                                                
}                                                                                  
                                                                                   
type OccupationInfo struct {                                                       
  Name string                                                                      
  code uint                                                                        
}                                                                                  
                                                                                   
type S struct {                                                                    
  Name       string                                                                
  age        int8                                                                  
  Occupation OccupationInfo                                                        
  Employee                                                                         
}                                                                                  
                                                                                   
func main() {                                                                      
  s := S{"Dan", 50, OccupationInfo{"Boston", 3}, Employee{"Literature", "Writer"}} 
                                                                                   
  dmpstruct.Init(os.Stdout, &logrus.TextFormatter{}, logrus.DebugLevel)            
                                                                                   
  dump, err := dmpstruct.Dump(&s)                                                   
  if err != nil {                                                                  
    logrus.Println("Dump error: ", err)                                            
  } else {                                                                         
    logrus.WithFields(dump).Println("Dumped successfully: ")                       
  }                                                                                

}                                                                                  
```
Results:
```
Dumped: 
INFO[0000] Dumped successfully:                          
Employee="map[department:Field \"department\" of type \"string\" is unexported position:Field \"position\" of type \"string\" is unexported]" Name=Dan Occupation="map[Name:Boston code:Field \"code\" of type \"uint\" is unexported]" age="Field \"age\" of type \"int8\" is unexported"
```

## TBD

1. Should tag be dumped?
2. Should cmd be able to dump package by cmd-line url?


## Testing 
From project root run 
```
>go test
```
