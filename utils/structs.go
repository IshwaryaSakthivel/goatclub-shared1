package utils

import (
	//"internal/syscall/windows"
	"time"
	// "github.com/aws/aws-sdk-go/service/codebuild"
)

type PaymentDetails struct {
	Id string `json:"id,omitempty"`
	//StripePaymentId              string    `json:"stripePaymentId"`
	UserId                       string  `json:"userId"`
	PaymentAmount                string  `json:"paymentAmount"`
	Plan                         string  `json:"plan"`
	IsUserFreeTrialExpired       bool    `json:"isFreeTrialExpired"`
	TransactionType              string  `json:"transactionType"`
	ProSubscriptionCreated       float64 `json:"proSubscriptionCreated"`
	ProSubscriptionExpiredOn     float64 `json:"proSubscriptionExpiredOn"`
	SubscriptionPaymentfrequency string  `json:"paymentFrequency"`
	//PaymentIntentParams	PaymentIntentParams `json:"paymentIntentParams"`
	// SubscriptionStatus      string    `json:"subscriptionStatus"`
	// SubscriptionCancelledOn time.Time `json:"subscriptionCancelledOn"`
	TotalPoints       int    `json:"totalPoints"`
	Currency          string `json:"currency"`
	Points            int    `json:"points"`
	IsAutoPayEnrolled bool   `json:"isAutoPayEnrolled"`
}

type PaymentStripe struct {
	Amount       string   `json:"amount"`
	Currency     string   `json:"currency"`
	PaymentType  []string `json:"payment_method_types"`
	ReceiptEmail string   `json:"receiptEmail"`
	UserMailId   string   `json:"userMailId"`
}

type PaymentIntentParams struct {
	Amount          int64    `json:"amount"`
	Currency        string   `json:"currency"`
	CustomerId      string   `json:"customer_id"`
	PaymentMethodId string   `json:"payment_method_id"`
	Status          []string `json:"status"`
	ReceiptEmail    string   `json:"receipt_email"`
	PaymentIntentId string   `json:"payment_intent_id"`
	//SetupFutureUsage       string    `json:"setup_future_usage"`
	AutomaticPaymentMethod string `json:"automatic_payment_method"`
	// SubscriptionId         string    `json:"subscription_id"`
	LiveMode    string    `json:"livemode"`
	EmailId     string    `json:"emailId"`
	CreatedDate time.Time `json:"createdDate,omitempty"`
}

type UserStruct struct {
	Username      string `json:"username"`
	EmailAddress  string `json:"emailAddress"`
	Password      string `json:"password"`
	Guid          string `json:"guid,omitempty"`
	AccountStatus string `json:"accountStatus"`
	MobileNumber  string `json:"mobileNumber"`
	//UserDetails   UserDetailsStruct `json:"userDetails"`
	//UserDetails     map[string]interface{} `json:"userDetails"`
	SportPreference          []string  `json:"sportPreference"`
	MyGoats                  []string  `json:"myGoats"`
	MyFans                   []string  `json:"myFans"`
	Following                []string  `json:"following"`
	CreatedDate              time.Time `json:"createdDate,omitempty"`
	Sort                     int       `json:"sort"`
	AccountBalance           string    `json:"accountBalance"`
	IsProUser                bool      `json:"isProUser"`
	TotalPoints              int       `json:"totalPoints"`
	ProfileAvatarUrl         string    `json:"profileAvatarUrl"`
	AvatarUrl                string    `json:"avatarUrl"`
	WinningUnits             int       `json:"winningUnits"`
	WinningPercentage        float64   `json:"winningPercentage"`
	Last30DaysWU             int       `json:"last30DaysWU"`
	Win                      int       `json:"win"`
	Loss                     int       `json:"loss"`
	GoatStatusCode           string    `json:"goatStatusCode"`
	FollowersCount           int       `json:"followersCount"`
	IsEmailNotificationOpted bool      `json:"isEmailNotificationOpted"`
	ProSubscriptionPlan      string    `json:"proSubscriptionPlan"`
	ProSubscriptionExpiredOn int64     `json:"proSubscriptionExpiredOn"`
	IsFreeTrialExpired       bool      `json:"isFreeTrialExpired"`
	FreeTrialExpireOn        time.Time `json:"freeTrialExpireOn"`
}

