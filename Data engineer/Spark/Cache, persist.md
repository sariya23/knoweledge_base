[[Spark]]
В PySpark есть возможность закэшировать результат вычислений в оперативной памяти (через метод `cache`) или использовать более гибкое сохранение методом `persist`:
- `MEMORY_ONLY` - сохранение только в ОЗУ
- `MEMORY_AND_DISK` - кэширование в памяти и на диске, если не хватило места в памяти
- `DISK_ONLY` - кэширование только на диске
- `MEMORY_ONLY_SER` - кэширование в сериализованном виде. Уменьшает размер, но увеличивает время доступа
- `MEMORY_AND_DISK_SER` - кэширование в сериализованном виде на диске, и в памяти
- `OFF_HEAP` - кэширование вне кучи
```python
from pyspark import StorageLevel
from pyspark.sql import SparkSession
from pyspark.sql.functions import col

spark = SparkSession.builder \
    .appName("Persist Example") \
    .getOrCreate()

data = [(i, f"Name_{i % 5}") for i in range(1000)]
df = spark.createDataFrame(data, ["id", "name"])

filtered_df = df.filter(col("id") % 2 == 0).persist(StorageLevel.MEMORY_AND_DISK)

# Первое действие вызывает выполнение и кэширование
filtered_df.count()  # Это действие выполнит фильтрацию и закэширует результат в памяти и на диске

filtered_df.show()  # Это действие извлечет данные из кэша и отобразит их

spark.stop()
```
Для того, чтобы высвободить ресурсы, нужно использовать метод `unpersist`