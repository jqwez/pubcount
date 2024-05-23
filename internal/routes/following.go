package routes

import (
	"net/http"
)

type UserData struct {
	User   string
	Goals  []*Goal
	Status bool
}

func NewUserData(u string, gs []*Goal) *UserData {
	ud := &UserData{User: u, Goals: gs}
	status := ud.GetStatus()
	ud.Status = status
	return ud
}

func (ud *UserData) GetStatus() bool {
	status := true
	if len(ud.Goals) < 1 {
		status = false
	}
	for _, goal := range ud.Goals {
		if goal.Status == false {
			status = false
			break
		}
	}
	return status
}

type Goal struct {
	Title  string
	Status bool
}

func GetHandleViewAccountData() HandleViewAccountsData {
	usera := NewUserData("DefaultUser", []*Goal{
		{Title: "foo", Status: false},
		{Title: "bar", Status: false},
	})
	userb := NewUserData("OtherUser", []*Goal{
		{Title: "baz", Status: false},
		{Title: "spam", Status: false},
	})
	data := HandleViewAccountsData{Users: []*UserData{usera, userb}}
	return data
}

type HandleViewAccountsData struct {
	Users []*UserData
}

func HandleFollowing(w http.ResponseWriter, r *http.Request) {
	_ = r.Body

	data := GetHandleViewAccountData()
	err := HandleTemplate(w, "following.tmpl", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
