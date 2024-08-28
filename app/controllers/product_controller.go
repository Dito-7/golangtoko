package controllers

import (
	"net/http"

	"github.com/Dito-7/golangtoko/app/models"
	"github.com/unrolled/render"
)

func (s *Server) products(w http.ResponseWriter, r *http.Request) {
	render := render.New(render.Options{
		Layout: "layout",
	})

	ProductModel := models.Product{}

	products, err := ProductModel.GetProduct(s.DB)
	if err != nil {
		return
	}

	_ = render.HTML(w, http.StatusOK, "products", map[string]interface{}{
		"products": products,
	})
}
