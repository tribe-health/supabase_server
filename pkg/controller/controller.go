package controller

// Controller holds the controllers for the entire across application
type Controller struct {
	Project interface{ Project }
	Api     interface{ Api }
}
