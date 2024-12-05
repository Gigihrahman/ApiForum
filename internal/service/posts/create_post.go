package posts

import (
	"context"
	"forumapp-restapi/internal/model/posts"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

func (s *service) CreatePost(ctx context.Context,userID int64, req posts.CreatePostRequest) error{
	postHastags:= strings.Join( req.PostHastags, ",")
	now:= time.Now()
	model:=posts.PostModel{
		UserID: userID ,
		PostTitle: req.PostTitle,
		PostContent: req.PostContent,
		PostHastags: postHastags,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: strconv.FormatInt(userID,10),
		UpdatedBy: strconv.FormatInt(userID,10),
	}
	err:= s.postRepo.CreatePost(ctx,model)
	if err != nil{
		log.Error().Err(err).Msg("error create")
		return err

	
	}
	return nil
}