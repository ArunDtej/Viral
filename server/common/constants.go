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
	"how-many-kids-will-you-have":        "How many kids will you have?",
	"which-country-will-you-live-in":     "Which country will you live in?",
	"your-2025-fortune-reading":          "Your 2025 fortune reading",
}
var FuturePredictions = []string{
    "Your future wife divorces you because your mom still cuts your hair.",
    "You finally get a girlfriend but she breaks up after meeting your aunties for 5 minutes.",
    "Your wedding photos are 90% relatives you’ve never seen since 2009.",
    "Your future kid gets 99/100 on a test and the family group chat calls an emergency meeting.",
    "You become a millionaire but your dad still asks when you’re getting a stable job.",
    "Your mom starts an Instagram to post childhood photos the day you get your first girlfriend.",
    "You bring your girlfriend home and your grandma asks if she can cook instant noodles properly.",
    "Your future son brings home a white girlfriend and suddenly everyone forgets you exist.",
    "You get promoted to manager and your relatives still call you 'the lazy one'.",
    "Your wedding budget gets hijacked because 'we must invite the entire village'.",
    "You save up for a PS5 but end up paying for your cousin’s tuition instead.",
    "Your future daughter gets into Harvard and your dad brags that it’s because of his genes, not yours.",
    "Your mom finds your secret second phone and thinks you’re selling drugs.",
    "You finally move out at 35 and your mom rents your room on Airbnb the next day.",
    "Your future girlfriend asks to follow your mom on IG and gets blocked for having a bikini pic.",
    "Your tombstone just says 'He was a good boy but never became doctor'.",
    "You die and your mom uses your funeral money to buy iPhone 27 Pro Max for the relatives.",
    "Your future kid gets a B+ and the family WhatsApp changes its name to 'Disappointment GC'.",
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
    "Mom already picked your wedding date using lunar calendar. You’re still single.",
    "Auntie sends girl’s biodata: 162cm, fair skin, can cook, parents own 3 flats in Shanghai.",
    "First date question: “Do you have a flat in the city or still paying mortgage?”",
    "She asks your annual bonus before asking your favorite color.",
    "Mom cries on WeChat voice note: ‘I just want to carry grandson before I die.’",
    "You’re 28 and your mom still calls you ‘the small one’ in front of girls.",
    "Girl likes your photos until she sees you standing next to your 5’7” friend.",
    "Future mother-in-law: ‘Can put my name on the flat title first?’",
    "She says ‘simple wedding’, then sends guest list of 600 people.",
    "Date brings her own Tupperware to pack leftovers at hotpot.",
    "Mom sees your Hinge: ‘Why swipe when I have premium auntie network?’",
    "Girl asks if you’re ‘marriage type’ while the bubble tea is still sweating.",
    "You bring white girlfriend home. Suddenly every relative speaks perfect English.",
    "Grandma: ‘She pretty but can she give birth to a doctor?’",
    "Your red packet to her parents costs more than your PS5 fund.",
    "Future kid gets 99/100. Mom: ‘Where’s the last mark?’",
    "Cousin born in 2002 already has second kid. You get tagged in baby photos daily.",
    "Lunar New Year 2026 opening line: ‘Still no girlfriend or still no hope?’",
    "Mom started a ‘Pray for son to marry’ WhatsApp group. You’re not in it.",
    "Your future wedding WeChat group has 312 members. You’ve met 7.",
}