[[Airflow]]
Операторы сенсоры ждут наступления определенного события, чтобы продолжить выполнение таски. Если событие сенсора не наступает, задача перезапускается. Сенсоры применяются не для обработки данных, а для проверки наступления того или иного события
## PythonSensor
```python
sensor = PythonSensor(
    task_id='task',
    mode='poke',
    python_callable=func,
    poke_interval=1, # Через какое время перезапускаться
    timeout=10,# Время до принудительного падения
    soft_fail=True # Пропустить, скипнуть, задачу если она закончится неудачей
)
```
Сенсор ожидает, что функция будет возвращать True или False, None будет расцениваться как False
Параметр `mode` принимает одно из нескольких значений. `poke` - режим по умолчанию. В этом режиме сенсор помешается в сенсор пул и не освобождает его пока не выполнится. `reschedule` режим при котором сенсор, в случае неудачи, будет перезапускаться, навремя освобождая пул (перезапускаться будет каждый `poke_interval`)
`timeout` - время, через которое сенсор прекратит попытки выполнения. `soft_fail` - параметр, который отвечает за статус задачи при неуспешном выполнении. Если `soft_fail=True` и сенсор завершился с ошибкой, то статус задачи будет `skipped`
Все время в секундах
## HttpSensor
Позволяет отправить хттп запрос и проверять какие-то условия
```python
response_check = lambda response: response.status_code == 200

sensor = HttpSensor(
    task_id='http_sensor',          
    http_conn_id='http_default',    
    endpoint='',                 
    response_check=response_check, 
    poke_interval=10,              
    dag=dag                         
)
```
## FileSensor
Этот сенсор проверяет есть ли файл в файловой системе
```python
file_sensor = FileSensor(
    task_id='check_for_file',
    filepath='/path/to/file.txt',  
    fs_conn_id='fs_default',      
    poke_interval=60,             
    timeout=1800,                 
    mode='poke',                
    poke_mode=True,             
    dag=dag
)
```
Важным параметром является `fs_conn_id`. Тут указывается ssh сервера, к чей файловой системе мы хотим подключится
## ClickHouseSensor
Если запрос возвращает не пустую таблицу, то сенсор считается успешно выполненым
```python
sql= "SELECT count(*) FROM my_table WHERE my_column = 'some_value'"

clickhouse_sensor = ClickHouseSqlSensor(
    task_id='check_if_data_exists',
    sql=sql,                                 
    clickhouse_conn_id='clickhouse_default', 
    poke_interval=10,                         
    timeout=600,                              
    mode='poke',                              
    dag=dag                                   
)
```