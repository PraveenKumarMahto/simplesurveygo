package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	mgo "gopkg.in/mgo.v2"
)

type Movie struct {
	Title    string `json:"title"`
	Year     int    `json:"year"`
	Director string `json:"director"`
	Cast     string `json:"cast"`
	Genre    string `json:"genre"`
	Notes    string `json:"notes"`
}

func (p Movie) toString() string {
	return toJson(p)
}

func toJson(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		//os.Exit(1)
		return ""
	}
	return string(bytes)
}

func getMovies() []Movie {
	raw, err := ioutil.ReadFile("./data.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []Movie
	json.Unmarshal(raw, &c)
	DumpMovieDataInMongo(c)
	fmt.Println(c)
	return c

}

func DumpMovieDataInMongo(mov []Movie) {
	session, err := mgo.Dial("localhost,localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("simplesurveys").C("movies")

	length := len(mov)
	for i := 0; i < length; i++ {
		singleMovie := mov[i]
		movie := &Movie{
			Title:    singleMovie.Title,
			Year:     singleMovie.Year,
			Director: singleMovie.Director,
			Cast:     singleMovie.Cast,
			Genre:    singleMovie.Genre,
			Notes:    singleMovie.Notes,
		}

		err := c.Insert(movie)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	pages := getMovies()
	for _, p := range pages {
		fmt.Println(p.toString())
	}
	length := len(pages)
	go DumpMovieDataInMongo(pages[0 : length/4])
	go DumpMovieDataInMongo(pages[length/4 : length/2])
	go DumpMovieDataInMongo(pages[length/2 : (length*3)/4])
	go DumpMovieDataInMongo(pages[(length*3)/4 : length])
}
