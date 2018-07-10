package main

import (
  "fmt"
  "os"
  "sort"
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
    var n,m int64
    ls :=make([]int,m*n)
    seen:= make([]bool,m*n)

    _, err = fmt.Fscanln(fileHandle, &n, &m)
    if err != nil{
      fmt.Println(err)
    }
    for j:=0; j<int(n);j++{
      x:=make([]int,m)
      y := make([]interface{}, len(x))
      for i := range x {
        y[i] = &x[i]
      }
      _, err = fmt.Fscanln(fileHandle, y...)
      if err != nil{
        fmt.Println(err)
      }
      z := make([]bool,m)
      sort.Ints(x)
      ls= append(ls,x...)
      seen=append(seen,z...)
    }

    ans, tmp, ind, in := int64(0),int64(1000000000010),int64(-1),int64(-1)
    for l:=int64(0); l<n ;l++{
      var j int64
    for j=int64(0); j<=l; j++ {
      for k := int64(0); k < n; k++ {
        if seen[j*m+k] == false {
          if tmp > (k+1)*(k+1)-k*k+int64(ls[j*m+k]) {
            ind = k
            in = j
            tmp = (k+1)*(k+1) - k*k + int64(ls[j*m+k])
          }
          break
        }
      }
    }
      ans+=tmp
      seen[in*m+ind]=true
      tmp = 1000000000010
      ind = -1
    }

    _, err = fmt.Fprintf(f, "Case #%v: ",i+1)
    _, err = fmt.Fprintln(f, ans)
    if err != nil{
      fmt.Println(err)
    }
  }
}