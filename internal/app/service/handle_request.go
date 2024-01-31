package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/atharvi07/gin_practice/internal/app/dto"
	"github.com/atharvi07/gin_practice/internal/app/repository"
)

type UserService interface {
	GetAllUsers() ([]dto.Data, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (userService *userService) GetAllUsers() ([]dto.Data, error) {
	var wg sync.WaitGroup
	endpoint := "/api/user1"
	url := endpoint + "?page=1"
	// fmt.Println(">>>>>>>>>url is ", url)
	body, err := userService.userRepository.FetchApi(url)
	if err != nil {
		fmt.Println("Internal server error")
		return []dto.Data{}, err
	}
	response, err := parseData(body)
	if err != nil {
		fmt.Println("Json parse error")
		return []dto.Data{}, err
	}
	fmt.Println(">>>>>>>>>Response ", response)
	wg.Add(response.Total_pages-1)

	var allUsers []dto.Data
	allUsers = append(allUsers, response.Data...)

	if response.Total_pages > 1 {
		for i := 1; i < response.Total_pages; i++ {
			url := fmt.Sprintf("%s?page=%d", endpoint, i+1)
			// fmt.Println(">>>>>>>>>url is ", url)
			go func() {
				defer wg.Done()
				body, _ := userService.userRepository.FetchApi(url)
				response, _ := parseData(body)
				// fmt.Println(">>>>>>>>>Response in go ", response)
				allUsers = append(allUsers, response.Data...)
			}()
		}
		wg.Wait()
	}
	return allUsers, nil
}

func parseData(body []byte) (dto.Response, error) {
	var response dto.Response
	if err := json.Unmarshal(body, &response); err != nil {
		return dto.Response{}, errors.New("unable to parse the json data")
	}
	return response, nil
}

// func (userRepository *userRepository) GetAllUsers() ([]dto.Data, error) {
	

// 	return allUsers, err
// }