package main

import (
	"fmt"
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
		FightPropMap struct {
			Num1    float64 `json:"1"`
			Num2    float64 `json:"2"`
			Num3    float64 `json:"3"`
			Num4    float64 `json:"4"`
			Num5    int     `json:"5"`
			Num6    float64 `json:"6"`
			Num7    float64 `json:"7"`
			Num8    float64 `json:"8"`
			Num9    float64 `json:"9"`
			Num20   float64 `json:"20"`
			Num21   int     `json:"21"`
			Num22   float64 `json:"22"`
			Num23   float64 `json:"23"`
			Num26   int     `json:"26"`
			Num27   int     `json:"27"`
			Num28   float64 `json:"28"`
			Num29   int     `json:"29"`
			Num30   int     `json:"30"`
			Num40   int     `json:"40"`
			Num41   int     `json:"41"`
			Num42   int     `json:"42"`
			Num43   float64 `json:"43"`
			Num44   int     `json:"44"`
			Num45   int     `json:"45"`
			Num46   int     `json:"46"`
			Num50   int     `json:"50"`
			Num51   int     `json:"51"`
			Num52   int     `json:"52"`
			Num53   int     `json:"53"`
			Num54   int     `json:"54"`
			Num55   int     `json:"55"`
			Num56   int     `json:"56"`
			Num73   int     `json:"73"`
			Num1003 float64 `json:"1003"`
			Num1010 float64 `json:"1010"`
			Num2000 float64 `json:"2000"`
			Num2001 float64 `json:"2001"`
			Num2002 float64 `json:"2002"`
			Num2003 int     `json:"2003"`
			Num2004 int     `json:"2004"`
			Num3045 int     `json:"3045"`
			Num3046 int     `json:"3046"`
		} `json:"fightPropMap"`
		SkillDepotID           int   `json:"skillDepotId"`
		InherentProudSkillList []int `json:"inherentProudSkillList"`
		SkillLevelMap          struct {
			Num10731 int `json:"10731"`
			Num10732 int `json:"10732"`
			Num10735 int `json:"10735"`
		} `json:"skillLevelMap"`
		EquipList []struct {
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
			Flat0 struct {
				NameTextMapHash string `json:"nameTextMapHash"`
				RankLevel       int    `json:"rankLevel"`
				ItemType        string `json:"itemType"`
				Icon            string `json:"icon"`
				WeaponStats     []struct {
					AppendPropID string `json:"appendPropId"`
					StatValue    int    `json:"statValue"`
				} `json:"weaponStats"`
			} `json:"flat,omitempty"`
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

var ArtifactStatTypeMap = map[string]ArtifactStatType{
	"FIGHT_PROP_HP":                Health,
	"FIGHT_PROP_HP_PERCENT":        HealthPercent,
	"FIGHT_PROP_ATTACK":            Attack,
	"FIGHT_PROP_ATTACK_PERCENT":    AttackPercent,
	"FIGHT_PROP_DEFENSE":           Defence,
	"FIGHT_PROP_DEFENSE_PERCENT":   DefencePercent,
	"FIGHT_PROP_CRITICAL":          CriticalRate,
	"FIGHT_PROP_CRITICAL_HURT":     CriticalDamage,
	"FIGHT_PROP_ELEMENT_MASTERY":   ElementalMastery,
	"FIGHT_PROP_CHARGE_EFFICIENCY": ChargeEfficiency,
}

type ArtifactStatus struct {
	Type  ArtifactStatType
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
	fmt.Printf("%#v\n", characterStatusList)
	return characterStatusList
}
