### 题目：

[266. Palindrome Permutation](https://leetcode.com/problems/palindrome-permutation/)

Given a string, determine if a permutation of the string could form a palindrome.

**Example 1:**

```
Input: "code"
Output: false
```

**Example 2:**

```
Input: "aab"
Output: true
```

**Example 3:**


```
Input: "carerac"
Output: true
```

**Hint:**

Consider the palindromes of odd vs even length. What difference do you notice?
Count the frequency of each character.
If each character occurs even number of times, then it must be a palindrome. How about character which occurs odd number of times?



[266. 回文排列](https://leetcode-cn.com/problems/palindrome-permutation/)

描述
中文
English
给定一个字符串，判断字符串是否存在一个排列是回文排列。

**样例1**

```
输入: s = "code"
输出: False
解释:
没有合法的回文排列
```

**样例2**


```
输入: s = "aab"
输出: True
解释:
"aab" --> "aba"
```

**样例3**

```
输入: s = "carerac"
输出: True
解释:
"carerac" --> "carerac"
```

### 思路：

每个字母出现次数数是偶数，可以形成回文串；出现一个奇数次的字母，并且字母总数长度为奇数，也可以形成回文串。
