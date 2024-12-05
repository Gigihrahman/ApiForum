package posts

import (
	"context"
	"forumapp-restapi/internal/model/posts"

	"github.com/rs/zerolog/log"
)

func (s *service) GetPostByID(ctx context.Context,postId int64) (*posts.GetPostResponse, error){
	postDetail, err := s.postRepo.GetPostByID(ctx, postId)
	if err != nil{
		log.Error().Err(err).Msg("error get post by id to database")
		return nil, err
	
	}
	likeCount, err := s.postRepo.CountLikeByPostID(ctx, postId)
	if err != nil{
		log.Error().Err(err).Msg("error count like to database")
		return nil, err
	
	}
	comments, err:=s.postRepo.GetCommentBypostId(ctx,postId)
	if err != nil{
		log.Error().Err(err).Msg("error get comment to database")
		return nil, err
	
	}

	return &posts.GetPostResponse{
		PostDetail: posts.Post{
			ID: postDetail.ID,
			UserID: postDetail.UserID,
			Username: postDetail.Username,
			PostTitle: postDetail.PostTitle,
			PostContent: postDetail.PostContent,
			PostHastags: postDetail.PostHastags,
			IsLiked: postDetail.IsLiked,
		},
		LikeCount: likeCount,
		Comments: comments,

	}, nil

}