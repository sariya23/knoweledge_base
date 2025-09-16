[[Spark]]
## Avro
Запись данных в avro
```python
spark = SparkSession.builder.appName("Spark APP")\
            .config("spark.master", "local[*]")\
            .config("spark.jars.packages", "org.apache.spark:spark-avro_2.12:4.0.0")\
            .getOrCreate()

data = [("Alice", 25), ("Bob", 30), ("Cathy", 29)]
df = spark.createDataFrame(data, ["Name", "Age"])
df.write.format("avro").save("./avro_files")
```
Там проблема с зависимостями, поэтому нужно в конфиг добавить строку с понижением пакета
Чтение из avro:
```python
df_avro = spark.read.format("avro").load("./avro_files")
df_avro.show()
```

У остальных логика такая же
