### 题目：

[267. Palindrome Permutation II](https://leetcode.com/problems/palindrome-permutation-ii/)

Given a string s, return all the palindromic permutations (without duplicates) of it. Return an empty list if no palindromic permutation could be form.

**For example:**


```
Given s = "aabb", return ["abba", "baab"].

Given s = "abc", return [].
```


**Hint:**

If a palindromic permutation exists, we just need to generate the first half of the string.
To generate all distinct permutations of a (half of) string, use a similar approach from: Permutations II or Next Permutation.


[267. 回文排列II
](https://leetcode-cn.com/problems/palindrome-permutation-ii/)


**描述**

给定一个字符串s，返回所有回文排列(不重复)。如果没有回文排列，则返回空列表。

**样例1**

```
输入: s = "aabb"
输出: ["abba","baab"]
```

**样例2**

```
输入: "abc"
输出: []
```

### 思路：

如果回文全排列存在，只需要生成前半段字符串即可，后面的直接根据前半段得到。

进一步，由于回文字符串有奇偶两种情况，偶数回文串例如abba，可以分成前（ab）中（‘’）后（ba）三段，中间一段为空字符串；而奇数回文串例如abcba，分成前（ab）中（c）后（ba）三段，需要注意的是中间部分只能是一个字符。

由此可以分析得出，如果一个字符串的回文字符串要存在，那么奇数个的字符只能有0个或1个，其余的必须是偶数个。

所以我们可以用哈希表来记录所有字符的出现个数，然后找出出现奇数次数的字符加入mid中，如果有两个或两个以上的奇数个数的字符，那么返回空集。

对于每个字符，不管其奇偶，都将其个数除以2的个数的字符加入path中，这样做的原因是如果是偶数个，那么将其一般加入t中，如果是奇数，如果有1个，那么除以2是0，不会有字符加入path，如果是3个，那么除以2是1，取一个加入path。

等获得了path之后，path是就是前半段字符。path加上mid和该path的逆序列就是一种所求的回文字符串。

这样就可得到所有的回文全排列。

