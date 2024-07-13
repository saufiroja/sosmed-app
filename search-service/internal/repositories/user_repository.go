package repositories

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/saufiroja/sosmed/search-service/internal/models"
	"github.com/saufiroja/sosmed/search-service/pkg/database"
)

type userRepository struct {
	elastic *database.Elasticsearch
}

func NewUserRepository(elastic *database.Elasticsearch) UserRepositoryInterface {
	return &userRepository{elastic: elastic}
}

func (u *userRepository) InsertUser(ctx context.Context, user *models.User) error {
	db := u.elastic.DBConnecion()

	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(user); err != nil {
		return err
	}

	req := esapi.IndexRequest{
		Index:      "users",
		DocumentID: user.UserID,
		Body:       &buf,
		Refresh:    "true",
	}

	res, err := req.Do(ctx, db)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return nil
}

func (u *userRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	db := u.elastic.DBConnecion()
	req := esapi.GetRequest{
		Index:      "users",
		DocumentID: username,
	}

	res, err := req.Do(ctx, db)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var user models.User
	if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) GetAllUsers(ctx context.Context, page, limit *int32) ([]models.User, error) {
	db := u.elastic.DBConnecion()

	query := map[string]interface{}{
		"from": &page,
		"size": &limit,
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}

	res, err := db.Search(
		db.Search.WithContext(ctx),
		db.Search.WithIndex("users"),
		db.Search.WithBody(&buf),
		db.Search.WithTrackTotalHits(true),
	)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	// hits > hits
	var hits map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&hits); err != nil {
		return nil, err
	}

	var users []models.User
	for _, hit := range hits["hits"].(map[string]interface{})["hits"].([]interface{}) {
		source := hit.(map[string]interface{})["_source"]
		user := models.User{
			UserID:   source.(map[string]interface{})["user_id"].(string),
			Username: source.(map[string]interface{})["username"].(string),
			FullName: source.(map[string]interface{})["full_name"].(string),
		}

		users = append(users, user)
	}

	return users, nil
}
