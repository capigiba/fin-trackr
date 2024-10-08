CREATE TABLE IF NOT EXISTS transactions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    userid INTEGER,
    year INTEGER,
    month INTEGER,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
    description TEXT,
    amount DECIMAL(8,2),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(userid) REFERENCES users(id)
);
