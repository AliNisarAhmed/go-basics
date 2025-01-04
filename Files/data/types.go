package data

// type alias
type Json = map[string]string

type distance float64
type distanceKm float64

// Method on type distance
func (miles distance) ToKm() distanceKm {
	return distanceKm(1.60934 * miles)
}

// Method on type distanceKm
func (km distanceKm) ToMiles() distance {
	return distance(km / 1.60934)
}

func Test(miles distance) distance {
	d := distance(miles)
	dKm := d.ToKm()
	dMiles := dKm.ToMiles()
	return dMiles
}

type User struct {
	id   int
	name string
}

func TestUser() User {
	u1 := User{id: 1, name: "Ali Ahmed"}
	return u1
}

func (u User) Print() string {
	return string(u.id) + ": " + u.name
}
