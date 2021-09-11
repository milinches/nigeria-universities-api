package uni

type Universities struct {
	UniveristyID int
	UniversityName string
	UniversityAbbrv string
	UniversityWebsite string
}

func uni() []*Universities{
	return []*Universities {
		&Universities{UniveristyID: 1, UniversityName: "Abia State Univeristy", UniversityAbbrv: "ABSU", UniversityWebsite: "https://absu.go.ng"}
	}
}