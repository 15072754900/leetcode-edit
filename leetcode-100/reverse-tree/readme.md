这种题目处理方法
1.递归
使用栈，压入所有节点，对节点进行处理，判断无子节点就出栈

2.迭代
构建队列，将节点加入到队列中，如果该结点包含左右子节点，进行翻转，完成后推出队列。