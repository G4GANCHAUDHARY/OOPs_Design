package main

import (
	"fmt"
	"sync"
)

var (
	UserId  = 0
	MovieId = 0
)

type Users struct {
	id int
}

func NewUser() *Users {
	UserId += 1
	return &Users{id: UserId}
}

func (u *Users) GetId() int {
	return u.id
}

type Movie struct {
	id        int
	avgRating float64
	mu        sync.Mutex
}

func NewMovie() *Movie {
	MovieId += 1
	return &Movie{
		id:        MovieId,
		avgRating: 0,
		mu:        sync.Mutex{},
	}
}

func (m *Movie) GetId() int {
	return m.id
}

func (m *Movie) GetRating() float64 {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.avgRating
}

func (m *Movie) SetRating(rating float64) {
	m.avgRating = rating
}

func (m *Movie) RateMovie(users *Users, system *RecommendationSystem, rating int) error {

	if rating < 1 || rating > 5 {
		return fmt.Errorf("invalid rating for a movie : %v, plz rate from 1 to 5", rating)
	}

	m.SetRating((m.GetRating() + float64(rating)) / 2)

	userRatings := system.GetUserRatings()
	if _, found := userRatings[users.GetId()]; !found {
		userRatings[users.GetId()] = map[int]int{m.GetId(): rating}
	} else {
		userRatings[users.GetId()][m.GetId()] = rating
	}

	movieRatingMap := system.GetMovieUserRatingMap()
	if _, found := movieRatingMap[m.GetId()]; !found {
		movieRatingMap[m.GetId()] = make([][]int, 5)
		movieRatingMap[m.GetId()][rating-1] = []int{users.GetId()}
	} else {
		movieRatingMap[m.GetId()][rating-1] = append(movieRatingMap[m.GetId()][rating-1], users.GetId())
	}

	return nil
}

type RecommendationSystem struct {
	userRatings        map[int]map[int]int
	movieUserRatingMap map[int][][]int
	mu                 sync.Mutex
}

func NewRecommendationSystem() *RecommendationSystem {
	return &RecommendationSystem{
		userRatings:        make(map[int]map[int]int, 0),
		movieUserRatingMap: make(map[int][][]int, 0),
		mu:                 sync.Mutex{},
	}
}

func (r *RecommendationSystem) GetUserRatings() map[int]map[int]int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.userRatings
}

func (r *RecommendationSystem) GetMovieUserRatingMap() map[int][][]int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.movieUserRatingMap
}

func (r *RecommendationSystem) SetUserRatings(ratingsMap map[int]map[int]int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.userRatings = ratingsMap
}

func (r *RecommendationSystem) SetMovieUserRatingMap(movieMap map[int][][]int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.movieUserRatingMap = movieMap
}

func (r *RecommendationSystem) MovieRecommendations(user *Users) []int {

	allUserRatings := r.GetUserRatings()
	userMovieMap := allUserRatings[user.GetId()]
	allMovieRatingMap := r.GetMovieUserRatingMap()

	similarUserList := make(map[int]int, 0)

	for movie, rating := range userMovieMap {
		for mov, ratingUserList := range allMovieRatingMap {
			if movie == mov {
				for _, userId := range ratingUserList[rating-1] {
					if userId != user.GetId() {
						similarUserList[userId] = 1
					}
				}
			}
		}
	}

	movieList := make([]int, 0)
	for userId := range similarUserList {
		for movie := range allUserRatings[userId] {
			if _, found := userMovieMap[movie]; !found {
				movieList = append(movieList, movie)
			}
		}
	}

	return movieList
}

func main() {
	user1 := NewUser()
	user2 := NewUser()
	user3 := NewUser()

	movie1 := NewMovie()
	movie2 := NewMovie()
	movie3 := NewMovie()

	system := NewRecommendationSystem()
	movie1.RateMovie(user1, system, 5)
	movie2.RateMovie(user1, system, 2)
	movie2.RateMovie(user2, system, 2)
	movie3.RateMovie(user2, system, 4)

	for _, value := range system.MovieRecommendations(user1) {
		fmt.Printf("user 1 : %v", value)
	}

	for _, value := range system.MovieRecommendations(user2) {
		fmt.Printf("user 2 : %v", value)
	}

	for _, value := range system.MovieRecommendations(user3) {
		fmt.Printf("user 3 : %v", value)
	}

}
