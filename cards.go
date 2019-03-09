package glo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jackmcguire1/go-glo/internal/utils"
)

// Card contains information related to a card
type Card struct {
	ID                 string          `json:"id"`
	Name               string          `json:"name"`
	Position           int             `json:"position"`
	Description        *Description    `json:"Description"`
	BoardID            string          `json:"board_id"`
	ColumnID           string          `json:"column_id"`
	CreatedDate        string          `json:"created_date"`
	UpdatedDate        string          `json:"updated_date"`
	ArchivedDate       string          `json:"archived_date"`
	Assignees          []*PartialUser  `json:"assignees"`
	Labels             []*PartialLabel `json:"labels"`
	DueDate            string          `json:"due_date"`
	CommentCount       int             `json:"comment_count"`
	AttachmentCount    int             `json:"attachment_count"`
	CompletedTaskCount int             `json:"completed_task_count"`
	TotalTaskCount     int             `json:"total_task_count"`
	CreatedBy          *PartialUser    `json:"created_by"`
}

// Description contains information related to a card
type Description struct {
	Text        string       `json:"text"`
	CreatedDate string       `json:"created_date"`
	UpdatedDate string       `json:"updated_date"`
	CreatedBy   *PartialUser `json:"created_by"`
	UpdatedBy   *PartialUser `json:"updated_by"`
}

// CardsResp contains a list of cards and
// if further pagination calls are available
type CardsResp struct {
	Cards   []*Card `json:"cards"`
	HasMore bool    `json:"has_more"`
}

// CardsInput contains information used
// to create or edit a card
type CardsInput struct {
	Name        string                `json:"name"`
	Position    int                   `json:"position"`
	Description *MinimizedDescription `json:"description"`
	ColumnID    string                `json:"column_id"`
	Assignees   []*PartialUser        `json:"assignees"`
	Labels      []*PartialLabel       `json:"labels"`
	DueDate     string                `json:"due_date"`
}

var cardFields = []string{
	"archived_date",
	"assignees",
	"attachment_count",
	"board_id",
	"column_id",
	"comment_count",
	"completed_task_count",
	"created_by",
	"created_date",
	"due_date",
	"description",
	"labels",
	"name",
	"total_task_count",
	"updated_date",
}

// GetCards Get a list of Cards
// https://gloapi.gitkraken.com/v1/docs/#/Cards/get_boards__board_id__cards
func (a *Glo) GetCards(
	boardID string,
	page int,
	limit int,
	sortDesc bool,
	archived bool,
) (
	cardsResp *CardsResp,
	err error,
) {
	addr := fmt.Sprintf("%s/boards/%s/cards", a.BaseURI, boardID)

	q := utils.AddFields(cardFields)
	q.Set("page", fmt.Sprint(page))
	q.Set("per_page", fmt.Sprint(limit))

	if sortDesc {
		q.Set("sort", "desc")
	}

	if archived {
		q.Set("archived", "true")
	}

	cardsResp = &CardsResp{}
	resp, headers, err := a.jsonReq(http.MethodGet, addr, nil, q)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &cardsResp.Cards)
	if err != nil {
		return
	}

	hasMore, err := strconv.ParseBool(headers.Get("has-more"))
	if err != nil {
		return
	}
	cardsResp.HasMore = hasMore

	return
}

// CreateCard Creates a Card
// https://gloapi.gitkraken.com/v1/docs/#/Cards/post_boards__board_id__cards
func (a *Glo) CreateCard(
	boardID string,
	cardInput *CardsInput,
) (
	card *Card,
	err error,
) {
	addr := fmt.Sprintf("%s/boards/%s/cards", a.BaseURI, boardID)

	resp, _, err := a.jsonReq(http.MethodPost, addr, utils.ToRawMessage(cardInput), nil)
	if err != nil {
		return
	}

	card = &Card{}
	err = json.Unmarshal(resp, &card)

	return
}

// EditCard Edits a Card
// https://gloapi.gitkraken.com/v1/docs/#/Cards/post_boards__board_id__cards__card_id_
func (a *Glo) EditCard(
	boardID string,
	cardID string,
	cardInput *CardsInput,
) (
	card *Card,
	err error,
) {
	addr := fmt.Sprintf("%s/boards/%s/cards/%s", a.BaseURI, boardID, cardID)

	resp, _, err := a.jsonReq(http.MethodPost, addr, utils.ToRawMessage(cardInput), nil)
	if err != nil {
		return
	}

	card = &Card{}
	err = json.Unmarshal(resp, &card)

	return
}

// GetCard Get a Card by ID
// https://gloapi.gitkraken.com/v1/docs/#/Cards/get_boards__board_id__cards__card_id_
func (a *Glo) GetCard(
	boardID string,
	cardID string,
) (
	card *Card,
	err error,
) {
	addr := fmt.Sprintf("%s/boards/%s/cards/%s", a.BaseURI, boardID, cardID)

	q := utils.AddFields(cardFields)

	resp, _, err := a.jsonReq(http.MethodGet, addr, nil, q)
	if err != nil {
		return
	}

	card = &Card{}
	err = json.Unmarshal(resp, &card)

	return
}

// DeleteCard Deletes a card
// https://gloapi.gitkraken.com/v1/docs/#/Cards/delete_boards__board_id__cards__card_id_
func (a *Glo) DeleteCard(
	boardID string,
	cardID string,
) (
	err error,
) {
	addr := fmt.Sprintf("%s/boards/%s/cards/%s", a.BaseURI, boardID, cardID)

	_, _, err = a.jsonReq(http.MethodDelete, addr, nil, nil)

	return
}

// CardsByColumn Get a list of Cards by Column
// https://gloapi.gitkraken.com/v1/docs/#/Cards/get_boards__board_id__columns__column_id__cards
func (a *Glo) CardsByColumn(
	boardID string,
	columnID string,
	page int,
	limit int,
	sortDesc bool,
	archived bool,
) (
	cardsResp *CardsResp,
	err error,
) {
	addr := fmt.Sprintf("%s/boards/%s/columns/%s/cards/", a.BaseURI, boardID, columnID)

	q := utils.AddFields(cardFields)

	q.Set("page", fmt.Sprint(page))
	q.Set("per_page", fmt.Sprint(limit))
	if sortDesc {
		q.Set("sort", "desc")
	}
	if archived {
		q.Set("archived", "true")
	}

	resp, headers, err := a.jsonReq(http.MethodGet, addr, nil, q)
	if err != nil {
		return
	}

	cardsResp = &CardsResp{}
	err = json.Unmarshal(resp, &cardsResp.Cards)
	if err != nil {
		return
	}

	hasMore, err := strconv.ParseBool(headers.Get("has-more"))
	if err != nil {
		return
	}
	cardsResp.HasMore = hasMore

	return
}
