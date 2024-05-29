CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE
);

-- Insert some initial data
INSERT INTO users (first_name,last_name, email) VALUES ('John', 'Doe', 'john.doe@example.com');
INSERT INTO users (first_name,last_name, email) VALUES ('Jane', 'Smith', 'jane.smith@example.com');