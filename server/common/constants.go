package common

// Error messages
const (
	ErrCannotParseJSON          = "Cannot parse JSON"
	ErrInvalidPageType          = "Invalid page_type"
	ErrFailedToCreatePrediction = "Failed to create prediction data"
	ErrFailedToSavePrediction   = "Failed to save prediction"
	ErrInvalidSlug              = "Invalid slug"
	ErrPredictionNotFound       = "Prediction not found"
	ErrFailedToParsePrediction  = "Failed to parse prediction data"
)

// Page types
const (
	PageTypeViral            = "viral"
	PageTypeFuturePrediction = "future_prediction"
)

// Default values
const (
	DefaultPredictionTitle = "Viral Prediction"
)

// Prediction Titles
var PredictionTitles = map[string]string{
	"future-prediction":                  "Future Prediction",
	"when-will-you-get-married":          "When will you get married?",
	"when-will-you-become-a-millionaire": "When will you become a millionaire?",
	"your-future-childs-face":            "Your future child's face",
	"which-country-will-you-live-in":     "Which country will you live in?",
	"your-2025-fortune-reading":          "Your 2025 fortune reading",
}
var FuturePredictions = []string{
	"You’ll bring home a girlfriend. Mom asks about wedding venues before she sits down.",
	"Auntie will ask your salary loud enough for the neighbors to update their group chat.",
	"You’ll score 99%. Mom asks where the other 1% went.",
	"You’ll come home after 10 p.m. Parents waiting on the couch like it’s a crime drama.",
	"Mom will forward you a girl’s full résumé, blood type, and childhood piano videos.",
	"You’ll say ‘I’m full.’ Mom: ‘Eat or I stop eating forever.’",
	"You’ll get into a top uni. Relatives: ‘Still not number one, sad.’",
	"You’ll cough once. Grandma attacks with every herb known to mankind.",
	"You’ll reach for the bill. Three uncles body-check you for honor.",
	"You’ll introduce your girlfriend. First question: ‘What do her parents do for work?’",
	"You’ll get a big bonus. Mom already spent it on your future house down payment.",
	"You’ll wear shorts in winter. Mom wraps you in three jackets like a burrito.",
	"You’ll sneeze. Family performs full exorcism with oils and hot towels.",
	"You’ll say you’re tired. Dad: ‘Sleep is for after you’re successful.’",
	"Next family dinner question: ‘When girlfriend? When house? When baby?’ In that order.",
	"You’ll get a 98. Report card disappears into the void.",
	"You’ll bring home a friend. Grandma gives them 90% of the food ‘because guest.’",
	"You’ll try to leave the table. Mom packs enough leftovers for a month.",
	"You’ll hit 27. Every auntie starts a countdown timer in the group chat.",
	"You’ll buy a nice car. Cousin already bought a nicer one last week.",
}

// Base prediction strings for generators
const (
	PredictionBaseAmazingEvent  = "In the next %d days, something amazing will happen to %s."
	PredictionBaseMarried       = "You will get married in %d years."
	PredictionBaseMillionaire   = "You will become a millionaire in %d years."
	PredictionBaseFutureChild   = "Here is what your future child could look like."
	PredictionBaseLiveInCountry = "You will live in %s."
)

// Fortune Readings
var FortuneReadings = []string{
	"Mom already printed the wedding invites. You haven’t met the girl yet.",
	"Your blind date is a dentist from Shanghai. She’s 26 and ‘running out of time.’",
	"She asks your household income before asking your name.",
	"Future wife chosen by mom. You just handle the ring and the mortgage.",
	"Auntie sends girl’s WeChat QR code + height + degree + parents’ job title.",
	"You’re 5’10” and still the shortest guy at family dinner.",
	"Date cancels because your horoscope ‘not compatible’ with hers.",
	"Mom cries because you’re 29 and still no baby photo to send relatives.",
	"She likes you until she hears you rent in Flushing, not own in Great Neck.",
	"Future mother-in-law asks if you can transfer property title before engagement.",
	"Girl says ‘I want a simple wedding.’ Budget starts at 88 tables.",
	"Your red envelope to her parents heavier than your yearly bonus.",
	"Mom sees your Tinder: ‘Why pay for app when I have 47 girls in my phone?’",
	"She asks if you’re the ‘marry type’ on date #1.",
	"Future kid must play piano AND violin or ‘wasted my suffering.’",
	"Date brings calculator to KBBQ to split the bill by who ate what.",
	"You finally bring girlfriend home. Grandma asks why she’s not a doctor.",
	"She ghosts after finding out your dad isn’t a factory owner.",
	"Lunar New Year question of the year: ‘Got girlfriend yet or still got time?’",
	"Future wedding emcee already booked. You’re still looking for the bride.",
}
