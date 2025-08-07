[[Airflow]]
`DAG` - ключевой объект при работе с Airflow
```python
from airflow import DAG
```
У DAG есть уникальное имя, по которому его можно найти. Создать DAG с таким же именем не получится. Чтобы создать DAG:
```python
dag = DAG( 
	'0_Examples_3.1_1_introduction', 
	schedule_interval='@daily', 
	start_date=days_ago(1),
	...)
```
В рамках дага мы может определять ему задачи. Задачи выполняются операторами, они бывают разные. У каждого оператора есть уникальное имя в рамках дага. Рассмотрим `PythonOperator`
```python
from airflow.operators.python import PythonOperator

task_extract = PythonOperator(
    task_id='extract_data',         
    python_callable=extract_data,    
    op_args=['http://www.cbr.ru/scripts/XML_daily.asp', '01/01/2022', './extracted_data.xml'],
    dag=dag,                         
)
```
`task_id` - уникальное имя задачи
`python_callable` - функция для вызова
`op_args` - список позиционных параметров, которые передадутся в функцию
`dag` - передаем инстанс дага, чтобы связать с ним задачу
Также можно передать параметры функции по именам через словарь
```python
task_transform = PythonOperator(
    task_id='transform_data',
    python_callable=transform_data,
    op_kwargs = {
        's_file': './extracted_data.xml', 
        'csv_file': './transformed_data.csv', 
        'date': '01/01/2022'},
    dag=dag,
)
```
