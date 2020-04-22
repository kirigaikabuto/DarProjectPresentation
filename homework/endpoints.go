package homework

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Endpoints interface {
	AddHomework() func(w http.ResponseWriter,r *http.Request)
	GetHomeWork() func(w http.ResponseWriter,r *http.Request)
}

type endpointsFactory struct {
	homeWorkInternship HomeWorkInternship
}

func NewEndpointsFactory(hwinter HomeWorkInternship) Endpoints{
	return &endpointsFactory{homeWorkInternship:hwinter}
}
func(ef *endpointsFactory) AddHomework() func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter,r *http.Request){
		data,err:=ioutil.ReadAll(r.Body)
		if err!=nil{
			respondJSON(w,http.StatusInternalServerError,err)
			return
		}
		fmt.Println(data)
		hw:=&HomeWork{}

		if err:= json.Unmarshal(data,&hw);err!=nil{
			respondJSON(w,http.StatusBadRequest,err.Error())
			return
		}
		newhw,err:=ef.homeWorkInternship.AddHomework(hw)
		respondJSON(w,http.StatusOK,newhw)
	}
}
func(ef *endpointsFactory)GetHomeWork() func(w http.ResponseWriter,r *http.Request){
	return func(w http.ResponseWriter,r *http.Request){

	}
}
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}
