DROP TABLE IF EXISTS locations;

CREATE TABLE locations(
    id serial primary key,
    address varchar(255),
    city varchar(255),
    postal_code varchar(15),
    position point,
    is_a_place boolean,
    last_record_date timestamp
);
