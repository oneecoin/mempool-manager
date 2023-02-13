package blocks

type BlocksQuery struct {
	Page int `form:"page" binding:"required"`
}
