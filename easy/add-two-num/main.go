package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	//Dummy For ans ListNode not use this value
	head := ListNode{0, nil}
	//Set to current
	current := &head

	//Set Pointer of function
	p := l1
	q := l2
	//Set carry
	carry := 0

	for {
		//Value of this loop
		v1 := 0
		v2 := 0

		//Check null if not null set value and set next pointer for next calculation
		if p != nil {
			v1 = p.Val
			p = p.Next
		}

		if q != nil {
			v2 = q.Val
			q = q.Next
		}

		//Sum of this loop
		sum := carry + v1 + v2
		//Set carry
		carry = sum / 10
		//Set Carry Value for checking later
		current.Next = &ListNode{sum % 10, nil}

		fmt.Println(sum % 10)
		//Set current Link to Next linkNode
		current = current.Next

		//If pointer is null then is final of calculation
		if p == nil && q == nil {
			break
		}
	}

	//Case of Carry over NextNode
	//ex 99 + 1 = -->[1]00
	if carry != 0 {
		current.Next = &ListNode{carry, nil}
	}

	return head.Next
}

func addTwoNumbersTwo(l1 *ListNode, l2 *ListNode) *ListNode {
	var result []*ListNode

	sum := l1.Val + l2.Val
	tempVal := 0
	if sum > 9 {
		tempVal = 1
		sum = sum - 10
	}
	result = append(result, &ListNode{
		Val: sum,
	})

	tempLinkedList1 := l1.Next
	tempLinkedList2 := l2.Next
	for tempLinkedList1 != nil || tempLinkedList2 != nil {
		if tempLinkedList1 == nil {
			tempSum := tempLinkedList2.Val + tempVal
			tempVal = 0
			if tempSum > 9 {
				tempVal = 1
				tempSum = tempSum - 10
			}
			result = append(result, &ListNode{
				Val: tempSum,
			})
			tempLinkedList2 = tempLinkedList2.Next
		} else if tempLinkedList2 == nil {
			tempSum := tempLinkedList1.Val + tempVal
			tempVal = 0
			if tempSum > 9 {
				tempVal = 1
				tempSum = tempSum - 10
			}
			result = append(result, &ListNode{
				Val: tempSum,
			})
			tempLinkedList1 = tempLinkedList1.Next
		} else {
			tempSum := tempLinkedList1.Val + tempLinkedList2.Val + tempVal
			tempVal = 0
			if tempSum > 9 {
				tempVal = 1
				tempSum = tempSum - 10
			}
			result = append(result, &ListNode{
				Val: tempSum,
			})
			tempLinkedList1 = tempLinkedList1.Next
			tempLinkedList2 = tempLinkedList2.Next
		}
	}

	if tempVal > 0 {
		result = append(result, &ListNode{
			Val: tempVal,
		})
	}

	for i := 0; i < len(result)-1; i++ {
		result[i].Next = result[i+1]
	}
	//     for _, linkedList := range result {
	//         fmt.Printf("%d\n", linkedList.Val)

	//     }
	return result[0]
}

func main() {

	a := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	b := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	ans := addTwoNumbers(a, b)
	for {
		fmt.Printf("%d", ans.Val)
		ans = ans.Next

		if ans == nil {
			break
		}
	}

}
