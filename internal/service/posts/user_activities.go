package posts

import (
	"context"
	"errors"
	"forumapp-restapi/internal/model/posts"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

func (s *service) UpsertUserActivity(ctx context.Context,postId,userId int64,request posts.UserActivityRequest) error{
	now:= time.Now()
	model := posts.UserActivityModel{
		PostID: postId,
		UserID: userId,
		IsLiked: request.IsLiked,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: strconv.FormatInt(userId, 10) ,
		UpdatedBy: strconv.FormatInt(userId, 10)  ,

	}
	userActivity, err:= s.postRepo.GetUserActivity(ctx, model )

	if err!= nil{
		log.Error().Err(err).Msg(" errpr get user Activity")
		return err
	}
	if userActivity == nil{
		if !request.IsLiked {
			return errors.New("anda belum like")
		}
		err = s.postRepo.CreateUserActivity(ctx, model)

	} else{
		err= s.postRepo.UpdatedUserActivity(ctx, model)

	}
	if err!=nil{
		log.Error().Err(err).Msg("errror create or update user activity to database")
		return err
	}
	return nil
}