package react

import (
	tele "github.com/jxo-me/gfbot"
)

type Reaction = tele.Reaction

func React(r ...Reaction) tele.ReactionOptions {
	return tele.ReactionOptions{Reactions: r}
}

// Currently available emojis.
var (
	ThumbUp                   = Reaction{Emoji: "👍"}
	ThumbDown                 = Reaction{Emoji: "👎"}
	Heart                     = Reaction{Emoji: "❤"}
	Fire                      = Reaction{Emoji: "🔥"}
	HeartEyes                 = Reaction{Emoji: "😍"}
	ClappingHands             = Reaction{Emoji: "👏"}
	GrinningFace              = Reaction{Emoji: "😁"}
	ThinkingFace              = Reaction{Emoji: "🤔"}
	ExplodingHead             = Reaction{Emoji: "🤯"}
	ScreamingFace             = Reaction{Emoji: "😱"}
	SwearingFace              = Reaction{Emoji: "🤬"}
	CryingFace                = Reaction{Emoji: "😢"}
	PartyPopper               = Reaction{Emoji: "🎉"}
	StarStruck                = Reaction{Emoji: "🤩"}
	VomitingFace              = Reaction{Emoji: "🤮"}
	PileOfPoo                 = Reaction{Emoji: "💩"}
	PrayingHands              = Reaction{Emoji: "🙏"}
	OkHand                    = Reaction{Emoji: "👌"}
	DoveOfPeace               = Reaction{Emoji: "🕊"}
	ClownFace                 = Reaction{Emoji: "🤡"}
	YawningFace               = Reaction{Emoji: "🥱"}
	WoozyFace                 = Reaction{Emoji: "🥴"}
	Whale                     = Reaction{Emoji: "🐳"}
	HeartOnFire               = Reaction{Emoji: "❤‍🔥"}
	MoonFace                  = Reaction{Emoji: "🌚"}
	HotDog                    = Reaction{Emoji: "🌭"}
	HundredPoints             = Reaction{Emoji: "💯"}
	RollingOnTheFloorLaughing = Reaction{Emoji: "🤣"}
	Lightning                 = Reaction{Emoji: "⚡"}
	Banana                    = Reaction{Emoji: "🍌"}
	Trophy                    = Reaction{Emoji: "🏆"}
	BrokenHeart               = Reaction{Emoji: "💔"}
	FaceWithRaisedEyebrow     = Reaction{Emoji: "🤨"}
	NeutralFace               = Reaction{Emoji: "😐"}
	Strawberry                = Reaction{Emoji: "🍓"}
	Champagne                 = Reaction{Emoji: "🍾"}
	KissMark                  = Reaction{Emoji: "💋"}
	MiddleFinger              = Reaction{Emoji: "🖕"}
	EvilFace                  = Reaction{Emoji: "😈"}
	SleepingFace              = Reaction{Emoji: "😴"}
	LoudlyCryingFace          = Reaction{Emoji: "😭"}
	NerdFace                  = Reaction{Emoji: "🤓"}
	Ghost                     = Reaction{Emoji: "👻"}
	Engineer                  = Reaction{Emoji: "👨‍💻"}
	Eyes                      = Reaction{Emoji: "👀"}
	JackOLantern              = Reaction{Emoji: "🎃"}
	NoMonkey                  = Reaction{Emoji: "🙈"}
	SmilingFaceWithHalo       = Reaction{Emoji: "😇"}
	FearfulFace               = Reaction{Emoji: "😨"}
	Handshake                 = Reaction{Emoji: "🤝"}
	WritingHand               = Reaction{Emoji: "✍"}
	HuggingFace               = Reaction{Emoji: "🤗"}
	Brain                     = Reaction{Emoji: "🫡"}
	SantaClaus                = Reaction{Emoji: "🎅"}
	ChristmasTree             = Reaction{Emoji: "🎄"}
	Snowman                   = Reaction{Emoji: "☃"}
	NailPolish                = Reaction{Emoji: "💅"}
	ZanyFace                  = Reaction{Emoji: "🤪"}
	Moai                      = Reaction{Emoji: "🗿"}
	Cool                      = Reaction{Emoji: "🆒"}
	HeartWithArrow            = Reaction{Emoji: "💘"}
	HearMonkey                = Reaction{Emoji: "🙉"}
	Unicorn                   = Reaction{Emoji: "🦄"}
	FaceBlowingKiss           = Reaction{Emoji: "😘"}
	Pill                      = Reaction{Emoji: "💊"}
	SpeaklessMonkey           = Reaction{Emoji: "🙊"}
	Sunglasses                = Reaction{Emoji: "😎"}
	AlienMonster              = Reaction{Emoji: "👾"}
	ManShrugging              = Reaction{Emoji: "🤷‍♂️"}
	PersonShrugging           = Reaction{Emoji: "🤷"}
	WomanShrugging            = Reaction{Emoji: "🤷‍♀️"}
	PoutingFace               = Reaction{Emoji: "😡"}
)
