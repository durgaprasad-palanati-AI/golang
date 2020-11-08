package main
import "fmt"

func main(){
var a=17
for i:=2;i<a;i++{
if a%i==0{
fmt.Printf("not prime\n")
break
}else if i==a-1{
fmt.Printf("PRIME \n")
}
}
}
