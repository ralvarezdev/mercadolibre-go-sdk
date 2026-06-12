package mercadolibre

import (
	"context"
	"net/url"
)

type (
	// Message is a post-sale message (new messaging architecture).
	Message struct {
		Subject                  *string         `json:"subject,omitempty"`
		MessageModeration        *MsgModeration  `json:"message_moderation,omitempty"`
		MessageDate              *MessageDate    `json:"message_date,omitempty"`
		Status                   string          `json:"status,omitempty"`
		ID                       string          `json:"id"`
		SiteID                   string          `json:"site_id,omitempty"`
		Text                     string          `json:"text"`
		MessageResources         []IDName        `json:"message_resources,omitempty"`
		MessageAttachments       []MessageAttach `json:"message_attachments,omitempty"`
		From                     MessageParty    `json:"from"`
		To                       MessageParty    `json:"to"`
		ClientID                 int64           `json:"client_id,omitempty"`
		ConversationFirstMessage bool            `json:"conversation_first_message,omitempty"`
	}

	// MessageParty identifies a message sender/recipient.
	MessageParty struct {
		UserID int64 `json:"user_id"`
	}

	// MessageDate holds the message lifecycle timestamps.
	MessageDate struct {
		Read      *Time `json:"read,omitempty"`
		Received  Time  `json:"received,omitempty"`
		Available Time  `json:"available,omitempty"`
		Notified  Time  `json:"notified,omitempty"`
		Created   Time  `json:"created,omitempty"`
	}

	// MsgModeration holds moderation status for a message.
	MsgModeration struct {
		ModerationDate Time    `json:"moderation_date,omitempty"`
		Reason         *string `json:"reason,omitempty"`
		Status         string  `json:"status,omitempty"`
		Source         string  `json:"source,omitempty"`
	}

	// MessageAttach is an attachment reference on a message.
	MessageAttach struct {
		Filename     string `json:"filename,omitempty"`
		OriginalName string `json:"original_filename,omitempty"`
		Type         string `json:"type,omitempty"`
		Size         int64  `json:"size,omitempty"`
	}

	// ConversationStatus describes the state of a pack conversation.
	ConversationStatus struct {
		StatusDate          Time    `json:"status_date,omitempty"`
		ClaimID             *int64  `json:"claim_id,omitempty"`
		ShippingID          *int64  `json:"shipping_id,omitempty"`
		Substatus           *string `json:"substatus,omitempty"`
		Path                string  `json:"path,omitempty"`
		Status              string  `json:"status,omitempty"`
		StatusUpdateAllowed bool    `json:"status_update_allowed,omitempty"`
	}

	// MessagesResponse is the response of listing a pack's messages.
	MessagesResponse struct {
		ConversationStatus *ConversationStatus `json:"conversation_status,omitempty"`
		Paging             *Paging             `json:"paging,omitempty"`
		Messages           []Message           `json:"messages"`
		SellerMaxLength    int                 `json:"seller_max_message_length,omitempty"`
		BuyerMaxLength     int                 `json:"buyer_max_message_length,omitempty"`
	}

	// SendMessageRequest is the body for sending a post-sale message. Omit
	// Attachments when there is no file.
	SendMessageRequest struct {
		Text        string       `json:"text"`
		Attachments []string     `json:"attachments,omitempty"`
		From        MessageParty `json:"from"`
		To          MessageParty `json:"to"`
	}

	// MessagingService accesses post-sale messaging. All calls use tag=post_sale.
	MessagingService struct{ c *Client }
)

func postSaleQuery(extra url.Values) url.Values {
	q := url.Values{string(QueryParamTag): {"post_sale"}}
	for k, v := range extra {
		q[k] = v
	}
	return q
}

// PackMessages lists the messages of a pack/order conversation
// (GET /messages/packs/{pack_id}/sellers/{seller_id}?tag=post_sale). When the
// order has no pack, pass the order ID as packID. By default this marks messages
// as read; pass markAsRead=false to avoid that.
func (s *MessagingService) PackMessages(
	ctx context.Context,
	packID, sellerID int64,
	opts ListOptions,
	markAsRead bool,
) (*MessagesResponse, error) {
	extra := opts.Values()
	if !markAsRead {
		extra.Set(string(QueryParamMarkAsRead), "false")
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
