package main

import (
  "fmt"
  "os"
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
    ans :=int64(0)
    var x,y,z string
    var n int64
    _, err = fmt.Fscanln(fileHandle, &n)
    _, err = fmt.Fscanln(fileHandle, &x)
    _, err = fmt.Fscanln(fileHandle, &y)
    _, err = fmt.Fscanln(fileHandle, &z)

    if n%2==0 && y[n-1]=='.' && z[n-1]=='.' && x[0]=='.' && y[0]=='.'{
      ans = 1
          for j:=int64(1) ; j+1<n-1; j=j+2{
            anss := int64(0)
            if y[j]=='#' || y[j+1]=='#'{
              ans = 0
              break
            }
            if x[j]=='.' && x[j+1]=='.'{
              anss++
            }
            if z[j]=='.' && z[j+1]=='.'{
              anss++
            }
            if anss == 0{
              ans = 0
              break
            }
            ans = (ans * anss)%1000000007
          }
    }

    if err != nil{
      fmt.Println(err)
    }
    _, err = fmt.Fprintf(f, "Case #%v: ",i+1)
    _, err = fmt.Fprintln(f, ans)
    if err != nil{
      fmt.Println(err)
    }
  }

}