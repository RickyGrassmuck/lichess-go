package lichess

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) Tournament(ctx context.Context, id string, params *TournamentParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewTournamentRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewTournamentRequest generates requests for Tournament
func NewTournamentRequest(server string, id string, params *TournamentParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/tournament/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.Page != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) ApiTeamArena(ctx context.Context, teamId string, params *ApiTeamArenaParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewApiTeamArenaRequest(c.Server, teamId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// ApiTeamArenaWithResponse request returning *ApiTeamArenaResponse
func (c *ClientWithResponses) ApiTeamArenaWithResponse(ctx context.Context, teamId string, params *ApiTeamArenaParams, reqEditors ...RequestEditorFn) (*ApiTeamArenaResponse, error) {
	rsp, err := c.ApiTeamArena(ctx, teamId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseApiTeamArenaResponse(rsp)
}

// NewApiTeamArenaRequest generates requests for ApiTeamArena
func NewApiTeamArenaRequest(server string, teamId string, params *ApiTeamArenaParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "teamId", runtime.ParamLocationPath, teamId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/team/%s/arena", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.Max != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "max", runtime.ParamLocationQuery, *params.Max); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) ApiTournament(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewApiTournamentRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewApiTournamentRequest generates requests for ApiTournament
func NewApiTournamentRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/tournament")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) ApiTournamentPostWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewApiTournamentPostRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewApiTournamentPostRequestWithBody generates requests for ApiTournamentPost with any type of body
func NewApiTournamentPostRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/tournament")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) ApiTournamentTeamBattlePostWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewApiTournamentTeamBattlePostRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewApiTournamentTeamBattlePostRequestWithBody generates requests for ApiTournamentTeamBattlePost with any type of body
func NewApiTournamentTeamBattlePostRequestWithBody(server string, id string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/tournament/team-battle/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}



func (c *Client) ApiTournamentUpdateWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewApiTournamentUpdateRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewApiTournamentUpdateRequestWithBody generates requests for ApiTournamentUpdate with any type of body
func NewApiTournamentUpdateRequestWithBody(server string, id string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/tournament/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) GamesByTournament(ctx context.Context, id string, params *GamesByTournamentParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGamesByTournamentRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGamesByTournamentRequest generates requests for GamesByTournament
func NewGamesByTournamentRequest(server string, id string, params *GamesByTournamentParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/tournament/%s/games", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.Player != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "player", runtime.ParamLocationQuery, *params.Player); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Moves != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "moves", runtime.ParamLocationQuery, *params.Moves); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.PgnInJson != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "pgnInJson", runtime.ParamLocationQuery, *params.PgnInJson); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Tags != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "tags", runtime.ParamLocationQuery, *params.Tags); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Clocks != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "clocks", runtime.ParamLocationQuery, *params.Clocks); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Evals != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "evals", runtime.ParamLocationQuery, *params.Evals); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Opening != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "opening", runtime.ParamLocationQuery, *params.Opening); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) ApiTournamentJoinWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewApiTournamentJoinRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewApiTournamentJoinRequestWithBody generates requests for ApiTournamentJoin with any type of body
func NewApiTournamentJoinRequestWithBody(server string, id string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/tournament/%s/join", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) ResultsByTournament(ctx context.Context, id string, params *ResultsByTournamentParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewResultsByTournamentRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewResultsByTournamentRequest generates requests for ResultsByTournament
func NewResultsByTournamentRequest(server string, id string, params *ResultsByTournamentParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/tournament/%s/results", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.Nb != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "nb", runtime.ParamLocationQuery, *params.Nb); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) TeamsByTournament(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewTeamsByTournamentRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewTeamsByTournamentRequest generates requests for TeamsByTournament
func NewTeamsByTournamentRequest(server string, id string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/tournament/%s/teams", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) ApiTournamentTerminate(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewApiTournamentTerminateRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewApiTournamentTerminateRequest generates requests for ApiTournamentTerminate
func NewApiTournamentTerminateRequest(server string, id string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/tournament/%s/terminate", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) ApiTournamentWithdraw(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewApiTournamentWithdrawRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewApiTournamentWithdrawRequest generates requests for ApiTournamentWithdraw
func NewApiTournamentWithdrawRequest(server string, id string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/tournament/%s/withdraw", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}


// ApiTournamentWithResponse request returning *ApiTournamentResponse
func (c *ClientWithResponses) ApiTournamentWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ApiTournamentResponse, error) {
	rsp, err := c.ApiTournament(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseApiTournamentResponse(rsp)
}

// ApiTournamentPostWithBodyWithResponse request with arbitrary body returning *ApiTournamentPostResponse
func (c *ClientWithResponses) ApiTournamentPostWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*ApiTournamentPostResponse, error) {
	rsp, err := c.ApiTournamentPostWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseApiTournamentPostResponse(rsp)
}

// ApiTournamentTeamBattlePostWithBodyWithResponse request with arbitrary body returning *ApiTournamentTeamBattlePostResponse
func (c *ClientWithResponses) ApiTournamentTeamBattlePostWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*ApiTournamentTeamBattlePostResponse, error) {
	rsp, err := c.ApiTournamentTeamBattlePostWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseApiTournamentTeamBattlePostResponse(rsp)
}

// TournamentWithResponse request returning *TournamentResponse
func (c *ClientWithResponses) TournamentWithResponse(ctx context.Context, id string, params *TournamentParams, reqEditors ...RequestEditorFn) (*TournamentResponse, error) {
	rsp, err := c.Tournament(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseTournamentResponse(rsp)
}

// ApiTournamentUpdateWithBodyWithResponse request with arbitrary body returning *ApiTournamentUpdateResponse
func (c *ClientWithResponses) ApiTournamentUpdateWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*ApiTournamentUpdateResponse, error) {
	rsp, err := c.ApiTournamentUpdateWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseApiTournamentUpdateResponse(rsp)
}

// GamesByTournamentWithResponse request returning *GamesByTournamentResponse
func (c *ClientWithResponses) GamesByTournamentWithResponse(ctx context.Context, id string, params *GamesByTournamentParams, reqEditors ...RequestEditorFn) (*GamesByTournamentResponse, error) {
	rsp, err := c.GamesByTournament(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGamesByTournamentResponse(rsp)
}

// ApiTournamentJoinWithBodyWithResponse request with arbitrary body returning *ApiTournamentJoinResponse
func (c *ClientWithResponses) ApiTournamentJoinWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*ApiTournamentJoinResponse, error) {
	rsp, err := c.ApiTournamentJoinWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseApiTournamentJoinResponse(rsp)
}

// ResultsByTournamentWithResponse request returning *ResultsByTournamentResponse
func (c *ClientWithResponses) ResultsByTournamentWithResponse(ctx context.Context, id string, params *ResultsByTournamentParams, reqEditors ...RequestEditorFn) (*ResultsByTournamentResponse, error) {
	rsp, err := c.ResultsByTournament(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseResultsByTournamentResponse(rsp)
}

// TeamsByTournamentWithResponse request returning *TeamsByTournamentResponse
func (c *ClientWithResponses) TeamsByTournamentWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*TeamsByTournamentResponse, error) {
	rsp, err := c.TeamsByTournament(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseTeamsByTournamentResponse(rsp)
}

// ApiTournamentTerminateWithResponse request returning *ApiTournamentTerminateResponse
func (c *ClientWithResponses) ApiTournamentTerminateWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*ApiTournamentTerminateResponse, error) {
	rsp, err := c.ApiTournamentTerminate(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseApiTournamentTerminateResponse(rsp)
}

// ApiTournamentWithdrawWithResponse request returning *ApiTournamentWithdrawResponse
func (c *ClientWithResponses) ApiTournamentWithdrawWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*ApiTournamentWithdrawResponse, error) {
	rsp, err := c.ApiTournamentWithdraw(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseApiTournamentWithdrawResponse(rsp)
}

// ParseApiTournamentResponse parses an HTTP response from a ApiTournamentWithResponse call
func ParseApiTournamentResponse(rsp *http.Response) (*ApiTournamentResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ApiTournamentResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ArenaTournaments
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseApiTournamentPostResponse parses an HTTP response from a ApiTournamentPostWithResponse call
func ParseApiTournamentPostResponse(rsp *http.Response) (*ApiTournamentPostResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ApiTournamentPostResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ArenaTournament
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}

// ParseApiTournamentTeamBattlePostResponse parses an HTTP response from a ApiTournamentTeamBattlePostWithResponse call
func ParseApiTournamentTeamBattlePostResponse(rsp *http.Response) (*ApiTournamentTeamBattlePostResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ApiTournamentTeamBattlePostResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ArenaTournament
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}

// ParseTournamentResponse parses an HTTP response from a TournamentWithResponse call
func ParseTournamentResponse(rsp *http.Response) (*TournamentResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &TournamentResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ArenaTournament
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseApiTournamentUpdateResponse parses an HTTP response from a ApiTournamentUpdateWithResponse call
func ParseApiTournamentUpdateResponse(rsp *http.Response) (*ApiTournamentUpdateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ApiTournamentUpdateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ArenaTournament
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}

// ParseGamesByTournamentResponse parses an HTTP response from a GamesByTournamentWithResponse call
func ParseGamesByTournamentResponse(rsp *http.Response) (*GamesByTournamentResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GamesByTournamentResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseApiTournamentJoinResponse parses an HTTP response from a ApiTournamentJoinWithResponse call
func ParseApiTournamentJoinResponse(rsp *http.Response) (*ApiTournamentJoinResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ApiTournamentJoinResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Ok
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}

// ParseResultsByTournamentResponse parses an HTTP response from a ResultsByTournamentWithResponse call
func ParseResultsByTournamentResponse(rsp *http.Response) (*ResultsByTournamentResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ResultsByTournamentResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}


// ParseApiTeamArenaResponse parses an HTTP response from a ApiTeamArenaWithResponse call
func ParseApiTeamArenaResponse(rsp *http.Response) (*ApiTeamArenaResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ApiTeamArenaResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []ArenaTournament
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseTeamsByTournamentResponse parses an HTTP response from a TeamsByTournamentWithResponse call
func ParseTeamsByTournamentResponse(rsp *http.Response) (*TeamsByTournamentResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &TeamsByTournamentResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseApiTournamentTerminateResponse parses an HTTP response from a ApiTournamentTerminateWithResponse call
func ParseApiTournamentTerminateResponse(rsp *http.Response) (*ApiTournamentTerminateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ApiTournamentTerminateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Ok
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}

// ParseApiTournamentWithdrawResponse parses an HTTP response from a ApiTournamentWithdrawWithResponse call
func ParseApiTournamentWithdrawResponse(rsp *http.Response) (*ApiTournamentWithdrawResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ApiTournamentWithdrawResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Ok
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}

type ApiTournamentResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ArenaTournaments
}

// Status returns HTTPResponse.Status
func (r ApiTournamentResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ApiTournamentResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ApiTournamentPostResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ArenaTournament
	JSON400      *Error
}

// Status returns HTTPResponse.Status
func (r ApiTournamentPostResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ApiTournamentPostResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ApiTournamentTeamBattlePostResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ArenaTournament
	JSON400      *Error
}

// Status returns HTTPResponse.Status
func (r ApiTournamentTeamBattlePostResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ApiTournamentTeamBattlePostResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
type ApiTournamentUpdateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ArenaTournament
	JSON400      *Error
}

// Status returns HTTPResponse.Status
func (r ApiTournamentUpdateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ApiTournamentUpdateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GamesByTournamentResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r GamesByTournamentResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GamesByTournamentResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ApiTournamentJoinResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Ok
	JSON400      *Error
}

// Status returns HTTPResponse.Status
func (r ApiTournamentJoinResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ApiTournamentJoinResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ResultsByTournamentResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r ResultsByTournamentResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ResultsByTournamentResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type TeamsByTournamentResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *interface{}
}

// Status returns HTTPResponse.Status
func (r TeamsByTournamentResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r TeamsByTournamentResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ApiTournamentTerminateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Ok
	JSON400      *Error
}

// Status returns HTTPResponse.Status
func (r ApiTournamentTerminateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ApiTournamentTerminateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ApiTournamentWithdrawResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Ok
	JSON400      *Error
}

// Status returns HTTPResponse.Status
func (r ApiTournamentWithdrawResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ApiTournamentWithdrawResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
type TournamentResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ArenaTournament
}

type ApiTeamArenaResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]ArenaTournament
}

// Status returns HTTPResponse.Status
func (r ApiTeamArenaResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ApiTeamArenaResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// Status returns HTTPResponse.Status
func (r TournamentResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r TournamentResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}