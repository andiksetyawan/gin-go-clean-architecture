package entity

type User struct {
	Guid     string `form:"guid" json:"guid" bson:"guid"`
	Email    string `form:"email" json:"email" bson:"email"`
	Password string `form:"password" json:"password" bson:"password"`
	Name     string `form:"name" json:"name" bson:"name"`
	Address  string `form:"address" json:"address" bson:"address"`
	Photo    string `json:"photo" bson:"photo"`
}
