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
	PageTypeCats             = "cats"
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
	"how-many-kids-will-you-have":        "How many kids will you have?",
	"which-country-will-you-live-in":     "Which country will you live in?",
	"your-2025-fortune-reading":          "Your 2025 fortune reading",
	"what-cat-are-you":                   "What Cat Are You?",
}
var FuturePredictions = []string{
	"You will make a life-changing decision within the next 30 days.",
	"A huge opportunity you once thought you lost will unexpectedly return.",
	"Your career will take a turning point that pushes you into a whole new level.",
	"You will meet someone who drastically changes the direction of your life.",
	"A long-term desire you’ve kept quiet about will finally start becoming reality.",
	"You will soon move to a new city or environment that transforms you completely.",
	"Your next big risk will turn out to be one of the smartest moves of your life.",
	"You will cut ties with someone who has been quietly holding you back.",
	"A moment of sudden clarity will help you solve a problem you've struggled with for years.",
	"Someone important from your past will re-enter your life with a completely new role.",
	"You will soon begin a project that influences more people than you expect.",
	"A future relationship of yours will change your views on love entirely.",
	"You will experience a turning point that completely shifts your priorities.",
	"Your talent will finally get noticed by someone influential.",
	"Within the next year, you will experience a glow-up — emotionally, mentally, or physically.",
	"You will soon break a cycle that has repeated in your life for far too long.",
	"A major challenge you face will end more easily than you expect.",
	"You will unexpectedly become the 'main character' of a huge event in your social circle.",
	"Something you thought was impossible for you will become achievable.",
	"You will discover a hidden strength that changes your confidence forever.",
	"A big chapter of your life will close soon — but what opens next will be much better.",
	"You will reach a point where you finally feel aligned with your future.",
	"You will surprise everyone — including yourself — with your next big move.",
	"Your future self will look back at this year as the start of everything.",
	"You will end up in a place you once wished for but didn’t think you’d reach.",
	"Your life will soon shift in a way that makes everything make sense.",
	"A dream you buried will find its way back into your life unexpectedly.",
	"You will soon make a connection that elevates your life dramatically.",
	"Your biggest 'what if' will finally get an answer.",
	"You will experience a turning point that feels like fate stepping in.",
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
	"The universe is rearranging itself to give you what you asked for — even what you forgot you wanted.",
	"Your silence has been louder than your words; soon, someone will understand exactly what you meant.",
	"A powerful shift is coming. You will not recognize your life a few months from now — in the best way.",
	"Your energy is being protected. What left your life was removed for your elevation.",
	"A new chapter begins when you finally release what you’ve been holding onto out of fear.",
	"Your patience is about to pay off in a way that feels almost unreal.",
	"Someone you never expected will play a major role in your destiny.",
	"The path you’re on is leading you to something you thought you’d never have.",
	"Your next breakthrough will come from a direction you’re not even looking at.",
	"Your intuition has been right about someone — the truth will reveal itself soon.",
	"You are entering a period where everything you touch begins to align.",
	"A blessing you thought was delayed is actually arriving right on time.",
	"You will soon realize why the universe said ‘not yet’ instead of ‘no.’",
	"Your story is about to take a turn that makes every struggle worth it.",
	"Someone is being guided toward you — your paths are meant to cross.",
	"You are preparing for a version of your life that once felt too big to imagine.",
	"The next person who enters your world will teach you something life-changing.",
	"Your destiny is shifting; something you lost will return in a new form.",
	"You will soon receive clarity that feels like a weight lifting off your chest.",
	"Karma is moving in your favor — a long-awaited balance is finally coming.",
	"You are stepping into the era where your effort meets your fate.",
	"The universe is closing old doors because your new beginning is overdue.",
	"Your glow will intimidate some, but it will inspire many more.",
	"A truth you feared will set you free instead of breaking you.",
	"Your heart is preparing for something it has been waiting for a long time.",
	"A powerful realization is coming — one that changes everything.",
	"The person you will become is someone your younger self prayed for.",
	"You’re being guided toward a life that matches your worth.",
	"You are entering a phase where you finally feel chosen — by people, by opportunities, by life itself.",
	"Your future self is already grateful for a decision you’re about to make.",
}

