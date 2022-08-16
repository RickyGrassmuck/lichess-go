package lichess

import (
	"context"
	"io"
	"net/http"
	"strings"
)

// ApiStreamEvent defines model for ApiStreamEvent.
type ApiStreamEvent struct {
	Type *string `json:"type,omitempty"`
}

// ArenaTournament defines model for ArenaTournament.
type ArenaTournament = interface{}

// ArenaTournaments defines model for ArenaTournaments.
type ArenaTournaments struct {
	Created  *[]ArenaTournament `json:"created,omitempty"`
	Finished *[]ArenaTournament `json:"finished,omitempty"`
	Started  *[]ArenaTournament `json:"started,omitempty"`
}

// BroadcastRound defines model for BroadcastRound.
type BroadcastRound = interface{}

// BroadcastTour defines model for BroadcastTour.
type BroadcastTour = interface{}

// BulkPairing defines model for BulkPairing.
type BulkPairing = interface{}

// ChallengeJson defines model for ChallengeJson.
type ChallengeJson struct {
	Challenger    ChallengeUser           `json:"challenger"`
	Color         ChallengeJsonColor      `json:"color"`
	DeclineReason *string                 `json:"declineReason,omitempty"`
	DestUser      ChallengeUser           `json:"destUser"`
	Direction     *ChallengeJsonDirection `json:"direction,omitempty"`
	Id            string                  `json:"id"`
	InitialFen    *string                 `json:"initialFen,omitempty"`
	Perf          struct {
		Icon *string `json:"icon,omitempty"`
		Name *string `json:"name,omitempty"`
	} `json:"perf"`
	Rated       bool                `json:"rated"`
	Speed       Speed               `json:"speed"`
	Status      ChallengeJsonStatus `json:"status"`
	TimeControl TimeControl         `json:"timeControl"`
	Url         string              `json:"url"`
	Variant     Variant             `json:"variant"`
}

// ChallengeJsonColor defines model for ChallengeJson.Color.
type ChallengeJsonColor string

// ChallengeJsonDirection defines model for ChallengeJson.Direction.
type ChallengeJsonDirection string

// ChallengeJsonStatus defines model for ChallengeJson.Status.
type ChallengeJsonStatus string

// ChallengeOpenJson defines model for ChallengeOpenJson.
type ChallengeOpenJson = interface{}

// ChallengeUser defines model for ChallengeUser.
type ChallengeUser struct {
	Id     *string `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
	Patron *bool   `json:"patron,omitempty"`
	Title  *Title  `json:"title,omitempty"`
}

// Count defines model for Count.
type Count struct {
	Ai       *int `json:"ai,omitempty"`
	All      *int `json:"all,omitempty"`
	Bookmark *int `json:"bookmark,omitempty"`
	Draw     *int `json:"draw,omitempty"`
	DrawH    *int `json:"drawH,omitempty"`
	Import   *int `json:"import,omitempty"`
	Loss     *int `json:"loss,omitempty"`
	LossH    *int `json:"lossH,omitempty"`
	Me       *int `json:"me,omitempty"`
	Playing  *int `json:"playing,omitempty"`
	Rated    *int `json:"rated,omitempty"`
	Win      *int `json:"win,omitempty"`
	WinH     *int `json:"winH,omitempty"`
}

// Crosstable defines model for Crosstable.
type Crosstable = interface{}

// Error defines model for Error.
type Error struct {
	// The cause of the error.
	Error *string `json:"error,omitempty"`
}

// GameChat defines model for GameChat.
type GameChat = interface{}

// GameJson defines model for GameJson.
type GameJson struct {
	Analysis *[]struct {
		// Best move in UCI notation (only if played move was inaccurate)
		Best *string `json:"best,omitempty"`

		// Evaluation in centipawns
		Eval float32 `json:"eval"`

		// Judgment annotation (only if played move was inaccurate)
		Judgment *struct {
			Comment *string                       `json:"comment,omitempty"`
			Name    *GameJsonAnalysisJudgmentName `json:"name,omitempty"`
		} `json:"judgment,omitempty"`

		// Best variation in SAN notation (only if played move was inaccurate)
		Variation *string `json:"variation,omitempty"`
	} `json:"analysis,omitempty"`
	Clock *struct {
		Increment *float32 `json:"increment,omitempty"`
		Initial   *float32 `json:"initial,omitempty"`
		TotalTime *float32 `json:"totalTime,omitempty"`
	} `json:"clock,omitempty"`
	CreatedAt   float32  `json:"createdAt"`
	DaysPerTurn *float32 `json:"daysPerTurn,omitempty"`
	Id          string   `json:"id"`
	InitialFen  *string  `json:"initialFen,omitempty"`
	LastMoveAt  float32  `json:"lastMoveAt"`
	Moves       *string  `json:"moves,omitempty"`
	Opening     *struct {
		Eco  *string  `json:"eco,omitempty"`
		Name *string  `json:"name,omitempty"`
		Ply  *float32 `json:"ply,omitempty"`
	} `json:"opening,omitempty"`
	Perf    string  `json:"perf"`
	Pgn     *string `json:"pgn,omitempty"`
	Players struct {
		Black *GameUser `json:"black,omitempty"`
		White *GameUser `json:"white,omitempty"`
	} `json:"players"`
	Rated bool  `json:"rated"`
	Speed Speed `json:"speed"`

	// Game status code. https://github.com/ornicar/scalachess/blob/0a7d6f2c63b1ca06cd3c958ed3264e738af5c5f6/src/main/scala/Status.scala#L16-L28
	Status     GameStatus      `json:"status"`
	Swiss      *string         `json:"swiss,omitempty"`
	Tournament *string         `json:"tournament,omitempty"`
	Variant    VariantKey      `json:"variant"`
	Winner     *GameJsonWinner `json:"winner,omitempty"`
}

// GameJsonAnalysisJudgmentName defines model for GameJson.Analysis.Judgment.Name.
type GameJsonAnalysisJudgmentName string

// GameJsonWinner defines model for GameJson.Winner.
type GameJsonWinner string

// GamePgn defines model for GamePgn.
type GamePgn = interface{}

// Game status code. https://github.com/ornicar/scalachess/blob/0a7d6f2c63b1ca06cd3c958ed3264e738af5c5f6/src/main/scala/Status.scala#L16-L28
type GameStatus string

// GameStreamEvent defines model for GameStreamEvent.
type GameStreamEvent struct {
	Type *string `json:"type,omitempty"`
}

// GameUser defines model for GameUser.
type GameUser struct {
	AiLevel  *float32 `json:"aiLevel,omitempty"`
	Analysis *struct {
		Acpl       float32 `json:"acpl"`
		Blunder    float32 `json:"blunder"`
		Inaccuracy float32 `json:"inaccuracy"`
		Mistake    float32 `json:"mistake"`
	} `json:"analysis,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Provisional *bool      `json:"provisional,omitempty"`
	Rating      *float32   `json:"rating,omitempty"`
	RatingDiff  *float32   `json:"ratingDiff,omitempty"`
	Team        *string    `json:"team,omitempty"`
	User        *LightUser `json:"user,omitempty"`
}

