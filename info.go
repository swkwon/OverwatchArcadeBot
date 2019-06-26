package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// ArcadeItem ...
type ArcadeItem struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Players string `json:"players"`
	Code    string `json:"code"`
}

// ArcadeInfo ...
type ArcadeInfo struct {
	UpdateAt      string     `json:"updated_at"`
	TileLarge     ArcadeItem `json:"tile_large"`
	TileWeekly1   ArcadeItem `json:"tile_weekly_1"`
	TileDaily     ArcadeItem `json:"tile_daily"`
	TileWeekly2   ArcadeItem `json:"tile_weekly_2"`
	TilePermanent ArcadeItem `json:"tile_permanent"`
}

const (
	overwatchArcade = "https://overwatcharcade.today/api/today"
	telegramMsg     = `CHANGES DAILY\n%s\n%s\n\nCHANGES DAILY\n%s\n%s\n\nCHANGES WEEKLY\n%s\n%s\n\nCHANGES WEEKLY\n%s\n%s\n\nPERMANENT\n%s\n%s`
)

var translateMap map[string]string

func init() {
	translateMap["1v1"] = "1대1"
	translateMap["3v3"] = "3대3"
	translateMap["4v4"] = "4대4"
	translateMap["5v1"] = "5대1"
	translateMap["6v6"] = "6대6"
	translateMap["8 Player FFA"] = "8인 개별전투"
	translateMap["Capture the Flag"] = "깃발뺏기"
	translateMap["Ch\u00e2teau Deathmatch"] = "샤토 데스매치"
	translateMap["Competitive CTF"] = "깃발뺏기 경쟁전"
	translateMap["Competitive Deathmatch"] = "데스매치 경쟁전"
	translateMap["Competitive Elimination"] = "섬멸전 경쟁전"
	translateMap["Competitive Team Deathmatch"] = "팀 데스매치 경쟁전"
	translateMap["Co-op"] = "협동"
	translateMap["Copa L\u00facioball"] = "루시우볼 경쟁전"
	translateMap["CTF: Ayutthaya Only"] = "깃발뺏기 아유타야"
	translateMap["CTF: Busan"] = "깃발뺏기 부산"
	translateMap["Deathmatch"] = "데스매치"
	translateMap["Elimination"] = "섬멸전"
	translateMap["Havana"] = "하바나"
	translateMap["Junkenstein's Revenge"] = "정켄슈타인의 복수"
	translateMap["L\u00facioball"] = "루시우볼"
	translateMap["Limited Duel"] = "진검승부"
	translateMap["Lockout Elimination"] = "승자제외 섬멸전"
	translateMap["Low Gravity"] = "저중력"
	translateMap["Mei's Snowball Offensive"] = "메이의 눈싸움 대작전"
	translateMap["Mission Archives"] = "임부기록 보관소"
	translateMap["Mystery Deathmatch"] = "수수께끼 데스매치"
	translateMap["Mystery Duel"] = "수수께끼의 결투"
	translateMap["Mystery Heroes"] = "수수께끼의 영웅"
	translateMap["No Limits"] = "똑같은 영웅도 환영"
	translateMap["Paris"] = "파리"
	translateMap["Petra Deathmatch"] = "페트라 데스매치"
	translateMap["Retribution"] = "응징의 날"
	translateMap["Retribution (All Heroes)"] = "응징의 날: 모든 영웅"
	translateMap["Retribution (Story)"] = "응징의 날: 스토리"
	translateMap["Storm Rising (All Heroes)"] = "폭풍의 서막: 모든 영웅"
	translateMap["Storm Rising (Story)"] = "폭풍의 서막: 스토리"
	translateMap["Team Deathmatch"] = "팀 데스매치"
	translateMap["Total Mayhem"] = "완전 난장판"
	translateMap["Uprising"] = "옴닉의 반란"
	translateMap["Uprising (All Heroes)"] = "옴닉의 반란: 모든 영웅"
	translateMap["Uprising (Story)"] = "옴닉의 반란: 스토리"
	translateMap["Yeti Hunter"] = "예티 사냥꾼"

}

func getArcadeInfo() (*ArcadeInfo, error) {
	res, err := http.Get(overwatchArcade)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	info := &ArcadeInfo{}
	if err := json.Unmarshal(data, info); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return info, nil
}

func translate(origin string) string {
	if v, ok := translateMap[origin]; ok == true {
		return v
	}
	return origin
}

func makeText(info *ArcadeInfo) string {
	return fmt.Sprintf(telegramMsg,
		translate(info.TileLarge.Players), translate(info.TileLarge.Name),
		translate(info.TileDaily.Players), translate(info.TileDaily.Name),
		translate(info.TileWeekly1.Players), translate(info.TileWeekly1.Name),
		translate(info.TileWeekly2.Players), translate(info.TileWeekly2.Name),
		translate(info.TilePermanent.Players), translate(info.TilePermanent.Name))
}
