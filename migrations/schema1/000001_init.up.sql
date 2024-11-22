CREATE TABLE users (
    id INT PRIMARY KEY NOT NULL,
    username VARCHAR(255) NOT NULL
);
CREATE TABLE tests (
    id INT PRIMARY KEY NOT NULL,
    user_id INT,
    theme VARCHAR(255) NOT NULL,
    score INT,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);