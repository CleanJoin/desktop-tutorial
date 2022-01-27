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

	return name.Age >= 18

}
func (name *NameT) IsAdult() bool {

	return name.Age >= 18 && name.Idiot == "YES"
}
func (name *Name) AgeSmall18() bool {

	return name.Age < 18
}
