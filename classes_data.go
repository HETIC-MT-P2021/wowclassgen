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

var TankAssociation = []string{
	"DeathKnight-Blood",
	"DemonHunter-Vengeance",
	"Druid-Guardian",
	"Monk-Brewmaster",
	"Paladin-Protection",
	"Warrior-Protection",
}

func contains(haystack []string, needle string) bool {
	for _, str := range haystack {
		if needle == str {
			return true
		}
	}
	return false
}
