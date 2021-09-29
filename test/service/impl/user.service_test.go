package service

import (
	"github.com/stretchr/testify/assert"
	service "rpolnx.com.br/golang-hex/domain/service/impl"
	"testing"
)

func TestFindAllSuccess(t *testing.T) {
	mongoRepository := new(mongoRepositorySuccess)
	cacheRepository := new(cacheRepositorySuccess)

	serviceImpl := service.NewUserService(mongoRepository, cacheRepository)

	res, err := serviceImpl.GetAll()

	assert.NotNil(t, res, "Result should not be nil")
	assert.Equal(t, 2, len(res), "Result should not be equal")
	assert.Nil(t, err, "Error should be nil")
}

func TestFindUniqueSettingCacheError(t *testing.T) {
	mongoRepository := new(mongoRepositorySuccess)
	cacheRepository := new(cacheRepositorySetFail)

	serviceImpl := service.NewUserService(mongoRepository, cacheRepository)

	res, err := serviceImpl.Get("Alice")

	assert.Nil(t, res, "Result should be nil")

	assert.NotNil(t, err, "Err should not be nil")

	expectedErrorMsg := "set error"
	assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)
}
