package users

type createParams struct {
	Username string `form:"username" json:"username" mold:"trim" binding:"required,token,max=20"`
}
