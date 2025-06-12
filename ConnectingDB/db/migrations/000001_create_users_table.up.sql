create table users (
    id serial primary key,
    first_name varchar(50) not null,
    last_name varchar(50) not null,
    email varchar(100) not null unique,
    created_at timestamp default current_timestamp
);

insert into users (first_name, last_name, email) values ('Ansel', 'Buckmaster', 'abuckmaster0@sbwire.com');
insert into users (first_name, last_name, email) values ('Chilton', 'Camps', 'ccamps1@constantcontact.com');
insert into users (first_name, last_name, email) values ('Elvera', 'Gunthorp', 'egunthorp2@businessweek.com');
insert into users (first_name, last_name, email) values ('Elli', 'Ambler', 'eambler3@blogger.com');
insert into users (first_name, last_name, email) values ('Kilian', 'Ivins', 'kivins4@businessweek.com');
insert into users (first_name, last_name, email) values ('Jervis', 'McGrale', 'jmcgrale5@opensource.org');
insert into users (first_name, last_name, email) values ('Helenelizabeth', 'Geistbeck', 'hgeistbeck6@wunderground.com');
insert into users (first_name, last_name, email) values ('Wyndham', 'Pavia', 'wpavia7@taobao.com');
insert into users (first_name, last_name, email) values ('Ase', 'Tather', 'atather8@squarespace.com');
insert into users (first_name, last_name, email) values ('Ferdinande', 'Batkin', 'fbatkin9@is.gd');
