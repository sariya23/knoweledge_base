	[[Spark]]
Основные компоненты Spark SQL:
- DataFrame API - API датафрейма позволяет делать те же операции, что и в чистом SQL: сортировка, фильтрация, etc
- SQL API - позволяет выполнять SQL запросы на DataFrame
## DataFrame API
```
+-----+---+-----------+
|name |age|department |
+-----+---+-----------+
|John |30 |HR         |
|Doe  |25 |Finance    |
|Jane |35 |HR         |
|Mark |40 |Finance    |
|Smith|23 |Engineering|
+-----+---+-----------+
```
```python
df_result.groupby("department").agg({"age": "avg"}).withColumnRenamed("avg(age)", "Avg age by dep").show()
```
Выглядит очень удобно (нет)
```
+-----------+--------------+
| department|Avg age by dep|
+-----------+--------------+
|Engineering|          23.0|
|         HR|          32.5|
|    Finance|          32.5|
+-----------+--------------+
```
## SQL API
С помощью SQL API мы можем выполнять операции над dataframe с помощью запросов на чистом SQL
Для того, чтобы это стало возможным, нужно зарегать датафрейм как вьюху. Она будет временная, на одну сессию (насколько я понимаю)
```python
df_result.createOrReplaceTempView("people")
```
Все) Теперь мы можем обращаться к таблице `people`:
```python
spark.sql("select * from people order by age").show()
```
```
+-----+---+-----------+
| name|age| department|
+-----+---+-----------+
|Smith| 23|Engineering|
|  Doe| 25|    Finance|
| John| 30|         HR|
| Jane| 35|         HR|
| Mark| 40|    Finance|
+-----+---+-----------+
```