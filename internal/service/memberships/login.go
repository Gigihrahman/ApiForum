package memberships

import (
	"context"
	"errors"
	"forumapp-restapi/internal/model/memberships"
	"forumapp-restapi/pkg/jwt"
	tokenUtil "forumapp-restapi/pkg/token"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest) (string,string,error){
	user,err := s.membershipRepo.GetUser(ctx, req.Email,"",0)

	if err !=nil{
		log.Error().Err(err).Msg("failed to get user")
		return "","", err
	}
	if user == nil{
		return  "","", errors.New("email is not exist")

	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(req.Password))
	if err !=nil{
		return  "","", errors.New("password invalid")
	}
	token, err := jwt.CreateToken(user.ID,user.Username, s.cfg.Service.SecretJWT)
	if err != nil{
		return  "", "",err
	}

	existingRefreshToken, err := s.membershipRepo.GetRefreshToken(ctx,user.ID,time.Now())

	if err != nil{
		log.Error().Err(err).Msg("error get lates refresh token")
		return  "", "", err
		}

	if existingRefreshToken!=nil{
		return token, existingRefreshToken.RefreshToken, nil

		}
	refreshToken := tokenUtil.GenerateRefreshToken()
	if refreshToken == ""{
		return token, "",errors.New("failed to generate refresh token")
	}
	err = s.membershipRepo.InsertRefreshToken(ctx,memberships.RefreshTokenModel{
		UserID: user.ID,
		RefreshToken: refreshToken,
		ExpiredAt: time.Now().Add(10*24*time.Hour),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: strconv.FormatInt(user.ID, 10),
		UpdatedBy: strconv.FormatInt(user.ID, 10),

	})
	if err !=nil{
		log.Error().Err(err).Msg("error inserting database")
		return token, refreshToken, err
	}

	return token,refreshToken, nil
}