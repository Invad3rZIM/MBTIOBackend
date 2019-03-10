package main

const (
	pi float64 = .017453292519943295 //pi / 100
)

type User struct {
	Uid      int `json:"uid"`
	Pin      int
	Phone    string
	Name     string
	Bio      string
	MBTI     string
	Sex      string
	Interest string

	Height int //used for profile information
	inches int //used for compatibility calculations

	Age       int
	legal     bool //true if >= 18
	twentyOne bool //true if >= 21

	lat  float64
	long float64

	//these are used for priority-heap sorting by distance + tracking user relative-compatibility scores
	Score    float64
	Distance int
	index    int

	//only true if profile is ready to be matched with
	ProfileReady bool
}

//creates a new user
func NewUser(uid int, phone string) *User {
	return &User{
		Uid:          uid,
		Phone:        phone,
		Pin:          GenUserPin(),
		ProfileReady: false,
	}
}

//checks to see if user is ready for the dating pool by verifying all requisite data has been collected
func (u *User) EvaluateProfileReady() {
	switch empty := 0; empty {

	case len(u.Name):
		return
	case len(u.MBTI):
		return
	case len(u.Sex):
		return
	case len(u.Interest):
		return
	case u.Height:
		return
	case u.Age:
		return
	default:
		if u.lat != 0 && u.long != 0 {
			u.ProfileReady = true
		}
	}
}

//returns true if user pin matches input pin
func (u *User) CheckPin(pin int) bool {
	return u.Pin == pin
}

func (u *User) GetPin() int {
	return u.Pin
}

//generates a user pin serving as a password for the weak verification purposes, replace with JWT if time permits
func GenUserPin() int {
	return GenInt(0, 99999999)
}

//sets the name
func (u *User) SetName(name string) {
	u.Name = name
}

//sets the user mbti
func (u *User) SetMBTI(mbti string) {
	u.MBTI = mbti
}

//sets the user sex
func (u *User) SetSex(sex string) {
	u.Sex = sex
}

//sets the user interest
func (u *User) SetInterest(interest string) {
	u.Interest = interest
}

//sets the user height
func (u *User) SetHeight(height int) {
	u.Height = height
	u.inches = 12*(height/100) + height%100
}

//returns users height in inches for matching purposes
func (u *User) InchHeight() int {
	return u.inches
}

//sets the user age
func (u *User) SetAge(age int) {
	u.Age = age
	u.legal = age >= 18
	u.twentyOne = age >= 21
}

func (u *User) SetLat(lat float64) {
	u.lat = lat
}

func (u *User) SetLong(long float64) {
	u.long = long
}

func (u *User) SetBio(bio string) {
	u.Bio = bio
}
