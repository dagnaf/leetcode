# Merge k Sorted Lists

## Description

```
 Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.
```

link: https://leetcode.com/problems/merge-k-sorted-lists/description/

difficulty: Hard

## Discussion

```
归并排序，每个链表设置一个指针。每次选取一个指针加入输出结果，选取过程由优先队列维护。
题目中没有指明所有链表都以ASC排序，另外结果需要以ASC排序。
priority_queue 的排序依据，方便记忆：
  假如有优先队列q，其中包含数据{a_1, a_2, ... }。
  根据cmp函数对象调用后数据可以按照其所在下标放置。
  下标越小表示其优先级越低，下标越大表示其优先级越高。
  例如，按照less<int>()排序后，最小的数字放在0位置，其优先级最低。
priority_queue 模版参数<type, container, comparator>
  其中comparator可以是decltype(lambda), less<>, greater<>, function<>
```