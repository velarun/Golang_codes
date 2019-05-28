package main
  
import (
    "fmt"
)
  
func uniqueInt(intSlice []int) []int {
    keys := make(map[int]bool)
    list := []int{} 
    for _, entry := range intSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }    
    return list
}

func uniqueStr(strSlice []string) []string {
    keys := make(map[string]bool)
    list := []string{} 
    for _, entry := range strSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }    
    return list
}

func removeDuplicates(elements []int) []int {
    result := []int{}

    for i := 0; i < len(elements); i++ {
        // Scan slice for a previous element of the same value.
        exists := false
        for v := 0; v < i; v++ {
            if elements[v] == elements[i] {
                exists = true
                break
            }
        }
        // If no previous element exists, append this one.
        if !exists {
            result = append(result, elements[i])
        }
    }
    return result
}
  
func main() {

    intSlice := []int{1,5,3,6,9,9,4,2,3,1,5}
    strSlice := []string{"cat", "dog", "cat", "bird"}
    
    fmt.Println("Integer Array Before Removing Duplicates: ", intSlice) 
    uniqueSlice := uniqueInt(intSlice)
    fmt.Println("Integer Array After Removing Duplicates: ", uniqueSlice)

    fmt.Println("String Array Before Removing Duplicates: ", strSlice)
    uniqueStrSlice := uniqueStr(strSlice)
    fmt.Println("String Array After Removing Duplicates: ", uniqueStrSlice)

    fmt.Println("Integer Array Before Removing Duplicates: ", intSlice) 
    uniqueIntSlice := removeDuplicates(intSlice)
    fmt.Println("Integer Array After Removing Duplicates: ", uniqueIntSlice)

}
