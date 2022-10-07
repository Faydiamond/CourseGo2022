package main

import "fmt"

type Video struct {
	title string
	curso Curso
}

type Curso struct {
	name   string
	videos []Video
}

func main() {
	video01 := Video{title: "Introduccion a la programacion"}
	video02 := Video{title: "Analisis de problemas"}
	video03 := Video{title: "Principios algoritmia"}

	curso := Curso{
		name:   "Programacion basica",
		videos: []Video{video01, video02, video03},
	}

	//.Println(curso)
	video01.curso = curso
	video02.curso = curso
	video03.curso = curso
	fmt.Println(video02.curso.name)

	for _, v := range curso.videos {
		fmt.Println(v.title)
	}
}