type ProfileStruct struct {
	Username          string  `json:"displayName"`
	Guid              string  `json:"personGUID,omitempty"`
	ProfileDP         string  `json:"avatarUrl"`
	UserStatus        string  `json:"goatStatus"`
	WinningUnits      int     `json:"winningUnits"`
	FollowersCount    int     `json:"followersCount"`
	WinningPercentage float64 `json:"winPercentage"`
	Last30DaysWU      int     `json:"last30DaysPoints"`
	Win               int     `json:"win"`
	Loss              int     `json:"loss"`
}
type ChatStruct struct {
	MessageParentId string    `json:"messageParentId"`
	PersonsGUID     string    `json:"personsGUID"`
	ChatId          string    `json:"chatId"`
	RoomName        string    `json:"roomName"`
	Message         string    `json:"message"`
	DisplayName     string    `json:"displayName"`
	UserId          string    `json:"userId"`
	AvatarUrl       string    `json:"avatarUrl"`
	RankLevelCode   string    `json:"rankLevelCode"`
	TimeStamp       time.Time `json:"timeStamp"`
	LikesCount      string    `json:"likesCount"`
	CommentsCount   string    `json:"commentsCount"`
	UpdateMode      string    `json:"updateMode"`
	ReceiverGUID    string    `json:"receiverGUID"`
	SenderGUID      string    `json:"senderGUID"`
	CreatedDate     time.Time `json:"createdDate,omitempty"`
}

type PvtChatStruct struct {
	MessageParentId string    `json:"messageParentId"`
	PersonsGUID     string    `json:"personsGUID"`
	ChatId          string    `json:"chatId"`
	RoomName        string    `json:"roomName"`
	Message         string    `json:"message"`
	DisplayName     string    `json:"displayName"`
	UserId          string    `json:"userId"`
	AvatarUrl       string    `json:"avatarUrl"`
	RankLevelCode   string    `json:"rankLevelCode"`
	TimeStamp       time.Time `json:"timeStamp"`
	LikesCount      string    `json:"likesCount"`
	CommentsCount   string    `json:"commentsCount"`
	UpdateMode      string    `json:"updateMode"`
	ReceiverGUID    string    `json:"receiverGUID"`
	SenderGUID      string    `json:"senderGUID"`
	CreatedDate     time.Time `json:"createdDate,omitempty"`
}

type PickStruct struct {
	// PickId      string            `json:"pickId,omitempty"`
	UserId      string        `json:"userId,omitempty"`
	CreatedDate time.Time     `json:"createdDate,omitempty"`
	UpdateDate  time.Time     `json:"updateDate"`
	PickDetails []PickDetails `json:"pickDetails"`
}
type PickDetails struct {
	Picks           []Picks   `json:"picks"`
	OddDataUnits    int       `json:"oddDataUnits"`
	OddDataWins     int       `json:"oddDataWins"`
	PickId          string    `json:"pickId"`
	Gamestatus      []string  `json:"gamestatus"`
	PickStatus      string    `json:"pickStatus"`
	Picktype        string    `json:"picktype"`
	PickCreatedDate time.Time `json:"pickCreatedDate,omitempty"`
}
type Picks struct {
	EventId          string `json:"eventId"`
	TeamName         string `json:"teamName"`
	OppTeamName      string `json:"oppTeamName"`
	Sport            string `json:"sport"`
	OddDataMoneyLine int    `json:"oddDataMoneyLine"`
}
type AwsPick struct {
	GameStatus      bool          `json:"gamestatus.NULL"`
	OddDataUnits    int           `json:"oddDataUnits.N"`
	OddDataWins     int           `json:"oddDataWins.N"`
	PickCreatedDate time.Time     `json:"pickCreatedDate.S"`
	PickId          string        `json:"pickId.S"`
	PickStatus      string        `json:"pickStatus.S"`
	Picks           []AwsPickItem `json:"picks.L"`
	PickType        string        `json:"picktype.S"`
}
type PaymentDetailsApi struct {
	Amount             *int64    `json:"amount"`
	Currency           *string   `json:"currency"`
	PaymentMethodTypes []*string `json:"payment_method_types"`
	ReceiptEmail       *string   `json:"receipt_email"`
	SetupFutureUsage   *string   `json:"setup_future_usage"`
	// PaymentMethodOptions      *PaymentIntentPaymentMethodOptionsParams `json:"payment_method_options"`

}

//	 type PaymentIntentPaymentMethodOptionsCardInstallments struct {
//	    // AvailablePlans []*PaymentIntentPaymentMethodOptionsCardInstallmentsPlan `json:"available_plans"`
//	    Enabled        bool                                                     `json:"enabled"`
//	    // Plan           *PaymentIntentPaymentMethodOptionsCardInstallmentsPlan   `json:"plan"`
//	}
type AwsPickItem struct {
	EventId          string `json:"eventId.S"`
	OddDataMoneyLine int    `json:"oddDataMoneyLine.N"`
	Sport            string `json:"sport.S"`
	TeamName         string `json:"teamname.S"`
}

