package main
  
import (
    "fmt"
    "os"
    "io"
)

type BinaryNode struct {
    left  *BinaryNode
    right *BinaryNode
    data  int64
}

func (n *BinaryNode) insert(data int64) {
    if n == nil {
        return
    } else if data <= n.data {
        if n.left == nil {
            n.left = &BinaryNode{data: data, left: nil, right: nil}
        } else {
            n.left.insert(data)
        }
    } else {
        if n.right == nil {
            n.right = &BinaryNode{data: data, left: nil, right: nil}
        } else {
            n.right.insert(data)
        }
    }   
}

func print(w io.Writer, node *BinaryNode, ns int, ch rune) {
    if node == nil {
        return
    }
  
    for i := 0; i < ns; i++ {
        fmt.Fprint(w, " ")
    }
    fmt.Fprintf(w, "%c:%v\n", ch, node.data)
    print(w, node.left, ns+2, 'L')
    print(w, node.right, ns+2, 'R')
}

func printCommon(root1, root2 *BinaryNode) {

	s1 := []*BinaryNode{} 
	s2 := []*BinaryNode{} 

	for { 
		if root1 != nil { 
			s1 = append(s1, root1) 
			root1 = root1.left 
		} else if root2 != nil { 
			s2 = append(s2, root2) 
			root2 = root2.left 
		} else if len(s1) != 0 && len(s2) != 0 { 
			root1 = s1[len(s1)-1]  
			root2 = s2[len(s2)-1] 
			
            if root1.data == root2.data {  
                fmt.Print(root1.data, " ")  
                s1 = s1[:len(s1)-1]
				s2 = s2[:len(s2)-1] 
				
                root1 = root1.right  
				root2 = root2.right  
            } else if root1.data < root2.data {    
                s1 = s1[:len(s1)-1]
                root1 = root1.right  
				root2 = nil  
            } else if root1.data > root2.data {  
                s2 = s2[:len(s2)-1] 
                root2 = root2.right  
				root1 = nil
            }  
		} else { 
			break
		}
	}
	fmt.Println()
}

func main() {
	root1 := &BinaryNode{data: 5}
	root1.insert(1) 
	root1.insert(10) 
	root1.insert(0) 
	root1.insert(4) 
	root1.insert(7) 
	root1.insert(9) 
	
	fmt.Println("Tree:")
	print(os.Stdout, root1, 0, 'M')

	root2 := &BinaryNode{data: 10} 
	root2.insert(7) 
	root2.insert(20) 
	root2.insert(4) 
	root2.insert(9) 

	fmt.Println("Tree:")
	print(os.Stdout, root2, 0, 'M')

	fmt.Println("Common Elements between two Binary Search Tree:")
	printCommon(root1, root2)
}

