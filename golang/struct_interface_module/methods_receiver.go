package main

import "fmt"

type enrollment struct {
	semester string
	courses  []string
}

type aboutMe struct {
	info []string
}

func (e enrollment) courseAt(index uint) string {
	return e.courses[index]
}

func (e enrollment) addCourse(course string) {
	e.courses = append(e.courses, course) //สำเนารับมาทำให้ได้เป็นคนละตัว
}

func (a *aboutMe) addInfo(info string) {
	a.info = append(a.info, info) // pointer เป็นตัวเดียวกัน
}

func main() {
	e := enrollment{semester: "1/64", courses: []string{"java", "C#"}}

	result := e.courseAt(1) //สำเนาส่งไปยังเมดทอดทำให้ได้เป็นคนละตัว
	fmt.Println(result)

	e.addCourse("Go") //สำเนาส่งไปยังเมดทอดทำให้ได้เป็นคนละตัว
	fmt.Println(e.courses)	

	a := aboutMe{info: []string{"suphasin","fuse","SUT"}}
	fmt.Println(a)

	a.addInfo("Sisaket") // pointer เป็นตัวเดียวกัน
	fmt.Println(a)
}