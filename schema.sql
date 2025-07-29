CREATE TABLE IF NOT EXISTS students (
                                        id SERIAL PRIMARY KEY,
                                        name VARCHAR(100) NOT NULL,
    age INT NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL
    );
