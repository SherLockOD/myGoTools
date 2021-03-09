package main

import (
	"fmt"
)

/*
给出由小写字母组成的字符串 S，重复项删除操作会选择两个相邻且相同的字母，并删除它们。

在 S 上反复执行重复项删除操作，直到无法继续删除。

在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。

 

示例：

输入："abbaca"
输出："ca"
解释：
例如，在 "abbaca" 中，我们可以删除 "bb" 由于两字母相邻且相同，这是此时唯一可以执行删除操作的重复项。之后我们得到字符串 "aaca"，其中又只有 "aa" 可以执行重复项删除操作，所以最后的字符串为 "ca"。
 

提示：

1 <= S.length <= 20000
S 仅由小写英文字母组成。

分析：
使用栈的数据结构来解决，假设有字符串 abbc，遍历字符串每次都往栈内放入一个元素
1)    v
     | |
     | |
     |a|       stack = append(stack, a)
     |-|

2)    v
     | |
     |b|
     |a|       stack = append(stack, b)
     |-|

3) 放入的元素如果与栈的元素一致，就消除掉，有点消消乐的感觉。观察发现，每次新元素都是与栈顶的元素进行判断，
如果相同，就将栈顶元素取消，也就剩下栈顶以下的元素。

      v
     | |   --> b
     | |   --> b     stack = stack[:len(stack)-1]
     |a|
     |-|


4)    v
     | |
     |c|
     |a|       stack = append(stack, c)
     |-|

应用场景：
1. 反向思路：只留下相同字符
2. 括号配对
*/

func removeDuplicates(S string) string {
	stack := make([]byte,len(S)*8)
	for i := range S {
		if len(stack) > 0 && stack[len(stack)-1] == S[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, S[i])
		}
	}
	return string(stack)
}

func main() {
	s := "abc"
	fmt.Printf("%T\n", s[0])
	fmt.Println(removeDuplicates("abbcde"))
}
