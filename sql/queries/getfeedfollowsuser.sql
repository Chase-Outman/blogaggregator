-- name: GetFeedFollowsForUser :many
SELECT *,
    feeds.name AS feed_name
FROM feed_follows
INNER JOIN feeds
ON feed_follows.feed_id = feeds.id
WHERE feed_follows.user_id = $1;