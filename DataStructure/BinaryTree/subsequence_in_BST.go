package main
  
import (
    "fmt"
    "os"
    "io"
)

type BinaryNode struct {
    left  *BinaryNode
    right *BinaryNode
    data  int
}

func (n *BinaryNode) insert(data int) {
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

func seqExistUtil(ptr *BinaryNode, seq []int, index *int) { 
	if ptr == nil { 
		return
	}
 
	seqExistUtil(ptr.left, seq, index) 

	if ptr.data == seq[*index] { 
		*index += 1
	}

	seqExistUtil(ptr.right, seq, index) 
}

func seqExist(root *BinaryNode, seq []int, n int) bool { 

	index := 0 
	seqExistUtil(root, seq, &index) 

	if index == n { 
		return true
	} else { 
		return false
	}
}

func main() {
	root1 := &BinaryNode{data: 8}
	root1.insert(10) 
	root1.insert(3) 
	root1.insert(6) 
	root1.insert(1) 
	root1.insert(4) 
	root1.insert(7) 
	root1.insert(14)
	root1.insert(13)
	
	fmt.Println("Tree:")
	print(os.Stdout, root1, 0, 'M')

	seq1 := []int{4, 6, 8, 14} 
	n1 := len(seq1) 
	if seqExist(root1, seq1, n1) { 
		fmt.Println("Yes") 
	} else { 
		fmt.Println("No")
	} 

	seq2 := []int{4, 6, 8, 12, 13}
	n2 := len(seq2) 
	if seqExist(root1, seq2, n2) { 
		fmt.Println("Yes") 
	} else { 
		fmt.Println("No")
	} 

}