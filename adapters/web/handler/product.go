package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Sotnasjeff/hexagonal-architecture-api/adapters/dto"
	"github.com/Sotnasjeff/hexagonal-architecture-api/app"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

// const (
// 	ENABLED  = "enabled"
// 	DISABLED = "disabled"
// )

func MakeProductHandler(r *mux.Router, n *negroni.Negroni, service app.ProductServiceInterface) {
	r.Handle("/product/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/product", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS")

	r.Handle("/product/enable", n.With(
		negroni.Wrap(enableProduct(service)),
	)).Methods("PUT", "OPTIONS")

	r.Handle("/product/disable", n.With(
		negroni.Wrap(disableProduct(service)),
	)).Methods("PUT", "OPTIONS")
}

func getProduct(service app.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

	})
}

func createProduct(service app.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var productDTO dto.Product
		err := json.NewDecoder(r.Body).Decode(&productDTO)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		product, err := service.Create(productDTO.Name, productDTO.Price)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

	})
}

func enableProduct(service app.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var productUpdate dto.ProductUpdateStatus
		err := json.NewDecoder(r.Body).Decode(&productUpdate)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		product, err := service.Get(productUpdate.ID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(jsonError(err.Error()))
			return
		}

		updatedProduct, err := service.Enable(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(updatedProduct)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
	})
}

func disableProduct(service app.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var productUpdate dto.ProductUpdateStatus
		err := json.NewDecoder(r.Body).Decode(&productUpdate)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		product, err := service.Get(productUpdate.ID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(jsonError(err.Error()))
			return
		}

		updatedProduct, err := service.Disable(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(updatedProduct)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
	})
}
