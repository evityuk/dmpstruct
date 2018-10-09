# dmpstruct
[![GoDoc](https://godoc.org/github.com/evityuk/dmpstruct?status.svg)](https://godoc.org/github.com/evityuk/dmpstruct)

Go try by writing utility package and cmd for dumping structs to map[string]interface{}

```Go
import "github.com/evityuk/dmpstruct"
```

## Usage

See [cmd/main.go]() for details

```Go
package main                                     

import (                                                                           
  logrus "github.com/sirupsen/logrus"                                              
  "os"                                                                             
  "evityuk/dmpstruct"                                                              
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
                                                                                   
  dump, err := dmpstruct.Dump(s)                                                   
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

1. Add more test cases
5. Should tag be dumped?



## Testing 
From project root run 
```
>go test
```
