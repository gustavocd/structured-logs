package main

import (
	"os"

	"log/slog"
)

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// func (u *User) LogValue() slog.Value {
// 	return slog.StringValue(u.ID)
// }

// func (u *User) LogValue() slog.Value {
// 	return slog.GroupValue(
// 		slog.String("id", u.ID),
// 		slog.String("name", u.FirstName+" "+u.LastName),
// 	)
// }

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	u := &User{
		ID:        "user-12234",
		FirstName: "Jan",
		LastName:  "Doe",
		Email:     "jan@example.com",
		Password:  "pass-12334",
	}

	logger.Info("info", "user", u)
}
