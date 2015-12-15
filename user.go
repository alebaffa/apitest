package main

import "time"

type User struct {
	Name       string
	NativeCity string
	Birthday   time.Time
}

type Users []User
