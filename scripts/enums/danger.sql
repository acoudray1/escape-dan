DROP TYPE IF EXISTS danger;

CREATE TYPE danger AS ENUM(
    'none',
    'theft',
    'mugger',
    'fire',
    'ghb'
);
