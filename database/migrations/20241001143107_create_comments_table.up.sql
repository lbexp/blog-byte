CREATE TABLE IF NOT EXISTS comments(
    comment_id INT(11) PRIMARY KEY AUTO_INCREMENT NOT NULL,
    post_id INT(11) NOT NULL,
    author_name VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    FOREIGN KEY (post_id) REFERENCES posts(post_id)
)