// Leaderboard defines model for Leaderboard.
type Leaderboard = interface{}

// LightUser defines model for LightUser.
type LightUser struct {
	Id     *string `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
	Patron *bool   `json:"patron,omitempty"`
	Title  *Title  `json:"title,omitempty"`
}

// MasterGamePgn defines model for MasterGamePgn.
type MasterGamePgn = interface{}

// Move defines model for Move.
type Move struct {
	Category  *MoveCategory `json:"category,omitempty"`
	Checkmate *bool         `json:"checkmate,omitempty"`

	// Distance to mate (only for positions with not more than 5 pieces)
	Dtm *int `json:"dtm,omitempty"`

	// DTZ50'' with rounding or null if unknown
	Dtz                  *int    `json:"dtz,omitempty"`
	InsufficientMaterial *bool   `json:"insufficient_material,omitempty"`
	San                  *string `json:"san,omitempty"`
	Stalemate            *bool   `json:"stalemate,omitempty"`
	Uci                  *string `json:"uci,omitempty"`
	VariantLoss          *bool   `json:"variant_loss,omitempty"`
	VariantWin           *bool   `json:"variant_win,omitempty"`
	Zeroing              *bool   `json:"zeroing,omitempty"`
}

// MoveCategory defines model for Move.Category.
type MoveCategory string

// MoveStream defines model for MoveStream.
type MoveStream = interface{}

// NotFound defines model for NotFound.
type NotFound struct {
	Error *string `json:"error,omitempty"`
}

// Ok defines model for Ok.
type Ok struct {
	Ok *bool `json:"ok,omitempty"`
}

// OpeningExplorerJson defines model for OpeningExplorerJson.
type OpeningExplorerJson = interface{}

// OpeningExplorerPlayerJson defines model for OpeningExplorerPlayerJson.
type OpeningExplorerPlayerJson = interface{}

// Perf defines model for Perf.
type Perf struct {
	Games  *int  `json:"games,omitempty"`
	Prog   *int  `json:"prog,omitempty"`
	Prov   *bool `json:"prov,omitempty"`
	Rating *int  `json:"rating,omitempty"`
	Rd     *int  `json:"rd,omitempty"`
}

// PerfStat defines model for PerfStat.
type PerfStat = interface{}

// PerfType defines model for PerfType.
type PerfType string

// Perfs defines model for Perfs.
type Perfs struct {
	Atomic         *Perf      `json:"atomic,omitempty"`
	Blitz          *Perf      `json:"blitz,omitempty"`
	Bullet         *Perf      `json:"bullet,omitempty"`
	Chess960       *Perf      `json:"chess960,omitempty"`
	Classical      *Perf      `json:"classical,omitempty"`
	Correspondence *Perf      `json:"correspondence,omitempty"`
	Horde          *Perf      `json:"horde,omitempty"`
	KingOfTheHill  *Perf      `json:"kingOfTheHill,omitempty"`
	Puzzle         *Perf      `json:"puzzle,omitempty"`
	RacingKings    *Perf      `json:"racingKings,omitempty"`
	Rapid          *Perf      `json:"rapid,omitempty"`
	Storm          *StormPerf `json:"storm,omitempty"`
	UltraBullet    *Perf      `json:"ultraBullet,omitempty"`
}

// PlayTime defines model for PlayTime.
type PlayTime struct {
	Total *int `json:"total,omitempty"`
	Tv    *int `json:"tv,omitempty"`
}

// Profile defines model for Profile.
type Profile struct {
	Bio        *string `json:"bio,omitempty"`
	Country    *string `json:"country,omitempty"`
	EcfRating  *int    `json:"ecfRating,omitempty"`
	FideRating *int    `json:"fideRating,omitempty"`
	FirstName  *string `json:"firstName,omitempty"`
	LastName   *string `json:"lastName,omitempty"`
	Links      *string `json:"links,omitempty"`
	Location   *string `json:"location,omitempty"`
	UscfRating *int    `json:"uscfRating,omitempty"`
}

// PuzzleDashboardJson defines model for PuzzleDashboardJson.
type PuzzleDashboardJson = interface{}

// PuzzleJson defines model for PuzzleJson.
type PuzzleJson = interface{}

// PuzzleRoundJson defines model for PuzzleRoundJson.
type PuzzleRoundJson struct {
	Date         *float32 `json:"date,omitempty"`
	Id           *string  `json:"id,omitempty"`
	PuzzleRating *float32 `json:"puzzleRating,omitempty"`
	Win          *bool    `json:"win,omitempty"`
}

// RatingHistory defines model for RatingHistory.
type RatingHistory = interface{}

// Simul defines model for Simul.
type Simul = interface{}

// Speed defines model for Speed.
type Speed string

// StormDashboardJson defines model for StormDashboardJson.
type StormDashboardJson = interface{}

// StormPerf defines model for StormPerf.
type StormPerf struct {
	Runs  *int `json:"runs,omitempty"`
	Score *int `json:"score,omitempty"`
}

// StudyPgn defines model for StudyPgn.
type StudyPgn = interface{}

// SwissTournament defines model for SwissTournament.
type SwissTournament = interface{}

// SwissUnauthorisedEdit defines model for SwissUnauthorisedEdit.
type SwissUnauthorisedEdit struct {
	Error *string `json:"error,omitempty"`
}

// TablebaseJson defines model for TablebaseJson.
type TablebaseJson struct {
	// `cursed-win` and `blessed-loss` means the 50-move rule prevents
	// the decisive result.
	//
	// `maybe-win` and `maybe-loss` means exact result is unknown due to
	// [DTZ rounding](https://syzygy-tables.info/metrics#dtz), i.e., the
	// win or loss could also be prevented by the 50-move rule if
	// the user has deviated from the tablebase recommendation since the
	// last pawn move or capture.
	Category  *TablebaseJsonCategory `json:"category,omitempty"`
	Checkmate *bool                  `json:"checkmate,omitempty"`

	// Distance to mate (only for positions with not more than 5 pieces)
	Dtm *int `json:"dtm,omitempty"`

	// [DTZ50'' with rounding](https://syzygy-tables.info/metrics#dtz) or null if unknown
	Dtz                  *int  `json:"dtz,omitempty"`
	InsufficientMaterial *bool `json:"insufficient_material,omitempty"`

	// Information about legal moves, best first
	Moves     *[]Move `json:"moves,omitempty"`
	Stalemate *bool   `json:"stalemate,omitempty"`

	// Only in chess variants
	VariantLoss *bool `json:"variant_loss,omitempty"`

	// Only in chess variants
	VariantWin *bool `json:"variant_win,omitempty"`
}

// `cursed-win` and `blessed-loss` means the 50-move rule prevents
// the decisive result.
//
// `maybe-win` and `maybe-loss` means exact result is unknown due to
// [DTZ rounding](https://syzygy-tables.info/metrics#dtz), i.e., the
// win or loss could also be prevented by the 50-move rule if
// the user has deviated from the tablebase recommendation since the
// last pawn move or capture.
type TablebaseJsonCategory string

// Team defines model for Team.
type Team struct {
	Description *string      `json:"description,omitempty"`
	Id          *string      `json:"id,omitempty"`
	Leader      *LightUser   `json:"leader,omitempty"`
	Leaders     *[]LightUser `json:"leaders,omitempty"`
	Location    *string      `json:"location,omitempty"`
	Name        *string      `json:"name,omitempty"`
	NbMembers   *int         `json:"nbMembers,omitempty"`
	Open        *bool        `json:"open,omitempty"`
}

// TimeControl defines model for TimeControl.
type TimeControl struct {
	Type *string `json:"type,omitempty"`
}

// Title defines model for Title.
type Title string

// Top10s defines model for Top10s.
type Top10s = interface{}

// User defines model for User.
type User struct {
	CreatedAt    *int      `json:"createdAt,omitempty"`
	Disabled     *bool     `json:"disabled,omitempty"`
	Id           *string   `json:"id,omitempty"`
	Online       *bool     `json:"online,omitempty"`
	Patron       *bool     `json:"patron,omitempty"`
	Perfs        *Perfs    `json:"perfs,omitempty"`
	PlayTime     *PlayTime `json:"playTime,omitempty"`
	Profile      *Profile  `json:"profile,omitempty"`
	SeenAt       *int      `json:"seenAt,omitempty"`
	Title        *Title    `json:"title,omitempty"`
	TosViolation *bool     `json:"tosViolation,omitempty"`
	Username     *string   `json:"username,omitempty"`
	Verified     *bool     `json:"verified,omitempty"`
}

// UserExtended defines model for UserExtended.
type UserExtended struct {
	Blocking       *bool     `json:"blocking,omitempty"`
	CompletionRate *int      `json:"completionRate,omitempty"`
	Count          *Count    `json:"count,omitempty"`
	CreatedAt      *int      `json:"createdAt,omitempty"`
	Disabled       *bool     `json:"disabled,omitempty"`
	Followable     *bool     `json:"followable,omitempty"`
	Following      *bool     `json:"following,omitempty"`
	FollowsYou     *bool     `json:"followsYou,omitempty"`
	Id             *string   `json:"id,omitempty"`
	Online         *bool     `json:"online,omitempty"`
	Patron         *bool     `json:"patron,omitempty"`
	Perfs          *Perfs    `json:"perfs,omitempty"`
	PlayTime       *PlayTime `json:"playTime,omitempty"`
	Playing        *string   `json:"playing,omitempty"`
	Profile        *Profile  `json:"profile,omitempty"`
	SeenAt         *int      `json:"seenAt,omitempty"`
	Streaming      *bool     `json:"streaming,omitempty"`
	Title          *Title    `json:"title,omitempty"`
	TosViolation   *bool     `json:"tosViolation,omitempty"`
	Url            *string   `json:"url,omitempty"`
	Username       *string   `json:"username,omitempty"`
	Verified       *bool     `json:"verified,omitempty"`
}

// UserPreferences defines model for UserPreferences.
type UserPreferences struct {
	Animation     *int                       `json:"animation,omitempty"`
	AutoQueen     *int                       `json:"autoQueen,omitempty"`
	AutoThreefold *int                       `json:"autoThreefold,omitempty"`
	BgImg         *string                    `json:"bgImg,omitempty"`
	Blindfold     *int                       `json:"blindfold,omitempty"`
	Captured      *bool                      `json:"captured,omitempty"`
	Challenge     *int                       `json:"challenge,omitempty"`
	ClockBar      *bool                      `json:"clockBar,omitempty"`
	ClockSound    *bool                      `json:"clockSound,omitempty"`
	ClockTenths   *int                       `json:"clockTenths,omitempty"`
	ConfirmResign *int                       `json:"confirmResign,omitempty"`
	CoordColor    *int                       `json:"coordColor,omitempty"`
	Coords        *int                       `json:"coords,omitempty"`
	Dark          *bool                      `json:"dark,omitempty"`
	Destination   *bool                      `json:"destination,omitempty"`
	Follow        *bool                      `json:"follow,omitempty"`
	Highlight     *bool                      `json:"highlight,omitempty"`
	InsightShare  *int                       `json:"insightShare,omitempty"`
	Is3d          *bool                      `json:"is3d,omitempty"`
	KeyboardMove  *int                       `json:"keyboardMove,omitempty"`
	Message       *int                       `json:"message,omitempty"`
	Moretime      *int                       `json:"moretime,omitempty"`
	MoveEvent     *int                       `json:"moveEvent,omitempty"`
	PieceSet      *UserPreferencesPieceSet   `json:"pieceSet,omitempty"`
	PieceSet3d    *UserPreferencesPieceSet3d `json:"pieceSet3d,omitempty"`
	Premove       *bool                      `json:"premove,omitempty"`
	Replay        *int                       `json:"replay,omitempty"`
	RookCastle    *int                       `json:"rookCastle,omitempty"`
	SoundSet      *UserPreferencesSoundSet   `json:"soundSet,omitempty"`
	SubmitMove    *int                       `json:"submitMove,omitempty"`
	Takeback      *int                       `json:"takeback,omitempty"`
	Theme         *UserPreferencesTheme      `json:"theme,omitempty"`
	Theme3d       *UserPreferencesTheme3d    `json:"theme3d,omitempty"`
	Transp        *bool                      `json:"transp,omitempty"`
	Zen           *int                       `json:"zen,omitempty"`
}

// UserPreferencesPieceSet defines model for UserPreferences.PieceSet.
type UserPreferencesPieceSet string

// UserPreferencesPieceSet3d defines model for UserPreferences.PieceSet3d.
type UserPreferencesPieceSet3d string

// UserPreferencesSoundSet defines model for UserPreferences.SoundSet.
type UserPreferencesSoundSet string

// UserPreferencesTheme defines model for UserPreferences.Theme.
type UserPreferencesTheme string

// UserPreferencesTheme3d defines model for UserPreferences.Theme3d.
type UserPreferencesTheme3d string

// Variant defines model for Variant.
type Variant struct {
	Key   *VariantKey `json:"key,omitempty"`
	Name  *string     `json:"name,omitempty"`
	Short *string     `json:"short,omitempty"`
}

// VariantKey defines model for VariantKey.
type VariantKey string

// AccountKidPostParams defines parameters for AccountKidPost.
type AccountKidPostParams struct {
	// Kid mode status
	V bool `form:"v" json:"v"`
}

// ApiAccountPlayingParams defines parameters for ApiAccountPlaying.
type ApiAccountPlayingParams struct {
	// Max number of games to fetch
	Nb *int `form:"nb,omitempty" json:"nb,omitempty"`
}

// BoardGameMoveParams defines parameters for BoardGameMove.
type BoardGameMoveParams struct {
	// Whether to offer (or agree to) a draw
	OfferingDraw *bool `form:"offeringDraw,omitempty" json:"offeringDraw,omitempty"`
}

// BotGameMoveParams defines parameters for BotGameMove.
type BotGameMoveParams struct {
	// Whether to offer (or agree to) a draw
	OfferingDraw *bool `form:"offeringDraw,omitempty" json:"offeringDraw,omitempty"`
}

// ApiBotOnlineParams defines parameters for ApiBotOnline.
type ApiBotOnlineParams struct {
	// How many bot users to fetch
	Nb *int `form:"nb,omitempty" json:"nb,omitempty"`
}

// BroadcastIndexParams defines parameters for BroadcastIndex.
type BroadcastIndexParams struct {
	// Max number of broadcasts to fetch
	Nb *int `form:"nb,omitempty" json:"nb,omitempty"`
}

// ChallengeCancelParams defines parameters for ChallengeCancel.
type ChallengeCancelParams struct {
	// Optional `challenge:write` token of the opponent. If set, the game can be canceled even if both players have moved.
	OpponentToken *string `form:"opponentToken,omitempty" json:"opponentToken,omitempty"`
}

// ChallengeStartClocksParams defines parameters for ChallengeStartClocks.
type ChallengeStartClocksParams struct {
	// OAuth token of a player
	Token1 *string `form:"token1,omitempty" json:"token1,omitempty"`

	// OAuth token of the other player
	Token2 *string `form:"token2,omitempty" json:"token2,omitempty"`
}

// ApiCloudEvalParams defines parameters for ApiCloudEval.
type ApiCloudEvalParams struct {
	// FEN of the position
	Fen string `form:"fen" json:"fen"`

	// Number of variations
	MultiPv *float32 `form:"multiPv,omitempty" json:"multiPv,omitempty"`

	// Variant
	Variant *VariantKey `form:"variant,omitempty" json:"variant,omitempty"`
}

// ApiCrosstableParams defines parameters for ApiCrosstable.
type ApiCrosstableParams struct {
	// Whether to get the current match data, if any
	Matchup *bool `form:"matchup,omitempty" json:"matchup,omitempty"`
}

// ApiGamesUserParams defines parameters for ApiGamesUser.
type ApiGamesUserParams struct {
	// Download games played since this timestamp.
	Since *int `form:"since,omitempty" json:"since,omitempty"`

	// Download games played until this timestamp.
	Until *int `form:"until,omitempty" json:"until,omitempty"`

	// How many games to download. Leave empty to download all games.
	Max *int `form:"max,omitempty" json:"max,omitempty"`

	// [Filter] Only games played against this opponent
	Vs *string `form:"vs,omitempty" json:"vs,omitempty"`

	// [Filter] Only rated (`true`) or casual (`false`) games
	Rated *bool `form:"rated,omitempty" json:"rated,omitempty"`

	// [Filter] Only games in these speeds or variants.
	//
	// Multiple perf types can be specified, separated by a comma.
	//
	// Example: blitz,rapid,classical
	PerfType *interface{} `form:"perfType,omitempty" json:"perfType,omitempty"`

	// [Filter] Only games played as this color.
	Color *ApiGamesUserParamsColor `form:"color,omitempty" json:"color,omitempty"`

	// [Filter] Only games with or without a computer analysis available
	Analysed *bool `form:"analysed,omitempty" json:"analysed,omitempty"`

	// Include the PGN moves.
	Moves *bool `form:"moves,omitempty" json:"moves,omitempty"`

	// Include the full PGN within the JSON response, in a `pgn` field.
	// The response type must be set to `application/x-ndjson` by the request `Accept` header.
	PgnInJson *bool `form:"pgnInJson,omitempty" json:"pgnInJson,omitempty"`

	// Include the PGN tags.
	Tags *bool `form:"tags,omitempty" json:"tags,omitempty"`

	// Include clock comments in the PGN moves, when available.
	//
	// Example: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }`
	Clocks *bool `form:"clocks,omitempty" json:"clocks,omitempty"`

	// Include analysis evaluation comments in the PGN, when available.
	//
	// Example: `12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }`
	Evals *bool `form:"evals,omitempty" json:"evals,omitempty"`

	// Include the opening name.
	//
	// Example: `[Opening "King's Gambit Accepted, King's Knight Gambit"]`
	Opening *bool `form:"opening,omitempty" json:"opening,omitempty"`

	// Include ongoing games. The last 3 moves will be omitted.
	Ongoing *bool `form:"ongoing,omitempty" json:"ongoing,omitempty"`

	// Include finished games. Set to `false` to only get ongoing games.
	Finished *bool `form:"finished,omitempty" json:"finished,omitempty"`

	// URL of a text file containing real names and ratings, to replace Lichess usernames and ratings in the PGN.
	// Example: <https://gist.githubusercontent.com/ornicar/6bfa91eb61a2dcae7bcd14cce1b2a4eb/raw/768b9f6cc8a8471d2555e47ba40fb0095e5fba37/gistfile1.txt>
	Players *string `form:"players,omitempty" json:"players,omitempty"`

	// Sort order of the games.
	Sort *ApiGamesUserParamsSort `form:"sort,omitempty" json:"sort,omitempty"`
}

