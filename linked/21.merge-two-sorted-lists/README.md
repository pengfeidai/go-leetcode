# [21. Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/)

## 题目

Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example :

```
Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4

```

## 题目大意

合并 2 个有序链表

## 解题思路

1. 递归法
  时间复杂度 O（m+n）
  空间复杂度 O（m+n）

2. 迭代法——维护一个prehead哨兵节点
  时间复杂度 O（m+n）
  空间复杂度 O（1）