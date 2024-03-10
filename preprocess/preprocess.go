package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"slices"
)

type Pokemon struct {
	No   int    `json:"no"`
	Name string `json:"name"`
}

type Data struct {
	Data []Pokemon `json:"data"`
}

type Word struct {
	Pokemon
	Ruby   string   `json:"reading"`
	Start  string   `json:"start"`
	End    []string `json:"end"`
	IsLast bool     `json:"is_last"`
}

type Vowel string

const (
	VowelA Vowel = "ア"
	VowelI Vowel = "イ"
	VowelU Vowel = "ウ"
	VowelE Vowel = "エ"
	VowelO Vowel = "オ"
)

var vowelMap = map[string]Vowel{}

func init() {
	for _, v := range []string{"ア", "カ", "サ", "タ", "ナ", "ハ", "マ", "ヤ", "ラ", "ワ", "ガ", "ザ", "ダ", "バ", "パ", "ァ", "ャ"} {
		vowelMap[v] = VowelA
	}
	for _, v := range []string{"イ", "キ", "シ", "チ", "ニ", "ヒ", "ミ", "リ", "ギ", "ジ", "ヂ", "ビ", "ピ", "ィ"} {
		vowelMap[v] = VowelI
	}
	for _, v := range []string{"ウ", "ク", "ス", "ツ", "ヌ", "フ", "ム", "ユ", "ル", "グ", "ズ", "ヅ", "ブ", "プ", "ゥ", "ュ", "ッ"} {
		vowelMap[v] = VowelU
	}
	for _, v := range []string{"エ", "ケ", "セ", "テ", "ネ", "ヘ", "メ", "レ", "ゲ", "ゼ", "デ", "ベ", "ペ", "ェ"} {
		vowelMap[v] = VowelE
	}
	for _, v := range []string{"オ", "コ", "ソ", "ト", "ノ", "ホ", "モ", "ヨ", "ロ", "ヲ", "ゴ", "ゾ", "ド", "ボ", "ポ", "ォ", "ョ"} {
		vowelMap[v] = VowelO
	}
}

type SmallKana string

const (
	SmallYa  SmallKana = "ャ"
	SmallYu  SmallKana = "ュ"
	SmallYo  SmallKana = "ョ"
	SmallA   SmallKana = "ァ"
	SmallI   SmallKana = "ィ"
	SmallU   SmallKana = "ゥ"
	SmallE   SmallKana = "ェ"
	SmallO   SmallKana = "ォ"
	SmallTsu SmallKana = "ッ"
)

var largeKanaMap = map[SmallKana]string{}

func init() {
	largeKanaMap[SmallYa] = "ヤ"
	largeKanaMap[SmallYu] = "ユ"
	largeKanaMap[SmallYo] = "ヨ"
	largeKanaMap[SmallA] = "ア"
	largeKanaMap[SmallI] = "イ"
	largeKanaMap[SmallU] = "ウ"
	largeKanaMap[SmallE] = "エ"
	largeKanaMap[SmallO] = "オ"
	largeKanaMap[SmallTsu] = "ツ"
}

type VoicedSoundKana string

const (
	VoicedSoundGa VoicedSoundKana = "ガ"
	VoicedSoundGi VoicedSoundKana = "ギ"
	VoicedSoundGu VoicedSoundKana = "グ"
	VoicedSoundGe VoicedSoundKana = "ゲ"
	VoicedSoundGo VoicedSoundKana = "ゴ"
	VoicedSoundZa VoicedSoundKana = "ザ"
	VoicedSoundJi VoicedSoundKana = "ジ"
	VoicedSoundZu VoicedSoundKana = "ズ"
	VoicedSoundZe VoicedSoundKana = "ゼ"
	VoicedSoundZo VoicedSoundKana = "ゾ"
	VoicedSoundDa VoicedSoundKana = "ダ"
	VoicedSoundDi VoicedSoundKana = "ヂ"
	VoicedSoundDu VoicedSoundKana = "ヅ"
	VoicedSoundDe VoicedSoundKana = "デ"
	VoicedSoundDo VoicedSoundKana = "ド"
	VoicedSoundBa VoicedSoundKana = "バ"
	VoicedSoundBi VoicedSoundKana = "ビ"
	VoicedSoundBu VoicedSoundKana = "ブ"
	VoicedSoundBe VoicedSoundKana = "ベ"
	VoicedSoundBo VoicedSoundKana = "ボ"
	VoicedSoundPa VoicedSoundKana = "パ"
	VoicedSoundPi VoicedSoundKana = "ピ"
	VoicedSoundPu VoicedSoundKana = "プ"
	VoicedSoundPe VoicedSoundKana = "ペ"
	VoicedSoundPo VoicedSoundKana = "ポ"
	VoicedSoundVu VoicedSoundKana = "ヴ"
)

