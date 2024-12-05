package memberships

import (
	"context"
	"errors"
	"forumapp-restapi/internal/model/memberships"
	"forumapp-restapi/pkg/jwt"
	"time"

	"github.com/rs/zerolog/log"
)

func (s *service) ValidateRefreshToken(ctx context.Context,userID int64, request memberships.RefreshTokenRequest) (string, error){

	existingRefreshToken, err:= s.membershipRepo.GetRefreshToken(ctx,userID,time.Now())
	if err != nil{
		log.Error().Err(err).Msg("error get refresh token from database")	
		return "", err
	}
	if existingRefreshToken == nil{
		
		return "", errors.New("refresh token has expired")
	}
	//not matchig token from db
	if existingRefreshToken.RefreshToken != request.Token{
		return "", errors.New("refresh token")

	}

	user,err := s.membershipRepo.GetUser(ctx, "","",userID)

	if err !=nil{
		log.Error().Err(err).Msg("failed to get user")
		return "", err
	}
	if user == nil{
		return  "", errors.New("user is not exist")

	}

	token, err := jwt.CreateToken(user.ID,user.Username, s.cfg.Service.SecretJWT)
	if err != nil{
		return  "", err
	}
	return token, nil

}