package lichess

const (
	OAuth2Scopes = "OAuth2.Scopes"
)

// Defines values for ChallengeJsonColor.
const (
	ChallengeJsonColorBlack  ChallengeJsonColor = "black"
	ChallengeJsonColorRandom ChallengeJsonColor = "random"
	ChallengeJsonColorWhite  ChallengeJsonColor = "white"
)

// Defines values for ChallengeJsonDirection.
const (
	In  ChallengeJsonDirection = "in"
	Out ChallengeJsonDirection = "out"
)

// Defines values for ChallengeJsonStatus.
const (
	ChallengeJsonStatusAccepted ChallengeJsonStatus = "accepted"
	ChallengeJsonStatusCanceled ChallengeJsonStatus = "canceled"
	ChallengeJsonStatusCreated  ChallengeJsonStatus = "created"
	ChallengeJsonStatusDeclined ChallengeJsonStatus = "declined"
	ChallengeJsonStatusOffline  ChallengeJsonStatus = "offline"
)

// Defines values for GameJsonAnalysisJudgmentName.
const (
	Blunder    GameJsonAnalysisJudgmentName = "Blunder"
	Inaccuracy GameJsonAnalysisJudgmentName = "Inaccuracy"
	Mistake    GameJsonAnalysisJudgmentName = "Mistake"
)

// Defines values for GameJsonWinner.
const (
	GameJsonWinnerBlack GameJsonWinner = "black"
	GameJsonWinnerWhite GameJsonWinner = "white"
)

// Defines values for GameStatus.
const (
	GameStatusAborted       GameStatus = "aborted"
	GameStatusCheat         GameStatus = "cheat"
	GameStatusCreated       GameStatus = "created"
	GameStatusDraw          GameStatus = "draw"
	GameStatusMate          GameStatus = "mate"
	GameStatusNoStart       GameStatus = "noStart"
	GameStatusOutoftime     GameStatus = "outoftime"
	GameStatusResign        GameStatus = "resign"
	GameStatusStalemate     GameStatus = "stalemate"
	GameStatusStarted       GameStatus = "started"
	GameStatusTimeout       GameStatus = "timeout"
	GameStatusUnknownFinish GameStatus = "unknownFinish"
	GameStatusVariantEnd    GameStatus = "variantEnd"
)

// Defines values for MoveCategory.
const (
	MoveCategoryBlessedLoss MoveCategory = "blessed-loss"
	MoveCategoryCursedWin   MoveCategory = "cursed-win"
	MoveCategoryDraw        MoveCategory = "draw"
	MoveCategoryLoss        MoveCategory = "loss"
	MoveCategoryMaybeLoss   MoveCategory = "maybe-loss"
	MoveCategoryMaybeWin    MoveCategory = "maybe-win"
	MoveCategoryUnknown     MoveCategory = "unknown"
	MoveCategoryWin         MoveCategory = "win"
)

// Defines values for PerfType.
const (
	PerfTypeAntichess      PerfType = "antichess"
	PerfTypeAtomic         PerfType = "atomic"
	PerfTypeBlitz          PerfType = "blitz"
	PerfTypeBullet         PerfType = "bullet"
	PerfTypeChess960       PerfType = "chess960"
	PerfTypeClassical      PerfType = "classical"
	PerfTypeCorrespondence PerfType = "correspondence"
	PerfTypeCrazyhouse     PerfType = "crazyhouse"
	PerfTypeHorde          PerfType = "horde"
	PerfTypeKingOfTheHill  PerfType = "kingOfTheHill"
	PerfTypeRacingKings    PerfType = "racingKings"
	PerfTypeRapid          PerfType = "rapid"
	PerfTypeThreeCheck     PerfType = "threeCheck"
	PerfTypeUltraBullet    PerfType = "ultraBullet"
)

// Defines values for Speed.
const (
	SpeedBlitz          Speed = "blitz"
	SpeedBullet         Speed = "bullet"
	SpeedClassical      Speed = "classical"
	SpeedCorrespondence Speed = "correspondence"
	SpeedRapid          Speed = "rapid"
	SpeedUltraBullet    Speed = "ultraBullet"
)

