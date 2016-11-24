package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

func main() {

	cwd, _ := os.Getwd()
	templatePath := filepath.Join(cwd, "./template/")
	exportPath := filepath.Join(cwd, "../")

	files, err := ioutil.ReadDir(templatePath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		tmpl := template.Must(template.ParseFiles(filepath.Join(templatePath, file.Name())))
		f, err := os.Create(filepath.Join(exportPath, file.Name()))
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(f, getInfo())
		if err != nil {
			panic(err)
		}
	}
}

func getInfo() Info {
	return Info{
		Name:        "Herych Oleh",
		Position:    "Backend Developer",
		Phone:       "+380977629793",
		Skype:       "dieseldraft",
		Birth:       "08.04.1987",
		Description: "I'm Web Developer with 11 years experience. Most of time I worked with PHP. But in this moment I'm experimenting with Go (golang) and have small commercial experience. I like this technology and would like to continue work with it. Worked with small and big (~60 developers in one place). Scrum, Kanban and Waterfall too. Locally and remote from home.",
		OtherSkills: []string{"PHP", "GO", "Javascript", "MySQL", "CouchDB", "XML", "Linux", "TDD", "jQuery", "Twitter Bootstrap", "Yii", "Redmine", "Jenkins"},

		Experiences: []Experience{
			{
				From:        "Oct 2014",
				To:          "Now",
				Project:     "Vivino",
				Company:     "Symphony Solutions",
				Position:    "PHP/Golang Backend Developer",
				Description: "",
			},
			{
				From:        "Sep 2010",
				To:          "Now",
				Project:     "nedoma.com.ua",
				Company:     "",
				Position:    "Founder",
				Description: "We like active life. And we try to develop active sport in our country. Nedoma is blog about active life. Everybody can talk about this experience and his knowledge. I'm a founder, editor, and author of most of  topics. Also i main organizer of our trips. Most of them take part big groups (30 peoples ). All my participation only like volunteering and without pay for me.",
			},
			{
				From:        "Sep 2013",
				To:          "Sep 2015",
				Project:     "LC Parts",
				Company:     "",
				Position:    "Technical Consultant",
				Description: "Friends project. I'm something like technical consultant and requirements manager. My responsibilities are structure all business needs. From last time i helping with bugs fixing. This project take about one hour in week.",
			},
			{
				From:        "Aug 2011",
				To:          "Sep 2014",
				Project:     "Borsh.ch",
				Company:     "",
				Position:    "Founder",
				Description: "This is my first serious experience founding commercial project. My responsibilities are prepare finance model and business plan. Also i take part in development server and client side. Some part from project we give to freelance. I mean design and content",
			},
			{
				From:        "Jun 2014",
				To:          "Jul 2014",
				Project:     "SalesDoubler",
				Company:     "",
				Position:    "Remote Web Developer",
				Description: "My responsibility was fixing bugs and implementation new features. Communication by Skype and phone.",
			},
		},
		Skills: []Skill{
			{Progress: 55, Label: "PHP", Color: "progress-bar-danger"},
			{Progress: 10, Label: "Go (Golang)", Color: "progress-bar-success"},
			{Progress: 25, Label: "Javascript", Color: "progress-bar-warning"},
			{Progress: 10, Label: "...", Color: ""},
		},
		Languages: []Language{
			{Progress: 40, Label: "Pre Intermediate", Name: "English", Color: "progress-bar-success"},
		},
		Educations: []Education{
			{From: "2005", To: "2010", School: "Polytechnic National University", Name: "Computer engineering"},
		},
		Recommendations: []Recommendation{
			{
				Date:        "Aug, 2014",
				Author:      "Severin Kälin",
				Source:      "LinkedIn",
				Link:        "https://www.linkedin.com/in/severin-kälin-263b9019",
				Description: "Oleg developed multiple web applications for and with us. He produces clean and well structured code. He is very focused on the final product. He points out requirements and design flaws and suggests solutions.",
			},
			{
				Date:        "Oct, 2014",
				Author:      "Upwork Customer",
				Source:      "Upwork",
				Link:        "https://www.upwork.com/freelancers/~01f78c0f4ad0e5a854",
				Description: "Excellent developer. Always trying to understand the business problem and suggesting best possible solutions. Great time management skills and availability. Will be doing business again.",
			},
		},
	}
}

type Info struct {
	Name            string
	Position        string
	Phone           string
	Email           string
	Skype           string
	Birth           string
	Description     string
	OtherSkills     []string
	Experiences     []Experience
	Educations      []Education
	Languages       []Language
	Skills          []Skill
	Recommendations []Recommendation
}

type Experience struct {
	From        string
	To          string
	Project     string
	Company     string
	Position    string
	Description string
}

type Education struct {
	From   string
	To     string
	School string
	Name   string
}

type Skill struct {
	Color    string
	Progress int
	Label    string
}

type Language struct {
	Color    string
	Progress int
	Label    string
	Name     string
}

type Recommendation struct {
	Date        string
	Author      string
	Source      string
	Link        string
	Description string
}
