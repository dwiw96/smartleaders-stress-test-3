package delivery

import (
	"encoding/json"
	// "fmt"
	"net/http"

	// middleware "stress-test-3-2-go/middleware"
	rent "stress-test-3-2-go/features/movie-rent"
	responses "stress-test-3-2-go/utils/responses"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/julienschmidt/httprouter"
)

type movieRentDelivery struct {
	router   *httprouter.Router
	service  rent.ServiceInterface
	validate *validator.Validate
	trans    ut.Translator
}

func NewMovieRentDelivery(router *httprouter.Router, service rent.ServiceInterface) {
	handler := &movieRentDelivery{
		router:   router,
		service:  service,
		validate: validator.New(),
	}

	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(handler.validate, trans)
	handler.trans = trans

	router.POST("/api/rent_book", handler.rentBook)
	router.GET("/api/list_of_rented_books", handler.listOfRentedBook)
}

func translateError(trans ut.Translator, err error) (errTrans []string) {
	errs := err.(validator.ValidationErrors)
	a := (errs.Translate(trans))
	for _, val := range a {
		errTrans = append(errTrans, val)
	}

	return
}

func (d *movieRentDelivery) rentBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var request rentMovieRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		responses.ErrorJSON(w, http.StatusUnprocessableEntity, err.Error(), r.RemoteAddr)
		return
	}

	err = d.validate.Struct(request)
	if err != nil {
		errTranslated := translateError(d.trans, err)
		responses.ErrorJSON(w, 422, errTranslated, r.RemoteAddr)
		return
	}

	response, err := d.service.RentBook(rent.RentMovieRequest(request))
	if err != nil {
		responses.ErrorJSON(w, 500, err.Error(), r.RemoteAddr)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(responses.SuccessWithDataResponse(rentMovieResponse(*response), 200, "rent book success"))
}

func (d *movieRentDelivery) listOfRentedBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response, err := d.service.ListOfRentBook()
	if err != nil {
		responses.ErrorJSON(w, 500, err.Error(), r.RemoteAddr)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(responses.SuccessWithDataResponse(response, 200, "list of rented book success"))
}
