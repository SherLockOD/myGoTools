package main

import "fmt"

type Instance interface {
	GetInstanceSet()
}

type Query struct {
	SecretID string
	SecretKey string
	IP string
}

type Cvm struct {
	Query
}

type Mongo struct {
	Query
}

type Redis struct {
	Query
}

func (cvm Cvm) GetInstanceSet() {
	fmt.Println("I'm cvm")
}

func (mongo Mongo) GetInstanceSet() {
	fmt.Println("I'm mongo")
}

func (redis Redis) GetInstanceSet() {
	fmt.Println("I'm redis")
}

func main() {
	args := Query{"test", "test", "1.1.1.1"}
	serviceList := []Instance{
		Cvm{args},
		Mongo{args},
		Redis{args},
	}

	fmt.Println(len(serviceList))

	for _, instance := range serviceList {
		instance.GetInstanceSet()
	}
}