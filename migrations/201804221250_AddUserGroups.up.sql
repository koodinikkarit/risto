create table if not exists user_groups (
    id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    group_name VARCHAR(255) default "" not null,
    access_level int unsigned not null,
    created_at DATETIME,
	updated_at DATETIME NULL,
	deleted_at DATETIME NULL
);
create table if not exists user_user_groups (
    id int8 unsigned auto_increment primary key,
    user_id int8 unsigned not null,
    user_group_id int8 unsigned not null,
    created_at DATETIME
);