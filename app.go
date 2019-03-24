package main

import (
	"log"
	"net/http"

	"encoding/json"

	"github.com/boltdb/bolt"
	minio "github.com/minio/minio-go"
)

func apiGetMovies(writer http.ResponseWriter, req *http.Request) {
	var objects []minio.ObjectInfo
	// Create a done channel to control 'ListObjects' go routine.
	doneCh := make(chan struct{})
	// Indicate to our routine to exit cleanly upon return.
	defer close(doneCh)

	// List all objects from a bucket-name with a matching prefix.
	for object := range mc.ListObjectsV2(miniobucket, "", true, doneCh) {
		if object.Err != nil {
			log.Fatal(object.Err)
			return
		}
		objects = append(objects, object)
	}

	json, _ := json.Marshal(objects)
	writer.Write([]byte(json))
}

func apiGetPlaylists(writer http.ResponseWriter, req *http.Request) {
	var playlists []string
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Playlists"))
		c := bucket.Cursor()

		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			playlists = append(playlists, string(k))
		}
		return nil
	})

	json, _ := json.Marshal(playlists)
	writer.Write([]byte(json))
}