// ApiGamesUserParamsColor defines parameters for ApiGamesUser.
type ApiGamesUserParamsColor string

// ApiGamesUserParamsSort defines parameters for ApiGamesUser.
type ApiGamesUserParamsSort string

// ApiPuzzleActivityParams defines parameters for ApiPuzzleActivity.
type ApiPuzzleActivityParams struct {
	// How many entries to download. Leave empty to download all activity.
	Max *int `form:"max,omitempty" json:"max,omitempty"`
}

// ApiStormDashboardParams defines parameters for ApiStormDashboard.
type ApiStormDashboardParams struct {
	// How many days of history to return
	Days *int `form:"days,omitempty" json:"days,omitempty"`
}

// GamesByUsersParams defines parameters for GamesByUsers.
type GamesByUsersParams struct {
	// Include the already started games at the beginning of the stream.
	WithCurrentGames *bool `form:"withCurrentGames,omitempty" json:"withCurrentGames,omitempty"`
}

// StudyAllChaptersPgnParams defines parameters for StudyAllChaptersPgn.
type StudyAllChaptersPgnParams struct {
	// Include clock comments in the PGN moves, when available.
	//
	// Example: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }`
	Clocks *bool `form:"clocks,omitempty" json:"clocks,omitempty"`

	// Include analysis and annotator comments in the PGN moves, when available.
	//
	// Example: `12. Bxf6 { [%eval 0.23] } a3 { White is in a pickle. }`
	Comments *bool `form:"comments,omitempty" json:"comments,omitempty"`

	// Include non-mainline moves, when available.
	//
	// Example: `4. d4 Bb4+ (4... Nc6 5. Nf3 Bb4+ 6. Bd2 (6. Nbd2 O-O 7. O-O) 6... Bd6) 5. Nd2`
	Variations *bool `form:"variations,omitempty" json:"variations,omitempty"`
}

