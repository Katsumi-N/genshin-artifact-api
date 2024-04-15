package domain

type ArtifactStatType int

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
