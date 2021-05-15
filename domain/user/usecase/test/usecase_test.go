package test

import (
	"go-clean-architecture/domain/user"
	"go-clean-architecture/domain/user/usecase"
	"go-clean-architecture/mock"
	"go-clean-architecture/util/application"
	"reflect"
	"testing"
)

// duck type
var trxMock = mock.TrxMock{}
var userRepoMock = UserRepoMock{}
var userUsecase = usecase.New(trxMock, userRepoMock)

func TestNew(t *testing.T) {
	if userUsecase == nil {
		t.Fail()
	}
}

func TestList(t *testing.T) {
	result, err := userUsecase.List(application.Query{})

	if err != nil {
		t.Errorf("error is not expected")
	}

	if !reflect.DeepEqual(result, []user.User{UserDataMock}) {
		t.Errorf("List is not returning correct data")
	}
}

func TestCreate(t *testing.T) {
	result, err := userUsecase.Create(UserDataMock)

	if err != nil {
		t.Errorf("error is not expected")
	}

	if !reflect.DeepEqual(result, UserDataMock) {
		t.Errorf("Create is not returning correct data")
	}
}
