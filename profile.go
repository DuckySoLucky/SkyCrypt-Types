package skycrypttypes

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Profile struct {
	ProfileID         string             `json:"profile_id"`
	CuteName          string             `json:"cute_name"`
	Selected          bool               `json:"selected"`
	Members           map[string]Member  `json:"members"`
	GameMode          string             `json:"game_mode,omitempty"`
	Banking           *Banking           `json:"banking,omitempty"`
	CommunityUpgrades *CommunityUpgrades `json:"community_upgrades,omitempty"`
}

type Member struct {
	PlayerData          *PlayerData            `json:"player_data"`
	CoopInvitation      *CoopInvitation        `json:"coop_invitation"`
	Profile             *ProfileData           `json:"profile"`
	JacobsContest       JacobsContest          `json:"jacobs_contest,omitempty"`
	Pets                *Pets                  `json:"pets_data,omitempty"`
	Leveling            Leveling               `json:"leveling,omitempty"`
	Currencies          Currencies             `json:"currencies,omitempty"`
	FairySouls          *FairySouls            `json:"fairy_soul,omitempty"`
	Inventory           *Inventory             `json:"inventory,omitempty"`
	SharedInventory     SharedInventory        `json:"shared_inventory,omitempty"`
	Rift                Rift                   `json:"rift,omitempty"`
	AccessoryBagStorage AccessoryBagStorage    `json:"accessory_bag_storage,omitempty"`
	CrimsonIsle         CrimsonIsleData        `json:"nether_island_player_data,omitempty"`
	Mining              Mining                 `json:"mining_core,omitempty"`
	Objectives          *Objectives            `json:"objectives,omitempty"`
	GlaciteTunnels      *GlaciteData           `json:"glacite_player_data,omitempty"`
	Forge               Forge                  `json:"forge,omitempty"`
	Quests              Quests                 `json:"quests,omitempty"`
	Garden              GardenProfileData      `json:"garden_player_data,omitempty"`
	PlayerStats         PlayerStats            `json:"player_stats,omitempty"`
	TrophyFish          MemberTrophyFish       `json:"trophy_fish,omitempty"`
	Experimentation     ExperimentationData    `json:"experimentation,omitempty"`
	Dungeons            Dungeons               `json:"dungeons,omitempty"`
	Slayer              Slayer                 `json:"slayer,omitempty"`
	Bestiary            *Bestiary              `json:"bestiary,omitempty"`
	Collections         map[string]int         `json:"collection,omitempty"`
	ItemData            ItemData               `json:"item_data,omitempty"`
	WinterPlayerData    WinterPlayerIslandData `json:"winter_player_data,omitempty"`
	SackCounts          map[string]int         `json:"sack_counts"`
}

type WinterPlayerIslandData struct {
	RefinedJyrreUses int `json:"refined_bottle_of_jyrre_uses,omitempty"`
}

type ItemData struct {
	Soulflow               float64 `json:"soulflow,omitempty"`
	TeleporterPillConsumed bool    `json:"teleporter_pill_consumed,omitempty"`
}

type Races struct {
	ForagingRaceBestTime float64            `json:"foraging_race_best_time"`
	EndRaceBestTime      float64            `json:"end_race_best_time"`
	ChickenRaceBestTime2 float64            `json:"chicken_race_best_time_2"`
	DungeonHub           map[string]float64 `json:"dungeon_hub"`
	RiftRaceBestTime     float64            `json:"rift_race_best_time"`
}

type CoopInvitation struct {
	Confirmed bool `json:"confirmed,omitempty"`
}

type PlayerData struct {
	Experience         *Experience `json:"experience"`
	Minions            []string    `json:"crafted_generators"`
	ReaperPeppersEaten int         `json:"reaper_peppers_eaten,omitempty"`
}

