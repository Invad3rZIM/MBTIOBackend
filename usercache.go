package main

import (
	"errors"
	"math/rand"
	"time"
)

type UserCache struct {
	cache map[int]*User
}

func NewUserCache() *UserCache {
	uc := UserCache{
		cache: make(map[int]*User),
	}

	return &uc
}

//CacheUserFromPhone creates a user from a phone number and adds it to the user cache
func (uc *UserCache) CacheUserFromPhone(phone string) *User {
	user := NewUser(uc.GenUserID(), phone)

	uc.cache[user.Uid] = user

	return user
}

//AddUserToCache simply adds the user to the cache
func (uc *UserCache) AddUserToCache(user *User) error {
	if user == nil {
		return errors.New("error: user pointer reference nil")
	}

	//assess if user is ready upon intial cache add
	if !user.ProfileReady {
		user.EvaluateProfileReady()
	}

	uc.cache[user.Uid] = user

	return nil
}

//generates a userid that's not in the cache
func (uc UserCache) GenUserID() int {
	rand.Seed(time.Now().UnixNano())

	id := GenInt(0, 999999999)

	for uc.UserExists(id) {
		id = GenInt(0, 999999999)
	}

	return id
}

//GetUser returns the cached pointer to the user
func (uc *UserCache) GetUser(uid int) (*User, error) {
	user, ok := uc.cache[uid]

	//if user isn't in cache
	if !ok {
		return nil, errors.New("error: user not found")
	}

	if !user.ProfileReady {
		user.EvaluateProfileReady()
	}

	return user, nil
}

func (uc UserCache) UserExists(uid int) bool {
	_, ok := uc.cache[uid]

	return ok
}

//returns the size of the cache
func (uc *UserCache) Size() int {
	return len(uc.cache)
}
