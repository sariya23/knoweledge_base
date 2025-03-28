[[Postgres]]
При уровне изоляции READ COMMITTED все равно возможны некоторые аномалии. Например, фантомное чтение и неповторяемое чтение
### Аномалия неповторяемое чтение
Аномалия неповторяемого чтения возникает, когда в транзакции при первом запросе возвращаются одни данные, а при втором - другие. То есть во время выполнения первой транзакции, какая-то другая изменила данные и закоммитила их
Первая транзакция
```sql
db=> begin;
BEGIN
db=*> table test.account;
 id | balance
----+---------
  1 |    1000
  2 |    2000
(2 строки)

db=*> table test.account;
 id | balance
----+---------
  1 |    1000
  2 |    2000
(2 строки)

db=*> table test.account;
 id | balance
----+---------
  2 |    2000
  1 |     100
(2 строки)
```
Вторая транзакция
```sql
db=> begin;
BEGIN
db=*> update test.account set balance = 100 where id = 1;
UPDATE 1
db=*> commit;
COMMIT
db=>
```
Во второй транзакции мы поменяли поле `balance` и закоммитили это. Первая транзакция на уровне READ UNCOMMITTED это видит, так как вторая транзакция закоммитила эти данные и все так и должно быть
### Аномалия фантомное чтение
Эта аномалия возникает, когда во время выполнения первой транзакции, вторая добавляет или удаляет строку. Первая транзакция увидит эти изменения, так как они были закомичены во второй
Первая транзакция
```sql
db=> begin;
BEGIN
db=*> table test.account;
 id | balance
----+---------
  2 |    2000
  1 |    1000
(2 строки)

db=*> table test.account;
 id | balance
----+---------
  2 |    2000
  1 |    1000
  3 |    3000
(3 строки)
```
Вторая транзакция
```sql
db=> begin;
BEGIN
db=*> insert into test.account values (3, 3000);
INSERT 0 1
db=*> commit;
COMMIT
db=>
```
Так как изменения были закомичены, первая транзакция все видит.
## REPEATABLE READ
Этот уровень изоляции фиксит аномалии, описанные выше.
### Фикс неповторяемого чтения
Первая транзакция
```sql
db=> begin isolation level repeatable read;
BEGIN
db=*> table test.account;
 id | balance
----+---------
  2 |    2000
  1 |    1000
(2 строки)

db=*> table test.account;
 id | balance
----+---------
  2 |    2000
  1 |    1000
(2 строки)
```
Вторая транзакция
```sql
db=> update test.account set balance = 100 where id = 1;
UPDATE 1
```
Как видно, в первой транзакции не видны изменения данных во второй транзакции
### Фикс фантомного чтения
Первая транзакция
```sql
db=> begin isolation level repeatable read;
BEGIN
db=*> table test.account;
 id | balance
----+---------
  2 |    2000
  1 |     100
(2 строки)

db=*>
db=*> table test.account;
 id | balance
----+---------
  2 |    2000
  1 |     100
(2 строки)
```
Вторая транзакция
```sql
db=> insert into test.account values (3, 3000);
INSERT 0 1
```
Новая строка не видна в таблице даже после коммита первой транзакции
Решенные аномалии:
- Фантомное чтение
- Повторное чтение
Нерешенные аномалии:
- Аномалия сериализации
