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
	PageTypeViral           = "viral"
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

// Future Predictions
var FuturePredictions = []string{
	"You'll text your boss about your llama weekend plans by mistake.",
	"You'll trip, spill drinks, but somehow still look cool in front of your crush.",
	"Your pet will learn to talk, only to spill your embarrassing secrets.",
	"You'll set off the smoke alarm twice trying a fancy cooking move.",
	"You'll wear two different shoes to a big meeting, unnoticed until later.",
	"Stuck in a public restroom, your only escape is a tiny window.",
	"You'll sing loudly in a silent library, thinking you're alone.",
	"You'll accidentally call your teacher 'mom' or 'dad'.",
	"A dance move gone wrong will lead to an unintentional worm.",
	"You'll walk into the wrong house, interrupting an awkward family dinner.",
	"A bird will poop on your head mid-selfie.",
	"You'll 'reply all' to an email complaining about the sender.",
	"Head stuck in a sweater, someone important walks in.",
	"You'll forget a joke's punchline, leading to awkward silence.",
	"You'll order a giant item online with no idea where to put it.",
	"Trying to sneak out, you'll loudly trip over your own feet.",
	"You'll wear clothes inside out all day, oblivious.",
	"A high-five attempt will result in an awkward air-slap.",
	"Coffee will spill on your white shirt right before a big presentation.",
	"Leaning on nothing, you'll fall spectacularly.",
	"You'll accidentally like an old social media post from years ago.",
	"Caught singing off-key at a red light.",
	"You'll walk into a glass door, leaving a face imprint.",
	"Trying to pick something up, you'll do a full public squat.",
	"A toy will make an embarrassing noise during a work meeting.",
}

// Base prediction strings for generators
const (
	PredictionBaseAmazingEvent      = "In the next %d days, something amazing will happen to %s."
	PredictionBaseMarried           = "You will get married in %d years."
	PredictionBaseMillionaire       = "You will become a millionaire in %d years."
	PredictionBaseFutureChild       = "Here is what your future child could look like."
	PredictionBaseLiveInCountry     = "You will live in %s."
)

// Fortune Readings
var FortuneReadings = []string{
	"Your next big idea will come to you in the shower, and you'll forget it by the time you grab a towel.",
	"A mysterious stranger will offer you cryptic advice, which will turn out to be a coupon for a discount on socks.",
	"You will find a small, shiny object today that will bring you absolutely no luck whatsoever.",
	"Your future holds a moment of profound realization: you've been pronouncing a common word wrong your entire life.",
	"A sudden gust of wind will reveal your most embarrassing fashion choice to a crowd of onlookers.",
	"You will accidentally become a local legend for a completely mundane act, like perfectly stacking groceries.",
	"Your lucky number will lead you to a parking spot, but it will be for a unicycle.",
	"A fortune cookie will predict your future, but it will be in a language you don't understand.",
	"You will receive an unexpected inheritance, which turns out to be a collection of antique thimbles.",
	"Your attempt to be spontaneous will lead to you getting lost in a corn maze for an hour.",
	"A cosmic alignment will grant you a superpower: the ability to perfectly toast bread every time.",
	"You will discover a secret talent for competitive napping, winning a bronze medal.",
	"Your future self will send you a warning about a bad hair day, but it will arrive too late.",
	"You will accidentally invent a new dance move that involves tripping, and it will become a viral sensation.",
	"A talking animal will reveal a profound truth to you, but it will be about the best brand of cat food.",
	"Your lucky charm will only work when you're wearing mismatched socks.",
	"You will win a contest you don't remember entering, and the prize will be a year's supply of lukewarm tap water.",
	"A moment of destiny will arrive, and you'll be too busy looking for your phone to notice.",
	"Your future holds a brief, but intense, rivalry with a squirrel over a single acorn.",
	"You will finally understand the meaning of life, but then immediately forget it because you're hungry.",
}
