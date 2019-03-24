create database if not exists inventario;
use inventario;

create table if not exists cajero(
    id bigint unsigned not null auto_increment,
    usuario varchar(255) not null,
    pass varchar(255) not null,
    primary key(id)
);

create table if not exists supervisor(
    id bigint unsigned not null auto_increment,
    usuario varchar(255) not null,
    pass varchar(255) not null,
    primary key(id)
);

create table if not exists producto(
    id bigint unsigned not null auto_increment,
    nombre varchar(255) not null,
    marca varchar(255) not null,
    precio integer not null,
    cantidad integer not null,
    primary key(id)
);

create table if not exists persona(
    id bigint unsigned not null auto_increment,
    nombre varchar(255) not null,
    direccion varchar(255) not null,
    primary key(id)
);

INSERT INTO persona (id, nombre, direccion) VALUES (1, "Darwin", "San jose de Chirica");
INSERT INTO persona (id, nombre, direccion) VALUES (2, "yonaikel", "Vista al sol");
INSERT INTO persona (id, nombre, direccion) VALUES (3, "yefelshon", "25 de marzo");

INSERT INTO supervisor (id, usuario, pass) VALUES (1, "admin", "admin");

INSERT INTO cajero (id, usuario, pass) VALUES (1, "darwin", "123");
INSERT INTO cajero (id, usuario, pass) VALUES (2, "jose", "123");

INSERT INTO producto (id, nombre, marca, precio, cantidad) VALUES (1, "arroz", "primor", 4000, 2);
INSERT INTO producto (id, nombre, marca, precio, cantidad) VALUES (2, "arina", "pan", 4000, 2);
INSERT INTO producto (id, nombre, marca, precio, cantidad) VALUES (3, "Mayonesa", "Mavesa", 8000, 2);
INSERT INTO producto (id, nombre, marca, precio, cantidad) VALUES (4, "Mantequilla", "Mavesa", 6000, 2);
