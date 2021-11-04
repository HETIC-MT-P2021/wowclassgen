package main

type Data struct {
	GameData GameData `json:"gameData"`
}

type GameData struct {
	Classes []Classes `json:"classes"`
}

type Classes struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Specs []Specs `json:"specs"`
}

type Specs struct {
	Name string `json:"name"`
}

var HealAssociation = []string{
	"Monk-Mistweaver",
	"Druid-Restoration",
	"Paladin-Holy",
	"Priest-Discipline",
	"Priest-Holy",
	"Shaman-Restoration",
}

func Contains(haystack []string, needle string) bool {
	for _, str := range haystack {
		if needle == str {
			return true
		}
	}
	return false
}
