# [206. Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/description/)

## 题目

Reverse a singly linked list.

**Example:**

```go
Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
```

**Follow up:**

A linked list can be reversed either iteratively or recursively. Could you implement both?


## 题目大意

翻转单链表

## 解题思路

- 定义两个指针： prepre 和 curcur ；prepre 在前 curcur 在后。
- 每次让 prepre 的 nextnext 指向 curcur ，实现一次局部反转
- 局部反转完成之后，prepre 和 curcur 同时往前移动一个位置
- 循环上述过程，直至 prepre 到达链表尾部