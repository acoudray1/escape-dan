DROP TABLE IF EXISTS risk_places;

CREATE TABLE risk_places(
    id serial primary key,
    location_id int,
    number_of_records int
);
