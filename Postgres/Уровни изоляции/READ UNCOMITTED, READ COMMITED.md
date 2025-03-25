[[Postgres]]
## READ UNCOMMITTED

Уровень изоляции READ UNCOMMITTED является самым "мягким" в стандарте SQL. Он позволяет читать другим транзакциям данные, которые еще не были закоммичены
Вот пример аномалии dirtry read
Первая транзакция:
```sql
db=> begin;
BEGIN
db=*> update test.account set balance = balance - 900 where id = 1;
UPDATE 1
db=*>
```
Вторая транзакция
```sql
db=> begin;
BEGIN
db=*> update test.account set balance = balance - 100 where id = 1;

```
Тут первая транзакция меняет поле `balance`, вычитая из него 900, но не коммит изменения. Вторая транзакция отнимает еще 100. Для второй транзакции `balance` равен 100, так как изначально он был 1000. Но мы можем откатить первую транзакцию. И тогда может быть ситуация, что баланс в итоге будет 0, то есть как будто первая транзакция не откатывалась, а применилась. 
==В Postgres нет уровня READ UNCOMMITTED, так как тут есть MVCC, где указаны номера транзакций и их актуальность. В Postgres этот уровень работает так же как и READ COMMITTED==

## READ COMMITTED
Уровень READ COMMITTED не позволяет другой транзакции читать данные, которые были изменены. В примере выше во второй транзакции Postgres будет ждать завершения первой транзакции. Этот уровень изоляции в Postgres по умолчанию
Первая транзакция
```sql
db=> begin transaction isolation level read committed;
BEGIN
db=*> update test.account set balance = balance - 900 where id = 1;
UPDATE 1
db=*>
```
Вторая транзакция видит "старые" данные
```sql
db=> begin transaction isolation level read committed;
BEGIN
db=*> table test.account;
 id | balance
----+---------
  1 |    1000
  2 |    2000
```
Решенные аномалии:
- Грязное чтение
Допустимые аномалии:
- Неповторяемое чтение
- Фантомное чтение
- Аномалия сериализации