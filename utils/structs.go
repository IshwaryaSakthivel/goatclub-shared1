package utils

import (
	//"internal/syscall/windows"
	"time"
)

type UserStruct struct {
	Username                 string    `json:"username"`
	EmailAddress             string    `json:"emailAddress"`
	Password                 string    `json:"password"`
	Guid                     string    `json:"guid,omitempty"`
	AccountStatus            string    `json:"accountStatus"`
	MobileNumber             string    `json:"mobileNumber"`
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
	// UserDetails   UserDetailsStruct `json:"userDetails"`
	// UserDetails     map[string]interface{} `json:"userDetails"`
	InAppEntitlementID    string `json:"inAppEntitlementID"`
	InAppOfferingIdentifier string `json:"inAppOfferingIdentifier"`
	Description string `json:"description"`
	SubscriptionPaymentMode     string    `json:"subscriptionPaymentMode"`
	GoatSubscriptionMode  []string  `json:"goatSubscriptionMode"`
	IsPromoUser             bool     `json:"isPromoUser"`
	PromoCode               string   `json:"promoCode"`
}


type PromoCodeUser struct {
	Guid        string `json:"guid"`
	Email       string `json:"email"`
	IsPromoUser bool   `json:"isPromoUser"`
	PromoCode   string `json:"promoCode"`
}


type Profile struct {
	Guid              string  `json:"guid,omitempty"`
	Username          string  `json:"username"`
	AvatarUrl         string  `json:"avatarUrl"`
	GoatStatusCode    string  `json:"goatStatusCode"`
	WinningUnits      int     `json:"winningUnits"`
	FollowersCount    int     `json:"followersCount"`
	WinningPercentage float64 `json:"winningPercentage"`
	Last30DaysWU      int     `json:"last30DaysWU"`
	Win               int     `json:"win"`
	Loss              int     `json:"loss"`
	InAppEntitlementID    string `json:"inAppEntitlementID"`
	InAppOfferingIdentifier string `json:"inAppOfferingIdentifier"`
	Description string `json:"description"`
	GoatSubscriptionMode  []string  `json:"goatSubscriptionMode"`
}
type MyGoatsProfile struct {
	Guid              string  `json:"guid,omitempty"`
	Username          string  `json:"username"`
	AvatarUrl         string  `json:"avatarUrl"`
	GoatStatusCode    string  `json:"goatStatusCode"`
	WinningUnits      int     `json:"winningUnits"`
	FollowersCount    int     `json:"followersCount"`
	WinningPercentage float64 `json:"winningPercentage"`
	Last30DaysWU      int     `json:"last30DaysWU"`
	Win               int     `json:"win"`
	Loss              int     `json:"loss"`
	IsExpiringOn      int64   `json:"IsExpiringOn"`
	IsSubscribed      bool    `json:"isSubscribed"`
	IsFollowed        bool    `json:"isFollowed"`
}

type MyChatsProfile struct {
	Guid              string    `json:"guid,omitempty"`
	Username          string    `json:"username"`
	AvatarUrl         string    `json:"avatarUrl"`
	GoatStatusCode    string    `json:"goatStatusCode"`
	WinningUnits      int       `json:"winningUnits"`
	FollowersCount    int       `json:"followersCount"`
	WinningPercentage float64   `json:"winningPercentage"`
	Last30DaysWU      int       `json:"last30DaysWU"`
	Win               int       `json:"win"`
	Loss              int       `json:"loss"`
	LastMessageTime   time.Time `json:"lastMessageTime"`
}

type Persons struct {
	PersonsGUID string    `json:"personGuid"`
	CreatedDate time.Time `json:"createdDate"`
	Persons     []string  `json:"persons"`
}
type Guids struct {
	UserGuid   string `json:"userGuid"`
	PickerGuid string `json:"pickerGuid"`
}

type MyGoats struct {
	Guid    string   `json:"guid"`
	MyGoats []string `json:"myGoats"`
}

