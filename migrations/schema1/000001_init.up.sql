CREATE TABLE users (
    id INT PRIMARY KEY NOT NULL,
    username VARCHAR(255) NOT NULL
);
CREATE TABLE tests (
    id INT PRIMARY KEY NOT NULL,
    user_id INT,
    test_name VARCHAR(255) NOT NULL,
    score INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
CREATE TABLE questions (
    id INT PRIMARY KEY NOT NULL,
    test_id INT,
    question VARCHAR(255) NOT NULL,
    FOREIGN KEY (test_id) REFERENCES tests(id) ON DELETE CASCADE
);
CREATE TABLE answers (
    id INT PRIMARY KEY NOT NULL,
    question_id INT,
    answer VARCHAR(255) NOT NULL,
    is_correct BOOLEAN,
    FOREIGN KEY (question_id) REFERENCES questions(id) ON DELETE CASCADE
);