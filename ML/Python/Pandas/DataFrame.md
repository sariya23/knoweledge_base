## Создание DataFrame
### На основе словаря
Создать DataFrame можно на основе словаря. Внутри датафрейма по сути хранятся серии, просто объединенные. Поэтому при создании словаря будет так: 
- ключ: название колонки в DataFrame
- значение: значение. Часто делаю списком, чтобы был маппинг одного на другое
```python
d = {
    "username": ["qwe", "zxc", "nagibator"],
    "hours": [3, 5000, 567],
    "rank": ["herald", "titan", "archon"]
}
df = pd.DataFrame(d)
print(df)

    username  hours    rank
0        qwe      2  herald
1        zxc      2   titan
2  nagibator      2  archon
```
Также как и с серией можно поменять порядок колонок при создании, передав параметр `columns` и задать порядок
```python
df = pd.DataFrame(d, columns=["rank", "hours", "username"])
print(df)
     rank  hours   username
0  herald      2        qwe
1   titan      2        zxc
2  archon      2  nagibator
```
Если какая-то колонка будет написана с ошибкой, то эта колонка добавится, но все значения будут `NaN`
```python
df = pd.DataFrame(d, columns=["rank", "hours", "zxc"])
print(df)
     rank  hours  zxc
0  herald      2  NaN
1   titan      2  NaN
2  archon      2  NaN
```
Если не указать какую-то колонку, то она просто не добавится
```python
df = pd.DataFrame(d, columns=["rank", "hours"])
print(df)
     rank  hours
0  herald      2
1   titan      2
2  archon      2
```
### На основе вложенного словаря
Если в словаре хранится ключ - словарь, то ключи вложенного словаря становятся индексами
```python
d = {
    "qwe": {"hour": 3, "rank": "herald"},
    "zxc": {"hour": 5000, "rank": "titan"},
    "nagibator": {"hour": 567, "rank": "archon"},
}
df = pd.DataFrame(d)
print(df)
         qwe    zxc nagibator
hour       3   5000       567
rank  herald  titan    archon
```

## Получение колонок
Чтобы получить колонку надо передать ее название в паре квадратных скобок
```python
print(df["hours"])
0    2
1    2
2    2
```
## Получение строки по индексу
Чтобы получить строку по индексу нужно вызвать метод `loc` и передать в кв. скобки индекс. Если индекс - число, то передаем как число, если строка - как строку
```python
print(df.loc[0])
username       qwe
hours            2
rank        herald
Name: 0, dtype: object
```
```python
df.index = ["a", "b", "c"]
print(df.loc["a"])
username       qwe
hours            2
rank        herald
Name: a, dtype: object
```
## Добавление колонки
Либо передаем одно значение, тогда колонка будет с этим одним значением, либо список значений такой же длины как и колонка
```python
df["male"] = ["f", "m", "f"]
print(df)
    username  hours    rank male
0        qwe      2  herald    f
1        zxc      2   titan    m
2  nagibator      2  archon    f
```
```python
df["platform"] = "pc"
print(df)
    username  hours    rank male platform
0        qwe      2  herald    f       pc
1        zxc      2   titan    m       pc
2  nagibator      2  archon    f       pc
```
Также можно добавить колонку как серию. Добавление происходит по индексам. Если какого-то значение нет, то поставится `NaN`
```python
s = pd.Series([20, 18], index=[2, 1])
df["age"] = s
print(df)
    username  hours    rank male platform   age
0        qwe      2  herald    f       pc   NaN
1        zxc      2   titan    m       pc  18.0
2  nagibator      2  archon    f       pc  20.0
```
По индексу 2 он поставил 2, а по индексу 1 18. На индекс 0 значения нет, поэтому там `NaN`. Если какого-то индекса нет в DataFrame, то там будет `NaN`
```python
s = pd.Series([1, 2, 3], index=[10, 2, 1])
df["a"] = s
print(df)
    username  hours    rank    a
0        qwe      3  herald  NaN
1        zxc   5000   titan  3.0
2  nagibator    567  archon  2.0
```
Индекса 10 нет в DataFrame, поэтому значение 1 не добавится
## Удаление колонки
Удаление колонки идет через оператор `del`
```python
print(df)
    username  hours    rank    a
0        qwe      3  herald  NaN
1        zxc   5000   titan  3.0
2  nagibator    567  archon  2.0
del df["a"]
print(df)
    username  hours    rank
0        qwe      3  herald
1        zxc   5000   titan
2  nagibator    567  archon
```
## Транспонирование 
Если мы хотим сделать колонки индексами, а индексы колонками, то DataFrame можно транспонировать
```python
print(df)
         qwe    zxc nagibator
hour       3   5000       567
rank  herald  titan    archon
print(df.T)
           hour    rank
qwe           3  herald
zxc        5000   titan
nagibator   567  archon
```
## Изменение индексов и колонок
Названия колонок и индексов нельзя изменить по одиночке
```python
df.index[0] = "some"
df.columns[0] = "a"
```
Так как индексы и колонки лежат в неизменимой последовательности. Для того, чтобы изменить индексы или колонки, нужно переприсваивать все значения
```python
df.index = [1, 2, 3, 4]
df.columns = ["a", "b", "c", "d"]
```