// Defines values for TablebaseJsonCategory.
const (
	TablebaseJsonCategoryBlessedLoss TablebaseJsonCategory = "blessed-loss"
	TablebaseJsonCategoryCursedWin   TablebaseJsonCategory = "cursed-win"
	TablebaseJsonCategoryDraw        TablebaseJsonCategory = "draw"
	TablebaseJsonCategoryLoss        TablebaseJsonCategory = "loss"
	TablebaseJsonCategoryMaybeLoss   TablebaseJsonCategory = "maybe-loss"
	TablebaseJsonCategoryMaybeWin    TablebaseJsonCategory = "maybe-win"
	TablebaseJsonCategoryUnknown     TablebaseJsonCategory = "unknown"
	TablebaseJsonCategoryWin         TablebaseJsonCategory = "win"
)

// Defines values for Title.
const (
	BOT Title = "BOT"
	CM  Title = "CM"
	FM  Title = "FM"
	GM  Title = "GM"
	IM  Title = "IM"
	LM  Title = "LM"
	NM  Title = "NM"
	WCM Title = "WCM"
	WFM Title = "WFM"
	WGM Title = "WGM"
	WIM Title = "WIM"
	WNM Title = "WNM"
)

// Defines values for UserPreferencesPieceSet.
const (
	Alpha       UserPreferencesPieceSet = "alpha"
	California  UserPreferencesPieceSet = "california"
	Cardinal    UserPreferencesPieceSet = "cardinal"
	Cburnett    UserPreferencesPieceSet = "cburnett"
	Chess7      UserPreferencesPieceSet = "chess7"
	Chessnut    UserPreferencesPieceSet = "chessnut"
	Companion   UserPreferencesPieceSet = "companion"
	Dubrovny    UserPreferencesPieceSet = "dubrovny"
	Fantasy     UserPreferencesPieceSet = "fantasy"
	Fresca      UserPreferencesPieceSet = "fresca"
	Gioco       UserPreferencesPieceSet = "gioco"
	Governor    UserPreferencesPieceSet = "governor"
	Icpieces    UserPreferencesPieceSet = "icpieces"
	Kosal       UserPreferencesPieceSet = "kosal"
	Leipzig     UserPreferencesPieceSet = "leipzig"
	Letter      UserPreferencesPieceSet = "letter"
	Maestro     UserPreferencesPieceSet = "maestro"
	Merida      UserPreferencesPieceSet = "merida"
	Pirouetti   UserPreferencesPieceSet = "pirouetti"
	Pixel       UserPreferencesPieceSet = "pixel"
	Reillycraig UserPreferencesPieceSet = "reillycraig"
	Riohacha    UserPreferencesPieceSet = "riohacha"
	Shapes      UserPreferencesPieceSet = "shapes"
	Spatial     UserPreferencesPieceSet = "spatial"
	Staunty     UserPreferencesPieceSet = "staunty"
	Tatiana     UserPreferencesPieceSet = "tatiana"
)

// Defines values for UserPreferencesPieceSet3d.
const (
	UserPreferencesPieceSet3dBasic        UserPreferencesPieceSet3d = "Basic"
	UserPreferencesPieceSet3dCubesAndPi   UserPreferencesPieceSet3d = "CubesAndPi"
	UserPreferencesPieceSet3dExperimental UserPreferencesPieceSet3d = "Experimental"
	UserPreferencesPieceSet3dGlass        UserPreferencesPieceSet3d = "Glass"
	UserPreferencesPieceSet3dMetal        UserPreferencesPieceSet3d = "Metal"
	UserPreferencesPieceSet3dModernJade   UserPreferencesPieceSet3d = "ModernJade"
	UserPreferencesPieceSet3dModernWood   UserPreferencesPieceSet3d = "ModernWood"
	UserPreferencesPieceSet3dRedVBlue     UserPreferencesPieceSet3d = "RedVBlue"
	UserPreferencesPieceSet3dStaunton     UserPreferencesPieceSet3d = "Staunton"
	UserPreferencesPieceSet3dTrimmed      UserPreferencesPieceSet3d = "Trimmed"
	UserPreferencesPieceSet3dWood         UserPreferencesPieceSet3d = "Wood"
)

