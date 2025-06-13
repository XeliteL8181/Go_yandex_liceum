package final

import (
	"testing"
)

func TestParseLine(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		name, gStr, mStr, aStr := ParseLine("Роналдо 10 5 3")
		if name != "Роналдо" || gStr != "10" || mStr != "5" || aStr != "3" {
			t.Errorf("ParseLine returned %q, %q, %q, %q; want %q, %q, %q, %q",
				name, gStr, mStr, aStr,
				"Роналдо", "10", "5", "3",
			)
		}
	})

	t.Run("invalid input", func(t *testing.T) {
		name, gStr, mStr, aStr := ParseLine("TooFewFields")
		if name != "" || gStr != "" || mStr != "" || aStr != "" {
			t.Errorf("Ввод запаршен неправильно: %q, %q, %q, %q",
				name, gStr, mStr, aStr)
		}
	})
}

func TestNewPlayerAndCalculateRating(t *testing.T) {
	tests := []struct {
		name                    string
		goals, misses, assists int
		wantRating              float64
	}{
		{"Без промахов", 8, 0, 2, 8 + 2.0/2},
		{"С промахами", 10, 5, 4, (10 + 4.0/2) / 5},
	}

	for _, tt := range tests {
		p := NewPlayer(tt.name, tt.goals, tt.misses, tt.assists)
		if p.Name != tt.name {
			t.Errorf("NewPlayer: got Name=%q; want %q", p.Name, tt.name)
		}
		if p.Goals != tt.goals || p.Misses != tt.misses || p.Assists != tt.assists {
			t.Errorf("NewPlayer: fields mismatch; got %+v; want goals=%d, misses=%d, assists=%d",
				p, tt.goals, tt.misses, tt.assists)
		}
		if p.Rating != tt.wantRating {
			t.Errorf("calculateRating: for %+v got Rating=%.2f; want %.2f",
				p, p.Rating, tt.wantRating)
		}
	}
}

func TestGoalSort(t *testing.T) {
	players := []Player{
		NewPlayer("A", 5, 1, 0),
		NewPlayer("B", 8, 2, 0),
		NewPlayer("C", 3, 0, 1),
	}
	sorted := goalsSort(players)
	wantOrder := []string{"B", "A", "C"}
	for i, wantName := range wantOrder {
		if sorted[i].Name != wantName {
			t.Errorf("GoalSort position %d = %q; want %q", i, sorted[i].Name, wantName)
		}
	}
}

func TestRatingSort(t *testing.T) {
	players := []Player{
		NewPlayer("Low", 2, 2, 0),   // rating = 1
		NewPlayer("High", 10, 2, 0), // rating = 5
		NewPlayer("Mid", 4, 2, 0),   // rating = 2
	}
	sorted := ratingSort(players)
	wantOrder := []string{"High", "Mid", "Low"}
	for i, wantName := range wantOrder {
		if sorted[i].Name != wantName {
			t.Errorf("ratingSort position %d = %q; want %q", i, sorted[i].Name, wantName)
		}
	}
}

func TestGmSort(t *testing.T) {
	players := []Player{
		NewPlayer("ZeroMiss", 3, 0, 0), // gm = 3
		NewPlayer("LowRat", 4, 2, 0),   // gm = 2
		NewPlayer("HighRat", 5, 1, 0),  // gm = 5
	}
	sorted := gmSort(players)
	wantOrder := []string{"HighRat", "ZeroMiss", "LowRat"}
	for i, wantName := range wantOrder {
		if sorted[i].Name != wantName {
			t.Errorf("gmSort position %d = %q; want %q", i, sorted[i].Name, wantName)
		}
	}
}