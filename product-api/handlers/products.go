package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tduchars/goservcoffee/product-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.l.Println("GET", r.RequestURI)
		if r.RequestURI == "/products" {
			p.getProducts(rw, r)
			return
		} else if r.RequestURI == "/products/" {
			p.getProducts(rw, r)
			return
		}
	} else if r.Method == http.MethodPost {
		p.l.Println("POST", r.RequestURI)
	}

	//handle update request±

	//catch all
	p.l.Println("Method not allowed in products handler")
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("GET Products Handler on /")

	lp := data.GetProducts()
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) getSingleProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("GET single product Handler on /products/id")
	fmt.Println("GET params were:", r.URL.Query())

	// param1 := r.URL.Query().Get("param1")

	sp := data.GetSingleProduct(1)
	err := sp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}
}
