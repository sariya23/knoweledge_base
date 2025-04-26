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
## Иерархия чисел
Как и в математике в python сделали иерархию чисел `Number :> Complex :> Real :> Rational :> Integral`. Эти классы располагаются в модуле `numbers`. Вся дополнительная информация описана в PEP-3141.
==НЕ СТОИТ ИСПОЛЬЗОВАТЬ ЭТИ КЛАССЫ ДЛЯ АННОТАЦИИ==
## Исходники
Изначально в питоне было 2 инта - лонг и обычный. Даже теперь инт называется как лонг - `PyLong_Type`:
```c
PyTypeObject PyLong_Type = {
    PyVarObject_HEAD_INIT(&PyType_Type, 0)
    "int",                                      /* tp_name */
    offsetof(PyLongObject, long_value.ob_digit),  /* tp_basicsize */
    sizeof(digit),                              /* tp_itemsize */
    long_dealloc,
    ...
};             
```
https://github.com/python/cpython/blob/main/Objects/longobject.c#L6614C14-L6614C26
Если рассматривать как хранятся числа в C, то они хранятся как последовательность:
```c
typedef struct _PyLongValue {
    uintptr_t lv_tag; /* Number of digits, sign and flags */
    digit ob_digit[1];
} _PyLongValue;

struct _longobject {
    PyObject_HEAD
    _PyLongValue long_value;
};
```
https://github.com/python/cpython/blob/main/Include/cpython/longintrepr.h#L98
## sys.maxsize
В общем, тут что-то намудрено. Они там понасоздавали кучу каких-то ssize. `sys.maxsize` возвращается максимальное число, которое может вместить `Py_ssize_t`. 
Насколько я понял, этот `Py_ssize_t` используется везде для индексации.
```c
static PyObject *
list_insert_impl(PyListObject *self, Py_ssize_t index, PyObject *object)
/*[clinic end generated code: output=7f35e32f60c8cb78 input=b1987ca998a4ae2d]*/
{
    if (ins1(self, index, object) == 0) {
        Py_RETURN_NONE;
    }
    return NULL;
}
```
Большинство методов в качестве индекса принимают тип как раз `Py_ssize_t`.
Единственное, что я понял, это то, что максимальный инт в питоне это не `sys.maxsize`, он может быть больше, но индексы помещаются только в `Py_ssize_t`, то есть в `sys.maxsize`
## AST оптимизации
С числами есть много оптимизаций. Например
```python
import dis


def f():
    return 1 + 2

dis.dis(f)
  4           0 RESUME                   0

  5           2 LOAD_CONST               1 (3)
              4 RETURN_VALUE
```
Видно, что никакого сложения не происходит, а загружается константа. Но это работает только в случае, если происходит явное сложение, а не через переменные
```python
import dis


def g(a, b):
    def f():
        return a + b
    return f

f = g(1, 2)

dis.dis(f)

              0 COPY_FREE_VARS           2

  5           2 RESUME                   0

  6           4 LOAD_DEREF               0 (a)
              6 LOAD_DEREF               1 (b)
              8 BINARY_OP                0 (+)
             12 RETURN_VALUE
```
