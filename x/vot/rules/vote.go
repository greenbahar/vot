package rules

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	one_day         = time.Hour * 24
	vote_entity_sep = "|"
	vote_result_sep = "||"
	Key_val_sep     = "::"
)

type Voter struct{}

type ElectionOption struct {
	Option string
}

type Vote struct {
	Question              string
	ValidTimeToVoteInDays time.Time
	//Options 				[]ElectionOption
	VotingResult map[ElectionOption]int
}

func NewVote(question string, days int, options ...ElectionOption) *Vote {
	votingResult := make(map[ElectionOption]int)
	for _, item := range options {
		votingResult[item] = 0
	}

	return &Vote{
		Question:              question,
		ValidTimeToVoteInDays: GetNowUTC().AddDate(0, 0, days),
		//Options: options,
		VotingResult: votingResult,
	}
}

func (v *Vote) TimeToVoteIsValid() bool {
	return v.ValidTimeToVoteInDays.Sub(GetNowUTC()) > 0
}

func (v *Vote) VoteFor(option string) {
	if v.TimeToVoteIsValid() {
		selectedOption := ElectionOption{
			option,
		}

		v.VotingResult[selectedOption]++
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

func (v *Vote) GetVoteOptions() []ElectionOption {
	options := make([]ElectionOption, 0)
	for item := range v.VotingResult {
		options = append(options, item)
	}
	return options
}

func GetNowUTC() time.Time {
	return time.Now().UTC()
}

func GetTimeFormatString(t time.Time) string {
	return t.Format(time.RFC1123Z)
}

func ParseTime(t string) (time.Time, error) {
	return time.Parse(time.RFC1123Z, t)
}

func (v *Vote) String() string {
	var buf bytes.Buffer
	buf.WriteString(v.Question)
	buf.WriteString(vote_entity_sep)
	buf.WriteString(GetTimeFormatString(v.ValidTimeToVoteInDays))
	buf.WriteString(vote_entity_sep)

	for key, val := range v.VotingResult {
		buf.WriteString(key.Option)
		buf.WriteString(Key_val_sep)
		buf.WriteString(strconv.Itoa(val))
		buf.WriteString(vote_result_sep)
	}

	return buf.String()
}

func Parse(s string) (*Vote, error) {
	parsedData := strings.Split(s, vote_entity_sep)
	vote := &Vote{}
	vote.Question = parsedData[0]

	voteTimeUp, err := ParseTime(parsedData[1])
	if err != nil {
		return nil, errors.New(fmt.Sprintf("invalid vote, invalid metadata: %v", err))
	}
	vote.ValidTimeToVoteInDays = voteTimeUp

	votingResult := make(map[ElectionOption]int)
	electionOption := ElectionOption{}
	for _, item := range strings.Split(parsedData[2], vote_result_sep) {
		option := strings.Split(item, Key_val_sep)

		electionOption = ElectionOption{option[0]}
		optionCount, err := strconv.Atoi(option[1])
		if err != nil {
			return nil, errors.New(fmt.Sprintf("invalid vote, invalid metadata: %v", err))
		}

		votingResult[electionOption] = optionCount
	}
	vote.VotingResult = votingResult

	return vote, nil
}
