package main

import (
  "fmt"
  "os"
  "strings"
)

func min(x, y int) int {
  if x<y{
    return x
  }
  return y
}

func main() {

  fileHandle, err := os.Open("/Users/ankurpanwar/Downloads/ip.txt")
  if err != nil{
    fmt.Println(err)
  }
  defer fileHandle.Close()
  f, err := os.Create("/Users/ankurpanwar/Downloads/op.txt")
  if err != nil{
    fmt.Println(err)
  }
  defer f.Close()
  var t int
  _, err = fmt.Fscanln(fileHandle, &t)
  if err != nil{
    fmt.Println(err)
  }

  for i:=0 ; i<t; i++{
    var n,k,v int
    var ls []string
    _, err = fmt.Fscanln(fileHandle, &n,&k,&v)
    if err != nil{
      fmt.Println(err)
    }
    for j:=0; j< n;j++{
      var s string
      _, err = fmt.Fscanln(fileHandle, &s)
      if err != nil{
        fmt.Println(err)
      }
      ls=append(ls,s)
    }

    x := k*(v-1)
    x = x%n
    lss := ls[x:min(x+k,len(ls))]
    _, err = fmt.Fprintf(f, "Case #%v: ",i+1)
    if err != nil{
      fmt.Println(err)
    }
    if len(lss) < k{
     lss =  append(ls[:k-len(lss)], lss...)
    }
    _, err = fmt.Fprintln(f, strings.Join(lss," "))
    if err != nil{
      fmt.Println(err)
    }
  }
}