type Followers struct {
	Guid string   `json:"guid"`
	Fans []string `json:"myFans"`
}

type CheckMyGoat struct {
	IsGoatSubscribed bool `json:"isGoatSubscribed"`
}
type Updated struct {
	Updated string `json:"updated"`
}
type Response struct {
	Updated bool `json:"updated"`
}
type Expiry struct {
	Expiry            int64 `json:"expires_at"`
	IsAutoPayEnrolled bool  `json:"isAutoPayEnrolled"`
	CancelAtPeriodEnd bool  `json:"cancelAtPeriodEnd"`
}
type ExpiryInfo struct {
	Guid  string `json:"guid"`
	Type  string `json:"type"`
	Email string `json:"email"`
}
type Goats struct {
	Guid      string   `json:"guid"`
	Goats     []string `json:"myGoats"`
	Following []string `json:"following"`
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
	LikesCount      int  	  `json:"likesCount"`
	CommentsCount   int  	  `json:"commentsCount"`
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
	ProfileAvatarUrl string `json:"profileAvatarUrl"`
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
	SeenDate       int64     `json:"seenDate"`
	SubscriberGuid string    `json:"subscriberGuid"`
	StackId        string    `json:"stackId"`
}
type Websocket struct {
	ConnectionId string `json:"connectionId"`
}

