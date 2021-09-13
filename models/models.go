package models

type Universities struct{
	Name string `json:"name"`
	Abbrv string `json:"abbreviation"`
	WebsiteLink string `json:"websitelink"`
}

func Uni() []*Universities {
	return []*Universities{
		&Universities{Name: "Abia State University, Uturu", Abbrv: "ABSU", WebsiteLink: "https://abiastateuniversity.edu.ng/"},
		&Universities{Name: "Achievers University, Owo", Abbrv: "AUO", WebsiteLink: "https://www.achievers.edu.ng/"},
		&Universities{Name: "Adamawa State University", Abbrv: "ADSU", WebsiteLink: "https://adsu.edu.ng/"},
		&Universities{Name: "", Abbrv: "", WebsiteLink: ""},
	}
}