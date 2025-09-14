CREATE TABLE movies (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    year INT NOT NULL,
    genre TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);