type Experience struct {
	SkillFishing      float64 `json:"SKILL_FISHING"`
	SkillAlchemy      float64 `json:"SKILL_ALCHEMY"`
	SkillMining       float64 `json:"SKILL_MINING"`
	SkillFarming      float64 `json:"SKILL_FARMING"`
	SkillEnchanting   float64 `json:"SKILL_ENCHANTING"`
	SkillTaming       float64 `json:"SKILL_TAMING"`
	SkillForaging     float64 `json:"SKILL_FORAGING"`
	SkillSocial       float64 `json:"SKILL_SOCIAL"`
	SkillCarpentry    float64 `json:"SKILL_CARPENTRY"`
	SkillCombat       float64 `json:"SKILL_COMBAT"`
	SkillRunecrafting float64 `json:"SKILL_RUNECRAFTING"`
}

type ProfileData struct {
	DeletionNotice      *DeletionNotice `json:"deletion_notice"`
	FirstJoin           int64           `json:"first_join,omitempty"`
	BankAccount         float64         `json:"bank_account,omitempty"`
	PersonalBankUpgrade int             `json:"personal_bank_upgrade,omitempty"`
}

type DeletionNotice struct {
	Timestamp int64 `json:"timestamp,omitempty"`
}

type JacobsContest struct {
	Perks          *Perks                      `json:"perks,omitempty"`
	UniqueBrackets map[string][]string         `json:"unique_brackets,omitempty"`
	MedalsInv      map[string]int              `json:"medals_inv,omitempty"`
	Contests       map[string]JacobContestData `json:"contests,omitempty"`
}

type Perks struct {
	FarmingLevelCap int `json:"farming_level_cap,omitempty"`
}

type PetCare struct {
	PetTypesSacrificed []string `json:"pet_types_sacrificed,omitempty"`
}

type Leveling struct {
	Experience int `json:"experience,omitempty"`
}

type Currencies struct {
	CoinPurse  float64            `json:"coin_purse,omitempty"`
	MotesPurse float64            `json:"motes_purse,omitempty"`
	Essence    map[string]Essence `json:"essence,omitempty"`
}

type Essence struct {
	Current int `json:"current,omitempty"`
}

type Banking struct {
	Balance *float64 `json:"balance,omitempty"`
}

type FairySouls struct {
	TotalCollected int `json:"total_collected,omitempty"`
}

type Inventory struct {
	Inventory     EncodedItems            `json:"inv_contents"`
	Enderchest    EncodedItems            `json:"ender_chest_contents"`
	BackpackIcons map[string]EncodedItems `json:"backpack_icons"`
	Armor         EncodedItems            `json:"inv_armor"`
	Equipment     EncodedItems            `json:"equipment_contents"`
	PersonalVault EncodedItems            `json:"personal_vault_contents"`
	Backpack      map[string]EncodedItems `json:"backpack_contents"`
	Wardrobe      EncodedItems            `json:"wardrobe_contents"`
	BagContents   BagContents             `json:"bag_contents"`
	Sacks         map[string]int          `json:"sacks_counts"`
}

type SharedInventory struct {
	CandyInventory        EncodedItems `json:"candy_inventory_contents"`
	CarnivalMaskInventory EncodedItems `json:"carnival_mask_inventory_contents"`
}

type BagContents struct {
	PotionBag   EncodedItems `json:"potion_bag,omitempty"`
	TalismanBag EncodedItems `json:"talisman_bag,omitempty"`
	FishingBag  EncodedItems `json:"fishing_bag,omitempty"`
	SacksBag    EncodedItems `json:"sacks_bag,omitempty"`
	Quiver      EncodedItems `json:"quiver,omitempty"`
}

type Rift struct {
	Inventory  RiftInventory  `json:"inventory,omitempty"`
	Access     RiftAccess     `json:"access,omitempty"`
	DeadCats   DeadCats       `json:"dead_cats,omitempty"`
	Enigma     RiftEnigma     `json:"enigma,omitempty"`
	Castle     RiftCastle     `json:"castle,omitempty"`
	Gallery    RiftGallery    `json:"gallery,omitempty"`
	WitherCage RiftWitherCage `json:"wither_cage,omitempty"`
}

type RiftWitherCage struct {
	KilledEyes []string `json:"killed_eyes,omitempty"`
}

type RiftGallery struct {
	SecuredTrophies []RiftSecuredTrophy `json:"secured_trophies,omitempty"`
}

