package pet

// TestPet defines the struct to run test with
type TestPet struct {
	testName string // defines the name of the test
	pet      Pet    // represents the pet to test
	ok       bool   // defines if the pet has to be correctly created
}

// TestPets contains all the needed test to run
var TestPets = []TestPet{
	{
		"Working test",
		Pet{
			Name:    "Bruce0",
			Species: "Dog",
		},
		true,
	},
	{
		"Wrong breed test",
		Pet{
			Name:    "Tina0",
			Species: "Dog",
			Breed:   "Unknown",
		},
		false,
	},
	{
		"Working date test",
		Pet{
			Name:    "Tina1",
			Species: "Dog",
			Date:    "2006-May-30",
		},
		true,
	},
	{
		"Wrong date test",
		Pet{
			Name:    "Tina2",
			Species: "Dog",
			Date:    "Wrong date",
		},
		false,
	},
	{
		"Working size",
		Pet{
			Name:    "Kira0",
			Species: "Cat",
			Size:    "Medium",
		},
		true,
	},
	{
		"Working weight",
		Pet{
			Name:    "Kira1",
			Species: "Cat",
			Weight:  2.5,
		},
		true,
	},
	{
		"Wrong species",
		Pet{
			Name:    "Unknown0",
			Species: "Unknown",
		},
		false,
	},
}
