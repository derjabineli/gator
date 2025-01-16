-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetAllFeeds :many
SELECT * FROM feeds;

-- name: GetFeedByURL :one
SELECT * FROM feeds
WHERE feeds.url = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = $2,
    updated_at = $3
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT feed_follows.feed_id, feeds.*
FROM feed_follows
RIGHT JOIN feeds 
ON feed_follows.feed_id = feeds.id
WHERE feed_follows.user_id = $1
ORDER BY feeds.last_fetched_at ASC
NULLS FIRST
LIMIT 1;