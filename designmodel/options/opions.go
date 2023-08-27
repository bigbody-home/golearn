package main

import "fmt"

type School struct {
	Id      int
	Name    string
	Address string
	Manager string
}
type SetSchool func(*School)

func SetSchoolName(name string) SetSchool {
	return func(s *School) {
		s.Name = name
	}
}
func SetSchoolAddress(Address string) SetSchool {
	return func(s *School) {
		s.Address = Address
	}
}
func SetSchoolManager(Manager string) SetSchool {
	return func(s *School) {
		s.Manager = Manager
	}
}
func NewSchool(options ...SetSchool) *School {
	s := &School{}
	for _, v := range options {
		v(s)
	}
	return s
}
func main() {
	s := NewSchool(SetSchoolName("第二学校"), SetSchoolManager("李华"), SetSchoolAddress("雁塔区"))
	fmt.Println(s)
}
