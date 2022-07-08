package users

import (
	"errors"
	"learning-app/auth"
	"learning-app/common"
	"learning-app/database"
	"learning-app/email"

	"gorm.io/gorm"
)

type InvitationTokenPayload struct {
	id string
}

type UserAuthenticatePayload struct {
	Id string
}

func InviteUser(userInviteRequest UserInviteRequest) (*User, error) {
	db := database.GetInstance()

	exists := false
	db.Model(&User{}).
		Select("count(*) > 0").
		Where("email = ?", userInviteRequest.Email).
		Find(&exists)

	if exists {
		return nil, ErrEmailExists
	}

	user := User{Email: userInviteRequest.Email}

	err := db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	tokenCreationResult, err := auth.CreateToken(InvitationTokenPayload{user.ID.String()})
	if err != nil {
		return nil, err
	}

	db.Create(&UserInvitationToken{UserID: user.ID, Token: tokenCreationResult.Token})

	email.SendInvitationEmail(user.Email, tokenCreationResult.Token)

	return &user, nil
}

func ActivateUser(userActivateRequest UserActivateRequest) error {
	db := database.GetInstance()

	userInvitationToken := UserInvitationToken{}
	err := db.Joins("User").First(&userInvitationToken, "token = ?", userActivateRequest.Token).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrUserNotFound
	} else if err != nil {
		return err
	}

	_, err = auth.VerifyDecodeToken(userInvitationToken.Token)
	if err != nil {
		return auth.ErrTokenExpired
	}

	passwordHash, err := common.CreateHash(userActivateRequest.Password)
	if err != nil {
		return err
	}

	db.Model(&userInvitationToken.User).Update("is_active", true).Update("password", passwordHash)
	db.Where("token = ?", userInvitationToken.Token).Delete(&userInvitationToken)

	return nil
}

func Authenticate(userAuthenticateRequest UserAuthenticateRequest) (*auth.TokenCreationResult, error) {
	db := database.GetInstance()

	user := User{Email: userAuthenticateRequest.Email, IsActive: true}
	err := db.First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUnauthorized
		}
		return nil, err

	}

	passwordMatch := common.CompareStringWithHash(userAuthenticateRequest.Password, user.Password)
	if !passwordMatch {
		return nil, ErrUnauthorized
	}

	token, err := auth.CreateToken(&UserAuthenticatePayload{Id: user.ID.String()})
	if err != nil {
		return nil, err
	}

	return token, nil
}
