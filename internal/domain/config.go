// Copyright (c) 2021 - 2024, Ludvig Lundgren and the autobrr contributors.
// Code is modified for use with seasonpackarr
// SPDX-License-Identifier: GPL-2.0-or-later

package domain

type Client struct {
	Host          string `yaml:"host"`
	Port          int    `yaml:"port"`
	Username      string `yaml:"username"`
	Password      string `yaml:"password"`
	PreImportPath string `yaml:"preImportPath"`
}

type FuzzyMatching struct {
	SkipRepackCompare  bool `yaml:"skipRepackCompare"`
	SimplifyHdrCompare bool `yaml:"simplifyHdrCompare"`
}

type Notifications struct {
	NotificationLevel []string `yaml:"notificationLevel"`
	Discord           string   `yaml:"discord"`
	// Notifiarr string `yaml:"notifiarr"`
	// Shoutrrr  string `yaml:"shoutrrr"`
}

type Config struct {
	Version            string
	ConfigPath         string
	Host               string             `yaml:"host"`
	Port               int                `yaml:"port"`
	Clients            map[string]*Client `yaml:"clients"`
	LogPath            string             `yaml:"logPath"`
	LogLevel           string             `yaml:"logLevel"`
	LogMaxSize         int                `yaml:"logMaxSize"`
	LogMaxBackups      int                `yaml:"logMaxBackups"`
	SmartMode          bool               `yaml:"smartMode"`
	SmartModeThreshold float32            `yaml:"smartModeThreshold"`
	ParseTorrentFile   bool               `yaml:"parseTorrentFile"`
	FuzzyMatching      FuzzyMatching      `yaml:"fuzzyMatching"`
	APIToken           string             `yaml:"apiToken"`
	Notifications      Notifications      `yaml:"notifications"`
}
