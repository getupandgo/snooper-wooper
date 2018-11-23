package mock

type Token struct {
	Text string
	Count uint32
}

var tokens = []Token { 
    {
        "Vanya",
        50,
    },
    {
		"soset",
        30,
	},
	{
		"pisos",
        25,
	},
	{
		"big",
        20,
	},
	{
		"black",
        18,
	},
	{
		"white",
        15,
	},
	{
		"seven",
        14,
	},
	{
		"eight",
        12,
	},
	{
		"ten",
        10,
	},
	{
		"eleven",
        9,
	},
	{
		"twelve",
        98,
	},
	{
		"thirteen",
        7,
    },
}

func GetTokens(limit uint64) []Token {
	if limit == 0 {
		limit = 10
	}

	return tokens[:limit]
}