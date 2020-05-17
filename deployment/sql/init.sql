CREATE TABLE IF NOT EXISTS threads (
    id integer,
    board_id integer,
    title text
);

CREATE TABLE IF NOT EXISTS posts (
    id serial PRIMARY KEY,
    post text,
    replyTo integer
);

CREATE TABLE IF NOT EXISTS boards (
    id serial PRIMARY KEY,
    title text
);

