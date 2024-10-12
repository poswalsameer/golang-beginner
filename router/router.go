package router

import (
	"github.com/gorilla/mux"
	"github.com/poswalsameer/workingWithDB/controllers"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	// route to get all the videos
	router.HandleFunc("/api/videos", controllers.GetAllVIdeos).Methods("GET")

	// route to create a video
	router.HandleFunc("/api/createVideo", controllers.CreateVideo).Methods("POST")

	//route to update a video
	router.HandleFunc("/api/updateVideo/{id}", controllers.UpdateVideo).Methods("PUT")

	// route to delete single video
	router.HandleFunc("/api/deleteSingleVideo/{id}", controllers.DeleteOneVideo).Methods("DELETE")

	// route to delete all videos
	router.HandleFunc("/api/deleteAllVideos", controllers.DeleteAllVideos).Methods("DELETE")

	return router

}
