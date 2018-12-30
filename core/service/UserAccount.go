package service

import (
	"github.com/gsoultan/uam/core/domain"
	g "github.com/gsoultan/uam/core/repository/gorm"
	"github.com/gsoultan/uam/core/utils"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type userAccount struct {
	db *gorm.DB
}

type Login struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (this *userAccount) Registration(user domain.User) error {
	db := this.db.Begin()
	userRepository := g.NewUserGorm(db)
	accountRepository := g.NewAccountGorm(db)

	if err := accountRepository.Create(&user.Account); err != nil {
		db.Rollback()
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		db.Rollback()
		return err
	}

	user.Password = string(hash)
	if err := userRepository.Create(&user); err != nil {
		db.Rollback()
		return err
	}

	db.Commit()
	return nil

}

func (this *userAccount) Login(login Login) (bool, domain.User, error) {
	pass := []byte(login.Password)

	userRepository := g.NewUserGorm(this.db)
	user, err := userRepository.FindByUserName(login.UserName)
	if err != nil {
		return false, domain.User{}, err
	}

	err = bcrypt.CompareHashAndPassword(pass, []byte(user.Password))
	if err != nil {
		return false, domain.User{}, err
	}

	token, err := bcrypt.GenerateFromPassword([]byte(utils.RandomString(20)), bcrypt.MinCost)
	if err != nil {
		return false, domain.User{}, err
	}

	user.Token = string(token)
	if err := userRepository.SetTokenByID(user.Token, user.ID); err != nil {
		return false, domain.User{}, err
	}

	return true, user, nil
}

func (this *userAccount) LogOut(token string) (bool, error) {
	userRepository := g.NewUserGorm(this.db)
	user, err := userRepository.FindByToken(token)
	if err != nil {
		return false, err
	}
	u := domain.User{}
	u.ID = user.ID
	if err := userRepository.SetTokenByID(token, user.ID); err != nil {
		return false, err
	}
	return true, nil
}

func (this *userAccount) Authenticate(token string) (bool, error) {
	userRepository := g.NewUserGorm(this.db)
	user, err := userRepository.FindByToken(token)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(token), []byte(user.Token))
	if err != nil {
		return false, err
	}

	return true, nil
}
