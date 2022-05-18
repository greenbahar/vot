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
	vote_entity_sep  = "|"
	vote_result_sep  = ";;"
	key_val_sep      = "::"
	input_option_sep = "//"
)

type Voter string

type ElectionOption struct {
	Option string
	code   uint32
}

type OptionCounter struct {
	Option  string
	Counter int
}

type Vote struct {
	Question              string
	ValidTimeToVoteInDays time.Time
	VotingResult          map[ElectionOption]int
}

func NewVote(question string, days uint32, options ...ElectionOption) *Vote {
	votingResult := make(map[ElectionOption]int)
	for _, item := range options {
		votingResult[item] = 0
	}

	return &Vote{
		Question:              question,
		ValidTimeToVoteInDays: GetNowUTC().AddDate(0, 0, int(days)),
		VotingResult:          votingResult,
	}
}

func ParsOptions(options string) []ElectionOption {
	electionOptions := make([]ElectionOption, 0)

	items := strings.Split(options, input_option_sep)
	for index, item := range items {
		electionOptions = append(electionOptions, ElectionOption{Option: item, code: uint32(index + 1)})
	}

	return electionOptions
}

func (v *Vote) TimeToVoteIsValid() bool {
	return v.ValidTimeToVoteInDays.Sub(GetNowUTC()) > 0
}

func (v *Vote) VoteFor(optionCode uint32) error {
	if v.TimeToVoteIsValid() {
		for key := range v.VotingResult {
			if key.code == optionCode {
				v.VotingResult[key]++
				return nil
			}
		}

		return errors.New("this code option is not an election option")
	}

	return errors.New("voting time is expired")
}

func (v *Vote) WinnerOption() (*ElectionOption, error) {
	if GetNowUTC().Sub(v.ValidTimeToVoteInDays) < 0 {
		return nil, errors.New("the end of the poll is yet to come")
	}

	maxCount := 0
	var winner = new(ElectionOption)
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

func (v *Vote) GetVoteOptions() []*ElectionOption {
	options := make([]*ElectionOption, 0)
	for item := range v.VotingResult {
		options = append(options, &item)
	}
	return options
}

func (v *Vote) GetVoteOptionsWithCounter() []OptionCounter {
	var results []OptionCounter
	var item = new(OptionCounter)
	for key, val := range v.VotingResult {
		item.Option = key.Option
		item.Counter = val
		results = append(results, *item)
	}

	return results
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

func (v Vote) String() (string, error) {
	var buf bytes.Buffer
	buf.WriteString(v.Question)
	buf.WriteString(vote_entity_sep)
	buf.WriteString(GetTimeFormatString(v.ValidTimeToVoteInDays))
	buf.WriteString(vote_entity_sep)

	counter := len(v.VotingResult)
	for key, val := range v.VotingResult {
		buf.WriteString(key.Option)
		buf.WriteString(key_val_sep)
		buf.WriteString(strconv.Itoa(int(key.code)))
		buf.WriteString(key_val_sep)
		buf.WriteString(strconv.Itoa(val))
		counter--
		if counter != 0 {
			buf.WriteString(vote_result_sep)
		}
	}

	return buf.String(), nil
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
		option := strings.Split(item, key_val_sep)

		optionCode, codeErr := strconv.Atoi(option[1])
		if codeErr != nil {
			return nil, errors.New("parse error")
		}
		electionOption = ElectionOption{Option: option[0], code: uint32(optionCode)}
		optionCount, err := strconv.Atoi(option[2])
		if err != nil {
			return nil, errors.New(fmt.Sprintf("invalid vote, invalid metadata: %v", err))
		}

		votingResult[electionOption] = optionCount
	}
	vote.VotingResult = votingResult

	return vote, nil
}
