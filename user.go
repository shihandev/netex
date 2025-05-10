package main

type user struct {
	Name string `schema:"name,required"`
	Sex  bool   `schema:"sex"`
	Age  uint   `schema:"age"`
}
