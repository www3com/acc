package api

type user struct {
	Nickname  string `form:"nickname" binding:"required"`
	Username  string `form:"username" binding:"required"`
	Password  string `form:"password" binding:"required"`
	Agreement bool   `form:"agreement" binding:"required"`
}
