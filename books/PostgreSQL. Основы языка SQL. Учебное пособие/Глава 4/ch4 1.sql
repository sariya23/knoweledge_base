create temp table test_numeric (
	meas numeric(5, 2),
	descr text
);

insert into test_numeric values
(998.9999, 'qwd');

table test_numeric;

drop table test_numeric;

create temp table test_numeric (
	meas numeric,
	descr text
);

insert into test_numeric values 
(361237612637.435435, 'qwdqwd'),
(1.5, 'qwdqwd'),
(0.123456789, 'qwdqwd'),
(1234, 'qwd');


select 'NaN'::numeric > 10000;

create temp table test_serial (
	id serial,
	name text
);

insert into test_serial(name) values 
('qwd'),
('asd'),
('qwdwqd');

table test_serial;

insert into test_serial values
(10, 'qwd');

insert into test_serial(name) values
('qwe');

create temp table test_serial2 (
	id serial primary key,
	name text
);

insert into test_serial2 (name) values
('cherry');
table test_serial2;

insert into test_serial2 values 
(2, 'apple');

insert into test_serial2 (name) values
('qwd');

show datestyle;

select 'Feb 29, 2015'::date;

select pg_typeof('2016-09-16'::date - '2016-09-01'::date);

select '20:34:35'::time + '19:44:45'::time;

select '2016-01-31'::date + '30 day'::interval;
select '2016-02-29'::date + '30 day'::interval;

select '2016-09-16'::date - '2015-09-01'::date;
select '2016-09-16'::timestamp - '2015-09-01'::timestamp;

select '20:34:35'::time - '1 hour'::interval;
select '2016-09-16'::date - 1;

create temp table dbs (
	is_open bool,
	dbms_name text
);

insert into dbs values (true, 'psql');

select * from dbs where is_open <> '1';

create temp table tmp (
	a bool,
	b text 
);
table tmp;
select * from tmp where b = 'true';
insert into tmp values (true, 'yes');
insert into tmp values (yes, 'yes');
insert into tmp values ('yes', true);
insert into tmp values ('1', 'true');
insert into tmp values (1, 'yes');
insert into tmp values ('t', 'yes');
insert into tmp values (true, true);
insert into tmp values (1::bool, 'yes');
insert into tmp values (111::bool, 'yes');

create temp table births (
	person text not null,
	birth date not null
);

insert into births values
('Ken', '1955-03-23'::date),
('Ben', '1971-03-19'::date),
('Andy', '1987-08-12'::date);

select * from births where extract('month' from birth) = 3;
select * from births where current_date - birth >= 40;

create temp table pilots (
	pilot_name text,
	schedule int[],
	meal text[]
);
insert into pilots values
('Ivan', '{1, 3, 5, 6, 7}'::int[], '{"meat", "rice"}'::text[]),
('Petr', '{2, 3, 4, 5, 6}'::int[], '{"spagetti", "bowl"}'::text[]);

table pilots;

select * from pilots
where meal[2] = 'rice';