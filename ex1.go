package main

import _ "fmt"

type Book struct {
	ID         string
	Title      string
	Author     string
	isBorrowed bool
}
type Library struct {
	Books map[string]Book
}
