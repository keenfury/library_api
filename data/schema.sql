create table book (
    id serial,
    author varchar(100) not null,
    image_link varchar(250) null,
    lang varchar(50) null,
    link varchar(250) null,
    pages int not null,
    title varchar(250) not null,
    year int not null,
    primary key(id)
);