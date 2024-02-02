package db

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/giwty/switch-library-manager/settings"
)

type TitleAttributes struct {
	Id          string      `json:"id"`
	Name        string      `json:"name,omitempty"`
	Version     json.Number `json:"version,omitempty"`
	Region      string      `json:"region,omitempty"`
	ReleaseDate int         `json:"releaseDate,omitempty"`
	Publisher   string      `json:"publisher,omitempty"`
	IconUrl     string      `json:"iconUrl,omitempty"`
	Screenshots []string    `json:"screenshots,omitempty"`
	BannerUrl   string      `json:"bannerUrl,omitempty"`
	Description string      `json:"description,omitempty"`
	Size        int         `json:"size,omitempty"`
	IsDemo      bool        `json:"isDemo,omitempty"`
}

type BlacklistTitleAttributes struct {
	Id     string `json:"ID"`
	Name   string `json:"Name,omitempty"`
	Region string `json:"Region,omitempty"`
	Reason string `json:"Reason,omitempty"`
}

type SwitchTitle struct {
	Attributes TitleAttributes
	Updates    map[int]string
	Dlc        map[string]TitleAttributes
}

type SwitchTitlesDB struct {
	TitlesMap map[string]*SwitchTitle
}

var DemoTitles = make(map[string]TitleAttributes)
var BlacklistTitles = make(map[string]BlacklistTitleAttributes)

func CreateSwitchTitleDB(titlesFile, versionsFile io.Reader) (*SwitchTitlesDB, error) {
	baseFolder, err := os.Getwd()

	//parse the titles objects
	var titles = map[string]TitleAttributes{}
	err = decodeToJsonObject(titlesFile, &titles)
	if err != nil {
		return nil, err
	}

	//parse the blacklist objects
	var BlacklistTitlesFile *os.File = nil
	if _, err := os.Stat("blacklist.json"); err == nil {
		BlacklistTitlesFile, err = os.Open("blacklist.json")
		err = decodeToJsonObject(BlacklistTitlesFile, &BlacklistTitles)
		if err != nil {
			return nil, err
		}
	}

	//parse the whitelist objects
	var WhitelistTitles = map[string]TitleAttributes{}
	var WhitelistTitlesFile *os.File = nil
	if _, err := os.Stat("whitelist.json"); err == nil {
		WhitelistTitlesFile, err = os.Open("whitelist.json")
		err = decodeToJsonObject(WhitelistTitlesFile, &WhitelistTitles)
		if err != nil {
			return nil, err
		}
	}

	//parse the titles objects
	//titleID -> versionId-> release date
	var versions = map[string]map[int]string{}
	err = decodeToJsonObject(versionsFile, &versions)
	if err != nil {
		return nil, err
	}

	result := SwitchTitlesDB{TitlesMap: map[string]*SwitchTitle{}}
	for id, attr := range titles {
		if titles[id].IsDemo && settings.ReadSettings(baseFolder).IgnoreDemos {
			DemoTitles[id] = attr
			continue
		}

		if _, ok := BlacklistTitles[id]; ok {
			continue
		}

		id = strings.ToLower(id)

		//TitleAttributes id rules:
		//main TitleAttributes ends with 000
		//Updates ends with 800
		//Dlc adds 1 to 4th char starting from the right (always odd) and
		//    have a running counter (starting with 001) in the 3 last chars
		switchTitle := &SwitchTitle{Dlc: map[string]TitleAttributes{}}
		idPrefix := id[0 : len(id)-3]
		if !(strings.HasSuffix(id, "000") || strings.HasSuffix(id, "800")) {
			intVar, _ := strconv.ParseUint(id[len(id)-4:len(id)-3], 16, 64)
			h := fmt.Sprintf("%x", intVar-1)
			idPrefix = id[0:len(id)-4] + h
		}

		if t, ok := result.TitlesMap[idPrefix]; ok {
			switchTitle = t
		}
		result.TitlesMap[idPrefix] = switchTitle

		//process Updates
		if strings.HasSuffix(id, "800") {
			updates := versions[id[0:len(id)-3]+"000"]
			switchTitle.Updates = updates
			continue
		}

		//process main TitleAttributes
		if strings.HasSuffix(id, "000") {
			switchTitle.Attributes = attr
			continue
		}

		//not an update, and not main TitleAttributes, so treat it as a DLC
		switchTitle.Dlc[id] = attr

	}

	for WhitelistId, WhitelistAttr := range WhitelistTitles {
		WhitelistId = strings.ToLower(WhitelistId)
		WhitelistIdPrefix := WhitelistId[0 : len(WhitelistId)-3]
		WhitelistSwitchTitle := &SwitchTitle{Dlc: map[string]TitleAttributes{}}

		if _, ok := result.TitlesMap[WhitelistIdPrefix]; ok {
			WhitelistSwitchTitle.Updates = result.TitlesMap[WhitelistIdPrefix].Updates
			WhitelistSwitchTitle.Dlc = result.TitlesMap[WhitelistIdPrefix].Dlc
		}
		WhitelistSwitchTitle.Attributes = WhitelistAttr
		result.TitlesMap[WhitelistIdPrefix] = WhitelistSwitchTitle
	}

	return &result, nil
}
