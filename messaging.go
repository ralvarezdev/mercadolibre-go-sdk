package mercadolibre

import (
	"context"
	"net/url"
)

// Message is a post-sale message (new messaging architecture).
type Message struct {
	ID                     string           `json:"id"`
	SiteID                 string           `json:"site_id,omitempty"`
	ClientID               int64            `json:"client_id,omitempty"`
	From                   MessageParty     `json:"from"`
	To                     MessageParty     `json:"to"`
	Status                 string           `json:"status,omitempty"`
	Subject                *string          `json:"subject,omitempty"`
	Text                   string           `json:"text"`
	MessageDate            *MessageDate     `json:"message_date,omitempty"`
	MessageModeration      *MsgModeration   `json:"message_moderation,omitempty"`
	MessageAttachments     []MessageAttach  `json:"message_attachments,omitempty"`
	MessageResources       []IDName         `json:"message_resources,omitempty"`
	ConversationFirstMessage bool           `json:"conversation_first_message,omitempty"`
}

// MessageParty identifies a message sender/recipient.
type MessageParty struct {
	UserID int64 `json:"user_id"`
}

// MessageDate holds the message lifecycle timestamps.
type MessageDate struct {
	Received  Time  `json:"received,omitempty"`
	Available Time  `json:"available,omitempty"`
	Notified  Time  `json:"notified,omitempty"`
	Created   Time  `json:"created,omitempty"`
	Read      *Time `json:"read,omitempty"`
}

// MsgModeration holds moderation status for a message.
type MsgModeration struct {
	Status         string  `json:"status,omitempty"`
	Reason         *string `json:"reason,omitempty"`
	Source         string  `json:"source,omitempty"`
	ModerationDate Time    `json:"moderation_date,omitempty"`
}

// MessageAttach is an attachment reference on a message.
type MessageAttach struct {
	Filename     string `json:"filename,omitempty"`
	OriginalName string `json:"original_filename,omitempty"`
	Type         string `json:"type,omitempty"`
	Size         int64  `json:"size,omitempty"`
}

// ConversationStatus describes the state of a pack conversation.
type ConversationStatus struct {
	Path                string  `json:"path,omitempty"`
	Status              string  `json:"status,omitempty"`
	Substatus           *string `json:"substatus,omitempty"`
	StatusDate          Time    `json:"status_date,omitempty"`
	StatusUpdateAllowed bool    `json:"status_update_allowed,omitempty"`
	ClaimID             *int64  `json:"claim_id,omitempty"`
	ShippingID          *int64  `json:"shipping_id,omitempty"`
}

// MessagesResponse is the response of listing a pack's messages.
type MessagesResponse struct {
	Paging             *Paging             `json:"paging,omitempty"`
	ConversationStatus *ConversationStatus `json:"conversation_status,omitempty"`
	Messages           []Message           `json:"messages"`
	SellerMaxLength    int                 `json:"seller_max_message_length,omitempty"`
	BuyerMaxLength     int                 `json:"buyer_max_message_length,omitempty"`
}

// SendMessageRequest is the body for sending a post-sale message. Omit
// Attachments when there is no file.
type SendMessageRequest struct {
	From        MessageParty `json:"from"`
	To          MessageParty `json:"to"`
	Text        string       `json:"text"`
	Attachments []string     `json:"attachments,omitempty"`
}

// MessagingService accesses post-sale messaging. All calls use tag=post_sale.
type MessagingService struct{ c *Client }

func postSaleQuery(extra url.Values) url.Values {
	q := url.Values{"tag": {"post_sale"}}
	for k, v := range extra {
		q[k] = v
	}
	return q
}

// PackMessages lists the messages of a pack/order conversation
// (GET /messages/packs/{pack_id}/sellers/{seller_id}?tag=post_sale). When the
// order has no pack, pass the order ID as packID. By default this marks messages
// as read; pass markAsRead=false to avoid that.
func (s *MessagingService) PackMessages(ctx context.Context, packID, sellerID int64, opts ListOptions, markAsRead bool) (*MessagesResponse, error) {
	extra := opts.Values()
	if !markAsRead {
		extra.Set("mark_as_read", "false")
	}
	return GetQ[*MessagesResponse](ctx, s.c, EPMessagesPackSeller, postSaleQuery(extra), packID, sellerID)
}

// Get returns a single message by ID
// (GET /messages/{message_id}?tag=post_sale).
func (s *MessagingService) Get(ctx context.Context, messageID string) (*Message, error) {
	return GetQ[*Message](ctx, s.c, EPMessageByID, postSaleQuery(nil), messageID)
}

// Send posts a message to the buyer
// (POST /messages/packs/{pack_id}/sellers/{seller_id}?tag=post_sale). Note: for
// MLB/MLC the To.UserID must be the country's agent ID under the new messaging
// architecture.
func (s *MessagingService) Send(ctx context.Context, packID, sellerID int64, req SendMessageRequest) (*Message, error) {
	return Do[*Message, SendMessageRequest](ctx, s.c, "POST", Request[SendMessageRequest]{
		Endpoint: EPMessagesPackSeller,
		Args:     []any{packID, sellerID},
		Query:    postSaleQuery(nil),
		Body:     &req,
	})
}

// GetAttachment downloads the raw bytes of a message attachment
// (GET /messages/attachments/{id}?tag=post_sale).
// The content type matches the original file (image, PDF, etc.).
func (s *MessagingService) GetAttachment(ctx context.Context, attachmentID string) ([]byte, error) {
	return GetRaw(ctx, s.c, EPMessageAttachByID, postSaleQuery(nil), attachmentID)
}
