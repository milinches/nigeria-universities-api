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
		&Universities{Name: "Adekunle Ajasin University", Abbrv: "AAUA", WebsiteLink: "https://aaua.edu.ng/"},
		&Universities{Name: "Adeleke University, Ede", Abbrv: "AUE", WebsiteLink: "https://adelekeuniversity.edu.ng/"},
		&Universities{Name: "Admiralty University of Nigeria", Abbrv: "ADUN", WebsiteLink: "https://adun.edu.ng/"},
		&Universities{Name: "Afe Babalola University", Abbrv: "ABUAD", WebsiteLink: "https://www.abuad.edu.ng/"},
		&Universities{Name: "African Univeristy of Science and Technology", Abbrv: "AUST", WebsiteLink: "https://aust.edu.ng/"},
		&Universities{Name: "Ahmadu Bello University, Zaria", Abbrv: "ABUZARIA", WebsiteLink: "https://www.abu.edu.ng/"},
		&Universities{Name: "Ajayi Crowther University", Abbrv: "ACU", WebsiteLink: "https://acu.edu.ng/"},
		&Universities{Name: "Akwa Ibom State University", Abbrv: "AKSU", WebsiteLink: "https://aksu.edu.ng/newsite/"},
		&Universities{Name: "Alex Ekwueme Federal University, Ndufu-Alike", Abbrv: "FUNAI", WebsiteLink: "https://funai.edu.ng/"},
		&Universities{Name: "Al-Hikmah University", Abbrv: "HUI", WebsiteLink: "https://alhikmahuniversity.edu.ng/"},
		&Universities{Name: "Al-Qalam University, Katsina", Abbrv: "AUK", WebsiteLink: "https://auk.edu.ng/"},
		&Universities{Name: "Ambrose Alli University", Abbrv: "AAU", WebsiteLink: "https://www.aauekpoma.edu.ng/"},
		&Universities{Name: "American University of Nigeria", Abbrv: "AUN", WebsiteLink: "https://www.aun.edu.ng/"},
		&Universities{Name: "", Abbrv: "", WebsiteLink: ""},
		&Universities{Name: "", Abbrv: "", WebsiteLink: ""},
		&Universities{Name: "", Abbrv: "", WebsiteLink: ""},
		&Universities{Name: "", Abbrv: "", WebsiteLink: ""},
		&Universities{Name: "", Abbrv: "", WebsiteLink: ""},
		&Universities{Name: "", Abbrv: "", WebsiteLink: ""},
		&Universities{Name: "", Abbrv: "", WebsiteLink: ""},
		&Universities{Name: "", Abbrv: "", WebsiteLink: ""},
		&Universities{Name: "", Abbrv: "", WebsiteLink: ""},
	}
}