package src

import (
	"reflect"
	"testing"
)

func TestGetMatchesFromMatchFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want *Matches
	}{
		{name: "match/match.json" ,args: args{path:"../match/match.json"}},
		{name: "match/match_result.json" ,args: args{path:"../match/match_result.json"}},
		{name: "match/seven_cards_with_ghost.json" ,args: args{path:"../match/seven_cards_with_ghost.json"}},
		{name: "match/seven_cards_with_ghost.result.json" ,args: args{path:"../match/seven_cards_with_ghost.result.json"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMatchesFromMatchFile(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMatchesFromMatchFile() = %v, want %v", got, tt.want)
			}

		})
	}
}

