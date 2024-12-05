package posts

import (
	"context"
	"forumapp-restapi/internal/configs"
	"forumapp-restapi/internal/model/posts"
)

type postRepository interface {
	CreatePost(ctx context.Context, model posts.PostModel) error
	CreateComment(ctx context.Context, model posts.CommentModel) error
	GetUserActivity(ctx context.Context, model posts.UserActivityModel)(*posts.UserActivityModel, error) 
	CreateUserActivity(ctx context.Context,model  posts.UserActivityModel) error
	UpdatedUserActivity(ctx context.Context, model posts.UserActivityModel) error
	GetAllPost(ctx context.Context, limit, offset int) (posts.GetAllPostResponse, error)
	GetPostByID(ctx context.Context,id int64) (*posts.Post , error)
	CountLikeByPostID(ctx context.Context,postId int64)(int, error)
	GetCommentBypostId(ctx context.Context, postId int64) ([]posts.Comment, error)
}
type service struct {
	cfg      *configs.Config
	postRepo postRepository
}

func NewService(cfg *configs.Config, postRepo postRepository) *service {
	return &service{
		cfg:       cfg,
		postRepo: postRepo,
	}
}