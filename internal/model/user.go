package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID 					int
	Name 				string 			`json:"name"`
	Email		 		string			`json:"email"`
	Password 			string			`json:"password"`
	ProfileImage 		string			`json:"profileImage"`
	Role 				string			`json:"role"`
	CoolID				int 			`json:"coolID"`
	MemberID			int				`json:"memberID"`
	IsCoolMember		bool			`json:"isCoolMember"`
	State				string			`json:"state"`
	SessionToken		string			`json:"sessionToken"`
	CreatedAt 			time.Time		`json:"createdAt"`
	UpdatedAt 			time.Time		`json:"updatedAt"`
	DeletedAt			sql.NullTime	`json:"deletedAt"`
}