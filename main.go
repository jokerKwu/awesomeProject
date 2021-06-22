package main

import (
	"encoding/json"
	mux2 "github.com/gorilla/mux"
	"net/http"
	"sort"
	"strconv"
)

type Student struct{
	Id int
	Name string
	Age int
	Score int
}
var students map[int]Student
var lastId int
type Students []Student

func MakeWebHandler() http.Handler{
	mux := mux2.NewRouter()
	mux.HandleFunc("/students",GetStudentListHandler).Methods("GET") // 전체 학생 조회
	// -- 여기에 새로운 핸들러 등록
	mux.HandleFunc("/students/{id:[0-9]+}",GetStudentHandler).Methods("GET") // 특정 학생 조회
	mux.HandleFunc("/students",PostStudentHandler).Methods("POST") 	// 학생 추가
	mux.HandleFunc("/students/{id:[0-9]+}",DeleteStudentHandler).Methods("DELETE")

	students = make(map[int]Student)
	students[1] = Student{1,"aaa",16,87}
	students[2] = Student{2,"bbb",18,99}
	lastId = 2
	return mux
}
func (s Students) Len() int{
	return len(s)
}
func (s Students) Swap(i,j int){
	s[i], s[j] = s[i] , s[i]
}

func (s Students) Less(i, j int) bool{
	return s[i].Id < s[j].Id
}
//학생 제거
func DeleteStudentHandler(w http.ResponseWriter, r *http.Request){
	vars := mux2.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	_, ok := students[id]
	if !ok{
		w.WriteHeader(http.StatusNotFound)
		return
	}

	delete(students,id)
	w.WriteHeader(http.StatusOK)

}

//학생 추가
func PostStudentHandler(w http.ResponseWriter,r  *http.Request){
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lastId++
	student.Id = lastId
	students[lastId] = student
	w.WriteHeader(http.StatusCreated)
}

// 전체 학생 리스트 조회
func GetStudentListHandler(w http.ResponseWriter, r *http.Request){
	list := make(Students, 0)
	for _, student := range students{
		list = append(list, student)
	}
	sort.Sort(list)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(list)
}
// 학생 조회
func GetStudentHandler(w http.ResponseWriter,r *http.Request){
	vars := mux2.Vars(r) // id를 가져온다.
	id, _ := strconv.Atoi(vars["id"])
	student, ok := students[id]
	if !ok{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(student)
}



func main(){
	http.ListenAndServe(":3000",MakeWebHandler())
}