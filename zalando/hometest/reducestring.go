package hometest

type Node struct {
	Rune rune
	Next *Node
	Prev *Node
}

func areRemovable(first, second rune) bool {
	return first == 'A' && second =='B' ||
		first == 'B' && second == 'A' ||
		first == 'C' && second == 'D' ||
		first == 'D' && second == 'C'
}

func removeTwoNodes(first *Node) (*Node, *Node) {
	second := first.Next
	before := first.Prev
	after := second.Next
	if before != nil {
		before.Next = after
	}
	if after != nil {
		after.Prev = before
	}
	first.Prev, first.Next = nil, nil
	second.Prev, second.Next = nil, nil
	return before, after
}

func convertToString(dummyHead *Node) string {
	var runes []rune
	for node := dummyHead.Next; node != nil; node = node.Next {
		runes = append(runes, node.Rune)
	}
	return string(runes)
}

func ReduceString(S string) string {
	// 1. Build a bidirectional linked list of letters
	list := &Node {  // dummy node
		Rune: -1,
		Prev: nil,
		Next: nil,
	}
	// Queue of nodes in list which were somehow changed
	var changedQueue []*Node
	lastInList := list
	for _, r := range S {
		node := &Node{
			Rune: r,
			Prev: lastInList,
			Next: nil,
		}
		lastInList.Next = node
		lastInList = node
		changedQueue = append(changedQueue, node)
	}
	// Keep dummy node which is not removable
	// list = list.Next

	// 2. Traverse it, removing pairs, remembering new nodes in some
	//    other structure (e.g. slice/queue/stack)
	// 3. Try to check any new char in this stack for removal

	for len(changedQueue) > 0 {
		var node *Node
		node, changedQueue = changedQueue[0], changedQueue[1:]
		if node == nil {
			continue
		}
		if node.Prev != nil && areRemovable(node.Rune, node.Prev.Rune) {
			before, after := removeTwoNodes(node.Prev)
			changedQueue = append(changedQueue, before, after)
		}
		if node.Next != nil && areRemovable(node.Rune, node.Next.Rune) {
			before, after := removeTwoNodes(node)
			changedQueue = append(changedQueue, before, after)
		}
	}

	return convertToString(list)
}

