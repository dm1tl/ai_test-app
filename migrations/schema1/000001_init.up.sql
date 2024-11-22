CREATE TABLE users (
    id INT PRIMARY KEY NOT NULL,
    username VARCHAR(255) NOT NULL
);
CREATE TABLE tests (
    id SERIAL PRIMARY KEY,
    user_id INT,
    theme VARCHAR(255) NOT NULL,
    score INT DEFAULT 0,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT unique_user_theme UNIQUE (user_id, theme)
);