// GamesBySwissParams defines parameters for GamesBySwiss.
type GamesBySwissParams struct {
	// Include the PGN moves.
	Moves *bool `form:"moves,omitempty" json:"moves,omitempty"`

	// Include the full PGN within the JSON response, in a `pgn` field.
	PgnInJson *bool `form:"pgnInJson,omitempty" json:"pgnInJson,omitempty"`

	// Include the PGN tags.
	Tags *bool `form:"tags,omitempty" json:"tags,omitempty"`

	// Include clock comments in the PGN moves, when available.
	//
	// Example: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }`
	Clocks *bool `form:"clocks,omitempty" json:"clocks,omitempty"`

	// Include analysis evaluation comments in the PGN, when available.
	//
	// Example: `12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }`
	Evals *bool `form:"evals,omitempty" json:"evals,omitempty"`

	// Include the opening name.
	//
	// Example: `[Opening "King's Gambit Accepted, King's Knight Gambit"]`
	Opening *bool `form:"opening,omitempty" json:"opening,omitempty"`
}

// ResultsBySwissParams defines parameters for ResultsBySwiss.
type ResultsBySwissParams struct {
	// Max number of players to fetch
	Nb *int `form:"nb,omitempty" json:"nb,omitempty"`
}

// TeamAllParams defines parameters for TeamAll.
type TeamAllParams struct {
	Page *float32 `form:"page,omitempty" json:"page,omitempty"`
}

