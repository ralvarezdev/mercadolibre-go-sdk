package mercadolibre

type OrderStatus string

// Order status values — used in Order.Status and order search filters.
const (
	OrderStatusConfirmed        OrderStatus = "confirmed"
	OrderStatusPaymentRequired  OrderStatus = "payment_required"
	OrderStatusPaymentInProcess OrderStatus = "payment_in_process"
	OrderStatusPartiallyPaid    OrderStatus = "partially_paid"
	OrderStatusPaid             OrderStatus = "paid"
	OrderStatusCancelled        OrderStatus = "cancelled"
	OrderStatusInvalidated      OrderStatus = "invalidated"
	OrderStatusMediation        OrderStatus = "mediation"
	OrderStatusChargedBack      OrderStatus = "charged_back"
)

type ShipmentStatus string

// Shipment status values — used in Shipment.Status.
const (
	ShipmentStatusPendingAgreement ShipmentStatus = "pending_agreement"
	ShipmentStatusPendingPickup    ShipmentStatus = "pending_pickup"
	ShipmentStatusReadyToShip      ShipmentStatus = "ready_to_ship"
	ShipmentStatusHandling         ShipmentStatus = "handling"
	ShipmentStatusShipped          ShipmentStatus = "shipped"
	ShipmentStatusDelivered        ShipmentStatus = "delivered"
	ShipmentStatusNotDelivered     ShipmentStatus = "not_delivered"
	ShipmentStatusLost             ShipmentStatus = "lost"
	ShipmentStatusCancelled        ShipmentStatus = "cancelled"
)

type QuestionStatus string

// Question status values — used in Question.Status.
const (
	QuestionStatusUnanswered QuestionStatus = "UNANSWERED"
	QuestionStatusAnswered   QuestionStatus = "ANSWERED"
	QuestionStatusBanned     QuestionStatus = "BANNED"
	QuestionStatusDeleted    QuestionStatus = "DELETED"
	QuestionStatusDisabled   QuestionStatus = "DISABLED"
	QuestionStatusCloseWin   QuestionStatus = "CLOSE_WIN"
)

type ClaimStatus string

// Claim status values — used in Claim.Status and claim search filters.
const (
	ClaimStatusOpened   ClaimStatus = "opened"
	ClaimStatusResolved ClaimStatus = "resolved"
	ClaimStatusClosed   ClaimStatus = "closed"
)

type ClaimStage string

// Claim stage values — used in Claim.Stage and claim search filters.
const (
	ClaimStageInitial      ClaimStage = "initial"
	ClaimStageMercadoLibre ClaimStage = "ml"
)

type ClaimRole string

// Claim player role values — used in ClaimsService.SearchByPlayer.
const (
	ClaimRoleComplainant ClaimRole = "complainant"
	ClaimRoleRespondent  ClaimRole = "respondent"
)

type ItemStatus string

// Item status values — used in Item.Status and UpdateItemRequest.Status.
const (
	ItemStatusActive   ItemStatus = "active"
	ItemStatusPaused   ItemStatus = "paused"
	ItemStatusClosed   ItemStatus = "closed"
	ItemStatusDeleted  ItemStatus = "deleted"
	ItemStatusInactive ItemStatus = "inactive"
)

type ItemCondition string

// Item condition values — used in Item.Condition and CreateItemRequest.Condition.
const (
	ItemConditionNew          ItemCondition = "new"
	ItemConditionUsed         ItemCondition = "used"
	ItemConditionNotSpecified ItemCondition = "not_specified"
)

type BuyingMode string

// Item buying mode values — used in Item.BuyingMode.
const (
	BuyingModeBuyItNow   BuyingMode = "buy_it_now"
	BuyingModeAuction    BuyingMode = "auction"
	BuyingModeClassified BuyingMode = "classified"
)

type PaymentStatus string

// Payment status values — used in OrderPayment.Status.
const (
	PaymentStatusApproved    PaymentStatus = "approved"
	PaymentStatusPending     PaymentStatus = "pending"
	PaymentStatusInProcess   PaymentStatus = "in_process"
	PaymentStatusRejected    PaymentStatus = "rejected"
	PaymentStatusCancelled   PaymentStatus = "cancelled"
	PaymentStatusRefunded    PaymentStatus = "refunded"
	PaymentStatusChargedBack PaymentStatus = "charged_back"
	PaymentStatusInMediation PaymentStatus = "in_mediation"
)

type PromotionType string

// Promotion type values — used in Promotion.Type and PromotionsService.Items.
const (
	PromotionTypeMarketplaceCampaign PromotionType = "MARKETPLACE_CAMPAIGN"
	PromotionTypeVolume              PromotionType = "VOLUME"
	PromotionTypeDeal                PromotionType = "DEAL"
	PromotionTypeLoyaltyDiscount     PromotionType = "LOYALTY_DISCOUNT"
)
