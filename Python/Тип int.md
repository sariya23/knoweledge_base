[[Python]]

Для python >= 3.11

## Превращение объектов в `int`
Чтобы превратить какой-то объект в `int`, нужно реализовать метод `__int__`
```python
class MyInt:
    def __init__(self, number: int) -> None:
        self.number = number
    
    def __int__(self) -> int:
        return self.number


print(int(MyInt(7)))  # 7
```
Также важно, чтобы `__int__` возвращал именно `int`, иначе будет исключение
```python
class MyInt:
    def __init__(self, number: int) -> None:
        self.number = number
    
    def __int__(self) -> int:
        return float(self.number)


print(int(MyInt(7)))  # TypeError
```
Если мы хотим захинтить какую-то функцию, которая должна принимать что-то, что можно превратить в `int`, то можно использовать протокол `typing.SupportsInt`
```python
def f(i: SupportsInt) -> int:
    return int(i) + 2

print(f(MyInt(3)))
```
## Превращение строки в число
Можно превратить число из строкового представления с помощью функции `int()`
```python
>>> int("10")
10
```
Но также можно указать СС:
```python
>>> int("1010", base=10)
1010
>>> int("1010", base=2)
10
```
То есть указываем, из какой СС перевести число в десятичную СС.
Число в виде строки может быть не только арабским числом, но и каким-то восточным
## Другие способы получить `int`
Есть также еще 4 способа получить `int`:
- `math.ceil()` - округление вверх до ближайшего четного
- `math.floor()` - округление вниз до ближайшего четного
- `math.trunc()` - обрезка дробной части, либо просто `int()`
- `__index__` 
Магический метод `__index__` вызывается тогда, когда объект передается в качестве индекса в `[]`
```python
class MyInt:
    def __init__(self, number: int):
        self.number = number
    
    def __index__(self) -> int:
        return self.number

m = MyInt(1)
l = [1, 2, 3]
print(l[m])  # 2
```
Если у объекта не реализован метод `__int__`, то функции `int`, `float`, `complex` будут обращаться к `__index__`
```python
class MyInt:
    def __init__(self, number: int):
        self.number = number
    
    def __index__(self) -> int:
        return self.number

m = MyInt(1)
print(int(m))  # 1
print(float(m))  # 1.0
```

