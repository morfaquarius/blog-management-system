-- Создаем индексы
CREATE INDEX IF NOT EXISTS idx_posts_author_created_at
    ON posts(author_id, created_at DESC, id DESC);
CREATE INDEX IF NOT EXISTS idx_comments_post_created_at
    ON comments(post_id, created_at ASC, id ASC);
CREATE INDEX IF NOT EXISTS idx_posts_created_at
    ON posts(created_at DESC, id DESC);
