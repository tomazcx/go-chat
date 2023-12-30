package user

import	"github.com/tomazcx/go-chat-app/internal/infra/database/userdb"

type ValidateUserUseCase struct {
	repo *userdb.UserRepository
}

func (uc *ValidateUserUseCase) Execute(login, password string) bool{
	user, err := uc.repo.FindByLogin(login)

	if err != nil {
		return false
	}

	return user.ValidatePassword(password)

}

func NewValidateUserUseCase() *ValidateUserUseCase {
	repo := userdb.NewUserRepository()
	return &ValidateUserUseCase{repo:repo}
}
