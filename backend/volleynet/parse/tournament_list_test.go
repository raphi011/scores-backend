package parse

import (
	"os"
	"testing"
	"time"

	"github.com/raphi011/scores/test"
	"github.com/raphi011/scores/volleynet"
)

func mustParseDate(date string) time.Time {
	result, err := time.Parse("02.01.2006", date)

	if err != nil {
		panic("date could not be parsed")
	}

	return result
}

func TestTournamentList(t *testing.T) {
	response, _ := os.Open("../testdata/tournament-list-amateur.html")

	tournaments, err := TournamentList(response, "http://example.com")

	test.Check(t, "Tournaments() err: %s", err)
	test.Compare(t, "TournamentList() err: mismatch of tournament list\n%s", tournaments, tournamentListAmateur)
}

var tournamentListAmateur = []*volleynet.TournamentInfo{
	&volleynet.TournamentInfo{
		Name:             "Herren Beachvolley Wien Summer Opening",
		Start:            mustParseDate("21.04.2018"),
		End:              mustParseDate("21.04.2018"),
		ID:               21908,
		Season:           2018,
		RegistrationOpen: false,
		Gender:           "M",
		Status:           "upcoming",
		League:           "AMATEUR TOUR",
		Phase:            "ABV Tour AMATEUR 1",
		Link:             "http://example.com/beach/bewerbe/AMATEUR TOUR/phase/ABV Tour AMATEUR 1/sex/M/saison/2018/cup/21908",
	},
	&volleynet.TournamentInfo{
		Name:             "Herren Graz",
		Start:            mustParseDate("21.04.2018"),
		End:              mustParseDate("22.04.2018"),
		ID:               21880,
		Season:           2018,
		RegistrationOpen: false,
		Gender:           "M",
		Status:           "upcoming",
		League:           "AMATEUR TOUR",
		Phase:            "ABV Tour AMATEUR 1",
		Link:             "http://example.com/beach/bewerbe/AMATEUR TOUR/phase/ABV Tour AMATEUR 1/sex/M/saison/2018/cup/21880",
	},
	&volleynet.TournamentInfo{
		Name:             "Herren Beachvolley Grieskirchen",
		Start:            mustParseDate("06.05.2018"),
		End:              mustParseDate("06.05.2018"),
		ID:               22055,
		Season:           2018,
		RegistrationOpen: true,
		Gender:           "M",
		Status:           "upcoming",
		League:           "AMATEUR TOUR",
		Phase:            "ABV Tour AMATEUR 1",
		Link:             "http://example.com/beach/bewerbe/AMATEUR TOUR/phase/ABV Tour AMATEUR 1/sex/M/saison/2018/cup/22055",
		EntryLink:        "https://beach.volleynet.at/Anmelden/21617-22055-00",
	},
	&volleynet.TournamentInfo{
		Name:             "Herren Innsbruck",
		Start:            mustParseDate("01.05.2018"),
		End:              mustParseDate("01.05.2018"),
		ID:               21938,
		Season:           2018,
		RegistrationOpen: false,
		Gender:           "M",
		Status:           "canceled",
		League:           "AMATEUR TOUR",
		Phase:            "ABV Tour AMATEUR 1",
		Link:             "http://example.com/beach/bewerbe/AMATEUR TOUR/phase/ABV Tour AMATEUR 1/sex/M/saison/2018/cup/21938",
	},
}
