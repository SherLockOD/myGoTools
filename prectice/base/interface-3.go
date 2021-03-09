package main

import (
	"fmt"
	"sync"
)

type Instance interface {
	GetInstanceSet(wg *sync.WaitGroup)
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

func (cvm Cvm) GetInstanceSet(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("I'm cvm")
}

func (mongo Mongo) GetInstanceSet(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("I'm mongo")
}

func (redis Redis) GetInstanceSet(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("I'm redis")
}

func main() {
	var wg sync.WaitGroup
	args := Query{"test", "test", "1.1.1.1"}
	serviceList := []Instance{
		Cvm{args},
		Mongo{args},
		Redis{args},
	}

	fmt.Println(len(serviceList))
	wg.Add(len(serviceList))
	for _, instance := range serviceList {
		go instance.GetInstanceSet(&wg)
	}
	wg.Wait()
}