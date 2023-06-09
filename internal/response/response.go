package response

import (
	"GoGinStarter/internal/paginator"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response interface {
	general(ctx *gin.Context, statusCode int, succeed bool, message string, results any, metas any)
	Abort(ctx *gin.Context, statusCode int, succeed bool, message string, results any, metas any)
	OK(ctx *gin.Context)
	Success(ctx *gin.Context, results any, message string)
	SuccessWithMeta(ctx *gin.Context, message string, results any, metas any)
	CreateSuccess(ctx *gin.Context, results any)
	Error(ctx *gin.Context, status int, error string)
	BadRequest(ctx *gin.Context, error string, data any)
	NotFound(ctx *gin.Context, error string)
	Unauthorized(ctx *gin.Context, error string)
	Render(ctx *gin.Context, name string, obj any)
	WithPaginate(ctx *gin.Context, pagination *paginator.Paginator, message string)
}

type response struct {
	Succeed bool   `json:"succeed"`
	Message string `json:"message"`
	Results any    `json:"results"`
	Metas   any    `json:"metas"`
}

func (r *response) general(ctx *gin.Context, statusCode int, succeed bool, message string, results any, metas any) {
	ctx.JSON(
		statusCode,
		response{
			Succeed: succeed,
			Message: message,
			Results: results,
			Metas:   metas,
		},
	)
}

func (r *response) Abort(ctx *gin.Context, statusCode int, succeed bool, message string, results any, metas any) {
	ctx.AbortWithStatusJSON(statusCode,
		response{
			Succeed: succeed,
			Message: message,
			Results: results,
			Metas:   metas,
		})
}

func (r *response) OK(ctx *gin.Context) {
	r.general(ctx, http.StatusOK, true, "", nil, nil)
}

func (r *response) Success(ctx *gin.Context, results any, message string) {
	r.general(ctx, http.StatusOK, true, message, results, nil)
}

func (r *response) SuccessWithMeta(ctx *gin.Context, message string, results any, metas any) {
	r.general(ctx, http.StatusOK, true, message, results, metas)
}

func (r *response) CreateSuccess(ctx *gin.Context, results any) {
	r.general(ctx, http.StatusCreated, true, "", results, nil)
}

func (r *response) Error(ctx *gin.Context, status int, error string) {
	r.Abort(ctx, status, false, error, nil, nil)
}

func (r *response) BadRequest(ctx *gin.Context, error string, data any) {
	r.general(ctx, http.StatusBadRequest, false, error, data, nil)
}

func (r *response) NotFound(ctx *gin.Context, error string) {
	r.general(ctx, http.StatusNotFound, false, error, nil, nil)
}

func (r *response) Unauthorized(ctx *gin.Context, error string) {
	r.general(ctx, http.StatusUnauthorized, false, error, nil, nil)
}

func (r *response) Render(ctx *gin.Context, name string, obj any) {
	ctx.HTML(http.StatusOK, name, obj)
}

func (r *response) WithPaginate(ctx *gin.Context, pagination *paginator.Paginator, message string) {
	ctx.JSON(
		200,
		response{
			Succeed: true,
			Message: message,
			Results: pagination.Items,
			Metas: map[string]interface{}{
				"total":        pagination.Total,
				"per_page":     pagination.PerPage,
				"current_page": pagination.CurrentPage,
				"last_page":    pagination.LastPage,
			},
		},
	)
}

func ProvideResponse() Response {
	return &response{}
}
