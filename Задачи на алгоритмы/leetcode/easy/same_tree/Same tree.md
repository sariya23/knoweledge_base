[[easy]]

https://leetcode.com/problems/same-tree/description/

Ох уж эти деревья... Нагородил я знатно, какие-то указатели слайсов. Можно было без этого. Я был близок и пробовал это решение через вспомогательную функцию, но чет не дажал. Ну вот хотя бы первое попавшееся решение
```python
class Solution:
    def isSameTree(self, p, q):
        # If both nodes are None, they are identical
        if p is None and q is None:
            return True
        # If only one of the nodes is None, they are not identical
        if p is None or q is None:
            return False
        # Check if values are equal and recursively check left and right subtrees
        if p.val == q.val:
            return self.isSameTree(p.left, q.left) and self.isSameTree(p.right, q.right)
        # Values are not equal, they are not identical
        return False
```
По эффективности даже как будто бы тоже самое. У меня правда два раза дерево обходится и потом цикл еще. Нооо, что имеем