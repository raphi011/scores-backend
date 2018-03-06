package sqlite

import (
	"log"
	"testing"

	"github.com/raphi011/scores"
)

func TestSetPassword(t *testing.T) {
	s := createServices()
	defer Reset(s.db)

	email := "test@test.com"

	user, err := s.userService.Create(&scores.User{
		Email:           email,
		ProfileImageURL: "image.url",
	})

	pw := []byte("password")

	info, _ := s.pwService.HashPassword(pw)

	log.Printf("%v %v", info, user)

	err = s.userService.UpdatePasswordAuthentication(user.ID, info)

	if err != nil {
		t.Errorf("userService.UpdatePasswordAuthentication(), err: %s", err)
	}

	user, err = s.userService.User(user.ID)

	if err != nil {
		t.Errorf("userService.User(), err: %s", err)
	}

	if !s.pwService.ComparePassword(pw, &user.PasswordInfo) {
		t.Error("PasswordService.ComparePassword(), want true, got false")
	}
}

func TestCreateUser(t *testing.T) {
	s := createServices()
	defer Reset(s.db)

	email := "test@test.com"

	user, err := s.userService.Create(&scores.User{
		Email:           email,
		ProfileImageURL: "image.url",
	})

	if err != nil {
		t.Errorf("userService.Create() err: %s", err)
	}

	if user.ID == 0 {
		t.Errorf("userService.Create(), want ID != 0, got 0")
	}

	userID := user.ID

	user, _ = s.userService.ByEmail(email)

	if user.ID != userID {
		t.Errorf("userService.Create(), user not persisted")
	}
}

func TestUsers(t *testing.T) {
	s := createServices()
	defer Reset(s.db)

	s.userService.Create(&scores.User{
		Email: "test@test.at",
	})
	s.userService.Create(&scores.User{
		Email:           "test2@test.at",
		ProfileImageURL: "image.url",
	})

	users, err := s.userService.Users()

	if err != nil {
		t.Errorf("UserService.Users() err: %s", err)
	}

	userCount := len(users)
	if userCount != 2 {
		t.Errorf("len(UserService.Users()), want 2, got %d", userCount)
	}

}

func TestUpdateUser(t *testing.T) {
	db, _ := Open("file::memory:?mode=memory&cache=shared")
	defer Reset(db)

	email := "test@test.com"
	newEmail := "test2@test.com"

	userService := UserService{DB: db}
	user, _ := userService.Create(&scores.User{
		Email:           email,
		ProfileImageURL: "image.url",
	})

	user.Email = newEmail

	err := userService.Update(user)

	if err != nil {
		t.Errorf("UserService.Update() err: %s", err)
	}

	user, err = userService.ByEmail(newEmail)

	if err != nil || user.Email != newEmail {
		t.Error("UserService.Update(), user not updated")
	}
}