// Defines values for UserPreferencesSoundSet.
const (
	UserPreferencesSoundSetFuturistic UserPreferencesSoundSet = "futuristic"
	UserPreferencesSoundSetMusic      UserPreferencesSoundSet = "music"
	UserPreferencesSoundSetNes        UserPreferencesSoundSet = "nes"
	UserPreferencesSoundSetPiano      UserPreferencesSoundSet = "piano"
	UserPreferencesSoundSetRobot      UserPreferencesSoundSet = "robot"
	UserPreferencesSoundSetSfx        UserPreferencesSoundSet = "sfx"
	UserPreferencesSoundSetSilent     UserPreferencesSoundSet = "silent"
	UserPreferencesSoundSetSpeech     UserPreferencesSoundSet = "speech"
	UserPreferencesSoundSetStandard   UserPreferencesSoundSet = "standard"
)

// Defines values for UserPreferencesTheme.
const (
	UserPreferencesThemeBlue         UserPreferencesTheme = "blue"
	UserPreferencesThemeBlue2        UserPreferencesTheme = "blue2"
	UserPreferencesThemeBlue3        UserPreferencesTheme = "blue3"
	UserPreferencesThemeBlueMarble   UserPreferencesTheme = "blue-marble"
	UserPreferencesThemeBrown        UserPreferencesTheme = "brown"
	UserPreferencesThemeCanvas       UserPreferencesTheme = "canvas"
	UserPreferencesThemeGreen        UserPreferencesTheme = "green"
	UserPreferencesThemeGreenPlastic UserPreferencesTheme = "green-plastic"
	UserPreferencesThemeGrey         UserPreferencesTheme = "grey"
	UserPreferencesThemeIc           UserPreferencesTheme = "ic"
	UserPreferencesThemeLeather      UserPreferencesTheme = "leather"
	UserPreferencesThemeMaple        UserPreferencesTheme = "maple"
	UserPreferencesThemeMaple2       UserPreferencesTheme = "maple2"
	UserPreferencesThemeMarble       UserPreferencesTheme = "marble"
	UserPreferencesThemeMetal        UserPreferencesTheme = "metal"
	UserPreferencesThemeNewspaper    UserPreferencesTheme = "newspaper"
	UserPreferencesThemeOlive        UserPreferencesTheme = "olive"
	UserPreferencesThemePink         UserPreferencesTheme = "pink"
	UserPreferencesThemePurple       UserPreferencesTheme = "purple"
	UserPreferencesThemePurpleDiag   UserPreferencesTheme = "purple-diag"
	UserPreferencesThemeWood         UserPreferencesTheme = "wood"
	UserPreferencesThemeWood2        UserPreferencesTheme = "wood2"
	UserPreferencesThemeWood3        UserPreferencesTheme = "wood3"
	UserPreferencesThemeWood4        UserPreferencesTheme = "wood4"
)

// Defines values for UserPreferencesTheme3d.
const (
	BlackWhiteAluminium UserPreferencesTheme3d = "Black-White-Aluminium"
	BrushedAluminium    UserPreferencesTheme3d = "Brushed-Aluminium"
	ChinaBlue           UserPreferencesTheme3d = "China-Blue"
	ChinaGreen          UserPreferencesTheme3d = "China-Green"
	ChinaGrey           UserPreferencesTheme3d = "China-Grey"
	ChinaScarlet        UserPreferencesTheme3d = "China-Scarlet"
	ClassicBlue         UserPreferencesTheme3d = "Classic-Blue"
	GoldSilver          UserPreferencesTheme3d = "Gold-Silver"
	Jade                UserPreferencesTheme3d = "Jade"
	LightWood           UserPreferencesTheme3d = "Light-Wood"
	Marble              UserPreferencesTheme3d = "Marble"
	PowerCoated         UserPreferencesTheme3d = "Power-Coated"
	Rosewood            UserPreferencesTheme3d = "Rosewood"
	Wax                 UserPreferencesTheme3d = "Wax"
	Woodi               UserPreferencesTheme3d = "Woodi"
)

// Defines values for VariantKey.
const (
	VariantKeyAntichess     VariantKey = "antichess"
	VariantKeyAtomic        VariantKey = "atomic"
	VariantKeyChess960      VariantKey = "chess960"
	VariantKeyCrazyhouse    VariantKey = "crazyhouse"
	VariantKeyHorde         VariantKey = "horde"
	VariantKeyKingOfTheHill VariantKey = "kingOfTheHill"
	VariantKeyRacingKings   VariantKey = "racingKings"
	VariantKeyStandard      VariantKey = "standard"
	VariantKeyThreeCheck    VariantKey = "threeCheck"
)
