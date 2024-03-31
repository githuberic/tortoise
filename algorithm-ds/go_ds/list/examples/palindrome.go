package examples

import "algorithm-ds/go_ds/list"

/*
思路1：开一个栈存放链表前半段
时间复杂度：O(N)
空间复杂度：O(N)
*/
func isPalindrome1(l *list.SingleLinkedList) bool {
	lLen := l.Length
	if lLen == 0 {
		return false
	}
	if lLen == 1 {
		return true
	}

	s := make([]string, 0, lLen/2)
	cur := l.Head
	for i := uint(1); i <= lLen; i++ {
		cur = cur.Next
		if lLen%2 != 0 && i == (lLen/2+1) { //如果链表有奇数个节点，中间的直接忽略
			continue
		}
		if i <= lLen/2 { //前一半入栈
			s = append(s, cur.GetValue().(string))
		} else { //后一半与前一半进行对比
			if s[lLen-i] != cur.GetValue().(string) {
				return false
			}
		}
	}

	return true
}

/*
思路2
找到链表中间节点，将前半部分转置，再从中间向左右遍历对比
时间复杂度：O(N)
*/
func isPalindrome2(l *list.SingleLinkedList) bool {
	lLen := l.Length
	if lLen == 0 {
		return false
	}
	if lLen == 1 {
		return true
	}
	var isPalindrome = true
	step := lLen / 2
	var pre *list.Node = nil
	cur := l.Head.Next
	next := l.Head.Next.Next
	for i := uint(1); i <= step; i++ {
		tmp := cur.GetNext()
		cur.Next = pre
		pre = cur
		cur = tmp
		next = cur.GetNext()
	}
	mid := cur

	var left, right *list.Node = pre, nil
	if lLen%2 != 0 { //
		right = mid.GetNext()
	} else {
		right = mid
	}

	for nil != left && nil != right {
		if left.GetValue().(string) != right.GetValue().(string) {
			isPalindrome = false
			break
		}
		left = left.GetNext()
		right = right.GetNext()
	}

	//复原链表
	cur = pre
	pre = mid
	for nil != cur {
		next = cur.GetNext()
		cur.Next = pre
		pre = cur
		cur = next
	}
	l.Head.Next = pre

	return isPalindrome
}
