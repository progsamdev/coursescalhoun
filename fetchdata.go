package main

import "time"

type UserMock struct {
	Name       string
	Age        int
	IsAdmin    bool
	AccessData UserAccessData
}

type UserAccessData struct {
	NumberOfAccess      int
	LastAccess          time.Time
	FormattedLastAccess string
}

func getUserMock() *UserMock {
	access := time.Now().AddDate(0, 0, -5)
	us := UserMock{
		Name:    "Mock Silva",
		Age:     23,
		IsAdmin: false,
		AccessData: UserAccessData{
			NumberOfAccess:      12,
			LastAccess:          access,
			FormattedLastAccess: access.Format("01-02-2006"),
		},
	}
	return &us
}
