[[Airflow]]
XComs позволяет обмениваться данными между тасками (прям как канал в go). Каждая задача можно запушить данные в XCom (`xcom_push`), а другие задачи могут эти данные забрать (`xcom_pull`). Данные XCom хранятся в базе данных
Xcom используется только для передачи небольших данных - несколько мегабайт или килобайт.
## Xcom в PythonOperator
Для того, чтобы что-то положить к xcom или получить, нужно получить из контекста текущую задачу под ключом `ti`
```python
from airflow import DAG
from datetime import timedelta
from airflow.operators.python_operator import PythonOperator
from datetime import datetime
from airflow.utils.dates import days_ago


def push_function(**context):
    context['ti'].xcom_push(key='key_name', value="Text...")


def pull_function(**context):
    ti = context['ti'] 
    value_pulled = ti.xcom_pull(key='key_name') 
    print(value_pulled)

dag = DAG('0_Examples_5_1_1_xcom',schedule_interval=timedelta(days=1), start_date=days_ago(1))

push_task = PythonOperator(
    task_id='push_task',
    python_callable=push_function,
    dag=dag,
)

pull_task = PythonOperator(
    task_id='pull_task',
    python_callable=pull_function,
    dag=dag,
)

push_task >> pull_task
```
## Xcom в BashOperator
Тут доступ к контексту мы получаем через JInja
```python
from airflow import DAG
from datetime import timedelta
from airflow.utils.dates import days_ago
from airflow.operators.bash import BashOperator

dag = DAG('0_Examples_5_1_2_xcom_bash',schedule_interval=timedelta(days=1), start_date=days_ago(1))

downloading_data = BashOperator(
    task_id='downloading_data',
    bash_command='echo "Hello, I am a value!"',
    do_xcom_push=True,
    dag=dag
)

fetching_data = BashOperator(
    task_id='fetching_data',
    bash_command="echo 'XCom fetched: {{ ti.xcom_pull(task_ids=[\'downloading_data\']) }}'",
    dag=dag
)

downloading_data >> fetching_data
```
В новых версиях результат работы `echo` по дефолту будет класть значение в xcom