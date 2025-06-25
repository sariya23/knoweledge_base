Рассмотрим основные операции над rdd. Создание rdd:
```python
sc = spark.sparkContext
data = [1, 2, 3, 4, 5, 6]
rdd = sc.parallelize(data)
```
**Фильтраация**:
```python
rdd.filter(lambda x: x % 2 == 0).collect()
# [2, 4, 6]
```
**Маппинг**
```python
rdd.map(lambda x: x ** 2).collect()
# [1, 4, 9, 16, 25, 36]
```
**FlatMap**, не очень понял зачем он нужен. Да и как работает тоже. Наверное основные кейсы использования, это когда есть вложенные объекты, а нам нужно применить функцию именно к элементам коллекции, а потом результат сделать одномерным. Один из примеров - это разбивка техта на слова
```python
rdd.flatMap(lambda x: x.split(" ")).collect()
```
**Уникальные значения** получаются через метод `distinct`
```python
rdd2 = sc.parallelize([40, 50, 50])
rdd2.distinct().collect()
# [40, 50]
```
**Получение элементов** через метод `take(n)`
```python
rdd.take(4)
# [1, 2, 3, 4]
```
**foreach** применяет переданную функцию к каждому элементу
```python
rdd.foreach(lambda x: print(x + 1))
7
2
5
6
4
3
```
## Кейсы
### Подсчет кол-ва слов в строке
```python
rdd.flatMap(lambda x: x.split(" ")).map(lambda x: (x, 1)).reduceByKey(lambda a, b: a + b).collect()
```
Где rdd - это строка текста