// Payment
type PaymentStripe struct {
	Amount       string   `json:"amount"`
	Currency     string   `json:"currency"`
	PaymentType  []string `json:"payment_method_types"`
	ReceiptEmail string   `json:"receiptEmail"`
	UserMailId   string   `json:"userMailId"`
}
type ProductStruct struct {
	ProductId   string  `json:"id"`
	CreatedDate float64 `json:"created"`
	LiveMode    bool    `json:"livemode"`
	Active      bool    `json:"active"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	UpdatedDate float64 `json:"updated"`
}

type PriceStruct struct {
	PriceId     string  `json:"id"`
	CreatedDate float64 `json:"created"`
	LiveMode    bool    `json:"livemode"`
	Active      bool    `json:"active"`
	Name        string  `json:"nickname"`
	ProductId   string  `json:"product"`
	Currency    string  `json:"currency"`
	Amount      int64   `json:"unit_amount"`
	Type        string  `json:"type"`
}

type ProCancelSubscription struct {
	UserId       string `json:"userId"`
	Customer     string `json:"customer"`
	Subscription string `json:"subscription"`
}
type GoatCancelSubscription struct {
	UserId       string `json:"userId"`
	GoatGuid     string `json:"goatGuid"`
	Customer     string `json:"customer"`
	Subscription string `json:"subscription"`
	Email        string `json:"email"`
}
type CardDetails struct {
	UserId     string `json:"userId"`
	UserEmail  string `json:"userEmail"`
	CardId     string `json:"cardId"`
	CustomerId string `json:"customerId"`
}
type CardDetailsInfo struct {
	UserId      string `json:"userId"`
	UserEmail   string `json:"userEmail"`
	CardId      string `json:"cardId"`
	CardNumber  string `json:"cardNumber"`
	ExpiryYear  string `json:"expiryYear"`
	Cvc         string `json:"cvc"`
	ExpiryMonth string `json:"expiryMonth"`
	CustomerId  string `json:"customerId"`
	Name        string `json:"name"`
}
type ErrorMessage struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}
type ProMappingStruct struct {
	CustomerId string `json:"customerId"`
	UserId     string `json:"userId"`
	EmailId    string `json:"email"`
	PriceId    string `json:"priceId"`
}
type GoatMappingStruct struct {
	CustomerId     string    `json:"customerId"`
	UserId         string    `json:"guid"`
	EmailId        string    `json:"email"`
	PriceId        string    `json:"priceId"`
	GoatGuid       string    `json:"goatGuid"`
	SubscriptionId string    `json:"SubscriptionId"`
	GoatStatus     string    `json:"goatStatus"`
	CreatedDate    time.Time `json:"createdDate,omitempty"`
}
type CheckoutSessionParams struct {
	ID                           string  `json:"id"`
	InvoicecID                   string  `json:"invoice"`
	SubscriptionID               string  `json:"subscription"`
	Amount                       float64 `json:"amount_total"`
	Currency                     string  `json:"currency"`
	CustomerId                   string  `json:"customer"`
	PaymentStatus                string  `json:"payment_status"`
	Status                       string  `json:"status"`
	ReceiptEmail                 string  `json:"email"`
	LiveMode                     bool    `json:"livemode"`
	CreatedDate                  int64   `json:"created,omitempty"`
	ExpiredDate                  int64   `json:"expires_at,omitempty"`
	UserId                       string  `json:"userId"`
	Mode                         string  `json:"mode"`
	Name                         string  `json:"name"`
	SubscriptionPaymentfrequency string  `json:"paymentFrequency"`
	IsAutoPayEnrolled            bool    `json:"isAutoPayEnrolled"`
	Active                       bool    `json:"active"`
	CancelAtPeriodEnd            bool    `json:"cancelAtPeriodEnd"`
	// IsUserFreeTrialExpired       bool             `json:"isFreeTrialExpired"`
}
type CustomerSubscription struct {
	SubscriptionId          string `json:" id"`
	CustomerId              string `json:"customer"`
	SubscriptionCancelledAt string `json:"canceled_at"`
	SubscriptionCancelAt    string `json:"cancel_at"`
	Description             string `json:"description"`
	CurrentPeriodEnd        string `json:"current_period_end"`
	CurrentPeriodStart      string `json:"current_period_start"`
	DefaultPaymentMethod    string `json:"default_payment_method"`
	Email                   string `json:"email"`
	GoatGuid                string `json:"goatGuid"`
	Subscription            string `json:"SubscriptionId"`
	UserId                  string `json:"userId"`
	UserGuid                string `json:"guid"`
}
type GoatCheckoutSessionParams struct {
	ID                           string  `json:"id"`
	InvoicecID                   string  `json:"invoice"`
	SubscriptionID               string  `json:"subscription"`
	Amount                       float64 `json:"amount_total"`
	Currency                     string  `json:"currency"`
	CustomerId                   string  `json:"customer"`
	PaymentStatus                string  `json:"payment_status"`
	Status                       string  `json:"status"`
	ReceiptEmail                 string  `json:"email"`
	LiveMode                     bool    `json:"livemode"`
	CreatedDate                  int64   `json:"created,omitempty"`
	ExpiredDate                  int64   `json:"expires_at,omitempty"`
	UserId                       string  `json:"userId"`
	Mode                         string  `json:"mode"`
	Name                         string  `json:"name"`
	GoatStatus                   string  `json:"goatStatus"`
	GoatGuid                     string  `json:"GoatGuid"`
	SubscriptionPaymentfrequency string  `json:"paymentFrequency"`
	IsAutoPayEnrolled            bool    `json:"isAutoPayEnrolled"`
	Active                       bool    `json:"active"`
	CancelAtPeriodEnd            bool    `json:"cancelAtPeriodEnd"`
}
type SessionIDResponse struct {
	SessionID  string `"json:id"`
	PaymentUrl string `json:"url"`
}
type CheckoutSession struct {
	BillingAddressCollection *string   `json:"billing_address_collection"`
	CancelURL                *string   `json:"cancel_url"`
	ID                       *string   `json:"id"`
	Customer                 *string   `json:"customerId"`
	CustomerEmail            *string   `json:"userEmail"`
	Mode                     *string   `json:"mode"`
	PaymentMethodTypes       []*string `json:"payment_method_types"`
	SubmitType               *string   `json:"submit_type"`
	SuccessURL               *string   `json:"success_url"`
	UserId                   string    `json:"userGuid"`
	Plan                     string    `json:"plan"`
	PaymentAmount            string    `json:"paymentAmount"`
	SubscriptionId           string    `json:"subscription"`
	PriceId                  string    `json:"priceId"`
}
type TransactionDetails struct {
	Id              string    `json:"userId"`
	TransactionId   string    `json:"transactionId"`
	PaymentAmount   string    `json:"paymentAmount"`
	Plan            string    `json:"plan"`
	TransactionType string    `json:"transactionType"`
	TotalPoints     int       `json:"totalPoints"`
	Currency        string    `json:"currency"`
	Points          int       `json:"points"`
	CreatedDate     float64   `json:"createdDate,omitempty"`
	GoatStatus      string    `json:"goatStatus"`
	GoatGuid        string    `json:"GoatGuid"`
	Status          string    `json:"status"`
	TransactedAt    time.Time `json:"TransactedAt"`
}

type PaymentIntentFailureParams struct {
	ID              string  `json:"id"`
	Currency        string  `json:"currency"`
	CustomerId      string  `json:"customer"`
	PaymentMethodId string  `json:"payment_method"`
	Amount          string  `json:"amount"`
	Status          string  `json:"status"`
	ReceiptEmail    string  `json:"receipt_email"`
	SubscriptionId  string  `json:"subscription"`
	LiveMode        bool    `json:"livemode"`
	CreatedDate     float64 `json:"createdDate,omitempty"`
	UserId          string  `json:"userId"`
	Code            string  `json:"code"`
	DeclineCode     string  `json:"decline_code"`
	Message         string  `json:"message"`
}
type PaymentIntentParams struct {
	ID string `json:"id"`
	// PaymentId    string  `json:"paymentId"`
	Amount          float64 `json:"amount"`
	Currency        string  `json:"currency"`
	CustomerId      string  `json:"customer"`
	PaymentMethodId string  `json:"payment_method_id"`
	Status          string  `json:"status"`
	ReceiptEmail    string  `json:"receipt_email"`
	LiveMode        bool    `json:"livemode"`
	CreatedDate     float64 `json:"created,omitempty"`
	UserId          string  `json:"userId"`
}
type GoatsDetail struct {
	Guid                     string   `json:"guid"`
	MyFans                   []string `json:"myFans"`
	FollowersCount           int      `json:"followersCount"`
	Username                 string   `json:"username"`
	EmailAddress             string   `json:"emailAddress"`
	IsEmailNotificationOpted bool     `json:"isEmailNotificationOpted"`
	// MyGoats []string `json:"myGoats"`
}
type User struct {
	Guid                     string   `json:"guid"`
	MyGoats                  []string `json:"myGoats"`
	MyFans                   []string `json:"myFans"`
	Username                 string   `json:"username"`
	Following                []string `json:"following"`
	EmailAddress             string   `json:"emailAddress"`
	IsEmailNotificationOpted bool     `json:"isEmailNotificationOpted"`
}
type EventData struct {
	APIVersion string `json:"api_version"`
	Event      Eventrc  `json:"event"`
}

type Eventrc struct {
	Email                 string                         `json:"email"`
	UserId                string                         `json:"user_id"`
	AppuserId             string                         `json:"app_user_id"`
	ProductId             string                         `json:"product_id"`
	EventType             string                         `json:"type"`
	SubscriberAttributes  map[string]SubscriberAttribute `json:"subscriber_attributes"`
	Price                 float64                        `json:"price_in_purchased_currency"`
	Purchased_at_ms       float64                        `json:"purchased_at_ms"`
	Expires_at            float64                        `json:"expiration_at_ms"`
	Transaction_id        string                         `json:"transaction_id"`
	OriginalTransactionID string                         `json:"original_transaction_id"`
	Currency              string                         `json:"currency"`
	CancelReason          string                         `json:"cancel_reason"`
	Environment           string                         `json:"environment"`
}

type SubscriberAttribute struct {
	UpdatedAtMS int64  `json:"updated_at_ms"`
	Value       string `json:"value"`
}