package mercadolibre

import (
	"context"
	"net/url"
)

// Question is a buyer question on an item (GET /questions/{id}?api_version=4).
type Question struct {
	ID                int64         `json:"id"`
	SellerID          int64         `json:"seller_id,omitempty"`
	BuyerID           int64         `json:"buyer_id,omitempty"`
	ItemID            string        `json:"item_id"`
	Status            string        `json:"status"`
	Text              string        `json:"text"`
	DateCreated       Time          `json:"date_created"`
	LastUpdated       Time          `json:"last_updated,omitempty"`
	Hold              bool          `json:"hold,omitempty"`
	DeletedFromList   bool          `json:"deleted_from_listing,omitempty"`
	SuspectedSpam     bool          `json:"suspected_spam,omitempty"`
	AppID             *int64        `json:"app_id,omitempty"`
	Answer            *Answer       `json:"answer,omitempty"`
	From              *QuestionFrom `json:"from,omitempty"`
}

// Answer is a seller's answer to a question.
type Answer struct {
	Text        string `json:"text"`
	Status      string `json:"status"`
	DateCreated Time   `json:"date_created,omitempty"`
}

// QuestionFrom is the asker; vehicle/lead questions include contact details.
type QuestionFrom struct {
	ID               int64  `json:"id"`
	AnsweredQuestions int   `json:"answered_questions,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	Email            string `json:"email,omitempty"`
	Phone            *Phone `json:"phone,omitempty"`
}

// QuestionSearchResponse is the response of /questions/search. Note offset/limit
// live under filters in this resource.
type QuestionSearchResponse struct {
	Total           int            `json:"total"`
	Limit           int            `json:"limit"`
	Questions       []Question     `json:"questions"`
	Filters         map[string]any `json:"filters,omitempty"`
	AvailableFilters []Filter      `json:"available_filters,omitempty"`
	AvailableSorts  []string       `json:"available_sorts,omitempty"`
}

// QuestionsService accesses questions and answers. All reads pin api_version=4.
type QuestionsService struct{ c *Client }

func withAPIv4(params url.Values) url.Values {
	if params == nil {
		params = url.Values{}
	}
	params.Set("api_version", "4")
	return params
}

// Get returns a question by ID (GET /questions/{id}?api_version=4).
func (s *QuestionsService) Get(ctx context.Context, questionID int64) (*Question, error) {
	return GetQ[*Question](ctx, s.c, EPQuestionByID, withAPIv4(nil), questionID)
}

// Search lists questions (GET /questions/search). Typical params: item / item_id,
// seller_id, from, status, sort_fields, sort_types.
func (s *QuestionsService) Search(ctx context.Context, params url.Values) (*QuestionSearchResponse, error) {
	return GetQ[*QuestionSearchResponse](ctx, s.c, EPQuestionsSearch, withAPIv4(params))
}

// Ask posts a question on an item (POST /questions).
func (s *QuestionsService) Ask(ctx context.Context, itemID, text string) (*Question, error) {
	body := struct {
		ItemID string `json:"item_id"`
		Text   string `json:"text"`
	}{itemID, text}
	return Post[*Question](ctx, s.c, EPQuestions, body)
}

// Answer posts an answer to a question (POST /answers).
func (s *QuestionsService) Answer(ctx context.Context, questionID int64, text string) (*Answer, error) {
	body := struct {
		QuestionID int64  `json:"question_id"`
		Text       string `json:"text"`
	}{questionID, text}
	return Post[*Answer](ctx, s.c, EPAnswers, body)
}

// Delete deletes a question (DELETE /questions/{id}).
func (s *QuestionsService) Delete(ctx context.Context, questionID int64) error {
	_, err := Delete[struct{}](ctx, s.c, EPQuestionByID, questionID)
	return err
}

// QuestionResponseTime is the seller's average question-response time.
type QuestionResponseTime struct {
	UserID       int64   `json:"user_id"`
	AverageTime  float64 `json:"average_time"`
	AverageUnits string  `json:"average_units,omitempty"` // "hours", "minutes"
	LastUpdated  Time    `json:"last_updated,omitempty"`
}

// ResponseTime returns a seller's average question-response time
// (GET /users/{id}/questions/response_time).
func (s *QuestionsService) ResponseTime(ctx context.Context, userID int64) (*QuestionResponseTime, error) {
	return Get[*QuestionResponseTime](ctx, s.c, EPUserQuestionResponseTime, userID)
}
