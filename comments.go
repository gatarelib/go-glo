package glo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jackmcguire1/go-glo/internal/utils"
)

// Comment contains information related to a Comment.
type Comment struct {
	ID          string       `json:"id"`
	CardID      string       `json:"card_id"`
	BoardID     string       `json:"board_id"`
	CreatedDate string       `json:"created_date"`
	UpdatedDate string       `json:"updated_date"`
	CreatedBy   *PartialUser `json:"created_by"`
	UpdatedBy   *PartialUser `json:"updated_by"`
	Text        string       `json:"text"`
}

// CommentsResp contains a list of comments and
// if further pagination calls are available .
type CommentsResp struct {
	HasMore  bool
	Comments []*Comment
}

// CommentInput contains information used to
// create a comment
type CommentInput struct {
	Text string `json:"text"`
}

var commentFields = []string{
	"board_id",
	"card_id",
	"created_date",
	"created_by",
	"updated_by",
	"updated_date",
	"text",
}

// GetComments Get Comments for a Card
// https://gloapi.gitkraken.com/v1/docs/#/Comments/get_boards__board_id__cards__card_id__comments
func (a *Glo) GetComments(
	boardID string,
	cardID string,
	page int,
	limit int,
	sortDesc bool,
) (
	commentsResp *CommentsResp,
	err error,
) {
	addr := fmt.Sprintf("%s/boards/%s/cards/%s/comments", a.BaseURI, boardID, cardID)

	q := utils.AddFields(commentFields)
	q.Set("page", fmt.Sprint(page))
	q.Set("per_page", fmt.Sprint(limit))

	if sortDesc {
		q.Set("sort", "desc")
	}

	resp, headers, err := a.jsonReq(http.MethodGet, addr, nil, q)
	if err != nil {
		return
	}

	commentsResp = &CommentsResp{}
	err = json.Unmarshal(resp, commentsResp)
	if err != nil {
		return
	}

	hasMore, err := strconv.ParseBool(headers.Get("has-more"))
	if err != nil {
		return
	}
	commentsResp.HasMore = hasMore

	return
}

// CreateComment Creates Comment
// https://gloapi.gitkraken.com/v1/docs/#/Comments/post_boards__board_id__cards__card_id__comments
func (a *Glo) CreateComment(
	boardID string,
	cardID string,
	commentInput *Comment,
) (
	comment *Comment,
	err error,
) {
	addr := fmt.Sprintf("%s/boards/%s/cards/%s/comments", a.BaseURI, boardID, cardID)

	resp, _, err := a.jsonReq(http.MethodPost, addr, utils.ToRawMessage(commentInput), nil)
	if err != nil {
		return
	}

	comment = &Comment{}
	err = json.Unmarshal(resp, comment)
	return
}

// EditComment Edits Comment
// https://gloapi.gitkraken.com/v1/docs/#/Comments/post_boards__board_id__cards__card_id__comments__comment_id_
func (a *Glo) EditComment(
	boardID string,
	cardID string,
	commentID string,
	input *CommentInput,
) (
	comment *Comment,
	err error,
) {
	addr := fmt.Sprintf("%s/boards/%s/cards/%s/comments/%s", a.BaseURI, boardID, cardID, commentID)

	resp, _, err := a.jsonReq(http.MethodPost, addr, utils.ToRawMessage(input), nil)
	if err != nil {
		return
	}

	comment = &Comment{}
	err = json.Unmarshal(resp, comment)

	return
}

// DeleteComment Deletes a Comment
// https://gloapi.gitkraken.com/v1/docs/#/Comments/delete_boards__board_id__cards__card_id__comments__comment_id_
func (a *Glo) DeleteComment(
	boardID string,
	cardID string,
	commentID string,
) (
	err error,
) {
	addr := fmt.Sprintf("%s/boards/%s/cards/%s/comments/%s", a.BaseURI, boardID, cardID, commentID)

	_, _, err = a.jsonReq(http.MethodDelete, addr, nil, nil)

	return
}
