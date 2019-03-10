package main

func (h *Handler) CreateTestUsers(count int) {
	for i := 0; i < count; i++ {
		u := User{}
		u.MockLogistics()
		u.MockPersona()
		u.Uid = h.UserCache.GenUserID()
		u.Pin = GenUserPin()
		u.SetMBTI(h.RandomMBTI())

		h.AddUserToCache(&u)
	}

	u := User{}
	u.Uid = 0
	u.SetName("Kirk")
	u.SetAge(21)
	u.SetSex("M")
	u.SetInterest("F")
	u.MockGPS()
	u.SetMBTI("ENTJ")
	u.SetHeight(603)

	h.AddUserToCache(&u)
}

func (u *User) MockGPS() {
	u.lat = GenFloat(40, 40.5)
	u.long = GenFloat(60, 60.2)
}

func (u *User) MockPersona() {
	r := GenInt(0, 2)
	u.Name = GenName(r)

	if r == 0 {
		u.Sex = "F"
	} else {
		u.Sex = "M"
	}

	u.Interest = GenInterest(r)
}

//used for test objects
func (u *User) MockLogistics() {
	u.MockGPS()

	u.SetAge(GenInt(18, 22))
	if u.Sex == "F" {
		u.SetHeight(GenHeight(60, 71))
	} else {
		u.SetHeight(GenHeight(68, 76))
	}
}
