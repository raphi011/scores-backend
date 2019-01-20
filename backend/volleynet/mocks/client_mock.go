package mocks

import (
	"github.com/raphi011/scores/volleynet"
	"github.com/stretchr/testify/mock"
)

type ClientMock struct {
	mock.Mock
}

func (m *ClientMock) GetTournamentLink(t *volleynet.TournamentInfo) string {
	args := m.Called(t)

	return args.String(0)
}

func (m *ClientMock) GetAPITournamentLink(t *volleynet.TournamentInfo) string {
	args := m.Called(t)

	return args.String(0)
}

func (m *ClientMock) Login(username, password string) (*volleynet.LoginData, error) {
	args := m.Called(username, password)

	return args.Get(0).(*volleynet.LoginData), args.Error(1)
}

func (m *ClientMock) AllTournaments(gender, league string, year int) ([]*volleynet.TournamentInfo, error) {
	args := m.Called(gender, league, year)

	return args.Get(0).([]*volleynet.TournamentInfo), args.Error(1)
}

func (m *ClientMock) Ladder(gender string) ([]*volleynet.Player, error) {
	args := m.Called(gender)

	return args.Get(0).([]*volleynet.Player), args.Error(1)
}

func (m *ClientMock) ComplementMultipleTournaments(tournament []*volleynet.TournamentInfo) ([]*volleynet.Tournament, error) {
	args := m.Called(tournament)

	return args.Get(0).([]*volleynet.Tournament), args.Error(1)
}

func (m *ClientMock) ComplementTournament(tournament *volleynet.TournamentInfo) (*volleynet.Tournament, error) {
	args := m.Called(tournament)

	return args.Get(0).(*volleynet.Tournament), args.Error(1)
}

func (m *ClientMock) TournamentWithdrawal(tournamentID int) error {
	args := m.Called(tournamentID)

	return args.Error(0)
}

func (m *ClientMock) TournamentEntry(playerName string, playerID, tournamentID int) error {
	args := m.Called(playerName, playerID, tournamentID)

	return args.Error(0)
}

func (m *ClientMock) SearchPlayers(firstName, lastName, birthday string) ([]*volleynet.PlayerInfo, error) {
	args := m.Called(firstName, lastName, birthday)

	return args.Get(0).([]*volleynet.PlayerInfo), args.Error(1)
}