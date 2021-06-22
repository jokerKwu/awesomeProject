package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestJsonHandler(t *testing.T){
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET","/students",nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK,res.Code)

	var list []Student

	err := json.NewDecoder(res.Body).Decode(&list) // JSON 데이터를 list로 변환한다.
	assert.Nil(err)	//이렇게 변환한 객체의 값이다.
	assert.Equal("aaa",list[0].Name)
	assert.Equal(1,list[0].Id)
	assert.Equal(16,list[0].Age)
	assert.Equal(87,list[0].Score)
}

func TestJsonHandler2(t *testing.T){
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET","/students/1",nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)

	var stu Student
	err := json.NewDecoder(res.Body).Decode(&stu)
	assert.Nil(err)
	assert.Equal("aaa",stu.Name)
	assert.Equal(1,stu.Id)
	assert.Equal(16,stu.Age)
	assert.Equal(87,stu.Score)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET","/students/2",nil)
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK,res.Code)
	err = json.NewDecoder(res.Body).Decode(&stu)
	assert.Nil(err)
	assert.Equal("bbb",stu.Name)

}
func TestJsonHandler3(t *testing.T){
	assert := assert.New(t)
	var student Student
	mux := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST","/students",strings.NewReader(`{"Id":0,"Name":"ccc","Age":15,"Score":78}`))
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusCreated, res.Code)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET","/students/3",nil)

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("ccc",student.Name)
}

func TestJsonHandler4(t *testing.T){
	assert := assert.New(t)

	mux := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE","/students/1",nil)

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK,res.Code)
	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET","/students",nil)
	mux.ServeHTTP(res,req)

	assert.Equal(http.StatusOK,res.Code)
	var list []Student
	err := json.NewDecoder(res.Body).Decode(&list)
	assert.Nil(err)
	assert.Equal(1,len(list))
	assert.Equal("bbb",list[0].Name)
}