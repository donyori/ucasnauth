package main

type Response interface {
	IsSuccessful() bool
	GetMessage() string
}
