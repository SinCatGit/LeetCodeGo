### 题目：

[60. Permutation Sequence](https://leetcode.com/problems/permutation-sequence/)

The set [1,2,3,...,n] contains a total of n! unique permutations.

By listing and labeling all of the permutations in order, we get the following sequence for n = 3:


```
1. "123"
2. "132"
3. "213"
4. "231"
5. "312"
6. "321"
```

Given n and k, return the kth permutation sequence.

**Note:**

- Given n will be between 1 and 9 inclusive.
- Given k will be between 1 and n! inclusive.

**Example 1:**

```
Input: n = 3, k = 3
Output: "213"
```

**Example 2:**

```
Input: n = 4, k = 9
Output: "2314"
```


[60. 第k个排列](https://leetcode.com/problems/permutation-sequence/)

给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。

按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：

```
1. "123"
2. "132"
3. "213"
4. "231"
5. "312"
6. "321"
```
给定 n 和 k，返回第 k 个排列。

**说明：**

给定 n 的范围是 [1, 9]。
给定 k 的范围是[1,  n!]。

**示例 1:**


```
输入: n = 3, k = 3
输出: "213"
```

**示例 2:**


```
输入: n = 4, k = 9
输出: "2314"
```


### 思路：

举例来说明。

n = 4， 参与排列的数字「1， 2， 3， 4」

列出所有的排列

1 + (permutations of 2, 3, 4)

2 + (permutations of 1, 3, 4)

3 + (permutations of 1, 2, 4)

4 + (permutations of 1, 2, 3)

n个数字的排列数为n!,那么3个数的排列数为6。假如k=14，那么第14个排列落在

3 + (permutations of 1, 2, 4)

令k=14-1(减去1是因为程序中索引从0开始), k/(n-1)!= 13/(4-1)! = 2, 在数列「1， 2， 3， 4」中索引为2的数字为3，所以第一个数字为3。

那么问题进一步缩小为「1， 2， 4」数字的排列，求第 k= k%(n-1)!=13%(4-1)=1 个排列：

1 + (permutations of 2, 4)

2 + (permutations of 1, 4)

4 + (permutations of 1, 2)

在这一步中，2个数字排列有2!， 总共有3*2!=6个，我们寻找第一个排列，那么落在

1 + (permutations of 2, 4)

此时 k/(n-2)! = 1/(4-2)! = 0, 即「1， 2， 4」中索引0的数字1。目前我们知道前面两个数字3，1。剩下的数字依次类推。

「2, 4」

k = k % (n-2)! = 1%(4-2)! = 1，第三个数字在「2， 4」中的索引为 k/(n-3)!= 1/(4-3)! = 1，所以第三个数字为4

「2」

k = k % (n-3)! = 1%(4-3)! = 0，第四个数字在「2」中的索引为 k/(n-4)!= 0/(4-4)! = 0，所以第四个数字为2


参考：https://leetcode.com/problems/permutation-sequence/discuss/22507/%22Explain-like-I'm-five%22-Java-Solution-in-O(n)
