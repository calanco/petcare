package food

import "testing"

func TestGetFood(t *testing.T) {
	for _, test := range Tests {
		if getFood(test.Pet) != test.Food {
			t.Errorf("Test with %s doesn't pass", test.Pet.Name)
			continue
		}
		t.Logf("Test with %s passes", test.Pet.Name)
	}
}
