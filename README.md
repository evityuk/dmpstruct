# dmpstruct
Go try by writing utility package and cmd for dumping structs to map[string]interface{}

```Go
import "github.com/evityuk/dmpstruct"
```

## Usage

See [cmd/main.go]() for details

```Go
package main          
                      
import (              
  "fmt"               
  "yvitiuk/dmpstruct" 
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
                                                                                  
  dump, err := dmpstruct.Dump(s)                                                  
  if err != nil {                                                                 
    fmt.Println("Dump error: ", err)                                              
  } else {                                                                        
    fmt.Printf("Dumped: %q", dump)                                                
  }                                                                               
                                                                                  
```
Results:
```
Dumped: 
map["Name":"Dan" "age":"Field 'age' of type 'int8' is unexported" "Occupation":map["Name":"Boston" "code":"Field 'code' of type 'uint' is unexported"] "Employee":map["department":"Field 'department' of type 'string' is unexported" "position":"Field 'position' of type 'string' is unexported"]]
```



## TBD

1. Add more test cases
2. Add dump options: log level(or by env var), replace unexported message with pkg's const
3. Add documentation 
4. Change fmt.Println to golang.org/x/log 
5. Enhance cmd/main output(and possibly \_test.go output)
5. Should tag be dumped?



## Testing 
From project root run 
```
>go test
```
