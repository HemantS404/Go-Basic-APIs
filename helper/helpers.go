package helper

import (
	"context"
	"example/mongo-go/controller"
	"example/mongo-go/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection = controller.GetCollection()

func insertOneMovie(movie *model.Netflix) (*mongo.InsertOneResult, error) {
	result, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func updateOneMovie(movieID string) (*mongo.UpdateResult, error) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	update, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": bson.M{"watched": true}})
	if err != nil {
		return nil, err
	}
	return update, nil
}

func deleteOneMovie(movieID string) (*mongo.DeleteResult, error) {
	id, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		return nil, err
	}
	delete, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return nil, err
	}
	return delete, nil
}

func deleteAllMovie() (*mongo.DeleteResult, error) {
	delete, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	return delete, nil
}

func getAllMovies() ([]primitive.M, error) {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	var movies []primitive.M
	for cur.Next(context.Background()) {
		var movie bson.M
		if err := cur.Decode(&movie); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies, nil
}

func GetAllMovies(c *gin.Context) {
	movies, err := getAllMovies()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, movies)
}

func InsertOneMovie(c *gin.Context) {
	var movie model.Netflix
	if err := c.BindJSON(&movie); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	id, err := insertOneMovie(&movie)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	result := fmt.Sprintf("Inserted object with id : %v", id.InsertedID)
	c.IndentedJSON(http.StatusAccepted, gin.H{"message": result})
}

func UpdateByID(c *gin.Context) {
	update, err := updateOneMovie(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	result := fmt.Sprintf("Updated objects count : %v", update.ModifiedCount)
	c.IndentedJSON(http.StatusAccepted, gin.H{"message": result})
}

func DeleteByID(c *gin.Context) {
	delete, err := deleteOneMovie(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	result := fmt.Sprintf("Deleted objects count : %v", delete.DeletedCount)
	c.IndentedJSON(http.StatusAccepted, gin.H{"message": result})
}

func DeleteAllMovie(c *gin.Context) {
	delete, err := deleteAllMovie()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	result := fmt.Sprintf("Deleted objects count : %v", delete.DeletedCount)
	c.IndentedJSON(http.StatusAccepted, gin.H{"message": result})
}
