package main

import "fmt"

/*
题目：
请你仅使用两个栈实现先入先出队列。队列应当支持一般队列的支持的所有操作（push、pop、peek、empty）：

实现 MyQueue 类：

void push(int x) 将元素 x 推到队列的末尾
int pop() 从队列的开头移除并返回元素
int peek() 返回队列开头的元素
boolean empty() 如果队列为空，返回 true ；否则，返回 false
 

说明：

你只能使用标准的栈操作 —— 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
 

进阶：

你能否实现每个操作均摊时间复杂度为 O(1) 的队列？换句话说，执行 n 个操作的总时间复杂度为 O(n) ，即使其中一个操作可能花费较长时间。
 

示例：

输入：
["MyQueue", "push", "push", "peek", "pop", "empty"]
[[], [1], [2], [], [], []]
输出：
[null, null, null, 1, 1, false]

解释：
MyQueue myQueue = new MyQueue();
myQueue.push(1); // queue is: [1]
myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
myQueue.peek(); // return 1
myQueue.pop(); // return 1, queue is [2]
myQueue.empty(); // return false
 

提示：

1 <= x <= 9
最多调用 100 次 push、pop、peek 和 empty
假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）


分析：
1. 因为 golang 没有栈的数据结构，所以需要用 slice 来模拟栈，再用两个栈模拟队列，有点多费一道手的感觉，但是没办法题目要求。
2. 我们要用两个 slice 一个模拟入栈，一个模拟出栈，push 用于入栈，pop，peek，用于出栈，中间要加一步，循环从入栈slice转移到出栈slice的过程。
 */


type MyQueue struct {
    inStack, outStack []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
    return MyQueue{}
}


/** Push element x to the back of queue. */
func (q *MyQueue) Push(x int)  {
    q.inStack = append(q.inStack, x)
}

func (q *MyQueue) in2out() {
    for len(q.inStack) > 0 {
        q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
        q.inStack = q.inStack[:len(q.inStack)-1]
    }
}


/** Removes the element from in front of queue and returns that element. */
func (q *MyQueue) Pop() int {
	if len(q.outStack) == 0 {
	    q.in2out()
    }
	p := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]
	return p
}


/** Get the front element. */
func (q *MyQueue) Peek() int {
    if len(q.outStack) == 0 {
        q.in2out()
    }
    return q.outStack[len(q.outStack)-1]
}


/** Returns whether the queue is empty. */
func (q *MyQueue) Empty() bool {
    return len(q.inStack) == 0 && len(q.outStack) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {
    obj := Constructor()
    obj.Push(1)
    obj.Push(2)
    fmt.Println(obj.Peek())
    fmt.Println(obj.Pop())
    fmt.Println(obj.Empty())
    fmt.Println(obj.Pop())
    fmt.Println(obj.Empty())

}