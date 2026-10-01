package posts

import (
	"context"
	"database/sql"
	"forumapp-restapi/internal/model/posts"
)

func (r *repository) GetUserActivity(ctx context.Context, model posts.UserActivityModel) (*posts.UserActivityModel, error) {
	query := `SELECT id, post_id, user_id, is_liked, created_at, updated_at, created_by, updated_by FROM user_activities WHERE post_id = $1 AND user_id = $2`

	var response posts.UserActivityModel

	row := r.db.QueryRowContext(ctx, query, model.PostID, model.UserID)
	err := row.Scan(&response.ID, &response.PostID, &response.UserID, &response.IsLiked, &response.CreatedAt, &response.UpdatedAt, &response.CreatedBy, &response.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &response, nil

}

func (r *repository) CreateUserActivity(ctx context.Context, model posts.UserActivityModel) error {
	query := `INSERT INTO user_activities(post_id, user_id, is_liked, created_at, updated_at, created_by, updated_by) VALUES($1, $2, $3, $4, $5, $6, $7)`

	_, err := r.db.ExecContext(ctx, query, model.PostID, model.UserID, model.IsLiked, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdatedUserActivity(ctx context.Context, model posts.UserActivityModel) error {
	query := `UPDATE user_activities SET is_liked = $1, updated_at = $2, updated_by = $3 WHERE post_id = $4 AND user_id = $5`
	_, err := r.db.ExecContext(ctx, query, model.IsLiked, model.UpdatedAt, model.UpdatedBy, model.PostID, model.UserID)
	if err != nil {
		return err
	}
	return nil

}

func (r *repository) CountLikeByPostID(ctx context.Context, postId int64) (int, error) {
	query := `SELECT COUNT(id) FROM user_activities WHERE post_id = $1 AND is_liked = true`

	var response int

	row := r.db.QueryRowContext(ctx, query, postId)
	err := row.Scan(&response)
	if err != nil {
		if err == sql.ErrNoRows {
			return response, nil
		}
		return response, err
	}
	return response, nil

}
