[[Spark]]
## Чтение в RDD
```python
rdd = sc.textFile("./data.txt")
```
## Чтение DataFrame
```python
df = spark.read.csv("./data.csv", sep=";", header=True)
```
Параметры при чтении:
- `header` - содержит ли первая строка файла заголовки
- `inferSchema` - автоматически определять типы (только для csv и json)
- `schema` - явное указание схемы данных
- `sep` - разделитель в csv
## Чтение json
Если json сложный, то лучше создать его схему
```python
from pyspark.sql.types import StructField, IntegerType, StringType, StructType, ArrayType

schema = StructType([
    StructField("people", ArrayType(
        StructType([
            StructField("name", StringType(), True),
            StructField("age", IntegerType(), True),
            StructField("department", StringType(), True),
        ])
    ), True)
])

df = spark.read.schema(schema).option("multiline", "true").json("people2.json")

from pyspark.sql.functions import explode

df_exploded = df.select(explode("people").alias("person"))
df_result = df_exploded.select(
    "person.name",
    "person.age",
    "person.department"
)

df_result.show(truncate=False)
```
Не любой json можно преобразовать в "нормальный" dataframe. 