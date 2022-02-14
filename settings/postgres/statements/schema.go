package statements

const SCHEMA = `
DROP SCHEMA IF EXISTS cartesian CASCADE ;

CREATE SCHEMA cartesian;
ALTER SCHEMA cartesian OWNER TO postgres;

CREATE TABLE cartesian.point
(
    x        INT NOT NULL,
    y        INT NOT NULL,
);
`
