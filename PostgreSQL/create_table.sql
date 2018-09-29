create table users ( 
	id SERIAL,
	name varchar(20),
	line_id varchar(50),
	student_id varchar(8),
	primary key(id)
);
create table genres (
	id SERIAL,
	name varchar(30),
	primary key(id)
);
create table user_genre (
	id SERIAL,
	user_id integer references users(id),
	genre_id integer references genres(id),
	primary key(id)
);
insert into genres (name) values ('C');
insert into genres (name) values ('Java');
insert into genres (name) values ('C++');
insert into genres (name) values ('Ruby on Rails');
insert into genres (name) values ('Python');
insert into genres (name) values ('Swift');
insert into genres (name) values ('HTML');
insert into genres (name) values ('CSS');
insert into genres (name) values ('JavaScript');
insert into genres (name) values ('Kotlin');
insert into genres (name) values ('React');
insert into genres (name) values ('Vue');
insert into genres (name) values ('PHP');
insert into genres (name) values ('Go');
insert into genres (name) values ('SQL');
insert into genres (name) values ('Ruby');
insert into genres (name) values ('JavaScript');
insert into genres (name) values ('Unity');
insert into genres (name) values ('Perl');
insert into genres (name) values ('Git');
insert into genres (name) values ('Raspberry Pi');
insert into genres (name) values ('AWS');
insert into genres (name) values ('AI');
insert into genres (name) values ('Deep Learning');
insert into genres (name) values ('画像解析');