var voicedSoundMap = map[VoicedSoundKana]string{}

func init() {
	voicedSoundMap[VoicedSoundGa] = "カ"
	voicedSoundMap[VoicedSoundGi] = "キ"
	voicedSoundMap[VoicedSoundGu] = "ク"
	voicedSoundMap[VoicedSoundGe] = "ケ"
	voicedSoundMap[VoicedSoundGo] = "コ"
	voicedSoundMap[VoicedSoundZa] = "サ"
	voicedSoundMap[VoicedSoundJi] = "シ"
	voicedSoundMap[VoicedSoundZu] = "ス"
	voicedSoundMap[VoicedSoundZe] = "セ"
	voicedSoundMap[VoicedSoundZo] = "ソ"
	voicedSoundMap[VoicedSoundDa] = "タ"
	voicedSoundMap[VoicedSoundDi] = "チ"
	voicedSoundMap[VoicedSoundDu] = "ツ"
	voicedSoundMap[VoicedSoundDe] = "テ"
	voicedSoundMap[VoicedSoundDo] = "ト"
	voicedSoundMap[VoicedSoundBa] = "ハ"
	voicedSoundMap[VoicedSoundBi] = "ヒ"
	voicedSoundMap[VoicedSoundBu] = "フ"
	voicedSoundMap[VoicedSoundBe] = "ヘ"
	voicedSoundMap[VoicedSoundBo] = "ホ"
	voicedSoundMap[VoicedSoundPa] = "ハ"
	voicedSoundMap[VoicedSoundPi] = "ヒ"
	voicedSoundMap[VoicedSoundPu] = "フ"
	voicedSoundMap[VoicedSoundPe] = "ヘ"
	voicedSoundMap[VoicedSoundPo] = "ホ"
	voicedSoundMap[VoicedSoundVu] = "ウ"
}

// TODO: ニドランの♂と♀を処理する
func main() {
	f, err := os.Open("data/zukan.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var data Data
	json.NewDecoder(f).Decode(&data)

	words := []Word{}
	for _, p := range data.Data {
		var w Word
		w.Pokemon = p

		var ruby string
		switch p.Name {
		case "ニドラン♂":
			ruby = "ニドラオス"
		case "ニドラン♀":
			ruby = "ニドランメス"
		case "ポリゴン2":
			ruby = "ポリゴンツー"
		case "ポリゴンZ":
			ruby = "ポリゴンゼット"
		default:
			ruby = p.Name
		}
		w.Ruby = ruby

		r := []rune(w.Ruby)
		w.Start = string(r[:1])
		endIndex := len(r) - 1
		end := r[endIndex : endIndex+1]

		pattern := `[アイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤユヨラリルレロワンヴガギグゲゴザジズゼゾダヂヅデドバビブベボパピプペポ]`
		re, err := regexp.Compile(pattern)
		if err != nil {
			panic(err)
		}

		// ルールはhttps://w.atwiki.jp/ultimate/pages/16.htmlに準拠
		// 伸ばし棒･小文字･濁音･半濁音を再帰的に処理していく
		for i := endIndex; i > 0; i-- {
			back := false
			// 伸ばし棒で終わる場合
			if string(end) == "ー" {
				endIndex--
				end = r[endIndex : endIndex+1]
				w.End = append(w.End, string(vowelMap[string(end)])) // 伸ばし棒の直前の文字の母音を追加
				back = true
			}

			// 小文字で終わる場合
			if l, ok := largeKanaMap[SmallKana(end)]; ok {
				w.End = append(w.End, string(l)) //大文字に変換して追加

				endIndex--
				end = r[endIndex : endIndex+1]
				back = true
			}

			// 濁音･半濁音で終わる場合
			if l, ok := voicedSoundMap[VoicedSoundKana(end)]; ok {
				w.End = append(w.End, string(l)) // もとの文字に変換して追加
			}

			if !back {
				break
			}
		}
		w.End = append(w.End, string(end))

		w.End = slices.Compact(w.End)

		if len(w.End) == 0 {
			fmt.Println(p.Name)
			panic("end is empty")
		}

		for _, e := range w.End {
			if e == "ン" {
				w.IsLast = true
				break
			}

			if re.FindString(e) == "" {
				panic(fmt.Errorf("illegal char is contained in End words: %s", e))
			}
		}

		words = append(words, w)
	}

	j, err := json.Marshal(words)
	if err != nil {
		panic(err)
	}

	f, err = os.Create("data/words.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.Write(j)
}
