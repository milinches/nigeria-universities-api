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
		&Universities{Name: "Anchor University, Lagos", Abbrv: "AUL", WebsiteLink: "https://aul.edu.ng/"},
		&Universities{Name: "Arthur Jarvis University", Abbrv: "AJU", WebsiteLink: "https://arthurjarvisuniversity.edu.ng/"},
		&Universities{Name: "Atiba University", Abbrv: "ATU", WebsiteLink: "https://www.atibauniversity.edu.ng/"},
		&Universities{Name: "Augustine University", Abbrv: "AUI", WebsiteLink: "https://augustineuniversity.edu.ng/"},
		&Universities{Name: "Babcock University", Abbrv: "BU", WebsiteLink: "https://www.babcock.edu.ng/"},
		&Universities{Name: "Bauchi State University", Abbrv: "BASUG", WebsiteLink: "Bauchi State University"},
		&Universities{Name: "Bayero University Kano", Abbrv: "BUK", WebsiteLink: "https://www.buk.edu.ng/"},
		&Universities{Name: "Baze University", Abbrv: "BAZE", WebsiteLink: "https://www.bazeuniversity.edu.ng/"},
		&Universities{Name: "Bells University of Technology", Abbrv: "BELLSTECH", WebsiteLink: "https://www.bellsuniversity.edu.ng/"},
		&Universities{Name: "Benson Idahosa University", Abbrv: "BIU", WebsiteLink: "https://www.biu.edu.ng/"},
		&Universities{Name: "Benue State University", Abbrv: "BSUM", WebsiteLink: "https://www.bsum.edu.ng/"},
		&Universities{Name: "Bingham University", Abbrv: "BU", WebsiteLink: "https://binghamuni.edu.ng/v2/"},
		&Universities{Name: "Borno State University", Abbrv: "BOSU", WebsiteLink: "https://bosu.edu.ng/"},
		&Universities{Name: "Bowen University", Abbrv: "BOWEN", WebsiteLink: "https://bowen.edu.ng/"},
		&Universities{Name: "Caleb University", Abbrv: "CUL", WebsiteLink: "https://calebuniversity.edu.ng/"},
		&Universities{Name: "Caritas Univeristy", Abbrv: "CUA", WebsiteLink: "https://www.caritasuni.edu.ng/"},
		&Universities{Name: "Chrisland University", Abbrv: "CHRISLAND", WebsiteLink: "https://www.chrislanduniversity.edu.ng/"},
		&Universities{Name: "Christopher University", Abbrv: "UNICHRIS", WebsiteLink: "https://christopheruniversity.edu.ng/"},
		&Universities{Name: "Chukwuemeka Odumegwu Ojukwu University", Abbrv: "COOU", WebsiteLink: "https://coou.edu.ng/"},
		&Universities{Name: "Clifford University", Abbrv: "CLU", WebsiteLink: "https://clifforduni.edu.ng/"},
		&Universities{Name: "Coal City University", Abbrv: "CCU", WebsiteLink: "https://ccu.edu.ng/"},
		&Universities{Name: "Convenant University", Abbrv: "CU", WebsiteLink: "https://www.covenantuniversity.edu.ng/"},
		&Universities{Name: "Crawford University", Abbrv: "CRU", WebsiteLink: "http://crawforduniversity.edu.ng/crawford/"},
		&Universities{Name: "Crescent University, Abeokuta", Abbrv: "CUAB", WebsiteLink: ""},
		&Universities{Name: "", Abbrv: "", WebsiteLink: ""},
		&Universities{Name: "", Abbrv: "", WebsiteLink: ""},
		&Universities{Name: "", Abbrv: "", WebsiteLink: ""},
		&Universities{Name: "", Abbrv: "", WebsiteLink: ""},
	}
}