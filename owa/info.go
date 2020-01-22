package owa

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// ArcadeItem ...
type ArcadeItem struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Players string  `json:"players"`
	Code    string  `json:"code"`
	Label   *string `json:"label"`
}

// ArcadeInfo ...
type ArcadeInfo struct {
	CreatedAt string     `json:"created_at"`
	Tile1     ArcadeItem `json:"tile_1"`
	Tile2     ArcadeItem `json:"tile_2"`
	Tile3     ArcadeItem `json:"tile_3"`
	Tile4     ArcadeItem `json:"tile_4"`
	Tile5     ArcadeItem `json:"tile_5"`
	Tile6     ArcadeItem `json:"tile_6"`
	Tile7     ArcadeItem `json:"tile_7"`
}

const (
	overwatchArcade = "https://overwatcharcade.today/api/today"
)

var translateMap map[string]string

func init() {
	translateMap = make(map[string]string)
	translateMap["1v1"] = "1대1"
	translateMap["3v3"] = "3대3"
	translateMap["4v4"] = "4대4"
	translateMap["5v1"] = "5대1"
	translateMap["6v6"] = "6대6"
	translateMap["8 Player FFA"] = "8인 개별전투"
	translateMap["Capture the Flag"] = "\U0001f3c1깃발뺏기\U0001f3c1"
	translateMap["Ch\u00e2teau Deathmatch"] = "\U0001f3f0샤토 데스매치\U00002620"
	translateMap["Competitive CTF"] = "\U0001f3c1깃발뺏기 경쟁전\U0001f3c6"
	translateMap["Competitive Deathmatch"] = "\U00002620데스매치 경쟁전\U0001f3c6"
	translateMap["Competitive Elimination"] = "\U0001f52b섬멸전 경쟁전\U0001f3c6"
	translateMap["Competitive Team Deathmatch"] = "\U00002620팀 데스매치 경쟁전\U0001f3c6"
	translateMap["Co-op"] = "\U0001F91D협동\U0001F91D"
	translateMap["Copa L\u00facioball"] = "\U000026bd루시우볼 경쟁전\U0001f3c6"
	translateMap["CTF: Ayutthaya Only"] = "\U0001f3c1깃발뺏기 아유타야\U0001f1f9\U0001f1ed"
	translateMap["CTF: Busan"] = "\U0001f3c1깃발뺏기 부산\U0001f1f0\U0001f1f7"
	translateMap["Deathmatch"] = "\U00002620데스매치\U00002620"
	translateMap["Elimination"] = "\U0001f52b섬멸전\U0001f52b"
	translateMap["Havana"] = "\U0001f1e8\U0001f1fa하바나\U0001f1e8\U0001f1fa"
	translateMap["Junkenstein's Revenge"] = "\U0001f44a정켄슈타인의 복수\U0001f44a"
	translateMap["L\u00facioball"] = "\U000026bd루시우볼\U000026bd"
	translateMap["Limited Duel"] = "\U00002694진검승부\U00002694"
	translateMap["Lockout Elimination"] = "\U0001f52b승자제외 섬멸전\U0001f52b"
	translateMap["Low Gravity"] = "\U0001f92a저중력\U0001f628"
	translateMap["Mei's Snowball Offensive"] = "\U00002744메이의 눈싸움 대작전\U000026c4"
	translateMap["Mission Archives"] = "\U00002733임무기록 보관소\U0001f4dc"
	translateMap["Mystery Deathmatch"] = "\U0001f648수수께끼 데스매치\U00002620"
	translateMap["Mystery Duel"] = "\U0001f648수수께끼의 결투\U0001f44a"
	translateMap["Mystery Heroes"] = "\U0001f648수수께끼의 영웅\U0001f923"
	translateMap["No Limits"] = "\U0001f923똑같은 영웅도 환영\U0001f44b"
	translateMap["Paris"] = "\U0001f1eb\U0001f1f7파리\U0001f1eb\U0001f1f7"
	translateMap["Petra Deathmatch"] = "\U0001f1ef\U0001f1f4페트라 데스매치\U00002620"
	translateMap["Retribution"] = "\U0000270a응징의 날\U0001f4c5"
	translateMap["Retribution (All Heroes)"] = "\U0000270a응징의 날: 모든 영웅\U0001f923"
	translateMap["Retribution (Story)"] = "\U0000270a응징의 날: 스토리\U0000270a"
	translateMap["Storm Rising (All Heroes)"] = "\U0001f300폭풍의 서막: 모든 영웅\U0001f923"
	translateMap["Storm Rising (Story)"] = "\U0001f300폭풍의 서막: 스토리\U0001f923"
	translateMap["Team Deathmatch"] = "\U00002620팀 데스매치\U00002620"
	translateMap["Total Mayhem"] = "\U0001f635완전 난장판\U0001f4a5"
	translateMap["Uprising"] = "\U0001f916옴닉의 반란\U0001f6a9"
	translateMap["Uprising (All Heroes)"] = "\U0001f916옴닉의 반란: 모든 영웅\U0001f923"
	translateMap["Uprising (Story)"] = "\U0001f916옴닉의 반란: 스토리\U0001f916"
	translateMap["Yeti Hunter"] = "\U0001f649예티 사냥꾼\U0001f52b"
	translateMap["Mirrored Deathmatch"] = "\U0001f91c미러전 데스매치\U0001f91b"
	translateMap["Hero Gauntlet"] = "\U0001f923영웅 건틀릿\U0001f94a"
	translateMap["CTF Blitz"] = "\U0001f6a9깃발뺏기 속공전\U0001f6a9"
	translateMap["Quick Play Classic"] = "\U0001f938빠른대전 클래식\U0001f938"
}

// GetArcadeInfo ...
func GetArcadeInfo() ([]*ArcadeInfo, error) {
	res, err := http.Get(overwatchArcade)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var info []*ArcadeInfo
	if err := json.Unmarshal(data, &info); err != nil {
		log.Println(err, string(data))
		return nil, fmt.Errorf("got %s, error: %s", string(data), err.Error())
	}
	return info, nil
}

func translate(origin string) string {
	if v, ok := translateMap[origin]; ok == true {
		return v
	}
	return origin
}

func makeTileInfo(item *ArcadeItem) string {
	return translate(item.Players) + " " + translate(item.Name)
}

// MakeText ...
func MakeText(info []*ArcadeInfo) (string, error) {
	for _, v := range info {
		var created string
		if t, e := time.Parse("2006-01-02 15:04:05", v.CreatedAt); e != nil {
			created = v.CreatedAt
		} else {
			created = t.Format("2006.01.02")
		}

		return fmt.Sprintf(`%s %s\n\n%s\n%s\n%s\n%s\n%s\n%s\n%s`,
			created,
			"Today's Overwatch arcade 😍",
			makeTileInfo(&v.Tile1),
			makeTileInfo(&v.Tile2),
			makeTileInfo(&v.Tile3),
			makeTileInfo(&v.Tile4),
			makeTileInfo(&v.Tile5),
			makeTileInfo(&v.Tile6),
			makeTileInfo(&v.Tile7)), nil
	}
	return "", errors.New("no arcade information")
}