type RiftSecuredTrophy struct {
	Type      string `json:"type,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
}

type RiftCastle struct {
	GrubberStacks int `json:"grubber_stacks,omitempty"`
}

type RiftEnigma struct {
	FoundSouls []string `json:"found_souls,omitempty"`
}

type RiftInventory struct {
	Inventory  EncodedItems `json:"inv_contents"`
	Armor      EncodedItems `json:"inv_armor"`
	Enderchest EncodedItems `json:"ender_chest_contents"`
	Equipment  EncodedItems `json:"equipment_contents"`
}

type RiftAccess struct {
	ConsumedPrism bool `json:"consumed_prism,omitempty"`
}

type AccessoryBagStorage struct {
	SelectedPower string `json:"selected_power,omitempty"`
}

type CrimsonIsleData struct {
	Abiphone            Abiphone       `json:"abiphone,omitempty"`
	Kuudra              map[string]int `json:"kuudra_completed_tiers,omitempty"`
	Dojo                map[string]int `json:"dojo,omitempty"`
	SelectedFaction     string         `json:"selected_faction,omitempty"`
	MagesReputation     float64        `json:"mages_reputation,omitempty"`
	BarbarianReputation float64        `json:"barbarians_reputation,omitempty"`
}

type Abiphone struct {
	ActiveContacts []string `json:"active_contacts,omitempty"`
}

type DeadCats struct {
	FoundCats []string `json:"found_cats,omitempty"`
	Montezuma Pet      `json:"montezuma,omitempty"`
}

type Pet struct {
	Type       string  `json:"type,omitempty"`
	Experience float64 `json:"exp,omitempty"`
	Active     bool    `json:"active,omitempty"`
	Rarity     string  `json:"tier,omitempty"`
	HeldItem   string  `json:"heldItem,omitempty"`
	CandyUsed  int     `json:"candyUsed,omitempty"`
	Skin       string  `json:"skin,omitempty"`
}

type Pets struct {
	PetCare *PetCare `json:"pet_care,omitempty"`
	Pets    []Pet    `json:"pets,omitempty"`
}

type Mining struct {
	Nodes                  map[string]int     `json:"nodes,omitempty"`
	Experience             float64            `json:"experience,omitempty"`
	GreaterMinesLastAccess int64              `json:"greater_mines_last_access,omitempty"`
	LastReset              int64              `json:"last_reset,omitempty"`
	TokensSpent            int                `json:"tokens_spent,omitempty"`
	SelectedPickaxeAbility string             `json:"selected_pickaxe_ability,omitempty"`
	PowderMithril          int                `json:"powder_mithril,omitempty"`
	PowderMithrilTotal     int                `json:"powder_mithril_total,omitempty"`
	PowderSpentMithril     int                `json:"powder_spent_mithril,omitempty"`
	PowderGemstone         int                `json:"powder_gemstone,omitempty"`
	PowderGemstoneTotal    int                `json:"powder_gemstone_total,omitempty"`
	PowderSpentGemstone    int                `json:"powder_spent_gemstone,omitempty"`
	PowderGlacite          int                `json:"powder_glacite,omitempty"`
	PowderGlaciteTotal     int                `json:"powder_glacite_total,omitempty"`
	PowderSpentGlacite     int                `json:"powder_spent_glacite,omitempty"`
	Crystals               map[string]Crystal `json:"crystals,omitempty"`
	Biomes                 Biomes             `json:"biomes,omitempty"`
}

type Crystal struct {
	State       string `json:"state,omitempty"`
	TotalFound  int    `json:"total_found,omitempty"`
	TotalPlaced int    `json:"total_placed,omitempty"`
}

type Biomes struct {
	Precursor Precursor `json:"precursor,omitempty"`
}

type Precursor struct {
	PartsDelivered []string `json:"parts_delivered,omitempty"`
}

type Objectives struct {
	Tutorial []string `json:"tutorial,omitempty"`
}

type GlaciteData struct {
	FossilsDonated    []string       `json:"fossils_donated,omitempty"`
	FossilDust        float64        `json:"fossil_dust,omitempty"`
	CorpsesLooted     map[string]int `json:"corpses_looted,omitempty"`
	MineshaftsEntered int            `json:"mineshafts_entered,omitempty"`
}

type Forge struct {
	ForgeProcesses ForgeProcesses `json:"forge_processes"`
}

type ForgeProcesses struct {
	Forge map[string]ForgeProcess `json:"forge_1"`
}

type ForgeProcess struct {
	Id        string `json:"id"`
	StartTime int64  `json:"startTime"`
	Slot      int    `json:"slot"`
}

type Quests struct {
	TrapperQuest TrapperQuest `json:"trapper_quest,omitempty"`
}

type TrapperQuest struct {
	PeltCount int `json:"pelt_count,omitempty"`
}

type GardenProfileData struct {
	Copper        int `json:"copper,omitempty"`
	LarvaConsumed int `json:"larva_consumed,omitempty"`
}

type JacobContestData struct {
	Collected           int    `json:"collected"`
	ClaimedPosition     *int   `json:"claimed_position,omitempty"`
	ClaimedParticipants *int   `json:"claimed_participants,omitempty"`
	ClaimedMedal        string `json:"claimed_medal"`
}

type PlayerStats struct {
	Kills       map[string]float64 `json:"kills,omitempty"`
	Deaths      map[string]float64 `json:"deaths,omitempty"`
	ItemsFished struct {
		Total         float64 `json:"total,omitempty"`
		Normal        float64 `json:"normal,omitempty"`
		Treasure      float64 `json:"treasure,omitempty"`
		LargeTreasure float64 `json:"large_treasure,omitempty"`
		TrophyFish    float64 `json:"trophy_fish,omitempty"`
	} `json:"items_fished"`
	ShredderRod struct {
		Fished float64 `json:"fished,omitempty"`
		Bait   float64 `json:"bait,omitempty"`
	} `json:"shredder_rod"`
	Pets struct {
		Milestone struct {
			SeaCreaturesKilled float64 `json:"sea_creatures_killed,omitempty"`
			OresMined          float64 `json:"ores_mined,omitempty"`
		} `json:"milestone,omitempty"`
	} `json:"pets,omitempty"`
	Rift                  RiftPlayerData   `json:"rift,omitempty"`
	Races                 Races            `json:"races,omitempty"`
	Gifts                 Gifts            `json:"gifts"`
	WinterIslandData      WinterIslandData `json:"winter"`
	EndIsland             EndIsland        `json:"end_island"`
	HighestCriticalDamage float64          `json:"highest_critical_damage"`
	Mythos                Mythos           `json:"mythos"`
	Auctions              Auctions         `json:"auctions"`
}

type Auctions struct {
	Bids        float64            `json:"bids"`
	HighestBid  float64            `json:"highest_bid"`
	Won         float64            `json:"won"`
	TotalBought map[string]float64 `json:"total_bought"`
	GoldSpent   float64            `json:"gold_spent"`
	Created     float64            `json:"created"`
	Fees        float64            `json:"fees"`
	TotalSold   map[string]float64 `json:"total_sold"`
	GoldEarned  float64            `json:"gold_earned"`
	NoBids      float64            `json:"no_bids"`
}

type Mythos struct {
	Kills                 float64            `json:"kills"`
	BurrowsDugNext        map[string]float64 `json:"burrows_dug_next"`
	BurrowsDugCombat      map[string]float64 `json:"burrows_dug_combat"`
	BurrowsDugTreasure    map[string]float64 `json:"burrows_dug_treasure"`
	BurrowsChainsComplete map[string]float64 `json:"burrows_chains_complete"`
}

type EndIsland struct {
	DragonFight DragonFight `json:"dragon_fight"`
}

type DragonFight struct {
	EnderCrystalsDestroyed float64            `json:"ender_crystals_destroyed"`
	MostDamage             map[string]float64 `json:"most_damage"`
	FastestKill            map[string]float64 `json:"fastest_kill"`
}

type WinterIslandData struct {
	MostSnowballsHit     float64 `json:"most_snowballs_hit"`
	MostDamageDealt      float64 `json:"most_damage_dealt"`
	MostMagmaDamageDealt float64 `json:"most_magma_damage_dealt"`
	MostCannonballsHit   float64 `json:"most_cannonballs_hit"`
}

type Gifts struct {
	Given    float64 `json:"total_given"`
	Received float64 `json:"total_received"`
}

type RiftPlayerData struct {
	Visits                 float64 `json:"visits,omitempty"`
	LifetimeMotesCollected float64 `json:"lifetime_motes_earned,omitempty"`
	MotesOrbPickup         float64 `json:"motes_orb_pickup,omitempty"`
}

type MemberTrophyFish struct {
	Rewards     []int          `json:"rewards"`
	TotalCaught int            `json:"total_caught"`
	Extra       map[string]int `json:"-"`
}

func (t *MemberTrophyFish) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	t.Extra = make(map[string]int)

	for key, value := range raw {
		switch key {
		case "rewards":
			if err := json.Unmarshal(value, &t.Rewards); err != nil {
				return err
			}
		case "total_caught":
			if err := json.Unmarshal(value, &t.TotalCaught); err != nil {
				return err
			}
		default:
			var num int
			if err := json.Unmarshal(value, &num); err == nil {
				t.Extra[key] = num
			}
		}
	}

	return nil
}

type ExperimentationGame struct {
	LastAttempt int64            `json:"last_attempt,omitempty"`
	LastClaimed int64            `json:"last_claimed,omitempty"`
	BonusClicks int              `json:"bonus_clicks,omitempty"`
	Claimed     bool             `json:"claimed,omitempty"`
	Attempts    map[int]int      `json:"-"`
	Claims      map[int]int      `json:"-"`
	BestScores  map[int]int      `json:"-"`
	Raw         map[string]int64 `json:"-"`
}

func (e *ExperimentationGame) UnmarshalJSON(data []byte) error {
	type Alias ExperimentationGame
	aux := &struct {
		*Alias
	}{Alias: (*Alias)(e)}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	e.Attempts = make(map[int]int)
	e.Claims = make(map[int]int)
	e.BestScores = make(map[int]int)
	e.Raw = make(map[string]int64)

	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	for k, v := range raw {
		switch {
		case len(k) > 9 && k[:9] == "attempts_":
			var idx int
			fmt.Sscanf(k, "attempts_%d", &idx)
			e.Attempts[idx] = int(v.(float64))
		case len(k) > 7 && k[:7] == "claims_":
			var idx int
			fmt.Sscanf(k, "claims_%d", &idx)
			e.Claims[idx] = int(v.(float64))
		case len(k) > 11 && k[:11] == "best_score_":
			var idx int
			fmt.Sscanf(k, "best_score_%d", &idx)
			e.BestScores[idx] = int(v.(float64))
		}
	}
	return nil
}

type ExperimentationData struct {
	Simon                 *ExperimentationGame `json:"simon,omitempty"`
	Pairings              *ExperimentationGame `json:"pairings,omitempty"`
	Numbers               *ExperimentationGame `json:"numbers,omitempty"`
	ClaimsResets          *int64               `json:"claims_resets,omitempty"`
	ClaimsResetsTimestamp int64                `json:"claims_resets_timestamp,omitempty"`
	SerumsDrank           int                  `json:"serums_drank,omitempty"`
	ClaimedRetroactiveRng bool                 `json:"claimed_retroactive_rng,omitempty"`
	ChargeTrackTimestamp  int64                `json:"charge_track_timestamp,omitempty"`
}

type Dungeons struct {
	DungeonTypes         map[string]DungeonData `json:"dungeon_types,omitempty"`
	Classes              map[string]PlayerClass `json:"player_classes,omitempty"`
	SelectedDungeonClass string                 `json:"selected_dungeon_class,omitempty"`
	Secrets              float64                `json:"secrets,omitempty"`
}

type PlayerClass struct {
	Experience float64 `json:"experience,omitempty"`
}

type DungeonData struct {
	Experience float64 `json:"experience,omitempty"`

	HighestTierCompleted int                   `json:"highest_tier_completed,omitempty"`
	TimesPlayed          map[string]float64    `json:"times_played,omitempty"`
	TierCompletions      map[string]float64    `json:"tier_completions,omitempty"`
	MilestoneCompletions map[string]float64    `json:"milestone_completions,omitempty"`
	MobsKilled           map[string]float64    `json:"mobs_killed,omitempty"`
	MostMobsKilled       map[string]float64    `json:"most_mobs_killed,omitempty"`
	WatcherKills         map[string]float64    `json:"watcher_kills,omitempty"`
	MostDamageBerserk    map[string]float64    `json:"most_damage_berserk,omitempty"`
	MostDamageMage       map[string]float64    `json:"most_damage_mage,omitempty"`
	MostDamageHealer     map[string]float64    `json:"most_damage_healer,omitempty"`
	MostDamageArcher     map[string]float64    `json:"most_damage_archer,omitempty"`
	MostDamageTank       map[string]float64    `json:"most_damage_tank,omitempty"`
	MostHealing          map[string]float64    `json:"most_healing,omitempty"`
	FastestTime          map[string]float64    `json:"fastest_time,omitempty"`
	FastestTimeS         map[string]float64    `json:"fastest_time_s,omitempty"`
	FastestTimeSPlus     map[string]float64    `json:"fastest_time_s_plus,omitempty"`
	BestScore            map[string]float64    `json:"best_score,omitempty"`
	BestRuns             map[string]*[]BestRun `json:"best_runs,omitempty"`
}

type BestRun struct {
	Timestamp        int64   `json:"timestamp"`
	ScoreExploration int     `json:"score_exploration"`
	ScoreSpeed       int     `json:"score_speed"`
	ScoreSkill       int     `json:"score_skill"`
	ScoreBonus       int     `json:"score_bonus"`
	DungeonClass     string  `json:"dungeon_class"`
	ElapsedTime      int64   `json:"elapsed_time"`
	DamageDealt      float64 `json:"damage_dealt"`
	Deaths           int     `json:"deaths"`
	MobsKilled       int     `json:"mobs_killed"`
	SecretsFound     int     `json:"secrets_found"`
	DamageMitigated  float64 `json:"damage_mitigated"`
}

type Slayer struct {
	SlayerBosses map[string]SlayerBoss `json:"slayer_bosses,omitempty"`
}

type SlayerBoss struct {
	BossKillsTier0    int     `json:"boss_kills_tier_0,omitempty"`
	BossKillsTier1    int     `json:"boss_kills_tier_1,omitempty"`
	BossKillsTier2    int     `json:"boss_kills_tier_2,omitempty"`
	BossKillsTier3    int     `json:"boss_kills_tier_3,omitempty"`
	BossKillsTier4    int     `json:"boss_kills_tier_4,omitempty"`
	BossAttemptsTier0 int     `json:"boss_attempts_tier_0,omitempty"`
	BossAttemptsTier1 int     `json:"boss_attempts_tier_1,omitempty"`
	BossAttemptsTier2 int     `json:"boss_attempts_tier_2,omitempty"`
	BossAttemptsTier3 int     `json:"boss_attempts_tier_3,omitempty"`
	BossAttemptsTier4 int     `json:"boss_attempts_tier_4,omitempty"`
	Experience        float64 `json:"xp,omitempty"`
}

type CommunityUpgrades struct {
	UpgradeStates []CommunityUpgradeState `json:"upgrade_states,omitempty"`
}

type CommunityUpgradeState struct {
	Upgrade string `json:"upgrade,omitempty"`
	Tier    int    `json:"tier,omitempty"`
}
type Bestiary struct {
	Kills map[string]int `json:"kills,omitempty"`
}

func (b *Bestiary) UnmarshalJSON(data []byte) error {
	type Alias Bestiary
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(b),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		if !(json.Unmarshal(data, &map[string]interface{}{}) == nil && (err.Error() == "json: cannot unmarshal string into Go struct field .Alias.kills of type int" || err.Error() == "json: cannot unmarshal string into Go struct field .Alias.deaths of type int")) {
			return err
		}
	}

	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	b.Kills = make(map[string]int)
	if killsRaw, ok := raw["kills"]; ok {
		if killsMap, ok := killsRaw.(map[string]interface{}); ok {
			for k, v := range killsMap {
				switch val := v.(type) {
				case float64:
					b.Kills[k] = int(val)
				case int:
					b.Kills[k] = val
				case string:
					if i, err := strconv.Atoi(val); err == nil {
						b.Kills[k] = i
					}
				}
			}
		}
	}

	return nil
}