// TeamSearchParams defines parameters for TeamSearch.
type TeamSearchParams struct {
	Text *string  `form:"text,omitempty" json:"text,omitempty"`
	Page *float32 `form:"page,omitempty" json:"page,omitempty"`
}

// ApiTeamArenaParams defines parameters for ApiTeamArena.
type ApiTeamArenaParams struct {
	// How many tournaments to download.
	Max *int `form:"max,omitempty" json:"max,omitempty"`
}

// ApiTeamSwissParams defines parameters for ApiTeamSwiss.
type ApiTeamSwissParams struct {
	// How many tournaments to download.
	Max *int `form:"max,omitempty" json:"max,omitempty"`
}

// TournamentParams defines parameters for Tournament.
type TournamentParams struct {
	// Specify which page of player standings to view.
	Page *float32 `form:"page,omitempty" json:"page,omitempty"`
}

// GamesByTournamentParams defines parameters for GamesByTournament.
type GamesByTournamentParams struct {
	// Only games of a particular player. Leave empty to fetch games of all players.
	Player *string `form:"player,omitempty" json:"player,omitempty"`

	// Include the PGN moves.
	Moves *bool `form:"moves,omitempty" json:"moves,omitempty"`

	// Include the full PGN within the JSON response, in a `pgn` field.
	PgnInJson *bool `form:"pgnInJson,omitempty" json:"pgnInJson,omitempty"`

	// Include the PGN tags.
	Tags *bool `form:"tags,omitempty" json:"tags,omitempty"`

	// Include clock comments in the PGN moves, when available.
	//
	// Example: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }`
	Clocks *bool `form:"clocks,omitempty" json:"clocks,omitempty"`

	// Include analysis evaluation comments in the PGN, when available.
	//
	// Example: `12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }`
	Evals *bool `form:"evals,omitempty" json:"evals,omitempty"`

	// Include the opening name.
	//
	// Example: `[Opening "King's Gambit Accepted, King's Knight Gambit"]`
	Opening *bool `form:"opening,omitempty" json:"opening,omitempty"`
}

// ResultsByTournamentParams defines parameters for ResultsByTournament.
type ResultsByTournamentParams struct {
	// Max number of players to fetch
	Nb *int `form:"nb,omitempty" json:"nb,omitempty"`
}

// TvChannelGamesParams defines parameters for TvChannelGames.
type TvChannelGamesParams struct {
	// Number of games to fetch.
	Nb *float32 `form:"nb,omitempty" json:"nb,omitempty"`

	// Include the PGN moves.
	Moves *bool `form:"moves,omitempty" json:"moves,omitempty"`

	// Include the full PGN within the JSON response, in a `pgn` field.
	PgnInJson *bool `form:"pgnInJson,omitempty" json:"pgnInJson,omitempty"`

	// Include the PGN tags.
	Tags *bool `form:"tags,omitempty" json:"tags,omitempty"`

	// Include clock comments in the PGN moves, when available.
	//
	// Example: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }`
	Clocks *bool `form:"clocks,omitempty" json:"clocks,omitempty"`

	// Include the opening name.
	//
	// Example: `[Opening "King's Gambit Accepted, King's Knight Gambit"]`
	Opening *bool `form:"opening,omitempty" json:"opening,omitempty"`
}

// ApiUserCurrentGameParams defines parameters for ApiUserCurrentGame.
type ApiUserCurrentGameParams struct {
	// Include the PGN moves.
	Moves *bool `form:"moves,omitempty" json:"moves,omitempty"`

	// Include the full PGN within the JSON response, in a `pgn` field.
	PgnInJson *bool `form:"pgnInJson,omitempty" json:"pgnInJson,omitempty"`

	// Include the PGN tags.
	Tags *bool `form:"tags,omitempty" json:"tags,omitempty"`

	// Include clock comments in the PGN moves, when available.
	//
	// Example: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }`
	Clocks *bool `form:"clocks,omitempty" json:"clocks,omitempty"`

	// Include analysis evaluation comments in the PGN, when available.
	//
	// Example: `12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }`
	Evals *bool `form:"evals,omitempty" json:"evals,omitempty"`

	// Include the opening name.
	//
	// Example: `[Opening "King's Gambit Accepted, King's Knight Gambit"]`
	Opening *bool `form:"opening,omitempty" json:"opening,omitempty"`

	// Insert textual annotations in the PGN about the opening, analysis variations, mistakes, and game termination.
	//
	// Example: `5... g4? { (-0.98 → 0.60) Mistake. Best move was h6. } (5... h6 6. d4 Ne7 7. g3 d5 8. exd5 fxg3 9. hxg3 c6 10. dxc6)`
	Literate *bool `form:"literate,omitempty" json:"literate,omitempty"`

	// URL of a text file containing real names and ratings, to replace Lichess usernames and ratings in the PGN.
	// Example: <https://gist.githubusercontent.com/ornicar/6bfa91eb61a2dcae7bcd14cce1b2a4eb/raw/768b9f6cc8a8471d2555e47ba40fb0095e5fba37/gistfile1.txt>
	Players *string `form:"players,omitempty" json:"players,omitempty"`
}

// ApiUserNameTournamentCreatedParams defines parameters for ApiUserNameTournamentCreated.
type ApiUserNameTournamentCreatedParams struct {
	// Include tournaments in the given status: "Created" (10), "Started" (20), "Finished" (30)
	//
	// You can add this parameter more than once to include tournaments in different statuses.
	//
	// Example: `?status=10&status=20`
	Status *ApiUserNameTournamentCreatedParamsStatus `form:"status,omitempty" json:"status,omitempty"`
}

// ApiUserNameTournamentCreatedParamsStatus defines parameters for ApiUserNameTournamentCreated.
type ApiUserNameTournamentCreatedParamsStatus int

