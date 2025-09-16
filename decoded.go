package skycrypttypes

type EncodedItems struct {
	Type int    `json:"type"`
	Data string `json:"data"`
}

type Item struct {
	Count         *int   `nbt:"Count" json:"Count,omitempty"`
	Damage        *int   `nbt:"Damage" json:"Damage,omitempty"`
	ID            *int   `nbt:"id" json:"id,omitempty"`
	Tag           *Tag   `nbt:"tag" json:"tag,omitempty"`
	ContainsItems []Item `json:"containsItems,omitempty"`
}

type Tag struct {
	// HideFlags       int             `nbt:"HideFlags" json:"HideFlags,omitempty"`
	// Unbreakable     int             `nbt:"Unbreakable" json:"Unbreakable,omitempty"`
	// Enchantments    []Enchantment   `nbt:"ench" json:"ench,omitempty"`
	ExtraAttributes ExtraAttributes `nbt:"ExtraAttributes" json:"ExtraAttributes,omitempty"`
	Display         Display         `nbt:"display" json:"display"`
	SkullOwner      *SkullOwner     `nbt:"SkullOwner" json:"SkullOwner,omitempty"`
}

func (t *Tag) ToMap() TextureItemExtraAttributes {
	return TextureItemExtraAttributes{
		ExtraAttributes: t.ExtraAttributes.ToMap(),
		Display:         t.Display,
		SkullOwner:      t.SkullOwner,
	}
}

type TextureItemExtraAttributes struct {
	ExtraAttributes map[string]interface{} `nbt:"ExtraAttributes" json:"ExtraAttributes,omitempty"`
	Display         Display                `nbt:"display" json:"display"`
	SkullOwner      *SkullOwner            `nbt:"SkullOwner" json:"SkullOwner,omitempty"`
}

type ExtraAttributes struct {
	ID                 string         `nbt:"id" json:"id,omitempty"`
	UUID               string         `nbt:"uuid" json:"uuid,omitempty"`
	Timestamp          any            `nbt:"timestamp" json:"timestamp,omitempty"`
	Recombobulated     int            `nbt:"rarity_upgrades" json:"rarity_upgrades,omitempty"`
	Enchantments       map[string]int `nbt:"enchantments" json:"enchantments,omitempty"`
	Gems               map[string]any `nbt:"gems" json:"gems,omitempty"`
	HecatombSRuns      *int           `nbt:"hecatomb_s_runs" json:"hecatomb_s_runs,omitempty"`
	ChampionCombatXP   *float64       `nbt:"champion_combat_xp" json:"champion_combat_xp,omitempty"`
	FarmedCultivating  *int           `nbt:"farmed_cultivating" json:"farmed_cultivating,omitempty"`
	ExpertiseKills     *int           `nbt:"expertise_kills" json:"expertise_kills,omitempty"`
	CompactBlocks      *int           `nbt:"compact_blocks" json:"compact_blocks,omitempty"`
	Modifier           string         `nbt:"modifier" json:"modifier,omitempty"`
	Model              string         `nbt:"model" json:"model,omitempty"`
	TalismanEnrichment string         `nbt:"talisman_enrichment" json:"talisman_enrichment,omitempty"`
	Dye                string         `nbt:"dye_item" json:"dye_item,omitempty"`
}

func (t *ExtraAttributes) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                  t.ID,
		"uuid":                t.UUID,
		"timestamp":           t.Timestamp,
		"recombobulated":      t.Recombobulated,
		"enchantments":        t.Enchantments,
		"gems":                t.Gems,
		"hecatomb_s_runs":     t.HecatombSRuns,
		"champion_combat_xp":  t.ChampionCombatXP,
		"farmed_cultivating":  t.FarmedCultivating,
		"expertise_kills":     t.ExpertiseKills,
		"compact_blocks":      t.CompactBlocks,
		"modifier":            t.Modifier,
		"model":               t.Model,
		"talisman_enrichment": t.TalismanEnrichment,
	}
}

type Display struct {
	Name  string   `nbt:"Name" json:"Name,omitempty"`
	Lore  []string `nbt:"Lore" json:"Lore,omitempty"`
	Color int      `nbt:"color" json:"color,omitempty"`
}

type Enchantment struct {
	ID    int `nbt:"id" json:"id,omitempty"`
	Level int `nbt:"lvl" json:"lvl,omitempty"`
}

type SkullOwner struct {
	ID         string     `nbt:"Id" json:"Id,omitempty"`
	Properties Properties `nbt:"Properties" json:"Properties"`
}

type Properties struct {
	Textures []Texture `nbt:"textures" json:"textures,omitempty"`
}

type Texture struct {
	Value     string `nbt:"Value" json:"Value,omitempty"`
	Signature string `nbt:"Signature" json:"Signature,omitempty"`
}
