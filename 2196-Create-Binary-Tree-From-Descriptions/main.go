package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	mapDsc := map[int][][]int{}
	childs := map[int]struct{}{}
	for i := range descriptions {
		childs[descriptions[i][1]] = struct{}{}
		if _, ok := mapDsc[descriptions[i][0]]; ok {
			mapDsc[descriptions[i][0]] = append(mapDsc[descriptions[i][0]], descriptions[i])
			continue
		}
		mapDsc[descriptions[i][0]] = [][]int{descriptions[i]}
	}
	var treeHead int
	for k := range mapDsc {
		if _, ok := childs[k]; !ok {
			treeHead = k
		}
	}

	return addNode(&TreeNode{
		Val: treeHead,
	}, mapDsc)
}

func addNode(node *TreeNode, mapDesc map[int][][]int) *TreeNode {
	nodeDesc, ok := mapDesc[node.Val]
	if !ok {
		return node
	}
	for i := range nodeDesc {
		if nodeDesc[i][2] == 1 {
			node.Left = addNode(&TreeNode{Val: nodeDesc[i][1]}, mapDesc)

			continue
		}
		node.Right = addNode(&TreeNode{Val: nodeDesc[i][1]}, mapDesc)
	}

	return node
}
