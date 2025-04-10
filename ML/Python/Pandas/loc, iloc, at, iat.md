## loc
```
      name  age  clicks  balance  history
a   Andrey   25       4        0        1
b     Petr   40       9      250        8
c  Nikolay   19      12        0        1
d   Sergey   33       6      115        6
e   Andrey   38       2      152        4
f     Ilya   20      18        0        2
g     Igor   19       2        0        1
```
Это свойство используется для вытягивание колонок, строк. И он решает вот эту проблему:
```python
df[["name", "age"]]["c"]
```
То есть получить строку с лейблом `c` и колонки `name`, `age`, чтобы было так:
```python
name    Nikolay
age          19
Name: c, dtype: object
```
`loc` позволяет это сделать. В общем случае, `loc` - это позволяет делать индексацию как в матрицах `numpy`:
```python
print(df.loc["c", ["name", "age"]])
name    Nikolay
age          19
Name: c, dtype: object
```
Первое значение - какие строки берем, второе - какие столбцы.
### Вытянуть столбец
```python
print(df.loc[:, "age"])
a    25
b    40
c    19
d    33
e    38
f    20
g    19
Name: age, dtype: int64
```
То есть также как и в матрицах numpy, первое значение - какие строки, второе - какие колонки
### Вытянуть список строк
```python
print(df.loc[["a", "g"]])
     name  age  clicks  balance  history
a  Andrey   25       4        0        1
g    Igor   19       2        0        1
```
### Список строк и какой-то столбец
```python
print(df.loc[["a", "g"], "history"])
a    1
g    1
Name: history, dtype: int64
```
### Условия
Тут работаю также и условия
```python
print(df.loc[df["age"] > 30, :])
     name  age  clicks  balance  history
b    Petr   40       9      250        8
d  Sergey   33       6      115        6
e  Andrey   38       2      152        4
```
Это аналогично вот этому:
```python
df[df["age"] > 30]
```
## iloc
Делает то же самое, но работает с числовыми индексами
## at, iat
То же самое, только не могут принимать список колонок или строк. Использовать, когда нужно получить только одно значение