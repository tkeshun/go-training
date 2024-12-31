-- データベースの作成（postgresというデフォルトデータベースを使用する場合は不要）
-- CREATE DATABASE postgres;

-- 使用するデータベースを指定
\c postgres;

-- usersテーブルの作成
CREATE TABLE users (
    id SERIAL PRIMARY KEY,           -- 自動インクリメントのID
    name VARCHAR(100) NOT NULL,      -- ユーザー名（100文字以内、必須）
    email VARCHAR(255) UNIQUE NOT NULL, -- メールアドレス（一意、必須）
    password_hash TEXT NOT NULL,     -- パスワードハッシュ
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 作成日時（デフォルト値：現在時刻）
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- 更新日時（デフォルト値：現在時刻）
);

-- インデックスの作成（必要に応じて追加）
CREATE INDEX idx_users_email ON users (email);

-- ダミーデータの挿入（必要に応じて）
INSERT INTO users (name, email, password_hash) VALUES
('Alice', 'alice@example.com', 'hashed_password_1'),
('Bob', 'bob@example.com', 'hashed_password_2'),
('Charlie', 'charlie@example.com', 'hashed_password_3');
