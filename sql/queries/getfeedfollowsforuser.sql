-- name: GetFeedFollowsForUser :many
SELECT
    feed_follows.id AS feed_follow_id,
    feed_follows.created_at,
    feed_follows.updated_at,
    feeds.name AS feed_name,
    users.name AS user_name
FROM feed_follows
INNER JOIN feeds ON feed_follows.feed_id = feeds.id
INNER JOIN users ON feed_follows.user_id = users.id
WHERE feed_follows.user_id = $1;