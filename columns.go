package glo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jackmcguire1/go-glo/internal/utils"
)

// Column contains inforamtion related to a Column
type Column struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Position     int          `json:"position"`
	ArchivedDate string       `json:"archived_date"`
	CreatedDate  string       `json:"created_date"`
	CreatedBy    *PartialUser `json:"created_by"`
}

// ColumnInput contains information used
// to create or edit a column
type ColumnInput struct {
	Name     string `json:"name"`
	Position int    `json:"position"`
}

// CreateColumn Creates a Column
// https://gloapi.gitkraken.com/v1/docs/#/Columns/post_boards__board_id__columns
func (a *Glo) CreateColumn(
	boardID string,
	columnInput *ColumnInput,
) (
	col *Column,
	err error,
) {
	addr := fmt.Sprintf("%s/boards/%s/columns", a.BaseURI, boardID)

	resp, _, err := a.jsonReq(http.MethodPost, addr, utils.ToRawMessage(columnInput), nil)
	if err != nil {
		return
	}

	col = &Column{}
	err = json.Unmarshal(resp, &col)

	return
}

// EditColumn Edits a Column
// https://gloapi.gitkraken.com/v1/docs/#/Columns/post_boards__board_id__columns__column_id_
func (a *Glo) EditColumn(
	boardID,
	columnID string,
	columnInput *ColumnInput,
) (
	col *Column,
	err error,
) {
	addr := fmt.Sprintf("%s/boards/%s/columns/%s", a.BaseURI, boardID, columnID)

	resp, _, err := a.jsonReq(http.MethodPost, addr, utils.ToRawMessage(columnInput), nil)
	if err != nil {
		return
	}

	col = &Column{}
	err = json.Unmarshal(resp, &col)

	return
}

// DeteleColumn Deletes a Column
// https://gloapi.gitkraken.com/v1/docs/#/Columns/delete_boards__board_id__columns__column_id_
func (a *Glo) DeteleColumn(boardID, columnID string) (err error) {
	addr := fmt.Sprintf("%s/boards/%s/columns/%s", a.BaseURI, boardID, columnID)

	_, _, err = a.jsonReq(http.MethodDelete, addr, nil, nil)

	return
}
