// application.go

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Coditation struct {
	Router *mux.Router
	DB     *sql.DB
}

func (shop *Coditation) Initialize(User, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", User, password, dbname)
	fmt.Println(connectionString)
	var err error
	shop.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	shop.Router = mux.NewRouter()
	shop.initializeRoutes()
}

func (shop *Coditation) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, shop.Router))
}

func (shop *Coditation) initializeRoutes() {
	shop.Router.HandleFunc("/Categorys", shop.getCategorys).Methods("GET")
	shop.Router.HandleFunc("/Category", shop.createCategory).Methods("POST")
	shop.Router.HandleFunc("/Category/{Categoryname}", shop.getCategory).Methods("GET")
	shop.Router.HandleFunc("/Category/{id:[0-9]+}", shop.updateCategory).Methods("PUT")
	shop.Router.HandleFunc("/Category/{id:[0-9]+}", shop.deleteCategory).Methods("DELETE")

	shop.Router.HandleFunc("/Subcategories", shop.getVarients).Methods("GET")
	shop.Router.HandleFunc("/Subcategory", shop.createVarient).Methods("POST")
	shop.Router.HandleFunc("/Subcategory/{Subcategoryname}", shop.getVarient).Methods("GET")
	shop.Router.HandleFunc("/Subcategory/{id:[0-9]+}", shop.updateVarient).Methods("PUT")
	shop.Router.HandleFunc("/Subcategory/{id}", shop.deleteVarient).Methods("DELETE")

	shop.Router.HandleFunc("/Products", shop.getProducts).Methods("GET")
	shop.Router.HandleFunc("/Product", shop.createProduct).Methods("POST")
	shop.Router.HandleFunc("/Products/{ProductName}", shop.getProduct).Methods("GET")
	shop.Router.HandleFunc("/Products/{id}", shop.updateProduct).Methods("PUT")
	shop.Router.HandleFunc("/Products/{id}", shop.deleteProduct).Methods("DELETE")
	//getproductswithvar
	//shop.Router.HandleFunc("/apiofproduct", shop.getproductswithvar).Methods("GET")
}

func (shop *Coditation) getCategorys(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	products, err := getCategorys(shop.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}
func (shop *Coditation) getProducts(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	products, err := getProducts(shop.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

func (shop *Coditation) createCategory(w http.ResponseWriter, r *http.Request) {
	var category Category
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&category); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := category.createCategory(shop.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, category)
}

func (shop *Coditation) getCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)
	name := vars["Categoryname"]
	category := Category{Name: name}
	if err := category.getCategory(shop.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Category not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, category)
}

func (shop *Coditation) updateCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//fmt.Println(vars)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Category ID")
		return
	}

	var category Category
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&category); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	category.CategoryID = id

	if err := category.updateCategory(shop.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, category)
}

func (shop *Coditation) deleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Category ID")
		return
	}

	category := Category{CategoryID: id}
	if err := category.deleteCategory(shop.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "Coditationlication/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (shop *Coditation) createVarient(w http.ResponseWriter, r *http.Request) {
	var subcategory Subcategory
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&subcategory); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := subcategory.createVarient(shop.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, subcategory)
}

func (shop *Coditation) getVarient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//fmt.Println(vars)
	name := vars["Subcategoryname"]
	//fmt.Println(name)
	subcategory := Subcategory{
		SubcategoryID:   0,
		SubcategoryName: name,
		Products:        Product{},
	}
	if err := subcategory.getVarient(shop.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "var not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, subcategory)
}

func (shop *Coditation) updateVarient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	fmt.Println(id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid var ID")
		return
	}

	var subcategory Subcategory
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&subcategory); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	subcategory.SubcategoryID = id

	if err := subcategory.updateVarient(shop.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, subcategory)
}

func (shop *Coditation) deleteVarient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	subcategory := Subcategory{SubcategoryID: id}
	if err := subcategory.deleteVarient(shop.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (shop *Coditation) getVarients(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	varients, err := getVarients(shop.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, varients)
}

func (shop *Coditation) createProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := product.createProduct(shop.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, product)
}

func (shop *Coditation) getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["ProductName"]
	fmt.Println(name)
	product := Product{ProductName: name}
	if err := product.getProduct(shop.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Product not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, product)
}

func (shop *Coditation) updateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var product Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	product.ProductID = id

	if err := product.updateProduct(shop.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, product)
}

func (shop *Coditation) deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product := Product{ProductID: id}
	if err := product.deleteProduct(shop.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (shop *Coditation) getproductswithvar(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}
	productwithvar, err := getproductswithvar(shop.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, productwithvar)
}
