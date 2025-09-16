package skycrypttypes

type SkyblockMuseum struct {
	Value     int64                  `json:"value,omitempty"`
	Appraisal bool                   `json:"appraisal,omitempty"`
	Items     *map[string]MuseumItem `json:"items,omitempty"`
	Special   *[]MuseumItem          `json:"special,omitempty"`
}

type MuseumItem struct {
	DonatedTime  int64        `json:"donated_time,omitempty"`
	FeaturedSlot string       `json:"featured_slot,omitempty"`
	Borrowing    bool         `json:"borrowing,omitempty"`
	Items        EncodedItems `json:"items"`
}
