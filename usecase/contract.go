package usecase

import (
	"database/sql"
	"github.com/faizalnurrozi/phincon-catch/domain/view_models"
	"github.com/faizalnurrozi/phincon-catch/packages/jwe"
	"github.com/faizalnurrozi/phincon-catch/packages/jwt"
	"github.com/faizalnurrozi/phincon-catch/packages/redis"
	"github.com/faizalnurrozi/phincon-catch/packages/watermill"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type Contract struct {
	ReqID                string
	UserID               string
	RoleID               string
	App                  *fiber.App
	DB                   *sql.DB
	Mongo                *mongo.Database
	TX                   *sql.Tx
	RedisClient          redis.RedisClient
	JweCredential        jwe.Credential
	JwtCredential        jwt.JwtCredential
	Validate             *validator.Validate
	Translator           ut.Translator
	Kafka                watermill.Kafka
}

const (
	defaultLimit    = 10
	maxLimit        = 50
	defaultOrderBy  = "id"
	defaultSort     = "asc"
	defaultLastPage = 0

	// Default for product
	DefaultDaysOfNewProduct = 30
)

func (uc Contract) SetPaginationParameter(page, limit int, order, sort string) (int, int, int, string, string) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > maxLimit {
		limit = maxLimit
	}
	if order == "" {
		order = defaultOrderBy
	}
	if sort == "" {
		sort = defaultSort
	}
	offset := (page - 1) * limit

	return offset, limit, page, order, sort
}

func (uc Contract) SetPaginationResponse(page, limit, total int) (res view_models.PaginationVm) {
	var lastPage int

	if total > 0 {
		lastPage = total / limit

		if total%limit != 0 {
			lastPage = lastPage + 1
		}
	} else {
		lastPage = defaultLastPage
	}

	vm := view_models.NewPaginationVm()
	res = vm.Build(view_models.DetailPaginationVm{
		CurrentPage: page,
		LastPage:    lastPage,
		Total:       total,
		PerPage:     limit,
	})

	return res
}
