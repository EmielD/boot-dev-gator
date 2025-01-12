-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
USING feeds, users
WHERE feed_follows.feed_id = feeds.id AND feed_follows.user_id = users.id
AND users.name = $1 AND feeds.url = $2;