package article

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	router.HandleFunc("/v1/article/status/{id:[0-9]+}", bearerAuthMiddleware.VerifyBearer(handler.Edit)).Methods(http.MethodPut)

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

func (handler *ArticleHTTPHandler) Edit(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	var params EditArticleRequest
	var ctx = r.Context()
	path := mux.Vars(r)
	id := path["id"]

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		resp = response.Error(response.StatusUnprocessabelEntity, nil, err)
		resp.JSON(w)
		return
	}

	params.ID, err = strconv.ParseInt(id, 10, 64)
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

	resp = handler.Usecase.Edit(ctx, params)
	resp.JSON(w)
}
