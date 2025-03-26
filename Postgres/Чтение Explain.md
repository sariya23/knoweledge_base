[[Postgres]]

Таблица:
```sql
CREATE TABLE foo (c1 integer, c2 text);
INSERT INTO foo
  SELECT i, md5(random()::text)
  FROM generate_series(1, 1000000) AS i;
```

## Seq scan
Попробуем посмотреть шаги для запроса
```sql
select * from foo;
```
```sql
db=> explain select * from test.foo;
                          QUERY PLAN
--------------------------------------------------------------
 Seq Scan on foo  (cost=0.00..18334.00 rows=1000000 width=37)
(1 строка)

Время: 0,410 мс
```
Чтение из таблицы может происходить разными способами. В данном случае использовался `Seq Scan` - последовательное чтение.
`cost` - число в вакууме, которое говорит о трудозатратности операции. Первое значение до .., это стоимость на получение первой строки. В данном случае `0`, после .. стоимость получения всех строк. В данном случае `18334`
`rows` - планировщик пытается предугадать, сколько строк вернется при Seq Scan. 
`width` - средний размер одной строки в байтах
Делая просто `explain` мы не выполняем запрос, планировщик просто прикидывает. То есть делая insert, delete, данные не изменятся. Чтобы выполнить запрос реально нужно использовать `explain (analyse)`
## Explain Analyse
При `explain analyse` запрос выполняется. Сравним результаты
```sql
explain (analyse) select * from test.foo;
```
```sql
				                       QUERY PLAN
-----------------------------------------------------------------------------------
 Seq Scan on foo  (cost=0.00..18334.00 rows=1000000 width=37) (actual time=0.009..51.463 rows=1000000 loops=1)
 Planning Time: 0.057 ms
 Execution Time: 78.108 ms
(3 строки)

Время: 78,583 мс
```
Последовательное сканирование по таблице `foo`. 
Планируется:
- стоимость получения первой строки `0`
- получение всех `18334`.
- кол-во строк `10_000_000`
Факт:
- Получение всех строк заняло `51 мс` 
- Кол-во строк `10_000_000`
- `loops = 1`, сканирование пришлось выполнить 1 раз
- Время выполнения запроса `78 мс`
## Explain Buffers
Чтобы выводить информацию о кеше, нужно добавить `buffers`
```sql
db=> explain (analyze, buffers) select * from test.foo;
                                                  QUERY PLAN
-----------------------------------------------------------------------------------
 Seq Scan on foo  (cost=0.00..18334.00 rows=1000000 width=37) (actual time=0.026..55.280 rows=1000000 loops=1)
   Buffers: shared hit=8334
 Planning Time: 0.053 ms
 Execution Time: 85.070 ms
(4 строки)

Время: 85,481 мс

```
`Buffers: shared hit`:  кол-во блоков, считанных из кэша
## Условия в EXPLAIN
```sql
db=> explain (analyze) select * from test.foo where c1 > 500;
                                                 QUERY PLAN
-------------------------------------------------------------------------------------------------------------
 Seq Scan on foo  (cost=0.00..20834.00 rows=999544 width=37) (actual time=0.060..82.157 rows=999500 loops=1)
   Filter: (c1 > 500)
   Rows Removed by Filter: 500
 Planning Time: 0.056 ms
 Execution Time: 108.699 ms
(5 строк)

Время: 109,216 мс
```
У нас добавился Filter, который указана в `where`. Стоимость запроса увеличилась и выбираемое кол-во строк уменьшилось
## Index scan
Создадим индекс
```sql
CREATE INDEX ON test.foo(c1);
```
И посмотрим как поиск будет теперь
```sql
db=> explain (analyze) select * from test.foo where c1 > 500;
                                                 QUERY PLAN
-------------------------------------------------------------------------------------------------------------
 Seq Scan on foo  (cost=0.00..20834.00 rows=999504 width=37) (actual time=0.043..82.451 rows=999500 loops=1)
   Filter: (c1 > 500)
   Rows Removed by Filter: 500
 Planning Time: 0.078 ms
 Execution Time: 109.989 ms
(5 строк)

Время: 110,458 мс
```
Отфильтровалось 500 строк. Пришлось прочитать больше 99% таблицы. Индекс не использовался
Попробуем отключить Seq Scan
```sql
db=> SET enable_seqscan TO off;
SET
Время: 0,370 мс
db=> EXPLAIN (ANALYZE) SELECT * FROM test.foo WHERE c1 > 500;
                                                            QUERY PLAN
----------------------------------------------------------------------------------------------------------------------------------
 Index Scan using foo_c1_idx on foo  (cost=0.42..36800.75 rows=999504 width=37) (actual time=10.402..193.306 rows=999500 loops=1)
   Index Cond: (c1 > 500)
 Planning Time: 0.122 ms
 Execution Time: 220.883 ms
(4 строки)

Время: 221,357 мс
```
Теперь вместо Seq Scan Index Scan. Вместо Filter Index Cond. Но время выполнения запроса увеличилось, да и стоимость стала выше
Не забыть отключить
```sql
SET enable_seqscan TO on;
```
Изменим запрос
```sql
db=> explain (analyze) select * from test.foo where c1 < 500;
                                                      QUERY PLAN
----------------------------------------------------------------------------------
 Index Scan using foo_c1_idx on foo  (cost=0.42..25.09 rows=495 width=37) (actual time=0.009..0.070 rows=499 loops=1)
   Index Cond: (c1 < 500)
 Planning Time: 0.073 ms
 Execution Time: 0.097 ms
(4 строки)

Время: 10,592 мс
```
Теперь планировщик решил использовать индекс
## Bitmap index scan
Добавим условие с текстом
```sql
db=> explain (analyze) select * from test.foo where c1 < 500 and c2 like 'abcd%';
                                                    QUERY PLAN
------------------------------------------------------------------------------------------------------------------
 Index Scan using foo_c1_idx on foo  (cost=0.42..26.32 rows=1 width=37) (actual time=0.096..0.096 rows=0 loops=1)
   Index Cond: (c1 < 500)
   Filter: (c2 ~~ 'abcd%'::text)
   Rows Removed by Filter: 499
 Planning Time: 0.100 ms
 Execution Time: 0.108 ms
(6 строк)

Время: 0,494 мс

```
Для `c1<500` используется Idnex Cond, а для строки Filter. Но по прежнему Inex Scan.
Если же будет только текст, то Seq Scan
```sql
db=> explain (analyze) select * from test.foo where c2 like 'abcd%';
                                                    QUERY PLAN
-------------------------------------------------------------------------------------------------------------------
 Gather  (cost=1000.00..14552.33 rows=100 width=37) (actual time=16.046..51.361 rows=21 loops=1)
   Workers Planned: 2
   Workers Launched: 2
   ->  Parallel Seq Scan on foo  (cost=0.00..13542.33 rows=42 width=37) (actual time=9.051..27.095 rows=7 loops=3)
         Filter: (c2 ~~ 'abcd%'::text)
         Rows Removed by Filter: 333326
 Planning Time: 0.066 ms
 Execution Time: 51.383 ms
(8 строк)

Время: 51,817 мс

```
Но тут планировщик решил еще распараллелить задачу.
Построим индекс по `c2`
```sql
db=> create index on test.foo(c2);
```
И повторим
```sql
db=> explain (analyze) select * from test.foo where c2 like 'abcd%';
                                                    QUERY PLAN
-------------------------------------------------------------------------------------------------------------------
 Gather  (cost=1000.00..14552.33 rows=100 width=37) (actual time=17.845..66.644 rows=21 loops=1)
   Workers Planned: 2
   Workers Launched: 2
   ->  Parallel Seq Scan on foo  (cost=0.00..13542.33 rows=42 width=37) (actual time=6.452..34.044 rows=7 loops=3)
         Filter: (c2 ~~ 'abcd%'::text)
         Rows Removed by Filter: 333326
 Planning Time: 0.679 ms
 Execution Time: 66.702 ms
(8 строк)

Время: 68,765 мс
```
И все равно получаем Seq Scan. Так как мы используем поиск по шаблону, то нужно использовать другой класс операторов (https://postgrespro.ru/docs/postgrespro/9.5/indexes-opclass). 
```sql
db=> create index on test.foo(c2 text_pattern_ops);
```
```sql
db=> explain (analyze) select * from test.foo where c2 like 'abcd%';
                                                     QUERY PLAN
---------------------------------------------------------------------------------------------------------------------
 Index Scan using foo_c2_idx1 on foo  (cost=0.42..8.45 rows=100 width=37) (actual time=0.022..0.049 rows=21 loops=1)
   Index Cond: ((c2 ~>=~ 'abcd'::text) AND (c2 ~<~ 'abce'::text))
   Filter: (c2 ~~ 'abcd%'::text)
 Planning Time: 0.077 ms
 Execution Time: 0.063 ms
```
Дошли до Index Scan. А должен был быть Beatmap index scan. Хз
## Index only scan
Если выбирать только то поле, на котором есть индекс, то будет Index only scan
```sql
db=> EXPLAIN SELECT c1 FROM test.foo WHERE c1 < 500;
                                  QUERY PLAN
------------------------------------------------------------------------------
 Index Only Scan using foo_c1_idx on foo  (cost=0.42..17.09 rows=495 width=4)
   Index Cond: (c1 < 500)
(2 строки)
```
Самый быстрый
## Order by
```sql
db=> explain (analyze) select * from test.foo order by c1;
                                                            QUERY PLAN
----------------------------------------------------------------------------------------------------------------------------------
 Gather Merge  (cost=63789.50..161018.59 rows=833334 width=37) (actual time=91.245..260.070 rows=1000000 loops=1)
   Workers Planned: 2
   Workers Launched: 2
   ->  Sort  (cost=62789.48..63831.15 rows=416667 width=37) (actual time=75.998..105.490 rows=333333 loops=3)
         Sort Key: c1
         Sort Method: external merge  Disk: 19840kB
         Worker 0:  Sort Method: external merge  Disk: 13176kB
         Worker 1:  Sort Method: external merge  Disk: 12984kB
         ->  Parallel Seq Scan on foo  (cost=0.00..12500.67 rows=416667 width=37) (actual time=0.009..22.170 rows=333333 loops=3)
 Planning Time: 0.132 ms
 Execution Time: 295.604 ms
(11 строк)

Время: 296,200 мс

```
Сначала выполняется Seq Scan, а затем сортировка Sort. Чем больше отступ, тем раньше выполнялось действие.
`Sort key` - условие сортировки
`Sort method: external merge Disk` - при сортировки используется временный файл на диске указанного размера
Если же мы накинем индекс (который храни данные упорядоченно), то станет вообще конфетка
```sql
db=> create index on test.foo(c1);
CREATE INDEX
Время: 336,293 мс
db=> explain (analyze) select * from test.foo order by c1;
                                                            QUERY PLAN
-----------------------------------------------------------------------------------------------------------------------------------
 Index Scan using foo_c1_idx on foo  (cost=0.42..34317.43 rows=1000000 width=37) (actual time=0.040..133.003 rows=1000000 loops=1)
 Planning Time: 0.193 ms
 Execution Time: 163.065 ms
(3 строки)

Время: 163,732 мс
```
Стало быстрее
## Join
Еще одна таблица
```sql
CREATE TABLE bar (c1 integer, c2 boolean);
INSERT INTO bar
  SELECT i, i%2=1
  FROM generate_series(1, 500000) AS i;
ANALYZE bar;
```
Сделаем join
```sql
db=> explain (analyze) select * from test.foo join test.bar on foo.c1=bar.c1;
                                                      QUERY PLAN
-----------------------------------------------------------------------------------------------------------------------
 Hash Join  (cost=15417.00..60081.00 rows=500000 width=42) (actual time=98.246..589.664 rows=500000 loops=1)
   Hash Cond: (foo.c1 = bar.c1)
   ->  Seq Scan on foo  (cost=0.00..18334.00 rows=1000000 width=37) (actual time=0.006..69.687 rows=1000000 loops=1)
   ->  Hash  (cost=7213.00..7213.00 rows=500000 width=5) (actual time=98.093..98.094 rows=500000 loops=1)
         Buckets: 262144  Batches: 4  Memory Usage: 6562kB
         ->  Seq Scan on bar  (cost=0.00..7213.00 rows=500000 width=5) (actual time=0.006..31.115 rows=500000 loops=1)
 Planning Time: 0.135 ms
 Execution Time: 605.301 ms
(8 строк)

Время: 605,836 мс
```
Сначала последовательно просматривается таблица `bar` и для каждой строки вычисляется хэш (`Hash`). Далее последовательно сканируется таблица `foo` и для каждой строки вычисляется хэш. 
Добавим индекс
```sql
db=> create index on test.bar(c1);
CREATE INDEX
Время: 191,284 мс
db=> explain (analyze) select * from test.foo join test.bar on foo.c1=bar.c1;
                                                              QUERY PLAN
---------------------------------------------------------------------------------------------------------------------------------------
 Merge Join  (cost=2.33..39822.79 rows=500000 width=42) (actual time=0.032..286.742 rows=500000 loops=1)
   Merge Cond: (foo.c1 = bar.c1)
   ->  Index Scan using foo_c1_idx on foo  (cost=0.42..34317.43 rows=1000000 width=37) (actual time=0.010..64.833 rows=500001 loops=1)
   ->  Index Scan using bar_c1_idx on bar  (cost=0.42..15212.42 rows=500000 width=5) (actual time=0.018..94.317 rows=500000 loops=1)
 Planning Time: 0.321 ms
 Execution Time: 302.881 ms
(6 строк)

Время: 303,671 мс
```
Hash уже не используется. Теперь идет Index Scan, а потом Merge Join. Ускорили в 2 раза
