CREATE TABLE IF NOT EXISTS boards (
    id serial PRIMARY KEY,
    title text
);

CREATE TABLE IF NOT EXISTS threads (
    id serial PRIMARY KEY,
    title text,
    post text,
    board_id integer,
    created_at timestamptz DEFAULT now(),
    FOREIGN KEY (board_id) REFERENCES boards (id)
);

CREATE TABLE IF NOT EXISTS posts (
    id serial PRIMARY KEY,
    thread_id int,
    post text,
    created_at timestamptz DEFAULT now(),
    FOREIGN KEY (thread_id) REFERENCES threads (id)
);

INSERT INTO boards (title)
SELECT
    'technology'
WHERE
    NOT EXISTS (
        SELECT
            1
        FROM
            boards
        WHERE
            title ILIKE '%technology%');

INSERT INTO boards (title)
SELECT
    'literature'
WHERE
    NOT EXISTS (
        SELECT
            1
        FROM
            boards
        WHERE
            title ILIKE '%literature%');

INSERT INTO boards (title)
SELECT
    'news'
WHERE
    NOT EXISTS (
        SELECT
            1
        FROM
            boards
        WHERE
            title ILIKE '%news%');

