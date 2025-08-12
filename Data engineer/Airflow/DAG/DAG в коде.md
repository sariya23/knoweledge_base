[[Airflow]]
Создать инстанс дага через класс
```python
from airflow import DAG

dag = DAG('dag', schedule_interval=timedelta(days=1), start_date=days_ago(1))
```
Также можно создать даг через контекст манагер
```python
with DAG('dag', schedule_interval=timedelta(days=1), start_date=days_ago(1)) as dag: 
    pass 
```
## Основные параметры
- `dag_id` - уникальное имя дага. Оно отображается в UI Airflow
- `start_date` - дата начала выполнения дага
- `schedule` - интервал выполнения
Момент выполнения рассчитывается по такой формуле
```
ExecutionDate = start_date + schedule_interval
```
Передавать можно как и базовым `datetime` так и внутренними библиотеками Airflow. Даг начнет выполнятся, только если текущий момент больше или равен `ExecutionDate`
### Параметры тасок default args
Параметры для каждой таски в даге по дефолту ставятся через параметр `default_args`
```python
dag = DAG(
    'dag',
    default_args={"retries": 3}, 
    start_date=datetime(2024, 1, 1),
    schedule_interval=None, 
)
```
- `owner: str` - владелец или автор DAG
- `depends_on_past: bool` - определяет, будут ли выполнятся задачи, если в прошлом они упали (если True)
- `start_date: datetime` - время и дата, когда даг начнет первое выполнение
- `end_date: datetime` - время и дата, когда даг перестанет запускаться
- `retries: int` - кол-во ретраев в случае падения таски
- `retry_delay: timedelta` - врем ожидание повторного выполнения таски в случае падения
- `email: str` - email, на который будут отправляться статусы выполнения задач
- `email_on_failure: bool` - если True, то в случае падения таски на почту придет письмо
- `email_on_retry: bool` - если True, то в случае ретрая таски на почту придет письмо
- `schedule_interval: timedelta | cron_str` - интервал запуска дага. Либо timedelta, либо cron-строка
- `max_active_runs: int` - максимальное кол-во запущенных дагов
- `catchup: bool` - если True, то будут выполнятся задачи за пропущенные интервалы времени. То есть если есть выполнения на 1 февраля, а 1 февраля Airflow не работал, то он начнет выполнятся
- `on_success_callback: context` - вызов функции при успешном выполнении задачи
- `on_failure_callback: context` - вызов функции при неуспешном выполнении задачи
- `on_retry_callback: context` - вызов функции при ретаре задачи