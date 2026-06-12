package mercadolibre

import (
	"context"
	"net/url"
)

type (
	// Question is a buyer question on an item (GET /questions/{id}?api_version=4).
	Question struct {
		DateCreated     Time          `json:"date_created"`
		LastUpdated     Time          `json:"last_updated,omitempty"`
		From            *QuestionFrom `json:"from,omitempty"`
		Answer          *Answer       `json:"answer,omitempty"`
		AppID           *int64        `json:"app_id,omitempty"`
		Status          string        `json:"status"`
		ItemID          string        `json:"item_id"`
		Text            string        `json:"text"`
		SellerID        int64         `json:"seller_id,omitempty"`
		BuyerID         int64         `json:"buyer_id,omitempty"`
		ID              int64         `json:"id"`
		Hold            bool          `json:"hold,omitempty"`
		DeletedFromList bool          `json:"deleted_from_listing,omitempty"`
		SuspectedSpam   bool          `json:"suspected_spam,omitempty"`
	}

	// Answer is a seller's answer to a question.
	Answer struct {
		DateCreated Time   `json:"date_created,omitempty"`
		Text        string `json:"text"`
		Status      string `json:"status"`
	}

	// QuestionFrom is the asker; vehicle/lead questions include contact details.
	QuestionFrom struct {
		Phone             *Phone `json:"phone,omitempty"`
		Email             string `json:"email,omitempty"`
		FirstName         string `json:"first_name,omitempty"`
		LastName          string `json:"last_name,omitempty"`
		ID                int64  `json:"id"`
		AnsweredQuestions int    `json:"answered_questions,omitempty"`
	}

	// QuestionSearchResponse is the response of /questions/search. Note offset/limit
	// live under filters in this resource.
	QuestionSearchResponse struct {
		Filters          map[string]any `json:"filters,omitempty"`
		Questions        []Question     `json:"questions"`
		AvailableFilters []Filter       `json:"available_filters,omitempty"`
		AvailableSorts   []string       `json:"available_sorts,omitempty"`
		Total            int            `json:"total"`
		Limit            int            `json:"limit"`
	}

	// QuestionsService accesses questions and answers. All reads pin api_version=4.
	QuestionsService struct{ c *Client }
)

func withAPIv4(params url.Values) url.Values {
	if params == nil {
		params = url.Values{}
	}
	params.Set(string(QueryParamAPIVersion), "4")
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
		Text       string `json:"text"`
		QuestionID int64  `json:"question_id"`
	}{Text: text, QuestionID: questionID}
	return Post[*Answer](ctx, s.c, EPAnswers, body)
}

// Delete deletes a question (DELETE /questions/{id}).
func (s *QuestionsService) Delete(ctx context.Context, questionID int64) error {
	_, err := Delete[struct{}](ctx, s.c, EPQuestionByID, questionID)
	return err
}

// QuestionResponseTime is the seller's average question-response time.
type QuestionResponseTime struct {
	LastUpdated  Time    `json:"last_updated,omitempty"`
	AverageUnits string  `json:"average_units,omitempty"`
	AverageTime  float64 `json:"average_time"`
	UserID       int64   `json:"user_id"`
}

// ResponseTime returns a seller's average question-response time
// (GET /users/{id}/questions/response_time).
func (s *QuestionsService) ResponseTime(ctx context.Context, userID int64) (*QuestionResponseTime, error) {
	return Get[*QuestionResponseTime](ctx, s.c, EPUserQuestionResponseTime, userID)
}
