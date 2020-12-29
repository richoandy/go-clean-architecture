package usecase

import (
	"go-clean-architecture/domain/user"
	"go-clean-architecture/mock"
	"go-clean-architecture/util/application"
	"reflect"
	"testing"
)

// duck type
var userRepoMock = mock.UserRepoMock{}
var userUsecase = New(userRepoMock)
var acMock = mock.AcMock{}

func TestNew(t *testing.T) {
	if userUsecase == nil {
		t.Fail()
	}
}

func TestList(t *testing.T) {
	result, err := userUsecase.List(acMock, application.Query{})

	if err != nil {
		t.Errorf("error is not expected")
	}

	if !reflect.DeepEqual(result, []user.User{mock.UserData}) {
		t.Errorf("List is not returning correct data")
	}
}

func TestCreate(t *testing.T) {
	result, err := userUsecase.Create(acMock, mock.UserData)

	if err != nil {
		t.Errorf("error is not expected")
	}

	if !reflect.DeepEqual(result, mock.UserData) {
		t.Errorf("Create is not returning correct data")
	}
}
