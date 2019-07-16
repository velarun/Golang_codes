// A Java program to find 
// the deepest left leaf 
// in a binary tree 

// A Binary Tree node 
class Node 
{ 
	int data; 
	Node left, right; 

	// Constructor 
	public Node(int data) 
	{ 
		this.data = data; 
		left = right = null; 
	} 
} 

// Class to evaluate pass 
// by reference 
class Level 
{ 
	// maxlevel: gives the 
	// value of level of 
	// maximum left leaf 
	int maxlevel = 0; 
} 

class BinaryTree 
{ 
	Node root; 
	
	// Node to store resultant 
	// node after left traversal 
	Node result; 

	// A utility function to 
	// find deepest leaf node. 
	// lvl: level of current node. 
	// isLeft: A bool indicate 
	// that this node is left child 
	void deepestLeftLeafUtil(Node node, 
							int lvl, 
							Level level, 
							boolean isLeft) 
	{ 
		// Base case 
		if (node == null) 
			return; 

		// Update result if this node 
		// is left leaf and its level 
		// is more than the maxl level 
		// of the current result 
		if (isLeft != false && 
			node.left == null && 
			node.right == null && 
			lvl > level.maxlevel) 
		{ 
			result = node; 
			level.maxlevel = lvl; 
		} 

		// Recur for left and right subtrees 
		deepestLeftLeafUtil(node.left, lvl + 1, 
							level, true); 
		deepestLeftLeafUtil(node.right, lvl + 1, 
							level, false); 
	} 

	// A wrapper over deepestLeftLeafUtil(). 
	void deepestLeftLeaf(Node node) 
	{ 
		Level level = new Level(); 
		deepestLeftLeafUtil(node, 0, level, false); 
	} 
	
	// Driver program to test above functions 
	public static void main(String[] args) 
	{ 
		BinaryTree tree = new BinaryTree(); 
		tree.root = new Node(1); 
		tree.root.left = new Node(2); 
		tree.root.right = new Node(3); 
		tree.root.left.left = new Node(4); 
		tree.root.right.left = new Node(5); 
		tree.root.right.right = new Node(6); 
		tree.root.right.left.right = new Node(7); 
		tree.root.right.right.right = new Node(8); 
		tree.root.right.left.right.left = new Node(9); 
		tree.root.right.right.right.right = new Node(10); 

		tree.deepestLeftLeaf(tree.root); 
		if (tree.result != null) 
			System.out.println("The deepest left child"+ 
							" is " + tree.result.data); 
		else
			System.out.println("There is no left leaf in"+ 
							" the given tree"); 
	} 
} 

// This code has been contributed by Mayank Jaiswal(mayank_24) 
