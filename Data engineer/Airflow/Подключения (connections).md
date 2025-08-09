[[Airflow]]
Подключения позволяют во время таски дага обращаться к БД или другим сервисам. 
Например, можно сохранить подключение к постгресу и обращаться к нему в любой таске или скрипте Airflow
Создать также можно через UI или API
Есть 2 способа использования подключений. 
Первый - это нативное использование для некоторых операторов. 
```python
create_view = ClickHouseOperator(
    task_id=...,
    sql=... ,
    clickhouse_conn_id='clickhouse_default',
    dag=dag,
)
```
И "подход дикаря". Такой подход является антипаттерном.
```python
from airflow.hooks.base_hook import BaseHook 

host = BaseHook.get_connection("postgres_default").host
pass = BaseHook.get_connection("postgres_default").password
```
