package simple

import "errors"

//======================//
//contoh dependecy injection
//======================//

type SimpleRepository struct {
	Error bool
}

// provider
func NewSimpleRepository(isError bool) *SimpleRepository {
	return &SimpleRepository{
		Error: isError,
	}
}

type SimpleService struct {
	*SimpleRepository
}

// provider
func NewSimpleService(repository *SimpleRepository) (*SimpleService, error) {
	if repository.Error {
		return nil, errors.New("repository error")
	} else {
		return &SimpleService{
			SimpleRepository: repository, //dependency untuk return
		}, nil
	}
}
