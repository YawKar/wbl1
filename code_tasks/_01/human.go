package _01

type Human struct {
	name     string
	age      int
	jobTitle string
}

func (h *Human) Name() string {
	return h.name
}

func (h *Human) Age() int {
	return h.age
}

func (h *Human) JobTitle() string {
	return h.jobTitle
}

func (h Human) WithName(newName string) Human {
	h.name = newName
	return h
}
