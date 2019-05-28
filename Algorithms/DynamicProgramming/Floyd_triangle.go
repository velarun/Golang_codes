package main
import "fmt"
 
func main(){
    var rows int = 4
    var temp int = 1
 
    for i := 1; i <= rows; i++ { 
         
        for k := 1; k <= i; k++ {
 
            fmt.Printf(" %d",temp)              
            temp++
        }
        fmt.Println("")     
    }
 
}
