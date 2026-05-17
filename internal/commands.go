package internal

import (
	"encoding/json"
	"errors"
	"fmt"
)

type FetchData struct {
	StatusCode int
	Body string
	Meta string
	RedirectCount int
}

type ParseData struct {
	StatusCode int
	Meta string
	Body string
	Tokens []any
	RedirectCount int
}

type TabData struct {
	Id string
	Title string
	Url string
	Content any
}

func Fetch(url string) (FetchData, error) {
	req := fmt.Sprintf("fetch %v", url);
	resp, err := Request(req);
	if err != nil {
		return FetchData{}, err;
	}

	if !resp.Ok {
		return FetchData{}, errors.New(resp.Error);
	}

	var data FetchData;
	if err := json.Unmarshal(resp.Data, &data); err != nil {
		return FetchData{}, err;
	}

	return data, nil;
}

func Parse(url string) (ParseData, error) {
	req := fmt.Sprintf("parse %v", url);
	resp, err := Request(req);
	if err != nil {
		return ParseData{}, err;
	}

	if !resp.Ok {
		return ParseData{}, errors.New(resp.Error);
	}

	var data ParseData;
	if err := json.Unmarshal(resp.Data, &data); err != nil {
		return ParseData{}, err;
	}

	return data, nil;
}

func TabNew(url string) (TabData, error) {
	req := fmt.Sprintf("tab new %v", url);
	resp, err := Request(req);
	if err != nil {
		return TabData{}, err;
	}

	if !resp.Ok {
		return TabData{}, errors.New(resp.Error);
	}

	var data TabData;
	if err := json.Unmarshal(resp.Data, &data); err != nil {
		return TabData{}, err;
	}

	return data, nil;
}

func TabPut(url string) (TabData, error) {
	req := fmt.Sprintf("tab put %v", url);
	resp, err := Request(req);
	if err != nil {
		return TabData{}, err;
	}

	if !resp.Ok {
		return TabData{}, errors.New(resp.Error);
	}

	var data TabData;
	if err := json.Unmarshal(resp.Data, &data); err != nil {
		return TabData{}, err;
	}

	return data, nil;
}

func TabDel(id int) error {
	req := fmt.Sprintf("tab del %v", id);
	resp, err := Request(req);
	if err != nil {
		return err;
	}

	if !resp.Ok {
		return errors.New(resp.Error);
	}
	return nil;
}

func TabSel(id int) (TabData, error) {
	req := fmt.Sprintf("tab sel %v", id);
	resp, err := Request(req);
	if err != nil {
		return TabData{}, err;
	}

	if !resp.Ok {
		return TabData{}, errors.New(resp.Error);
	}

	var data TabData;
	if err := json.Unmarshal(resp.Data, &data); err != nil {
		return TabData{}, err;
	}

	return data, nil;
}

 func TabAll() ([]TabData, error) {
	resp, err := Request("tab all");
	if err != nil {
		return nil, err;
	}

	if !resp.Ok {
		return nil, errors.New(resp.Error);
	}

	var data []TabData;
	if err := json.Unmarshal(resp.Data, &data); err != nil {
		return nil, err;
	}

	return data, nil;
}

func TabGet() (TabData, error) {
	resp, err := Request("tab get");
	if err != nil {
		return TabData{}, err;
	}

	if !resp.Ok {
		return TabData{}, errors.New(resp.Error);
	}

	var data TabData;
	if err := json.Unmarshal(resp.Data, &data); err != nil {
		return TabData{}, err;
	}

	return data, nil;
}
