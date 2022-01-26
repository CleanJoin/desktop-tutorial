package pack

type Name struct {
	Name string
	Age  int
}
type NameT struct {
	Name  string
	Age   int
	Idiot string
}
type AgePerson interface {
	IsAdult() bool
	AgeSmall18() bool
}

func (name Name) IsAdult() bool {
	if name.Age >= 18 {
		return true
	}
	return false

}
func (name *NameT) IsAdult() bool {
	if name.Age >= 18 && name.Idiot == "YES" {
		return true
	}
	return false

}
func (name *Name) AgeSmall18() bool {
	if name.Age < 18 {
		return true
	}
	return false
}
