/**
 * Definition for a rope tree node.
 * type RopeTreeNode struct {
 * 	   len   int
 * 	   val   string
 * 	   left  *RopeTreeNode
 * 	   right *RopeTreeNode
 * }
 */
func getKthCharacter(root *RopeTreeNode, k int) byte {
    a := find(root)
    return a[k - 1]
}
func find(root *RopeTreeNode) string {
    if root == nil {
        return ""
    }
    if root.len > 0 {
        return find(root.left) + find(root.right)
    }
    return root.val
}