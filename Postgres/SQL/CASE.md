Эту штуку можно использовать в агрегатных функциях
```sql
select
	count(case when weekday(submit_date) in (5, 6) then 1 end) weekend_count,
	count(case when weekday(submit_date) not in (5, 6) then 1 end) working_count
from Actions;
```
И даже в `ORDER BY`
```sql
select suit, rankvalue
from Suits, Ranks
order by 
	case 
		when suit = 'Spades' then 1
		when suit = 'Clubs' then 2
		when suit = 'Diamonds' then 3
		else 4 
	end, 
	case 
		when rankvalue='Jack' then 11
		when rankvalue='Queen' then 12
		when rankvalue='King' then 13
		when rankvalue='Ace' then 14
		else rankvalue::int
	end;
```