##python实现二叉树排序
class Node(object):
    def __init__(self,value):
        self.value = value
        self.left = None
        self.right = None

class BinaryTree(object):
# root = Node(5)
# for i in [2,11,7,3,9,8,4,6,1]:
#     tree.insert(root,Node(i))
#这里第一个就是5,3 然后inert(5.left,3)，因为left是空，所以返回3，所以就变成 5,然后就是下一个11，是第二个，因为右边没有所以是 5  ，以此类推
    def insert(self, root, node):                                # 3                                         3 11
        if root is None:
            return node
        if node.value < root.value:
            root.left = self.insert(root.left, node)
        else:
            root.right = self.insert(root.right, node)

        return root

def pre_order_print(node):
    # self -- left -- right
    print(node.value,end=' ')
    if node.left:
        pre_order_print(node.left)
    if node.right:
        pre_order_print(node.right)

def mid_order_print(node):
#     5
#   3   6
# 1   4
    ###这里先取左发现有，再去做发现有，一直取到最左边一个值发现没有就输出，他没有右，然后退回倒数第二个函数，第一个判断做完成就输出自己，然后判断右，再退回上一个函数，以此类推
    # mid --self -- right
    if node.left:
        mid_order_print(node.left)
    print(node.value,end=' ')
    if node.right:
        mid_order_print(node.right)

def after_order_print(node):
    # left-- right--self
    if node.left:
        after_order_print(node.left)
    if node.right:
        after_order_print(node.right)
    print(node.value, end=' ')

root = Node(5)

tree = BinaryTree()

for i in [2,11,7,3,9,8,4,6,1]:
    tree.insert(root,Node(i))
mid_order_print(root)