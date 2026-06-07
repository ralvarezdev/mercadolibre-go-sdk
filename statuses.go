package mercadolibre

// Order status values — used in Order.Status and order search filters.
const (
	OrderStatusConfirmed        = "confirmed"
	OrderStatusPaymentRequired  = "payment_required"
	OrderStatusPaymentInProcess = "payment_in_process"
	OrderStatusPartiallyPaid    = "partially_paid"
	OrderStatusPaid             = "paid"
	OrderStatusCancelled        = "cancelled"
	OrderStatusInvalidated      = "invalidated"
	OrderStatusMediation        = "mediation"
	OrderStatusChargedBack      = "charged_back"
)

// Shipment status values — used in Shipment.Status.
const (
	ShipmentStatusPendingAgreement = "pending_agreement"
	ShipmentStatusPendingPickup    = "pending_pickup"
	ShipmentStatusReadyToShip      = "ready_to_ship"
	ShipmentStatusHandling         = "handling"
	ShipmentStatusShipped          = "shipped"
	ShipmentStatusDelivered        = "delivered"
	ShipmentStatusNotDelivered     = "not_delivered"
	ShipmentStatusLost             = "lost"
	ShipmentStatusCancelled        = "cancelled"
)

// Question status values — used in Question.Status.
const (
	QuestionStatusUnanswered = "UNANSWERED"
	QuestionStatusAnswered   = "ANSWERED"
	QuestionStatusBanned     = "BANNED"
	QuestionStatusDeleted    = "DELETED"
	QuestionStatusDisabled   = "DISABLED"
	QuestionStatusCloseWin   = "CLOSE_WIN"
)

// Claim status values — used in Claim.Status and claim search filters.
const (
	ClaimStatusOpened   = "opened"
	ClaimStatusResolved = "resolved"
	ClaimStatusClosed   = "closed"
)

// Claim stage values — used in Claim.Stage and claim search filters.
const (
	ClaimStageInitial      = "initial"
	ClaimStageMercadoLibre = "ml"
)

// Claim player role values — used in ClaimsService.SearchByPlayer.
const (
	ClaimRoleComplainant = "complainant" // typically the buyer
	ClaimRoleRespondent  = "respondent"  // typically the seller
)

// Item status values — used in Item.Status and UpdateItemRequest.Status.
const (
	ItemStatusActive   = "active"
	ItemStatusPaused   = "paused"
	ItemStatusClosed   = "closed"
	ItemStatusDeleted  = "deleted"
	ItemStatusInactive = "inactive"
)

// Item condition values — used in Item.Condition and CreateItemRequest.Condition.
const (
	ItemConditionNew          = "new"
	ItemConditionUsed         = "used"
	ItemConditionNotSpecified = "not_specified"
)

// Item buying mode values — used in Item.BuyingMode.
const (
	BuyingModeBuyItNow   = "buy_it_now"
	BuyingModeAuction    = "auction"
	BuyingModeClassified = "classified"
)

// Payment status values — used in OrderPayment.Status.
const (
	PaymentStatusApproved   = "approved"
	PaymentStatusPending    = "pending"
	PaymentStatusInProcess  = "in_process"
	PaymentStatusRejected   = "rejected"
	PaymentStatusCancelled  = "cancelled"
	PaymentStatusRefunded   = "refunded"
	PaymentStatusChargedBack = "charged_back"
	PaymentStatusInMediation = "in_mediation"
)

// Promotion type values — used in Promotion.Type and PromotionsService.Items.
const (
	PromotionTypeMarketplaceCampaign = "MARKETPLACE_CAMPAIGN"
	PromotionTypeVolume              = "VOLUME"
	PromotionTypeDeal                = "DEAL"
	PromotionTypeLoyaltyDiscount     = "LOYALTY_DISCOUNT"
)