type SportStruct struct {
	SportsId        string            `json:"sportsId,omitempty"`
	SportsName      string            `json:"sportsName"`
	UpcomingGames   map[string]string `json:"upcomingGames"`
	ChatroomHistory map[string]string `json:"chatroomHistory"`
}
type CheckUsernameResp struct {
	IsUserNameAlreadyExist bool `json:"isUserNameAlreadyExist"`
}
type CheckUserEmailResp struct {
	IsUserEmailAlreadyExist bool `json:"isUserEmailAlreadyExist"`
}
type OddsBySportParam struct {
	Sport string `json:"sport"`
}
type Score struct {
	Name  string `json:"name"`
	Score string `json:"score"`
}
type UpdateAvatar struct {
	ProfileDP string `json:"profileDp"`
}
type Game struct {
	ID           string    `json:"id"`
	SportKey     string    `json:"sport_key"`
	SportTitle   string    `json:"sport_title"`
	CommenceTime time.Time `json:"commence_time"`
	Completed    bool      `json:"completed"`
	HomeTeam     string    `json:"home_team"`
	AwayTeam     string    `json:"away_team"`
	Scores       []Score   `json:"scores"`
	LastUpdate   time.Time `json:"last_update"`
}

type Pricing struct {
	Code        string `json:"code"`
	Points      string `json:"points"`
	Price       string `json:"price"`
	PricePerDay string `json:"pricePerDay"`
}

type Plans struct {
	Code           string `json:"code"`
	TotalPrice     string `json:"totalPrice"`
	PerDayPrice    string `json:"perDayPrice"`
	BonusPoints    string `json:"bonusPoints"`
	SavePercentage string `json:"savePercentage"`
}
type CheckPickResp struct {
	IsPickExist bool `json:"isPickExist"`
}
type Avatar struct {
	Base64string string `json:"base64"`
}
type Statistics struct {
	PickerId         string         `json:"pickerId,omitempty"`
	Summary          PickStatistics `json:"summary"`
	NBASummary       PickStatistics `json:"nbasummary"`
	MLBSummary       PickStatistics `json:"mlbsummary"`
	TennisSummary    PickStatistics `json:"tennissummary"`
	Yesterday        PickStatistics `json:"yesterday"`
	NBAYesterday     PickStatistics `json:"nbayesterday"`
	MLBYesterday     PickStatistics `json:"mlbyesterday"`
	TennisYesterday  PickStatistics `json:"tennisyesterday"`
	Last7Days        PickStatistics `json:"last7Days"`
	NBALast7Days     PickStatistics `json:"nbalast7Days"`
	MLBLast7Days     PickStatistics `json:"mlblast7Days"`
	TennisLast7Days  PickStatistics `json:"tennislast7Days"`
	Last30Days       PickStatistics `json:"last30Days"`
	NBALast30Days    PickStatistics `json:"nbalast30Days"`
	MLBLast30Days    PickStatistics `json:"mlblast30Days"`
	TennisLast30Days PickStatistics `json:"tennislast30Days"`
	Last60Days       PickStatistics `json:"last60Days"`
	NBALast60Days    PickStatistics `json:"nbalast60Days"`
	MLBLast60Days    PickStatistics `json:"mlblast60Days"`
	TennisLast60Days PickStatistics `json:"tennislast60Days"`
	UpdateDate       time.Time      `json:"updateDate"`
}

type PickStatistics struct {
	Followers         int     `json:"followers"`
	PicksCount        int     `json:"picksCount"`
	WinningUnits      int     `json:"winningUnits"`
	WinningPercentage float64 `json:"winningPercentage"`
	Last30DaysWU      int     `json:"last30DaysWU"`
	Win               int     `json:"win"`
	Loss              int     `json:"loss"`
	WinLoss           string  `json:"winLoss"`
}

type OddsBySport struct {
	DataId string  `json:"dataId"`
	Data   []Event `json:"data"`
}
type StaticData struct {
	DataId string `json:"dataId"`
	Data   string `json:"data"`
}
type Outcome struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Market struct {
	Key        string    `json:"key"`
	LastUpdate string    `json:"last_update"`
	Outcomes   []Outcome `json:"outcomes"`
}

type Bookmaker struct {
	Key        string   `json:"key"`
	Title      string   `json:"title"`
	LastUpdate string   `json:"last_update"`
	Markets    []Market `json:"markets"`
}

type Event struct {
	ID           string      `json:"id"`
	SportKey     string      `json:"sport_key"`
	SportTitle   string      `json:"sport_title"`
	CommenceTime string      `json:"commence_time"`
	HomeTeam     string      `json:"home_team"`
	AwayTeam     string      `json:"away_team"`
	Bookmakers   []Bookmaker `json:"bookmakers"`
}

type NotificationsStack struct {
	NotificationId      string    `json:"notificationId"`
	NotificationType    string    `json:"notificationType"`
	UserId              string    `json:"userId"`
	CreatedDate         time.Time `json:"createdDate"`
	TeamsPicked         []string  `json:"teamsPicked"`
	NotificationMessage string    `json:"notificationMessage"`
}
type UserNotificationsActivities struct {
	NotificationId string    `json:"notificationId"`
	Seen           bool      `json:"seen"`
	SourceGuid     string    `json:"sourceGuid"`
	CreatedDate    time.Time `json:"createdDate"`
	SeenDate       time.Time `json:"seenDate"`
	SubscriberGuid string    `json:"subscriberGuid"`
}
type Websocket struct {
	ConnectionId string `json:"connectionId"`
}
