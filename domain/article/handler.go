package article

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sangianpatrick/devoria-article-service/middleware"
	"github.com/sangianpatrick/devoria-article-service/response"
)

type ArticleHTTPHandler struct {
	Validate *validator.Validate
	Usecase  ArticleUsecase
}

func NewArticleHTTPHandler(
	router *mux.Router,
	bearerAuthMiddleware middleware.RouteMiddlewareBearer,
	validate *validator.Validate,
	usecase ArticleUsecase,
) {
	handler := &ArticleHTTPHandler{
		Validate: validate,
		Usecase:  usecase,
	}

	router.HandleFunc("/v1/article", bearerAuthMiddleware.VerifyBearer(handler.Create)).Methods(http.MethodPost)

}

func (handler *ArticleHTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	var params CreateArticleRequest
	var ctx = r.Context()

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		resp = response.Error(response.StatusUnprocessabelEntity, nil, err)
		resp.JSON(w)
		return
	}

	err = handler.Validate.StructCtx(ctx, params)
	if err != nil {
		resp = response.Error(response.StatusInvalidPayload, nil, err)
		resp.JSON(w)
		return
	}

	resp = handler.Usecase.Create(ctx, params)
	resp.JSON(w)
}
