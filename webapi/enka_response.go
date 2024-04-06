package main

import (
	"strconv"
)

type EnkaUserInfo struct {
	PlayerInfo struct {
		Nickname             string `json:"nickname"`
		Level                int    `json:"level"`
		Signature            string `json:"signature"`
		WorldLevel           int    `json:"worldLevel"`
		NameCardID           int    `json:"nameCardId"`
		FinishAchievementNum int    `json:"finishAchievementNum"`
		TowerFloorIndex      int    `json:"towerFloorIndex"`
		TowerLevelIndex      int    `json:"towerLevelIndex"`
		ShowAvatarInfoList   []struct {
			AvatarID int `json:"avatarId"`
			Level    int `json:"level"`
		} `json:"showAvatarInfoList"`
		ProfilePicture struct {
			ID int `json:"id"`
		} `json:"profilePicture"`
	} `json:"playerInfo"`
	AvatarInfoList []struct {
		AvatarID int `json:"avatarId"`
		PropMap  struct {
			Level struct {
				Val string `json:"val"`
			} `json:"4001"`
		} `json:"propMap"`
		FightPropMap           map[string]float64 `json:"fightPropMap"`
		SkillDepotID           int                `json:"skillDepotId"`
		InherentProudSkillList []int              `json:"inherentProudSkillList"`
		SkillLevelMap          map[string]int     `json:"skillLevelMap"`
		EquipList              []struct {
			ItemID    int `json:"itemId"`
			Reliquary struct {
				Level            int   `json:"level"`
				Exp              int   `json:"exp"`
				MainPropID       int   `json:"mainPropId"`
				AppendPropIDList []int `json:"appendPropIdList"`
			} `json:"reliquary,omitempty"`
			Flat struct {
				NameTextMapHash    string `json:"nameTextMapHash"`
				RankLevel          int    `json:"rankLevel"`
				ItemType           string `json:"itemType"`
				Icon               string `json:"icon"`
				EquipType          string `json:"equipType"`
				SetNameTextMapHash string `json:"setNameTextMapHash"`
				ReliquarySubstats  []struct {
					AppendPropID string  `json:"appendPropId"`
					StatValue    float64 `json:"statValue"`
				} `json:"reliquarySubstats"`
				ReliquaryMainstat struct {
					MainPropID string  `json:"mainPropId"`
					StatValue  float64 `json:"statValue"`
				} `json:"reliquaryMainstat"`
			} `json:"flat,omitempty"`
			Weapon struct {
				Level        int `json:"level"`
				PromoteLevel int `json:"promoteLevel"`
				AffixMap     struct {
					Num114301 int `json:"114301"`
				} `json:"affixMap"`
			} `json:"weapon,omitempty"`
		} `json:"equipList"`
		FetterInfo struct {
			ExpLevel int `json:"expLevel"`
		} `json:"fetterInfo"`
	} `json:"avatarInfoList"`
	TTL int    `json:"ttl"`
	UID string `json:"uid"`
}

type ArtifactStatType int

const (
	Attack ArtifactStatType = iota
	AttackPercent
	ChargeEfficiency
	Health
	HealthPercent
	Defence
	DefencePercent
	CriticalRate
	CriticalDamage
	ElementalMastery
)

// var ArtifactStatTypeMap = map[string]ArtifactStatType{
// 	"FIGHT_PROP_HP":                Health,
// 	"FIGHT_PROP_HP_PERCENT":        HealthPercent,
// 	"FIGHT_PROP_ATTACK":            Attack,
// 	"FIGHT_PROP_ATTACK_PERCENT":    AttackPercent,
// 	"FIGHT_PROP_DEFENSE":           Defence,
// 	"FIGHT_PROP_DEFENSE_PERCENT":   DefencePercent,
// 	"FIGHT_PROP_CRITICAL":          CriticalRate,
// 	"FIGHT_PROP_CRITICAL_HURT":     CriticalDamage,
// 	"FIGHT_PROP_ELEMENT_MASTERY":   ElementalMastery,
// 	"FIGHT_PROP_CHARGE_EFFICIENCY": ChargeEfficiency,
// }

var ArtifactStatTypeMap = map[string]string{
	"FIGHT_PROP_HP":                "Health",
	"FIGHT_PROP_HP_PERCENT":        "HealthPercent",
	"FIGHT_PROP_ATTACK":            "Attack",
	"FIGHT_PROP_ATTACK_PERCENT":    "AttackPercent",
	"FIGHT_PROP_DEFENSE":           "Defence",
	"FIGHT_PROP_DEFENSE_PERCENT":   "DefencePercent",
	"FIGHT_PROP_CRITICAL":          "CriticalRate",
	"FIGHT_PROP_CRITICAL_HURT":     "CriticalDamage",
	"FIGHT_PROP_ELEMENT_MASTERY":   "ElementalMastery",
	"FIGHT_PROP_CHARGE_EFFICIENCY": "ChargeEfficiency",
}

