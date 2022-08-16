package lichess

import (
	"io"
	"context"
	"net/http"
)

func (c *Client) ApiUser(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewApiUserRequest(c.Server, username)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ApiUserActivity(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewApiUserActivityRequest(c.Server, username)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ApiUserCurrentGame(ctx context.Context, username string, params *ApiUserCurrentGameParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewApiUserCurrentGameRequest(c.Server, username, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ApiUserPerf(ctx context.Context, username string, perf PerfType, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewApiUserPerfRequest(c.Server, username, perf)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ApiUserRatingHistory(ctx context.Context, username string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewApiUserRatingHistoryRequest(c.Server, username)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ApiUserNameTournamentCreated(ctx context.Context, username string, params *ApiUserNameTournamentCreatedParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewApiUserNameTournamentCreatedRequest(c.Server, username, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ApiUsersWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewApiUsersRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ApiUsersStatus(ctx context.Context, params *ApiUsersStatusParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewApiUsersStatusRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
