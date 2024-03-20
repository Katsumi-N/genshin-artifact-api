package main

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
		AvatarID     int `json:"avatarId"`
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
					MainPropID string `json:"mainPropId"`
					StatValue  int    `json:"statValue"`
				} `json:"reliquaryMainstat"`
			} `json:"flat,omitempty"`
			Reliquary0 struct {
				Level            int   `json:"level"`
				MainPropID       int   `json:"mainPropId"`
				AppendPropIDList []int `json:"appendPropIdList"`
			} `json:"reliquary,omitempty"`
			Reliquary1 struct {
				Level            int   `json:"level"`
				MainPropID       int   `json:"mainPropId"`
				AppendPropIDList []int `json:"appendPropIdList"`
			} `json:"reliquary,omitempty"`
			Reliquary2 struct {
				Level            int   `json:"level"`
				MainPropID       int   `json:"mainPropId"`
				AppendPropIDList []int `json:"appendPropIdList"`
			} `json:"reliquary,omitempty"`
			Reliquary3 struct {
				Level            int   `json:"level"`
				MainPropID       int   `json:"mainPropId"`
				AppendPropIDList []int `json:"appendPropIdList"`
			} `json:"reliquary,omitempty"`
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
	characterStatusList := make([]CharacterStatus, 0, len(avatorInfoList))
	for i, avatorInfo := range avatorInfoList {
		enka_id := avatorInfo.AvatarID
		level := avatorInfo.FightPropMap.Num1

		characterStatusList[i] = CharacterStatus{

			EnkaID:  enka_id,
			Level:   level,
			Flower:  ArtifactPiece{},
			Plume:   ArtifactPiece{},
			Sands:   ArtifactPiece{},
			Goblet:  ArtifactPiece{},
			Circlet: ArtifactPiece{},
		}
	}

	return characterStatusList
}