// ApiUsersStatusParams defines parameters for ApiUsersStatus.
type ApiUsersStatusParams struct {
	// User IDs separated by commas. Up to 100 IDs.
	Ids string `form:"ids" json:"ids"`
}

// GamePgnParams defines parameters for GamePgn.
type GamePgnParams struct {
	// Include the PGN moves.
	Moves *bool `form:"moves,omitempty" json:"moves,omitempty"`

	// Include the full PGN within the JSON response, in a `pgn` field.
	PgnInJson *bool `form:"pgnInJson,omitempty" json:"pgnInJson,omitempty"`

	// Include the PGN tags.
	Tags *bool `form:"tags,omitempty" json:"tags,omitempty"`

	// Include clock comments in the PGN moves, when available.
	//
	// Example: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }`
	Clocks *bool `form:"clocks,omitempty" json:"clocks,omitempty"`

	// Include analysis evaluation comments in the PGN, when available.
	//
	// Example: `12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }`
	Evals *bool `form:"evals,omitempty" json:"evals,omitempty"`

	// Include the opening name.
	//
	// Example: `[Opening "King's Gambit Accepted, King's Knight Gambit"]`
	Opening *bool `form:"opening,omitempty" json:"opening,omitempty"`

	// Insert textual annotations in the PGN about the opening, analysis variations, mistakes, and game termination.
	//
	// Example: `5... g4? { (-0.98 → 0.60) Mistake. Best move was h6. } (5... h6 6. d4 Ne7 7. g3 d5 8. exd5 fxg3 9. hxg3 c6 10. dxc6)`
	Literate *bool `form:"literate,omitempty" json:"literate,omitempty"`

	// URL of a text file containing real names and ratings, to replace Lichess usernames and ratings in the PGN.
	// Example: <https://gist.githubusercontent.com/ornicar/6bfa91eb61a2dcae7bcd14cce1b2a4eb/raw/768b9f6cc8a8471d2555e47ba40fb0095e5fba37/gistfile1.txt>
	Players *string `form:"players,omitempty" json:"players,omitempty"`
}

// GamesExportIdsParams defines parameters for GamesExportIds.
type GamesExportIdsParams struct {
	// Include the PGN moves.
	Moves *bool `form:"moves,omitempty" json:"moves,omitempty"`

	// Include the full PGN within the JSON response, in a `pgn` field.
	PgnInJson *bool `form:"pgnInJson,omitempty" json:"pgnInJson,omitempty"`

	// Include the PGN tags.
	Tags *bool `form:"tags,omitempty" json:"tags,omitempty"`

	// Include clock comments in the PGN moves, when available.
	//
	// Example: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }`
	Clocks *bool `form:"clocks,omitempty" json:"clocks,omitempty"`

	// Include analysis evaluation comments in the PGN, when available.
	//
	// Example: `12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }`
	Evals *bool `form:"evals,omitempty" json:"evals,omitempty"`

	// Include the opening name.
	//
	// Example: `[Opening "King's Gambit Accepted, King's Knight Gambit"]`
	Opening *bool `form:"opening,omitempty" json:"opening,omitempty"`

	// URL of a text file containing real names and ratings, to replace Lichess usernames and ratings in the PGN.
	// Example: <https://gist.githubusercontent.com/ornicar/6bfa91eb61a2dcae7bcd14cce1b2a4eb/raw/768b9f6cc8a8471d2555e47ba40fb0095e5fba37/gistfile1.txt>
	Players *string `form:"players,omitempty" json:"players,omitempty"`
}

// OpeningExplorerLichessParams defines parameters for OpeningExplorerLichess.
type OpeningExplorerLichessParams struct {
	// Variant
	Variant *VariantKey `form:"variant,omitempty" json:"variant,omitempty"`

	// FEN of the root position
	Fen *string `form:"fen,omitempty" json:"fen,omitempty"`

	// Comma separated sequence of legal moves in UCI notation.
	// Play additional moves starting from `fen`.
	// Required to find an opening name, if `fen` is not an exact match
	// for a named position.
	Play *string `form:"play,omitempty" json:"play,omitempty"`

	// Comma separated list of game speeds to look for
	Speeds *[]Speed `form:"speeds,omitempty" json:"speeds,omitempty"`

	// Comma separated list of rating groups, ranging from their value to the next higher group
	Ratings *[]OpeningExplorerLichessParamsRatings `form:"ratings,omitempty" json:"ratings,omitempty"`

	// Include only games from this month or later
	Since *string `form:"since,omitempty" json:"since,omitempty"`

	// Include only games from this month or earlier
	Until *string `form:"until,omitempty" json:"until,omitempty"`

	// Number of most common moves to display
	Moves *float32 `form:"moves,omitempty" json:"moves,omitempty"`

	// Number of top games to display
	TopGames *float32 `form:"topGames,omitempty" json:"topGames,omitempty"`

	// Number of recent games to display
	RecentGames *float32 `form:"recentGames,omitempty" json:"recentGames,omitempty"`
}

// OpeningExplorerLichessParamsRatings defines parameters for OpeningExplorerLichess.
type OpeningExplorerLichessParamsRatings float32

// OpeningExplorerMasterParams defines parameters for OpeningExplorerMaster.
type OpeningExplorerMasterParams struct {
	// FEN of the root position
	Fen *string `form:"fen,omitempty" json:"fen,omitempty"`

	// Comma separated sequence of legal moves in UCI notation.
	// Play additional moves starting from `fen`.
	// Required to find an opening name, if `fen` is not an exact match
	// for a named position.
	Play *string `form:"play,omitempty" json:"play,omitempty"`

	// Include only games from this year or later
	Since *float32 `form:"since,omitempty" json:"since,omitempty"`

	// Include only games from thi year or earlier
	Until *float32 `form:"until,omitempty" json:"until,omitempty"`

	// Number of most common moves to display
	Moves *float32 `form:"moves,omitempty" json:"moves,omitempty"`

	// Number of top games to display
	TopGames *float32 `form:"topGames,omitempty" json:"topGames,omitempty"`
}

// OauthParams defines parameters for Oauth.
type OauthParams struct {
	// Must be `code`.
	ResponseType string `form:"response_type" json:"response_type"`

	// Arbitrary identifier that uniquely identifies your application.
	ClientId string `form:"client_id" json:"client_id"`

	// The absolute URL that the user should be redirected to with the authorization result.
	RedirectUri string `form:"redirect_uri" json:"redirect_uri"`

	// Must be `S256`.
	CodeChallengeMethod string `form:"code_challenge_method" json:"code_challenge_method"`

	// Compute `BASE64URL(SHA256(code_verifier))`.
	CodeChallenge string `form:"code_challenge" json:"code_challenge"`

	// Space separated list of requested OAuth scopes, if any.
	Scope *string `form:"scope,omitempty" json:"scope,omitempty"`

	// Arbitrary state that will be returned verbatim with the authorization result.
	State *string `form:"state,omitempty" json:"state,omitempty"`
}

