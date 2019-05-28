package main
  
import (
    "fmt"
    "os"
    "io"
)

var res []int64

type BinaryNode struct {
    left  *BinaryNode
    right *BinaryNode
    data  int64
}
  
type BinaryTree struct {
    root *BinaryNode
}
  
func (t *BinaryTree) insert(data int64) *BinaryTree {
    if t.root == nil {
        t.root = &BinaryNode{data: data, left: nil, right: nil}
    } else {
        t.root.insert(data)
    }
    return t
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

func inorderTraversal(node *BinaryNode, res []int64) []int64 {
    if node != nil {
       	res = inorderTraversal(node.left, res)
        res = append(res, node.data)
        res = inorderTraversal(node.right, res)
    }
    return res
}     

func preorderTraversal(node *BinaryNode, res []int64) []int64 {
    if node != nil {
        res = append(res, node.data)
        res = preorderTraversal(node.left, res)
        res = preorderTraversal(node.right, res)
    }
    return res
}

func postorderTraversal(node *BinaryNode, res []int64) []int64 {
    if node != nil {
        res = postorderTraversal(node.left, res)
        res = postorderTraversal(node.right, res)
        res = append(res, node.data)
    }
    return res
}

func main() {
    tree := &BinaryTree{}
    tree.insert(27).
        insert(14).
        insert(35).
        insert(10).
        insert(19).
        insert(31).
        insert(42)
    
    fmt.Println("Tree:")
    print(os.Stdout, tree.root, 0, 'M')
    fmt.Println("Inorder Traversal: ", inorderTraversal(tree.root, res))
    fmt.Println("Preorder Traversal: ", preorderTraversal(tree.root, res))
    fmt.Println("Postorder Traversal: ", postorderTraversal(tree.root, res))
}
