package crawler

import (
	"github.com/axiangcoding/ax-web/data/table"
	"strconv"
	"time"
)

var (
	SourceGaijin       = "gaijin"
	SourceThunderSkill = "thunderskill"
)

type GaijinData struct {
	Nick         string `json:"nick" mapstructure:"nick"`
	Clan         string `json:"clan" mapstructure:"clan"`
	ClanUrl      string `json:"clan_url" mapstructure:"clanUrl"`
	Banned       bool   `json:"banned" mapstructure:"banned"`
	RegisterDate string `json:"register_date" mapstructure:"register_date"`
	Title        string `json:"title" mapstructure:"title"`
	Level        string `json:"level" mapstructure:"level"`
	UserStat     struct {
		Ab map[string]string `json:"ab,omitempty" mapstructure:"ab"`
		Rb map[string]string `json:"rb,omitempty" mapstructure:"rb"`
		Sb map[string]string `json:"sb,omitempty" mapstructure:"sb"`
	} `json:"user_stat" mapstructure:"user_stat"`
	UserRate struct {
		Aviation struct {
			Ab map[string]string `json:"ab,omitempty" mapstructure:"ab"`
			Rb map[string]string `json:"rb,omitempty" mapstructure:"rb"`
			Sb map[string]string `json:"sb,omitempty" mapstructure:"sb"`
		} `json:"aviation" mapstructure:"aviation"`
		GroundVehicles struct {
			Ab map[string]string `json:"ab,omitempty" mapstructure:"ab"`
			Rb map[string]string `json:"rb,omitempty" mapstructure:"rb"`
			Sb map[string]string `json:"sb,omitempty" mapstructure:"sb"`
		} `json:"ground_vehicles" mapstructure:"ground_vehicles"`
		Fleet struct {
			Ab map[string]string `json:"ab,omitempty" mapstructure:"ab"`
			Rb map[string]string `json:"rb,omitempty" mapstructure:"rb"`
			Sb map[string]string `json:"sb,omitempty" mapstructure:"sb"`
		} `json:"fleet" mapstructure:"fleet"`
	} `json:"user_rate" mapstructure:"user_rate"`
}

func (d GaijinData) ToTableGameUser() table.GameUser {
	dateStr := d.RegisterDate
	parse, _ := time.Parse("2006-01-02", dateStr)
	level, _ := strconv.Atoi(d.Level)
	return table.GameUser{
		Nick:         d.Nick,
		Clan:         d.Clan,
		ClanUrl:      d.ClanUrl,
		Banned:       d.Banned,
		RegisterDate: parse,
		Title:        d.Title,
		Level:        level,
	}
}

type ThunderSkillData struct {
	Nick        string `json:"nick"`
	Rank        string `json:"rank"`
	LastStat    string `json:"last_stat"`
	PreLastStat string `json:"pre_last_stat"`
	A           struct {
		Kpd          float64 `json:"kpd"`
		Win          int     `json:"win"`
		Mission      int     `json:"mission"`
		Death        int     `json:"death"`
		Winrate      float64 `json:"winrate"`
		PrevWinrate  float64 `json:"prev_winrate"`
		Kb           float64 `json:"kb"`
		PrevKb       float64 `json:"prev_kb"`
		KbAir        float64 `json:"kb_air"`
		PrevKbAir    float64 `json:"prev_kb_air"`
		KbGround     float64 `json:"kb_ground"`
		PrevKbGround float64 `json:"prev_kb_ground"`
		Kd           float64 `json:"kd"`
		PrevKd       float64 `json:"prev_kd"`
		KdAir        float64 `json:"kd_air"`
		PrevKdAir    float64 `json:"prev_kd_air"`
		KdGround     float64 `json:"kd_ground"`
		PrevKdGround float64 `json:"prev_kd_ground"`
		Lifetime     int     `json:"lifetime"`
		PrevLifetime int     `json:"prev_lifetime"`
	} `json:"a"`
	R struct {
		Kpd          float64 `json:"kpd"`
		Win          int     `json:"win"`
		Mission      int     `json:"mission"`
		Death        int     `json:"death"`
		Winrate      float64 `json:"winrate"`
		PrevWinrate  float64 `json:"prev_winrate"`
		Kb           float64 `json:"kb"`
		PrevKb       float64 `json:"prev_kb"`
		KbAir        float64 `json:"kb_air"`
		PrevKbAir    float64 `json:"prev_kb_air"`
		KbGround     float64 `json:"kb_ground"`
		PrevKbGround float64 `json:"prev_kb_ground"`
		Kd           float64 `json:"kd"`
		PrevKd       float64 `json:"prev_kd"`
		KdAir        float64 `json:"kd_air"`
		PrevKdAir    float64 `json:"prev_kd_air"`
		KdGround     float64 `json:"kd_ground"`
		PrevKdGround float64 `json:"prev_kd_ground"`
		Lifetime     int     `json:"lifetime"`
		PrevLifetime int     `json:"prev_lifetime"`
	} `json:"r"`
	S struct {
		Kpd          float64 `json:"kpd"`
		Win          int     `json:"win"`
		Mission      int     `json:"mission"`
		Death        int     `json:"death"`
		Winrate      float64 `json:"winrate"`
		PrevWinrate  float64 `json:"prev_winrate"`
		Kb           float64 `json:"kb"`
		PrevKb       float64 `json:"prev_kb"`
		KbAir        float64 `json:"kb_air"`
		PrevKbAir    float64 `json:"prev_kb_air"`
		KbGround     float64 `json:"kb_ground"`
		PrevKbGround float64 `json:"prev_kb_ground"`
		Kd           float64 `json:"kd"`
		PrevKd       float64 `json:"prev_kd"`
		KdAir        float64 `json:"kd_air"`
		PrevKdAir    float64 `json:"prev_kd_air"`
		KdGround     float64 `json:"kd_ground"`
		PrevKdGround float64 `json:"prev_kd_ground"`
		Lifetime     int     `json:"lifetime"`
		PrevLifetime int     `json:"prev_lifetime"`
	} `json:"s"`
}

func (d ThunderSkillData) ToTableGameUser() table.GameUser {
	return table.GameUser{
		TsABRate: d.A.Kpd,
		TsSBRate: d.S.Kpd,
		TsRBRate: d.R.Kpd,
	}
}
