package rules

import (
	"errors"
	"time"
)

const (
	one_day = time.Hour * 24
)

type Voter struct {}

type ElectionOption struct {
	Option string
}

type Vote struct {
	Question 				string
	Options 				[]ElectionOption
	VotingResult			map[ElectionOption] int
	ValidTimeToVoteInDays 	time.Time
}

func NewVote(question string, days int, options ...ElectionOption) *Vote{
	return &Vote{
		Question: question,
		Options: options,
		VotingResult: make(map[ElectionOption]int),
		ValidTimeToVoteInDays: GetNowUTC().AddDate(0, 0, days),
	}
}

func (v *Vote) TimeToVoteIsValid() bool {
	return v.ValidTimeToVoteInDays.Sub(GetNowUTC()) > 0
}

func (v *Vote) VoteFor(option string)  {
	if v.TimeToVoteIsValid(){
		selectedOption := ElectionOption{
			option,
		}

		v.VotingResult[selectedOption] ++
	}
}

func (v *Vote) WinnerOption() (*ElectionOption, error) {
	if GetNowUTC().Sub(v.ValidTimeToVoteInDays) < 0 {
		return nil, errors.New("the end of the poll is yet to come")
	}

	maxCount := 0
	var winner *ElectionOption
	for option, count := range v.VotingResult {
		if count > maxCount {
			maxCount = count
			winner = &option
		}
	}

	if maxCount == 0 {
		return nil, errors.New("no one participated in the poll")
	}

	return winner, nil
}

func GetNowUTC () time.Time {
	return time.Now().UTC()
}