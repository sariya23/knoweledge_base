[[Python]]
Сопоставление паттерну (паттерн-матчинг) появился в 3.10. С виду он похож на `switch/case`, но у него чуть другой смысл - матчить паттерн последовательностей
```python
def beep(freq, times):
    ...

def color(color):
    ...

def handle_command(message):
    match message:
        case ["BEEP", freq, times]:
            beep(freq, times)
        case ["LED", color]:
            color(color)
        case _:
            raise ValueError(message)
```
То есть мы смотрим, та ли структура у объекта и можем также проверять конкретные значения. Кейс с `_` специальный - это аналог `default`. Вообще тут та же логика, что и при распаковки - надо прописать структуру вложенности
Паттерн матчинг поддерживают следующие структуры:
- `list`
- `tuple`
- `range`
- `memoryview`
- `array.array`
- `collections.deque`
Также любую часть паттерна можно взять как переменную, используя `as`:
```python
metro_areas = [
 ('Tokyo', 'JP', 36.933, (35.689722, 139.691667)),
 ('Delhi NCR', 'IN', 21.935, (28.613889, 77.208889)),
 ('Mexico City', 'MX', 20.142, (19.433333, -99.133333)),
 ('New York-Newark', 'US', 20.104, (40.808611, -74.020386)),
 ('São Paulo', 'BR', 19.649, (-23.547778, -46.635833)),
]

for record in metro_areas:
    match record:
        case [name, _, _, (lat, lon) as coord] if lat <= 0:
            print(coord)
```
Тут мы сказали, чтобы часть паттерна сохранилась в переменную `coord`. Также при матчинге можно писать какие-то условия. Это условие начинает вычисляться только если матчинг по паттерну прошел
В паттерне можно указывать уточнения по типу:
```python
case [str(name), _, _, (float(lat), float(lon))]:
```
Тут не происходит вызов конструктора, это именно проверка типа. Также такое можно провернуть и с пользовательскими классами
```python
class A:
    def __init__(self):
        pass
    

lst = [
    (A(), A(), 2),
    ("asd", "zxc", "qwe"),
]

for l in lst:
    match l:
        case [A() as a1, A() as a2, int(a)]:
            print("get it", a1)
```
Также можно матчить не только по типу класса, но и так, чтобы были нужные значения атрибутов. Доп. инфа: 
- https://realpython.com/structural-pattern-matching/#customizing-pattern-matching-of-class-instances
- https://peps.python.org/pep-0636/