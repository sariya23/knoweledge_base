`UNION` - объеденяет результаты нескольких запросов. Важно, чтобы кол-во колонок в этих запросах было одинаковым, а также совпадали их типы
```sql
SELECT id, surname, name
FROM Students

UNION

SELECT id, surname, name
FROM Teachers;
```
![[Pasted image 20250822181639.png]]
Алиасы определяются первым UNION. 
## DISTINCT, ALL
По умолчанию `UNION` удаляет повторяющиеся строки. Но можно добавить ключевое слово `ALL` и тогда дублирование будет
```sql
SELECT surname
FROM Students

union all

SELECT surname
FROM Teachers;
```
## Limit, order by
Чтобы использовать `limit` или `order by`, то нужно брать запрос в скобки
```sql
(SELECT id, name, surname
FROM Students
LIMIT 1)

UNION

(SELECT id, name, surname
FROM Teachers
LIMIT 1);


(SELECT id, name, surname
FROM Students
ORDER BY id)

UNION

(SELECT id, name, surname
FROM Teachers
ORDER BY id);
```
Чтобы применить `LIMIT` к общему результату, нужно указать его без скобок
```sql
(SELECT id, name, surname
FROM Students
LIMIT 1)

UNION

SELECT id, name, surname
FROM Teachers
LIMIT 1;
```
## UNION как таблица
Чтобы использовать результат UNION как таблицу для выборки
```sql
SELECT name, surname
FROM (SELECT name, surname FROM Students
      UNION
      SELECT name, surname FROM Teachers) AS StudentsTeachers;
```