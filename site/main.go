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
		Description: "Hi, I'm Senior Backend Developer and currently I'm looking new dream project. I have more that 11 years experience on developer position. My main technology is PHP and I know it very well. But currently, I migrate to Golang and I spend whole myself to growing this skill.",
		OtherSkills: []string{"PHP", "GO", "Javascript", "MySQL", "CouchDB", "XML", "Linux", "TDD", "jQuery", "Twitter Bootstrap", "Yii", "Redmine", "Jenkins"},

		Experiences: []Experience{
			{
				From:        "Oct 2014",
				To:          "Now",
				Project:     "Vivino",
				Company:     "Symphony Solutions",
				Position:    "PHP/Golang Backend Developer",
				Description: "Vivino is cool projects about wine. We have a huge database of wines. Currently, we have more than 20 million users and big loading to our servers. I'm working in API team. My responsibility is development new features and support old. Very often my tasks were improvements of performance. Mostly I'm working on PHP side but also I'm Golang developer. Technologies: PHP, JS, Mysql, GO (Golang), ElasticSearch, AWS, Kafka, Aerospike, Kanban",
			},
			{
				From:        "Sep 2010",
				To:          "Now",
				Project:     "nedoma.com.ua",
				Company:     "",
				Position:    "Founder",
				Description: "This is personal non-commercial projects. We like active sport and like write about this. First of all, we are blog but also we organized low price trip. Our goals were popularization of active recreation. On this project, I'm a founder, editor, main event manager. Currently, project is frozen",
			},
			{
				From:        "Sep 2013",
				To:          "Sep 2015",
				Project:     "LC Parts",
				Company:     "",
				Position:    "Technical Consultant",
				Description: "Small project for the friend. We looked prices for car parts on several sites. It's was something like prices aggregator. My responsibility was the server support, architecture and development of the system. Also, I interviewed new candidates to this project. Technologies: PHP, Zend Framework, HTML, jQuery, MySql",
			},
			{
				From:        "Aug 2011",
				To:          "Sep 2014",
				Project:     "Borsh.ch",
				Company:     "",
				Position:    "Founder",
				Description: "My startup. Unfortunately, I made all possible junior mistakes. I started without the team, without the vision of goals and without founded target auditory. I want to say \"thank you\" to it. This project taught me a lot and stole all my money.",
			},
			{
				From:        "Jun 2014",
				To:          "Jul 2014",
				Project:     "SalesDoubler",
				Company:     "",
				Position:    "Remote Web Developer",
				Description: "Remote job. The main my responsibility was the return to live old coupon project. I started from small refactoring and then fixed all bugs. The project was short but interesting. Technologies: PHP, HTML, jQuery, CSS, MySql",
			},
			{
				From:        "Jul 2012",
				To:          "Apr 2014",
				Project:     "WAS",
				Company:     "Intellias",
				Position:    "Software Engineer",
				Description: "I came to this company on Team Lead position. But unfortunately customer changes his plans and I worked alone two years. My responsibility was the implementation of  Extjs components for admin panel. Also, I implemented new (we don't found analog) algorithm of comparing two HTML files without breaking of syntax. Technologies: PHP, JS, ExtJs, jQuery, Twitter Bootstrap, CouchDB, MySql",
			},
			{
				From:        "Nov 2011",
				To:          "May 2012",
				Project:     "NextGen",
				Company:     "MalkosUA",
				Position:    "Web Developer",
				Description: "We analyzed data from many garbage cars. Each car gave us huge log file. We need this data for statistics and calculation of bonuses. My responsibility was working on the control panel for this system. Technologies: Yii, PHP, JavaSript, MySql, Scrum",
			},
			{
				From:        "Aug 2009",
				To:          "Nov 2011",
				Project:     "HireVue",
				Company:     "SoftServe",
				Position:    "Web developer",
				Description: "When I come to this company I was the second developer on the project. After two years our team grew to 10 developers and 2 QA. It was a great experience of teamwork. My responsibility was development online video interview system, deploying, hot maintenance. Technologies: PHP, Zend Framework, MySql, Jquery, PHPUnit",
			},
			{
				From:        "Mar 2007",
				To:          "Aug 2009",
				Project:     "InterLogic",
				Company:     "",
				Position:    "Web developer",
				Description: "Its was first my outsourcing experience. I had very different work. Some projects took 1-2 days other more than the year. I'm worked with PHP, ActiveScript, C#",
			},
			{
				From:        "May 2005",
				To:          "Dec 2006",
				Project:     "VashComfort",
				Company:     "OWBT",
				Position:    "Web developer",
				Description: "First commercial experience. In working in the small team. In the same time, I studied in University. Technologies: PHP, MySQL, JavaScript."},
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
