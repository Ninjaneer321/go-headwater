package headwater

import (
	"context"
	"fmt"
	"testing"
)

type User struct {
	Id   string
	Name string
}

type GetUserQuery struct {
	Id string
}

type GetUserReceiver = Receiver[*GetUserQuery, *User]

type FullMediator struct {
	GetUser GetUserReceiver
}

var id = "1"
var user = User{Id: id}

func CreateGetUserHandler(mediator *FullMediator) {
	var users = map[string]*User{id: &user}

	mediator.GetUser.SetHandler(func(context context.Context, request *GetUserQuery) (*User, error) {
		user := users[request.Id]
		if user == nil {
			return nil, fmt.Errorf("no user for id: %v", request.Id)
		}
		return user, nil
	})
}

func TestMediator(t *testing.T) {

	mediator := FullMediator{}

	CreateGetUserHandler(&mediator)

	result, error := mediator.GetUser.Send(context.TODO(), &GetUserQuery{Id: id})

	if error != nil {
		t.Errorf("Received error: %v", error)
	}

	want := &user

	if result != want {
		t.Errorf("Received: %v, Wanted: %v", result, want)
	}
}
