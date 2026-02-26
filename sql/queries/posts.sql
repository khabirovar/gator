-- name: CreatePost :one
INSERT INTO posts(title, url, description, published_at, feed_id)
VALUES($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetPostsForUser :many
SELECT posts.title, posts.url, posts.description FROM posts
JOIN feeds ON posts.feed_id = feeds.id
JOIN feed_follows ON feeds.id = feed_follows.feed_id
WHERE feed_follows.user_id = $1
ORDER BY posts.published_at DESC NULLS LAST, posts.created_at DESC
LIMIT $2;

