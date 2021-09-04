package services

import (
	"autocomplete/app/repositories"
)

type IAutocompleteService interface {
	Search(word string) (result []string, err error)
	Insert(word string) (err error)
	Delete() (err error)
}
type AutocompleteService struct {
	keyName         string
	searchLength    int64
	redisRepository repositories.IRedisRepository
}

func NewInstanceOfAutocompleteService(redisRepository repositories.IRedisRepository, keyName string, searchLength int64) AutocompleteService {
	return AutocompleteService{redisRepository: redisRepository, keyName: keyName, searchLength: searchLength}
}

func (a *AutocompleteService) Search(word string) ([]string, error) {
	result, err := a.redisRepository.Search(word, a.searchLength)
	if err != nil {
		return []string{}, err
	}
	return result, nil
}

func (a *AutocompleteService) Insert(word string) error {
	for i := 1; i <= len(word); i++ {
		var prefix []byte
		for j := 0; j < i; j++ {
			prefix = append(prefix, word[j])
		}
		if len(prefix) != 0 {
			err := a.redisRepository.Insert(word, string(prefix))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *AutocompleteService) Delete() error {
	err := a.redisRepository.Delete()
	return err
}