// PlayerTopNbPerfTypeParamsPerfType defines parameters for PlayerTopNbPerfType.
type PlayerTopNbPerfTypeParamsPerfType string

// OpeningExplorerPlayerParams defines parameters for OpeningExplorerPlayer.
type OpeningExplorerPlayerParams struct {
	// Variant
	Variant *VariantKey `form:"variant,omitempty" json:"variant,omitempty"`

	// FEN of the root position
	Fen *string `form:"fen,omitempty" json:"fen,omitempty"`

	// Comma separated sequence of legal moves in UCI notation.
	// Play additional moves starting from `fen`.
	// Required to find an opening name, if `fen` is not an exact match
	// for a named position.
	Play *string `form:"play,omitempty" json:"play,omitempty"`

	// Comma separated list of game speeds to look for
	Speeds *[]Speed `form:"speeds,omitempty" json:"speeds,omitempty"`

	// Comma separated list of modes
	Modes *[]OpeningExplorerPlayerParamsModes `form:"modes,omitempty" json:"modes,omitempty"`

	// Include only games from this month or later
	Since *string `form:"since,omitempty" json:"since,omitempty"`

	// Include only games from this month or earlier
	Until *string `form:"until,omitempty" json:"until,omitempty"`

	// Number of most common moves to display
	Moves *float32 `form:"moves,omitempty" json:"moves,omitempty"`

	// Number of recent games to display
	RecentGames *float32 `form:"recentGames,omitempty" json:"recentGames,omitempty"`
}

// OpeningExplorerPlayerParamsModes defines parameters for OpeningExplorerPlayer.
type OpeningExplorerPlayerParamsModes string

// TablebaseStandardParams defines parameters for TablebaseStandard.
type TablebaseStandardParams struct {
	// FEN of the position. Underscores allowed.
	Fen string `form:"fen" json:"fen"`
}

// StudyExportAllPgnParams defines parameters for StudyExportAllPgn.
type StudyExportAllPgnParams struct {
	// Include clock comments in the PGN moves, when available.
	//
	// Example: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }`
	Clocks *bool `form:"clocks,omitempty" json:"clocks,omitempty"`

	// Include analysis and annotator comments in the PGN moves, when available.
	//
	// Example: `12. Bxf6 { [%eval 0.23] } a3 { White is in a pickle. }`
	Comments *bool `form:"comments,omitempty" json:"comments,omitempty"`

	// Include non-mainline moves, when available.
	//
	// Example: `4. d4 Bb4+ (4... Nc6 5. Nf3 Bb4+ 6. Bd2 (6. Nbd2 O-O 7. O-O) 6... Bd6) 5. Nd2`
	Variations *bool `form:"variations,omitempty" json:"variations,omitempty"`
}

