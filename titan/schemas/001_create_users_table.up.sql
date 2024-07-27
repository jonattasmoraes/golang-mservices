CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(255) PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255),
    role VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

INSERT INTO users (id, first_name, last_name, email, password, role, created_at, updated_at)
VALUES 
    ('01J1P83E4KJ3EM3H73139AC5PQ', 'John', 'Lennon', 'John.Lennon@example.com', 'password123', 'user', NOW(), NOW()),
    ('01J1PATMVDPDQJKRC037EGDFET', 'Peter', 'Parker', 'Peter.Parker@example.com', 'password123', 'user', NOW(), NOW()),
    ('01J1R7YT8C3S33XW02EV76JVHM', 'Emily', 'Van', 'Emily.Van@example.com', 'password123', 'user', NOW(), NOW()),
    ('01J1RAYJA9HCJVNW9VN9PSQ0HM', 'Penney', 'Lenn', 'Penney.Lenn@example.com', 'password123', 'user', NOW(), NOW());