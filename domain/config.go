package domain

import (
	"database/sql"
	"github.com/faizalnurrozi/phincon-catch/packages/functioncaller"
	"github.com/faizalnurrozi/phincon-catch/packages/jwe"
	"github.com/faizalnurrozi/phincon-catch/packages/jwt"
	"github.com/faizalnurrozi/phincon-catch/packages/logruslogger"
	"github.com/faizalnurrozi/phincon-catch/packages/str"
	"github.com/faizalnurrozi/phincon-catch/usecase"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	idTranslations "github.com/go-playground/validator/v10/translations/id"
	jwtFiber "github.com/gofiber/jwt/v2"
	"github.com/joho/godotenv"
	db "gitlab.com/s2-backend/packages/mysql"
	"google.golang.org/grpc"
	"log"
	"os"
)

type Config struct {
	UcContract           *usecase.Contract
	DB                   *sql.DB
	JweCredential        jwe.Credential
	JwtCredential        jwt.JwtCredential
	JwtConfig            jwtFiber.Config
	Validator            *validator.Validate
	GrpClientConn        *grpc.ClientConn
}

var (
	ValidatorDriver *validator.Validate
	Uni             *ut.UniversalTranslator
	Translator      ut.Translator
)

func LoadConfig() (res Config, err error) {
	err = godotenv.Load("../../.env")
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "load-env")
	}

	//setup db connection
	dbInfo := db.Connection{
		Host:                    os.Getenv("DB_HOST"),
		DbName:                  os.Getenv("DB_NAME"),
		User:                    os.Getenv("DB_USERNAME"),
		Password:                os.Getenv("DB_PASSWORD"),
		Port:                    os.Getenv("DB_PORT"),
		DBMaxConnection:         str.StringToInt(os.Getenv("DB_MAX_CONNECTION")),
		DBMAxIdleConnection:     str.StringToInt(os.Getenv("DB_MAX_IDLE_CONNECTION")),
		DBMaxLifeTimeConnection: str.StringToInt(os.Getenv("DB_MAX_LIFETIME_CONNECTION")),
	}
	res.DB, err = dbInfo.DbConnect()
	if err != nil {
		log.Fatal(err.Error())
	}

	//jwe credential
	res.JweCredential = jwe.Credential{
		KeyLocation: os.Getenv("JWE_PRIVATE_KEY"),
		Passphrase:  os.Getenv("JWE_PRIVATE_KEY_PASSPHRASE"),
	}

	//jwt credential
	res.JwtCredential = jwt.JwtCredential{
		TokenSecret:         os.Getenv("SECRET"),
		ExpiredToken:        str.StringToInt(os.Getenv("TOKEN_EXP_TIME")),
		RefreshTokenSecret:  os.Getenv("SECRET_REFRESH_TOKEN"),
		ExpiredRefreshToken: str.StringToInt(os.Getenv("REFRESH_TOKEN_EXP_TIME")),
	}

	//jwt config
	res.JwtConfig = jwtFiber.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
		Claims:     &jwt.CustomClaims{},
	}

	res.Validator = ValidatorDriver

	return res, err
}

func ValidatorInit() {
	en := en.New()
	id := id.New()
	Uni = ut.New(en, id)

	transEN, _ := Uni.GetTranslator("en")
	transID, _ := Uni.GetTranslator("id")

	ValidatorDriver = validator.New()

	enTranslations.RegisterDefaultTranslations(ValidatorDriver, transEN)
	idTranslations.RegisterDefaultTranslations(ValidatorDriver, transID)

	switch os.Getenv("APP_LOCALE") {
	case "en":
		Translator = transEN
	case "id":
		Translator = transID
	}
}
