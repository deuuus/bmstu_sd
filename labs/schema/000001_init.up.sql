CREATE TABLE pacient
(
    p_id serial primary key,
    name varchar(255) not null,
    second_name varchar(255) not null
);

CREATE TABLE doctor
(
    d_id serial primary key,
    name varchar(255) not null,
    second_name varchar(255) not null,
    middle_name varchar(255) not null,
    specialization varchar(255) not null
);

CREATE TABLE record
(
    r_id serial primary key,
    date varchar(255) not null,
    doctor_id int references doctor(d_id) on delete cascade not null,
    pacient_id int references pacient(p_id) on delete cascade not null,
    diagnosis varchar(255) not null
);
