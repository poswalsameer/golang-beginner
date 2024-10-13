package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/poswalsameer/workingWithDB/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbConnectionString = "removed-the-string-HEHEHE"

const dbName = "MeTube"
const collectionName = "watchList"

var collections *mongo.Collection

// connect with mongodb
func init() {

	//provide client options
	clientOptions := options.Client().ApplyURI(dbConnectionString)

	// context.TODO() to keep the db running in the background without any deadline
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("Error while connecting to database: ", err)
	}

	fmt.Println("MongoDB connection successful")

	collections = client.Database(dbName).Collection(collectionName)

	//collection instance
	fmt.Println("Collection instance is ready: ")
}

//MongoDB helpers

// insert 1 record
func insertOneData(movie models.VideoLibrary) {
	inserted, err := collections.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal("Error while inserting one value: ", err)
	}

	fmt.Println("Inserted 1 video in the db: ", inserted.InsertedID)
}

// update 1 record
func updateOneData(videoId string) {
	id, _ := primitive.ObjectIDFromHex(videoId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collections.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal("error while updating the record: ", err)
	}

	fmt.Println("Updated the record: ", result)

}

// delete one record
func deleteOneData(videoId string) {
	id, _ := primitive.ObjectIDFromHex(videoId)
	filter := bson.M{"_id": id}
	deleteCount, err := collections.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal("Error while deleting a data: ", err)
	}

	fmt.Println("data got deleted with delete count: ", deleteCount)
}

// delete all record
func deleteAllData() int64 {
	filter := bson.D{{}}
	deleteResult, err := collections.DeleteMany(context.Background(), filter, nil) //nil is optional

	if err != nil {
		log.Fatal("Error while deleting all data: ", err)
	}

	fmt.Println("Number of videos deleted: ", deleteResult.DeletedCount)

	return deleteResult.DeletedCount
}

// get all collections from mongoDB
func getAllCollections() []primitive.M {
	cursor, err := collections.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal("Error while fetching all the data: ", err)
	}

	var allVideos []primitive.M

	for cursor.Next(context.Background()) {

		var video bson.M
		err := cursor.Decode(&video)
		if err != nil {
			log.Fatal("error while decodiing inside the loop: ", err)
		}
		allVideos = append(allVideos, video)

	}

	defer cursor.Close(context.Background())

	return allVideos

}

// Actual controller - in a different file

// function starting with capital is exported as default
func GetAllVIdeos(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	// getting all the videos from the getAllCollection() function
	allVideos := getAllCollections()

	// sending json response
	json.NewEncoder(w).Encode(allVideos)

}

func CreateVideo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var currVideo models.VideoLibrary
	json.NewDecoder(r.Body).Decode(&currVideo)

	// using the created function to add videos
	insertOneData(currVideo)

	json.NewEncoder(w).Encode(currVideo)
}

func UpdateVideo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "appliation/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	// var currVideo models.VideoLibrary
	params := mux.Vars(r)
	updateOneData(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneVideo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneData(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteAllVideos(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	count := deleteAllData()
	json.NewEncoder(w).Encode(count)

}

func main() {

}
