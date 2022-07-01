DROP TABLE IF EXISTS history_records;

CREATE TABLE history_records(
    user_id int,
    location_id int,
    record_date timestamp,
    danger_type danger[],
    nb_of_validations int
);
