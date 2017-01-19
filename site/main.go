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
		Description: "Hi, I'm Senior Backend Developer and currently I'm looking for a new dream project. I have more that 11 years experience on developer position. My main technology is PHP and I know it very good. But currently, I migrate to Golang and spend whole myself to grow this skill.",
		OtherSkills: []string{"PHP", "GO", "Javascript", "MySQL", "CouchDB", "XML", "Linux", "TDD", "jQuery", "Twitter Bootstrap", "Yii", "Redmine", "Jenkins"},

		Experiences: []Experience{
			{
				From:        "Oct 2014",
				To:          "Jan 2017",
				Project:     "Vivino",
				Company:     "Symphony Solutions",
				Position:    "PHP/Golang Backend Developer",
				Description: "Vivino is a cool project about wine, collecting the biggest number of wines in the whole world. Currently, it counts more than 20 million users. I'm working in API team. My responsibility is development of new features and support old ones. I've got great experience of high load optimization. Mostly I'm working on PHP side but also I was a Golang developer.<br/>Technology: GO (Golang), PHP, JS, Mysql, ElasticSearch, AWS, Kafka, Aerospike, Kanban",
			},
			{
				From:        "Sep 2010",
				To:          "Now",
				Project:     "nedoma.com.ua",
				Company:     "",
				Position:    "Founder",
				Description: "This is personal non-commercial project. We like active sport and like write about this. First of all, it's a blog but also we organized low-price trips. Our goals are popularization of active recreation. On this project, I'm a founder, editor and main event manager. Currently, project is frozen",
			},
			{
				From:        "Sep 2013",
				To:          "Sep 2015",
				Project:     "LC Parts",
				Company:     "",
				Position:    "Technical Consultant",
				Description: "Small project for friend. We were searching for prices on car spare parts on several sites. It was something like prices aggregator. My responsibility was the server support, architecture and development of the system. Also, I interviewed new candidates for this project.<br/>Technology: PHP, Zend Framework, HTML, jQuery, MySql",
			},
			{
				From:        "Aug 2011",
				To:          "Sep 2014",
				Project:     "Borsh.ch",
				Company:     "",
				Position:    "Founder",
				Description: "This is personal non-commercial project. We like active sport and like write about this. First of all, it's a blog but also we organized low-price trips. Our goals are popularization of active recreation. On this project, I'm a founder, editor and main event manager. Currently, project is frozen",
			},
			{
				From:        "Jun 2014",
				To:          "Jul 2014",
				Project:     "SalesDoubler",
				Company:     "",
				Position:    "Remote Web Developer",
				Description: "Remote job. My main responsibility was return to live old coupon project. I started from small refactoring and then fixed all bugs. The project was short but interesting.<br/>Technology: PHP, HTML, jQuery, CSS, MySql",
			},
			{
				From:        "Jul 2012",
				To:          "Apr 2014",
				Project:     "WAS",
				Company:     "Intellias",
				Position:    "Software Engineer",
				Description: "I came to this company on Team Lead position. But unfortunately customer changed his plans and I worked alone two years. My responsibility was implementation of Extjs components for admin panel. Also, I designed new (we don't found analog) algorithm of comparing two HTML files without breaking of syntax.<br/>Technology: PHP, JS, ExtJs, jQuery, Twitter Bootstrap, CouchDB, MySql",
			},
			{
				From:        "Nov 2011",
				To:          "May 2012",
				Project:     "NextGen",
				Company:     "MalkosUA",
				Position:    "Web Developer",
				Description: "We analyzed data from many garbage trucks. Each truck gave us a huge data log, which we needed for collecting statistics and calculation of bonuses. My responsibility was working on the control panel for this system.<br/>Technology: Yii, PHP, JavaSript, MySql, Scrum",
			},
			{
				From:        "Aug 2009",
				To:          "Nov 2011",
				Project:     "HireVue",
				Company:     "SoftServe",
				Position:    "Web developer",
				Description: "When I came to this company I was the second developer on the project. After two years our team grew to 10 developers and 2 QAs. It was a great experience of teamwork. My responsibility was development of online video interview system, deploying and hot maintenance. <br/>Technology: PHP, Zend Framework, MySql, Jquery, PHPUnit",
			},
			{
				From:        "Mar 2007",
				To:          "Aug 2009",
				Project:     "InterLogic",
				Company:     "",
				Position:    "Web developer",
				Description: "It was my first outsourcing experience. I had very different work. Some projects took 1-2 days other more than a year. I worked with PHP, ActiveScript and C#",
			},
			{
				From:        "May 2005",
				To:          "Dec 2006",
				Project:     "VashComfort",
				Company:     "OWBT",
				Position:    "Web developer",
				Description: "First commercial experience of working in a small team. At the same time, I continued studing in the University. <br/>Technology: PHP, MySQL, JavaScript."},
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
