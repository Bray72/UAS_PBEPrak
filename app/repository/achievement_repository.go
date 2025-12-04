package repository

import (
	"clean-arch/app/model"
	"context"
	"database/sql"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AchievementRepository interface {
	GetAchievementsByUserID(userID string) ([]*model.Achievement, error)
	GetAchievementByID(id string) (*model.Achievement, error)
	CreateAchievement(achievement *model.Achievement) (*model.Achievement, error)
	UpdateAchievement(id string, achievement *model.Achievement) error
	DeleteAchievement(id string) error
	SubmitAchievement(id string, userID string) error
	// PostgreSQL methods for reference
	CreateAchievementReference(ref *model.AchievementPostgres) error
	UpdateAchievementStatus(mongoID string, status string) error
}

type achievementRepository struct {
	mongoDB *mongo.Database
	pgDB    *sql.DB
}

func NewAchievementRepository(mongoDB *mongo.Database, pgDB *sql.DB) AchievementRepository {
	return &achievementRepository{
		mongoDB: mongoDB,
		pgDB:    pgDB,
	}
}

func (r *achievementRepository) GetAchievementsByUserID(userID string) ([]*model.Achievement, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := r.mongoDB.Collection("achievements")
	filter := bson.M{"user_id": userID}
	opts := options.Find().SetSort(bson.M{"created_at": -1})

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var achievements []*model.Achievement
	if err = cursor.All(ctx, &achievements); err != nil {
		return nil, err
	}

	if achievements == nil {
		achievements = make([]*model.Achievement, 0)
	}

	return achievements, nil
}

func (r *achievementRepository) GetAchievementByID(id string) (*model.Achievement, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid achievement id format")
	}

	collection := r.mongoDB.Collection("achievements")
	filter := bson.M{"_id": objID}

	var achievement *model.Achievement
	err = collection.FindOne(ctx, filter).Decode(&achievement)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("achievement not found")
		}
		return nil, err
	}

	return achievement, nil
}

func (r *achievementRepository) CreateAchievement(achievement *model.Achievement) (*model.Achievement, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	achievement.Status = "draft"
	achievement.CreatedAt = time.Now()
	achievement.UpdatedAt = time.Now()

	collection := r.mongoDB.Collection("achievements")
	result, err := collection.InsertOne(ctx, achievement)
	if err != nil {
		return nil, err
	}

	achievement.ID = result.InsertedID.(primitive.ObjectID)
	return achievement, nil
}

func (r *achievementRepository) UpdateAchievement(id string, achievement *model.Achievement) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid achievement id format")
	}

	achievement.UpdatedAt = time.Now()
	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": achievement,
	}

	collection := r.mongoDB.Collection("achievements")
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("achievement not found")
	}

	return nil
}

func (r *achievementRepository) DeleteAchievement(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid achievement id format")
	}

	collection := r.mongoDB.Collection("achievements")
	filter := bson.M{"_id": objID, "status": "draft"}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("achievement not found or cannot be deleted (only draft can be deleted)")
	}

	return nil
}

func (r *achievementRepository) SubmitAchievement(id string, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid achievement id format")
	}

	now := time.Now()
	filter := bson.M{"_id": objID, "user_id": userID, "status": "draft"}
	update := bson.M{
		"$set": bson.M{
			"status":      "submitted",
			"submit_date": now,
			"updated_at":  now,
		},
	}

	collection := r.mongoDB.Collection("achievements")
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("achievement not found or invalid status")
	}

	return nil
}

func (r *achievementRepository) CreateAchievementReference(ref *model.AchievementPostgres) error {
	query := `INSERT INTO achievements (id, user_id, mongo_id, title, status, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := r.pgDB.Exec(query,
		ref.ID,
		ref.UserID,
		ref.MongoID,
		ref.Title,
		ref.Status,
		ref.CreatedAt,
		ref.UpdatedAt,
	)

	return err
}

func (r *achievementRepository) UpdateAchievementStatus(mongoID string, status string) error {
	query := `UPDATE achievements SET status = $1, updated_at = $2 WHERE mongo_id = $3`
	_, err := r.pgDB.Exec(query, status, time.Now(), mongoID)
	return err
}
