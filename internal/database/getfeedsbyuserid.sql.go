// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: getfeedsbyuserid.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const getFeedsByUserId = `-- name: GetFeedsByUserId :many
SELECT id, created_at, updated_at, name, url, user_id, last_fetched_at FROM feeds
WHERE user_id = $1
`

func (q *Queries) GetFeedsByUserId(ctx context.Context, userID uuid.UUID) ([]Feed, error) {
	rows, err := q.db.QueryContext(ctx, getFeedsByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Feed
	for rows.Next() {
		var i Feed
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Url,
			&i.UserID,
			&i.LastFetchedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
