package controllers;

type SuccessResponse struct {
	Message string `json:"message"`
}

type People struct {
	Name string `json:name`
	Age int `json:age`
}

var peoples []People

func SendSuccess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	successResponse := SuccessResponse{ Message: "Success!!!!!" }
	json.NewEncoder(w).Encode(&successResponse)
}

func CreatePeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var people People
	_ = json.NewDecoder(r.Body).Decode(&people)
	peoples = append(peoples, people);
	json.NewEncoder(w).Encode(&peoples);
}

func GetPeoples(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&peoples);
}