// StudyChapterPgnParams defines parameters for StudyChapterPgn.
type StudyChapterPgnParams struct {
	// Include clock comments in the PGN moves, when available.
	//
	// Example: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }`
	Clocks *bool `form:"clocks,omitempty" json:"clocks,omitempty"`

	// Include analysis and annotator comments in the PGN moves, when available.
	//
	// Example: `12. Bxf6 { [%eval 0.23] } a3 { White is in a pickle. }`
	Comments *bool `form:"comments,omitempty" json:"comments,omitempty"`

	// Include non-mainline moves, when available.
	//
	// Example: `4. d4 Bb4+ (4... Nc6 5. Nf3 Bb4+ 6. Bd2 (6. Nbd2 O-O 7. O-O) 6... Bd6) 5. Nd2`
	Variations *bool `form:"variations,omitempty" json:"variations,omitempty"`
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// AntichessAtomic request
	AntichessAtomic(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AccountMe request
	AccountMe(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AccountEmail request
	AccountEmail(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AccountKid request
	AccountKid(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AccountKidPost request
	AccountKidPost(ctx context.Context, params *AccountKidPostParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiAccountPlaying request
	ApiAccountPlaying(ctx context.Context, params *ApiAccountPlayingParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Account request
	Account(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BoardGameStream request
	BoardGameStream(ctx context.Context, gameId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BoardGameAbort request
	BoardGameAbort(ctx context.Context, gameId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BoardGameChatGet request
	BoardGameChatGet(ctx context.Context, gameId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BoardGameChatPost request with any body
	BoardGameChatPostWithBody(ctx context.Context, gameId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BoardGameClaimVictory request
	BoardGameClaimVictory(ctx context.Context, gameId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BoardGameDraw request
	BoardGameDraw(ctx context.Context, gameId string, accept bool, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BoardGameMove request
	BoardGameMove(ctx context.Context, gameId string, move string, params *BoardGameMoveParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BoardGameResign request
	BoardGameResign(ctx context.Context, gameId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BoardGameTakeback request
	BoardGameTakeback(ctx context.Context, gameId string, accept bool, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiBoardSeek request with any body
	ApiBoardSeekWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BotAccountUpgrade request
	BotAccountUpgrade(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BotGameStream request
	BotGameStream(ctx context.Context, gameId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BotGameAbort request
	BotGameAbort(ctx context.Context, gameId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BotGameChat request with any body
	BotGameChatWithBody(ctx context.Context, gameId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BotGameMove request
	BotGameMove(ctx context.Context, gameId string, move string, params *BotGameMoveParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BotGameResign request
	BotGameResign(ctx context.Context, gameId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiBotOnline request
	ApiBotOnline(ctx context.Context, params *ApiBotOnlineParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BroadcastIndex request
	BroadcastIndex(ctx context.Context, params *BroadcastIndexParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BroadcastRoundPgn request
	BroadcastRoundPgn(ctx context.Context, broadcastRoundId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BroadcastAllRoundsPgn request
	BroadcastAllRoundsPgn(ctx context.Context, broadcastTournamentId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BulkPairingGet request
	BulkPairingGet(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BulkPairingCreate request with any body
	BulkPairingCreateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BulkPairingDelete request
	BulkPairingDelete(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BulkPairingStartClocks request
	BulkPairingStartClocks(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ChallengeList request
	ChallengeList(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ChallengeAi request with any body
	ChallengeAiWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ChallengeOpen request with any body
	ChallengeOpenWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ChallengeAccept request
	ChallengeAccept(ctx context.Context, challengeId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ChallengeCancel request
	ChallengeCancel(ctx context.Context, challengeId string, params *ChallengeCancelParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ChallengeDecline request with any body
	ChallengeDeclineWithBody(ctx context.Context, challengeId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ChallengeStartClocks request
	ChallengeStartClocks(ctx context.Context, gameId string, params *ChallengeStartClocksParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ChallengeCreate request with any body
	ChallengeCreateWithBody(ctx context.Context, username string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiCloudEval request
	ApiCloudEval(ctx context.Context, params *ApiCloudEvalParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiCrosstable request
	ApiCrosstable(ctx context.Context, user1 string, user2 string, params *ApiCrosstableParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiGamesUser request
	ApiGamesUser(ctx context.Context, username string, params *ApiGamesUserParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GameImport request with any body
	GameImportWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiPuzzleActivity request
	ApiPuzzleActivity(ctx context.Context, params *ApiPuzzleActivityParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiPuzzleDaily request
	ApiPuzzleDaily(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiPuzzleDashboard request
	ApiPuzzleDashboard(ctx context.Context, days int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FollowUser request
	FollowUser(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiUserFollowing request
	ApiUserFollowing(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UnfollowUser request
	UnfollowUser(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RoundAddTime request
	RoundAddTime(ctx context.Context, gameId string, seconds string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiSimul request
	ApiSimul(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiStormDashboard request
	ApiStormDashboard(ctx context.Context, username string, params *ApiStormDashboardParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BroadcastStreamRoundPgn request
	BroadcastStreamRoundPgn(ctx context.Context, broadcastRoundId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiStreamEvent request
	ApiStreamEvent(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// StreamGame request
	StreamGame(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GamesByUsers request with any body
	GamesByUsersWithBody(ctx context.Context, params *GamesByUsersParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// StudyAllChaptersPgn request
	StudyAllChaptersPgn(ctx context.Context, studyId string, params *StudyAllChaptersPgnParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiSwissNew request with any body
	ApiSwissNewWithBody(ctx context.Context, teamId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiSwissUpdate request with any body
	ApiSwissUpdateWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GamesBySwiss request
	GamesBySwiss(ctx context.Context, id string, params *GamesBySwissParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiSwissJoin request with any body
	ApiSwissJoinWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ResultsBySwiss request
	ResultsBySwiss(ctx context.Context, id string, params *ResultsBySwissParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiSwissTerminate request
	ApiSwissTerminate(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// TeamAll request
	TeamAll(ctx context.Context, params *TeamAllParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// TeamOfUsername request
	TeamOfUsername(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// TeamSearch request
	TeamSearch(ctx context.Context, params *TeamSearchParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// TeamShow request
	TeamShow(ctx context.Context, teamId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiTeamArena request
	ApiTeamArena(ctx context.Context, teamId string, params *ApiTeamArenaParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiTeamSwiss request
	ApiTeamSwiss(ctx context.Context, teamId string, params *ApiTeamSwissParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// TeamIdUsers request
	TeamIdUsers(ctx context.Context, teamId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiTokenDelete request
	ApiTokenDelete(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiToken request with any body
	ApiTokenWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AdminChallengeTokens request with any body
	AdminChallengeTokensWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiTournament request
	ApiTournament(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiTournamentPost request with any body
	ApiTournamentPostWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiTournamentTeamBattlePost request with any body
	ApiTournamentTeamBattlePostWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Tournament request
	Tournament(ctx context.Context, id string, params *TournamentParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiTournamentUpdate request with any body
	ApiTournamentUpdateWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GamesByTournament request
	GamesByTournament(ctx context.Context, id string, params *GamesByTournamentParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiTournamentJoin request with any body
	ApiTournamentJoinWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ResultsByTournament request
	ResultsByTournament(ctx context.Context, id string, params *ResultsByTournamentParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// TeamsByTournament request
	TeamsByTournament(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiTournamentTerminate request
	ApiTournamentTerminate(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiTournamentWithdraw request
	ApiTournamentWithdraw(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// TvChannels request
	TvChannels(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// TvFeed request
	TvFeed(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// TvChannelGames request
	TvChannelGames(ctx context.Context, channel string, params *TvChannelGamesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiUser request
	ApiUser(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiUserActivity request
	ApiUserActivity(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiUserCurrentGame request
	ApiUserCurrentGame(ctx context.Context, username string, params *ApiUserCurrentGameParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiUserPerf request
	ApiUserPerf(ctx context.Context, username string, perf PerfType, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiUserRatingHistory request
	ApiUserRatingHistory(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiUserNameTournamentCreated request
	ApiUserNameTournamentCreated(ctx context.Context, username string, params *ApiUserNameTournamentCreatedParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiUsers request with any body
	ApiUsersWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiUsersStatus request
	ApiUsersStatus(ctx context.Context, params *ApiUsersStatusParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// TablebaseAtomic request
	TablebaseAtomic(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BroadcastTourCreate request with any body
	BroadcastTourCreateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BroadcastRoundUpdate request with any body
	BroadcastRoundUpdateWithBody(ctx context.Context, broadcastRoundId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BroadcastPush request with any body
	BroadcastPushWithBody(ctx context.Context, broadcastRoundId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BroadcastTourUpdate request with any body
	BroadcastTourUpdateWithBody(ctx context.Context, broadcastTournamentId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BroadcastRoundCreate request with any body
	BroadcastRoundCreateWithBody(ctx context.Context, broadcastTournamentId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BroadcastRoundGet request
	BroadcastRoundGet(ctx context.Context, broadcastTournamentSlug string, broadcastRoundSlug string, broadcastRoundId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BroadcastTourGet request
	BroadcastTourGet(ctx context.Context, slug string, broadcastTournamentId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GamePgn request
	GamePgn(ctx context.Context, gameId string, params *GamePgnParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GamesExportIds request with any body
	GamesExportIdsWithBody(ctx context.Context, params *GamesExportIdsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// InboxUsername request with any body
	InboxUsernameWithBody(ctx context.Context, username string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// OpeningExplorerLichess request
	OpeningExplorerLichess(ctx context.Context, params *OpeningExplorerLichessParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// OpeningExplorerMasterGame request
	OpeningExplorerMasterGame(ctx context.Context, gameId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// OpeningExplorerMaster request
	OpeningExplorerMaster(ctx context.Context, params *OpeningExplorerMasterParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Oauth request
	Oauth(ctx context.Context, params *OauthParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Player request
	Player(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PlayerTopNbPerfType request
	PlayerTopNbPerfType(ctx context.Context, nb int, perfType PlayerTopNbPerfTypeParamsPerfType, reqEditors ...RequestEditorFn) (*http.Response, error)

	// OpeningExplorerPlayer request
	OpeningExplorerPlayer(ctx context.Context, params *OpeningExplorerPlayerParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// TablebaseStandard request
	TablebaseStandard(ctx context.Context, params *TablebaseStandardParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// StreamerLive request
	StreamerLive(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// StudyExportAllPgn request
	StudyExportAllPgn(ctx context.Context, username string, params *StudyExportAllPgnParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// StudyChapterPgn request
	StudyChapterPgn(ctx context.Context, studyId string, chapterId string, params *StudyChapterPgnParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// SwissTrf request
	SwissTrf(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// TeamIdJoin request with any body
	TeamIdJoinWithBody(ctx context.Context, teamId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// TeamIdKickUserId request
	TeamIdKickUserId(ctx context.Context, teamId string, userId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// TeamIdPmAll request with any body
	TeamIdPmAllWithBody(ctx context.Context, teamId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// TeamIdQuit request
	TeamIdQuit(ctx context.Context, teamId string, reqEditors ...RequestEditorFn) (*http.Response, error)
}