type ArtifactStatus struct {
	Type  string
	Value float64
}

// TODO: 適切な名前に変更
type ArtifactPiece struct {
	Main ArtifactStatus
	Sub  []ArtifactStatus
}

type CharacterStatus struct {
	EnkaID  int
	Level   int
	Flower  ArtifactPiece
	Plume   ArtifactPiece
	Sands   ArtifactPiece
	Goblet  ArtifactPiece
	Circlet ArtifactPiece
}

func (e *EnkaUserInfo) ExtractCharacterStatus() []CharacterStatus {
	// AvatorInfoListが公開キャラクター
	avatorInfoList := e.AvatarInfoList

	// avatorInfoListからCharacterStatusを抽出
	characterStatusList := make([]CharacterStatus, len(avatorInfoList))
	for i, avatorInfo := range avatorInfoList {
		enka_id := avatorInfo.AvatarID
		level, _ := strconv.Atoi(avatorInfo.PropMap.Level.Val)

		characterStatusList[i] = CharacterStatus{
			EnkaID: enka_id,
			Level:  level,
		}

		artfacts := avatorInfo.EquipList
		for _, artifact := range artfacts {
			if artifact.Flat.ItemType != "ITEM_RELIQUARY" {
				continue
			}

			switch artifact.Flat.EquipType {
			case "EQUIP_BRACER":
				characterStatusList[i].Flower = ArtifactPiece{
					Main: ArtifactStatus{
						Type:  ArtifactStatTypeMap[artifact.Flat.ReliquaryMainstat.MainPropID],
						Value: artifact.Flat.ReliquaryMainstat.StatValue,
					},
				}
				for _, substat := range artifact.Flat.ReliquarySubstats {
					characterStatusList[i].Flower.Sub = append(characterStatusList[i].Flower.Sub, ArtifactStatus{
						Type:  ArtifactStatTypeMap[substat.AppendPropID],
						Value: substat.StatValue,
					})
				}
			case "EQUIP_NECKLACE":
				characterStatusList[i].Plume = ArtifactPiece{
					Main: ArtifactStatus{
						Type:  ArtifactStatTypeMap[artifact.Flat.ReliquaryMainstat.MainPropID],
						Value: artifact.Flat.ReliquaryMainstat.StatValue,
					},
				}
				for _, substat := range artifact.Flat.ReliquarySubstats {
					characterStatusList[i].Plume.Sub = append(characterStatusList[i].Plume.Sub, ArtifactStatus{
						Type:  ArtifactStatTypeMap[substat.AppendPropID],
						Value: substat.StatValue,
					})
				}
			case "EQUIP_SHOES":
				characterStatusList[i].Sands = ArtifactPiece{
					Main: ArtifactStatus{
						Type:  ArtifactStatTypeMap[artifact.Flat.ReliquaryMainstat.MainPropID],
						Value: artifact.Flat.ReliquaryMainstat.StatValue,
					},
				}
				for _, substat := range artifact.Flat.ReliquarySubstats {
					characterStatusList[i].Sands.Sub = append(characterStatusList[i].Sands.Sub, ArtifactStatus{
						Type:  ArtifactStatTypeMap[substat.AppendPropID],
						Value: substat.StatValue,
					})
				}
			case "EQUIP_RING":
				characterStatusList[i].Goblet = ArtifactPiece{
					Main: ArtifactStatus{
						Type:  ArtifactStatTypeMap[artifact.Flat.ReliquaryMainstat.MainPropID],
						Value: artifact.Flat.ReliquaryMainstat.StatValue,
					},
				}
				for _, substat := range artifact.Flat.ReliquarySubstats {
					characterStatusList[i].Goblet.Sub = append(characterStatusList[i].Goblet.Sub, ArtifactStatus{
						Type:  ArtifactStatTypeMap[substat.AppendPropID],
						Value: substat.StatValue,
					})
				}
			case "EQUIP_DRESS":
				characterStatusList[i].Circlet = ArtifactPiece{
					Main: ArtifactStatus{
						Type:  ArtifactStatTypeMap[artifact.Flat.ReliquaryMainstat.MainPropID],
						Value: artifact.Flat.ReliquaryMainstat.StatValue,
					},
				}
				for _, substat := range artifact.Flat.ReliquarySubstats {
					characterStatusList[i].Circlet.Sub = append(characterStatusList[i].Circlet.Sub, ArtifactStatus{
						Type:  ArtifactStatTypeMap[substat.AppendPropID],
						Value: substat.StatValue,
					})
				}
			}
		}
	}

	return characterStatusList
}
