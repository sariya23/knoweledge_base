[[Spark]]
UDF - user defined function. Что-то вроде map для создание какие-то новых столбцов на основе старых
```python
from pyspark.sql import SparkSession
from pyspark.sql.functions import col, udf
from pyspark.sql.types import IntegerType, StringType

spark = SparkSession.builder \
    .appName("UDF Example") \
    .getOrCreate()

data = [(1, "Alice"), (2, "Bob"), (3, "Cathy"), (4, "David")]
df = spark.createDataFrame(data, ["id", "name"])

# Определяем пользовательскую функцию
def add_prefix(name):
    return "Name_" + name

# Регистрируем функцию как UDF
add_prefix_udf = udf(add_prefix, StringType())

# Применяем UDF к DataFrame
df_with_prefix = df.withColumn("prefixed_name", add_prefix_udf(col("name")))

df_with_prefix.show()

spark.stop()
```
Также UDF можно применять не только через API питона, но и в чистом sql
```python
from pyspark.sql import SparkSession
from pyspark.sql.types import IntegerType
spark = SparkSession.builder \
    .appName("Register UDF Example") \
    .getOrCreate()


data = [(1, "Alice"), (2, "Bob"), (3, "Cathy"), (4, "David")]
df = spark.createDataFrame(data, ["id", "name"])


def name_length(name):
    return len(name)

# Регистрируем функцию как UDF с использованием spark.udf.register
spark.udf.register("name_length_udf", name_length, IntegerType())

# Создаем временную таблицу для выполнения SQL-запросов
df.createOrReplaceTempView("people")

# Используем зарегистрированную UDF в SQL-запросе
result_df = spark.sql("SELECT id, name, name_length_udf(name) as name_length FROM people")


result_df.show()


spark.stop()
```