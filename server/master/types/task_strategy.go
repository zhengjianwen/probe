package types

import "errors"

type Strategy struct {
	TaskId         string      `json:"-",bson:"taskid"`
	NoteMax        uint64      `json:"max",bson:"max"`                 //Max notification num
	NodeTimeBucket *TimeBucket `json:"bucket",bson:"bucket"`           //notification time bucket
	All            bool        `json:"all",bson:"all"`                 //note condition: all workers down
	AnyN           uint16      `json:"anyN",bson:"anyN"`               //note condition: large than N down
	AnySpec        []string    `json:"anySpecList",bson:"anySpecList"` //note condition: any spec worker down
	AllSpec        []string    `json:"allSpecList",bson:"allSpecList"` //note condition: all spec worker down
}

func (s *Strategy) Validate() error {
	n := 0
	if s.All {
		n++
	}
	if s.AnyN > 0 {
		n++
	}
	if len(s.AnySpec) > 0 {
		n++
	}
	if len(s.AllSpec) > 0 {
		n++
	}

	if n > 1 || n == 0 {
		return errors.New("strategy type must specific one")
	}

	if s.NodeTimeBucket != nil {
		if s.NodeTimeBucket.Start >= 24 || s.NodeTimeBucket.Start < 0 {
			return errors.New("param bucket.start invalid")
		}
		if s.NodeTimeBucket.End >= 24 || s.NodeTimeBucket.End < 0 {
			return errors.New("param bucket.start invalid")
		}

		if s.NodeTimeBucket.Start > s.NodeTimeBucket.End {
			return errors.New("param bucket invalid")
		}
	}

	return nil
}

type TimeBucket struct {
	Start uint8
	End   uint8
}
