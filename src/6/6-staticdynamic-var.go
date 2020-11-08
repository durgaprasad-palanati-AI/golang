package main
import "fmt"
func main(){
var a, b, c = 3, 4, "foo"  
	
   fmt.Println("a=%d\n",a)
   fmt.Println("b=%d\n",b)
   fmt.Println("c=%d\n",c)
   
   fmt.Println("a=",a)
   fmt.Println("\nb=",b)
   fmt.Println("\nc=",c)
   
   fmt.Printf("a is of type %T\n", a)
   fmt.Printf("b is of type %T\n", b)
   fmt.Printf("c is of type %T\n", c)
}
