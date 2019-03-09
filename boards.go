package glo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jackmcguire1/go-glo/internal/utils"
)

// Board contains the information related to a board
type Board struct {
	ID              string         `json:"id"`
	Name            string         `json:"name"`
	Columns         []*Column      `json:"columns"`
	ArchivedColumns []*Column      `json:"archived_columns"`
	InvitedMembers  []*BoardMember `json:"invited_members"`
	Members         []*BoardMember `json:"members"`
	ArchivedDate    string         `json:"archived_date"`
	CreatedDate     string         `json:"created_date"`
	CreatedBy       *PartialUser   `json:"created_by"`
	Labels          []*Label       `json:"labels"`
}

// BoardMember contains information related to a Board Member
type BoardMember struct {
	ID       string `json:"id"`
	Role     string `json:"role"`
	Username string `json:"username,omitempty"`
}

type BoardInput struct {
	Name string `json:"name"`
}

// BoardsResp contains a list of boards and
// if further pagination calls are available
type BoardsResp struct {
	HasMore bool     `json:"has_more"`
	Boards  []*Board `json:"boards"`
}

var boardFields = []string{
	"archived_columns",
	"archived_date",
	"columns",
	"created_by",
	"created_date",
	"invited_members",
	"labels",
	"members",
	"name",
}

// CreateBoard Creates a Board
// https://gloapi.gitkraken.com/v1/docs/#/Boards/post_boards
func (a *Glo) CreateBoard(
	input *BoardInput,
) (
	board *Board,
	err error,
) {

	addr := fmt.Sprintf("%s/boards", a.BaseURI)

	resp, _, err := a.jsonReq(http.MethodPost, addr, utils.ToRawMessage(input), nil)
	if err != nil {
		return
	}

	board = &Board{}
	err = json.Unmarshal(resp, &board)

	return
}

// EditBoard Edits a Board
// https://gloapi.gitkraken.com/v1/docs/#/Boards/post_boards__board_id_
func (a *Glo) EditBoard(
	boardID string,
	input *BoardInput,
) (
	board *Board,
	err error,
) {

	addr := fmt.Sprintf("%s/boards/%s", a.BaseURI, boardID)

	resp, _, err := a.jsonReq(http.MethodPost, addr, utils.ToRawMessage(input), nil)
	if err != nil {
		return
	}

	board = &Board{}
	err = json.Unmarshal(resp, &board)

	return
}

// GetBoards Get a list of Boards
// https://gloapi.gitkraken.com/v1/docs/#/Boards/get_boards
func (a *Glo) GetBoards(
	page int,
	limit int,
	sortDesc bool,
	archived bool,
) (
	boardsResp *BoardsResp,
	err error,
) {

	addr := fmt.Sprintf("%s/boards", a.BaseURI)

	q := utils.AddFields(boardFields)
	q.Set("page", fmt.Sprint(page))
	q.Set("per_page", fmt.Sprint(limit))

	if sortDesc {
		q.Set("sort", "desc")
	}

	if archived {
		q.Set("archived", "true")
	}

	boardsResp = &BoardsResp{}
	resp, headers, err := a.jsonReq(http.MethodGet, addr, nil, q)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &boardsResp.Boards)
	if err != nil {
		return
	}

	hasMore, err := strconv.ParseBool(headers.Get("has-more"))
	if err != nil {
		return
	}
	boardsResp.HasMore = hasMore

	return
}

// GetBoard Get a Board by ID
// https://gloapi.gitkraken.com/v1/docs/#/Boards/get_boards__board_id_
func (a *Glo) GetBoard(
	boardID string,
) (
	board *Board,
	err error,
) {
	addr := fmt.Sprintf("%s/boards/%s", a.BaseURI, boardID)

	q := utils.AddFields(boardFields)

	resp, _, err := a.jsonReq(http.MethodGet, addr, nil, q)
	if err != nil {
		return
	}

	board = &Board{}
	err = json.Unmarshal(resp, &board)

	return
}
