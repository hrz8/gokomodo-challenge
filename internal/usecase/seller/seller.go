package seller

import (
	"os"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
	"github.com/hrz8/gokomodo-challenge/internal/model/dto"
	"github.com/hrz8/gokomodo-challenge/internal/model/entity"
	"github.com/hrz8/gokomodo-challenge/internal/repository/db"
	res "github.com/hrz8/gokomodo-challenge/pkg/util/response"
	"github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type (
	usecase struct {
		Repository *db.Repository
	}

	IUsecaseSeller interface {
		Register(body *dto.RegisterRequest) (*entity.Seller, error)
		FindById(id string) (*entity.Seller, error)
		Login(email string, password string) (*entity.Seller, error)
		GenerateToken(id string) (string, error)
	}
)

func (u *usecase) Register(body *dto.RegisterRequest) (*entity.Seller, error) {
	id, _ := uuid.NewV4()
	bytes, _ := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	data := &entity.Seller{
		ID:            id,
		Email:         body.Email,
		Name:          body.Name,
		Password:      string(bytes),
		PickupAddress: body.Address,
	}

	result, err := u.Repository.Seller.Create(data)
	if err != nil {
		if err.(sqlite3.Error).Code == 19 {
			return nil, res.ErrorBuilder(
				&res.ErrorConstant.BadRequest,
				err,
				"email already registered",
			)
		}
		return nil, res.ErrorBuilder(
			&res.ErrorConstant.InternalServerError,
			err,
		)
	}

	return result, nil
}

func (u *usecase) FindById(id string) (*entity.Seller, error) {
	result, err := u.Repository.Seller.FindById(id)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, res.ErrorBuilder(
				&res.ErrorConstant.NotFound,
				err,
			)
		}
		return nil, res.ErrorBuilder(
			&res.ErrorConstant.InternalServerError,
			err,
		)
	}

	return result, nil
}

func (u *usecase) Login(email string, password string) (*entity.Seller, error) {
	incorrectError := "email or password incorrect"

	result, err := u.Repository.Seller.FindByEmail(email)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, res.ErrorBuilder(
				&res.ErrorConstant.EmailOrPasswordIncorrect,
				err,
				incorrectError,
			)
		}
		return nil, res.ErrorBuilder(
			&res.ErrorConstant.InternalServerError,
			err,
		)
	}

	if result == nil {
		return nil, res.ErrorBuilder(
			&res.ErrorConstant.EmailOrPasswordIncorrect,
			err,
			incorrectError,
		)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password)); err != nil {
		return nil, res.ErrorBuilder(
			&res.ErrorConstant.EmailOrPasswordIncorrect,
			err,
			incorrectError,
		)
	}

	return result, nil
}

func (u *usecase) GenerateToken(id string) (string, error) {
	jwtKey := os.Getenv("JWT_SECRET")
	if jwtKey == "" {
		jwtKey = "secret"
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "gokomodo",
		"sub": id,
		"aud": "seller",
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", res.ErrorBuilder(
			&res.ErrorConstant.InternalServerError,
			err,
		)
	}

	return tokenString, nil
}

func NewUsecase(r *db.Repository) IUsecaseSeller {
	return &usecase{r}
}
