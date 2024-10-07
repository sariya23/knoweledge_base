create table public.student (
	record_book numeric(5) not null,
	name text not null,
	doc_ser numeric(4),
	doc_num numeric(6),
	primary key (record_book)
);

drop table public.student;

create table public.progress (
	record_book numeric(5) not null,
	subject text not null,
	acad_year text not null,
	term numeric(1) not null check (term = 1 or term = 2),
	mark numeric(1) not null check (mark >= 3 and mark <= 5) default 5,
	foreign key (record_book) references public.student (record_book)
	on delete cascade
	on update cascade
);
alter table public.progress alter column mark set default 5;
drop table public.student;
create table public.airport (
	airport_code char(3) not null,
	airport_name text not null,
	city text not null,
	long float not null,
	lat float not null,
	tz text not null,
	primary key (airport_code)
);

create table public.flight (
	flight_id serial not null,
	flight_no char(6) not null,
	schedule_departure timestamptz not null,
	schedule_arrival timestamptz not null,
	departure_airport char(3) not null,
	arrival_airport char(3) not null,
	status varchar(20) not null,
	aircraft_code char(3) not null,
	actual_departure timestamptz,
	actual_arrival timestamptz,
	check (schedule_arrival > schedule_departure),
	check (status in ('On time', 'Delayed', 'Departed', 'Arrived', 'Scheduled', 'Cancelled')),
	check (actual_arrival is null or (actual_departure is not null and actual_arrival is not null and actual_arrival > actual_departure)),
	primary key (flight_id),
	unique (flight_no, schedule_departure),
	foreign key (aircraft_code) references public.aircrafts (aircraft_code),
	foreign key (arrival_airport) references bookings.airports_data (airport_code),
	foreign key (departure_airport) references bookings.airports_data (airport_code)
);

alter table public.aircrafts add column speed int;
table public.aircrafts;

update public.aircrafts set speed = 807 where aircraft_code='733';
update public.aircrafts set speed = 856 where aircraft_code='763';
update public.aircrafts set speed = 900 where aircraft_code='773';
update public.aircrafts set speed = 765 where aircraft_code='CR2';
update public.aircrafts set speed = 678 where aircraft_code='CN1';
update public.aircrafts set speed = 984 where aircraft_code='SU9';
update public.aircrafts set speed = 888 where aircraft_code in ('319', '320', '321');

alter table public.aircrafts alter column speed set not null;
alter table public.aircrafts add check(speed >= 300);

select aircraft_code, fare_conditions, count(*)
from public.seats
group by aircraft_code, fare_conditions
order by aircraft_code, fare_conditions;

create view seats_by_fare_cond as
	select aircraft_code, fare_conditions, count(*)
	from public.seats
	group by aircraft_code, fare_conditions
	order by aircraft_code, fare_conditions;
select * from seats_by_fare_cond;


create table public.student (
	record_book numeric(5) not null,
	name text not null,
	doc_ser numeric(4),
	doc_num numeric(6),
	who_add text default current_user,
	when_add timestamp default current_timestamp,
	primary key (record_book)
);
drop table student cascade;
insert into public.student (record_book, name, doc_ser, doc_num) values 
(12345, 'qwe', 1234, 123456);
table public.student;

alter table public.progress alter column mark drop not null;

insert into public.progress (record_book, subject, acad_year, term, mark)
values 
(12345, 'qwe', 'asd', 1, null);
table public.progress;

insert into public.progress (record_book, subject, acad_year, term)
values 
(12345, 'f', 'asd', 1);

alter table public.student add constraint unuqie_doc_ser unique(doc_ser, doc_num);
insert into public.student (record_book, name, doc_ser, doc_num) values
(23456, 'qwe', 2345, 888888);
insert into public.student (record_book, name, doc_ser, doc_num) values
(23458, 'qwe', 2345, null);
insert into public.student (record_book, name, doc_ser, doc_num) values
(23258, 'qwe', null, null);

table public.student;