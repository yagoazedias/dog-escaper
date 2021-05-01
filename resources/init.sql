create table if not exists port
(
    is_open boolean not null,
    timestamp timestamp not null,
    id serial not null
    constraint port_pk
    primary key
);

alter table port owner to postgres;

create unique index if not exists port_id_uindex
	on port (id);

create unique index if not exists port_timestamp_uindex
	on port (timestamp);