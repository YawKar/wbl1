package _01

type Action struct {
	Human                    // embedding
	humanAsNonAnonymousField Human
	actionName               string
	name                     string // duplicate field from human
	age                      int    // duplicate field from human
}
