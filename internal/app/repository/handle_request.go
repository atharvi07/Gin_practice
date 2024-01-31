package repository

import (
	"io/ioutil"
	"log"
	"net/http"
)

type UserRepository interface {
	FetchApi(url string) ([]byte, error)
}

type userRepository struct {
	baseUrl string
}

func NewUserRepository(baseUrl string) UserRepository {
	return &userRepository{
		baseUrl: baseUrl,
	}
}


func (userRepository *userRepository) FetchApi(url string) ([]byte, error) {
	res, err := http.Get(userRepository.baseUrl+url)
	if err != nil {
		// fmt.Println(">>>>>>>>>url in fetchApi is ", url)
		log.Fatal("unable to get response from api")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	return body, nil
}

