create table if not exists users (
    id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_name VARCHAR(255) DEFAULT "" NOT NULL,
    password_hash VARCHAR(255) DEFAULT "" NOT NULL,
    is_admin boolean not null
);