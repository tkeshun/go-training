CREATE SCHEMA IF NOT EXISTS group_a;
SET search_path = group_a;

-- Userテーブル
CREATE TABLE Users (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    password_hash VARCHAR NOT NULL,
    role VARCHAR NOT NULL DEFAULT 'REGULAR_USER' -- 'ADMIN' or 'REGULAR_USER'
);

-- Todoテーブル
CREATE TABLE IF NOT EXISTS Todos (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title VARCHAR NOT NULL,
    description TEXT,
    status VARCHAR NOT NULL DEFAULT 'UNFINISHED', -- 'UNFINISHED' or 'COMPLETED'
    priority VARCHAR NOT NULL DEFAULT 'MEDIUM', -- 'LOW', 'MEDIUM', 'HIGH'
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- UserとTodoの交差エンティティ
CREATE TABLE IF NOT EXISTS UserTodos (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES Users(id) ON DELETE CASCADE,
    todo_id BIGINT NOT NULL REFERENCES Todos(id) ON DELETE CASCADE,
    CONSTRAINT user_todos_user_todo_unique UNIQUE (user_id, todo_id) -- user_idとtodo_idのペアは一意
);

-- Tagテーブル
CREATE TABLE IF NOT EXISTS Tags (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR UNIQUE NOT NULL
);

-- TodoとTagの交差エンティティ
CREATE TABLE IF NOT EXISTS TodoTags (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    todo_id BIGINT NOT NULL REFERENCES Todos(id) ON DELETE CASCADE,
    tag_id BIGINT NOT NULL REFERENCES Tags(id) ON DELETE CASCADE,
    CONSTRAINT todos_tags_todo_tag_unique UNIQUE (todo_id, tag_id) -- todo_idとtag_idのペアは一意
);
