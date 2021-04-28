package service

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go-ms-session/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type AccessTokenService struct {
	mongoDb *repository.MongoRepository
	ctx     context.Context
}

func NewVerifyAccessTokenService(mongoDb *repository.MongoRepository) *AccessTokenService {
	return &AccessTokenService{
		mongoDb: mongoDb,
		ctx:     context.TODO(),
	}
}

type Message struct {
	Token string `json:"Token"`
}

func (s *AccessTokenService) VerifyToKen(c echo.Context) error {
	m := Message{}
	if err := c.Bind(&m); err != nil {
		log.Info(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	cursor := s.mongoDb.ProfileCollection.FindOne(s.ctx, bson.M{
		"token": m.Token,
	})

	doc := struct {
		Id    primitive.ObjectID `json:"_id"`
		Token string             `json:"token"`
	}{}

	if e := cursor.Decode(&doc); e != nil {
		log.Info(e)
	}

	log.Info(doc.Token)
	if doc.Token == "" {
		log.Info("token not found: ", m.Token)
		return c.JSON(http.StatusBadRequest, nil)
	}else {
		return c.JSON(http.StatusOK, m.Token)
	}

}
