-- Create Users table
CREATE TABLE IF NOT EXISTS Users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(100) NOT NULL
);

-- Create Departments table
CREATE TABLE IF NOT EXISTS Departments (
    dept_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-- Create JobPositions table
CREATE TABLE IF NOT EXISTS JobPositions (
    job_id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    dept_id INT NOT NULL,
    base_salary NUMERIC(10, 2) NOT NULL,
    FOREIGN KEY (dept_id) REFERENCES Departments(dept_id)
);

-- Create Reviews table
CREATE TABLE IF NOT EXISTS Reviews (
    review_id SERIAL PRIMARY KEY,
    job_id INT NOT NULL,
    user_id INT NOT NULL,
    rating INT NOT NULL CHECK (rating >= 1 AND rating <= 5),
    text_review TEXT NOT NULL,
    raises_opportunity BOOLEAN NOT NULL,
    raise_explanation TEXT,
    FOREIGN KEY (job_id) REFERENCES JobPositions(job_id),
    FOREIGN KEY (user_id) REFERENCES Users(user_id)
);

-- Create ReviewFeedback table
CREATE TABLE IF NOT EXISTS ReviewFeedback (
    feedback_id SERIAL PRIMARY KEY,
    review_id INT NOT NULL,
    user_id INT NOT NULL,
    is_useful BOOLEAN NOT NULL,
    FOREIGN KEY (review_id) REFERENCES Reviews(review_id),
    FOREIGN KEY (user_id) REFERENCES Users(user_id)
);


