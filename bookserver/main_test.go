package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var jsondata = []byte(`{"id":1,"title":"The Catcher in the Rye","author":"J.D. Salinger","quantity":10}`)

func connect(t *testing.T, f func(http.ResponseWriter, *http.Request), w *httptest.ResponseRecorder, req *http.Request) []byte {
	w.Header().Add("Content-Type", "application/json")
	f(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Error: %v", "Status Error")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Errorf("Unable to close body: %v", err)
		}
	}(resp.Body)
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	//fmt.Println(string(data))
	return data
}

func Test_Hello(t *testing.T) {
	name := "Naveen"
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/hello?name="+name, nil)
	data := connect(t, handleHello, w, req)
	if string(data) != ("Hello " + name) {
		t.Errorf("Expected 'Hello %s' but got %v", name, string(data))
	} else {
		fmt.Println("Test Hello Passed")
	}
}

func TestStartServer(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(jsondata))
	w := httptest.NewRecorder()
	connect(t, handleAddBook, w, req)
}

func Test_books(t *testing.T) {
	var expected = []byte(`[{"id":1,"title":"The Catcher in the Rye","author":"J.D. Salinger","quantity":10}]`)
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	w := httptest.NewRecorder()
	data := connect(t, handleBooks, w, req)
	if string(data[0]) != string(expected[0]) {
		t.Errorf("Expected %v but got %v", string(expected), string(data))
	} else {
		fmt.Println("Test Books Pass")
	}
}

func Test_add_books(t *testing.T) {
	expected := []byte(`[{"id":1,"title":"The Catcher in the Rye","author":"J.D. Salinger","quantity":10},{'id':4,'title':'This is test Book','author':'Test Author','quantity':10}]`)
	book := []byte(`{"id":4,"title":"This is test Book","author":"Test Author","quantity":10}`)
	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(book))
	w := httptest.NewRecorder()
	data := connect(t, handleAddBook, w, req)
	if string(data[0]) != string(expected[0]) {
		t.Errorf("Expected %v but got %v", string(expected), string(data))
	} else {
		fmt.Println("Test Add Books Pass")
	}
}

func Test_delete_books(t *testing.T) {
	expected := []byte(`[{"id":1,"title":"The Catcher in the Rye","author":"J.D. Salinger","quantity":10}]`)
	book := []byte(`{"id":4}`)
	req := httptest.NewRequest(http.MethodPost, "/books/delete", bytes.NewBuffer(book))
	w := httptest.NewRecorder()
	data := connect(t, handleDeleteBook, w, req)
	if string(data[0]) != string(expected[0]) {
		t.Errorf("Expected %v but got %v", string(expected), string(data))
	} else {
		fmt.Println("Test Delete Books Pass")
	}
}

func Test_insert_books(t *testing.T) {
	expected := []byte(`{"id":1,"title":"The Catcher in the Rye","author":"J.D. Salinger","quantity":11}`)
	book := []byte(`{"id":1}`)
	req := httptest.NewRequest(http.MethodPost, "/books/insert", bytes.NewBuffer(book))
	w := httptest.NewRecorder()
	data := connect(t, handleIncBook, w, req)
	if string(data[0]) != string(expected[0]) {
		t.Errorf("Expected %v but got %v", string(expected), string(data))
	} else {
		fmt.Println("Test Insert Books Pass")
	}
}

func Test_remove_books(t *testing.T) {
	expected := []byte(`{"id":1,"title":"The Catcher in the Rye","author":"J.D. Salinger","quantity":10}`)
	book := []byte(`{"id":1}`)
	req := httptest.NewRequest(http.MethodPost, "/books/borrow", bytes.NewBuffer(book))
	w := httptest.NewRecorder()
	data := connect(t, handleBorrowBook, w, req)
	if string(data[0]) != string(expected[0]) {
		t.Errorf("Expected %v but got %v", string(expected), string(data))
	} else {
		fmt.Println("Test remove Books Pass")
	}
}
