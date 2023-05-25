package wirepod_ttr

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/kercre123/chipper/pkg/vars"
	// "github.com/kercre123/chipper/pkg/logger"
)

var sub10 = map[int]string{
	0: "serraw",
	1: "yeah denn",
	2: "dvah",
	3: "tshi",
	4: "chtehrri",
	5: "pienntch",
	6: "sheshtch",
	7: "shehdehm",
	8: "oshehm",
	9: "jevvienntch",
}

var teen = map[int]string{
	10: "jehshienntch",
	11: "yeah dennashtch yeah",
	12: "dvahnashtch yeah",
	13: "tshinashtch yeah",
	14: "chternashtch yeah",
	15: "piehttnashtch yeah",
	16: "shessnashtch yeah",
	17: "shehdehmnashtch yeah",
	18: "oshehmnashtch yeah",
	19: "jevvienntnashtch yeah",
}

var dec = map[int]string{
	20:   "dvahjehshtchiah",
	30:   "tshijehshtchee",
	40:   "chterjehshtchee",
	50:   "piennjehshawnt",
	60:   "sheshjehshawnt",
	70:   "shehdehmjehshawnt",
	80:   "oshehmjehshawnt",
	90:   "jevviehnjehshawnt",
	100:  "staw",
	200:  "dveaszcie",
	300:  "tshista",
	400:  "chtersta",
	500:  "piennsett",
	600:  "sheshsett",
	700:  "shehdehmset",
	800:  "oshehmsett",
	900:  "jevviehnsett",
	1000: "tyhsionc",
	2000: "dvah tyhsioncy",
	3000: "tshi tyhsioncy",
	4000: "chtehrri tyhsioncy",
	5000: "pienntch tyhsioncy",
	6000: "sheshtch tyhsioncy",
	7000: "shehdehm tyhsioncy",
	8000: "oshehm tyhsioncy",
	9000: "jevvienntch tyhsioncy",
}

// ~40 lines of regexp
var repolon = []map[string]string{
	{`(\b[wz])(\s)`: `${1}`}, // lone letters
	{"rz": "ż"},              //upraszczam polski
	// --> {"([^pt])ch": "\1h",				//ph i th nie mogą powstać
	{"ó": "u"},
	{"ęł": "eł"},
	{"ął": "oł"},
	{`([ea])u`: `${1}ł`},
	{`ci(\B|[^eaouęą])`: `ći${1}`}, // ć, ś, dź i ź w różnych sytuacjach
	{`ci([^e])`: `ć${1}`},
	{`ci`: `ć`},
	{`si(\B|[^eaouęą])`: `śi${1}`},
	{`si([^e])`: `ś${1}`},
	{`si`: `ś`},
	{`dzi(\B|[^eaouęą])`: `dźi${1}`},
	{`dzi([^e])`: `dź${1}`},
	{`dzi`: `dź`},
	{`zi(\B|[^eaouęą])`: `źi${1}`},
	{`zi([^e])`: `ź${1}`},
	{`zi`: `ź`},
	{`ni(\B|[^eaouęą])`: `ńi${1}`}, // i ń
	{`ni([^e])`: `ń${1}`},
	{`ni`: `ń`},
	{`([śptkh])ż(\w+|\p{L})`: `${1}sz${2}`}, // ubezdźwięcznienie
	{`([śptksh])w(\w+|\p{L})`: `${1}f${2}`},
	{`([śptkh])d(\w+|\p{L})`: `${1}t${2}`},
	{`([^cs])z([kptf])`: `${1}s${2}`},
	{`w([kptshf])`: `f${1}`},
	{`ż([kptsfh])`: `sz${1}`},
	{`ż\B`: "sz"},
	{`(\S)w\b`: `${1}f`},
	{`\Bb\b`: "p"},
	{`\Bg\b`: "k"},
	{`\Bd\b`: "t"},
	{`([^cs\s])z(\b|\p{L})`: `${1}s${2}`}, // Widzę -> Widsę
	{`ź\B`: "ś"},
	{`dź\B`: "ć"},
	{`dż\B`: "cz"},
	// --> {"ń": "n",
	// --> {"ą\\B": "om",  // nasals
	{`ę\B`: "e"},
	{`ę([zscśćtdkgż]|cz|sz)`: `en${1}`},
	{`ą([zscśćtdgkż]|cz|sz)`: `on${1}`},
	{`ą([pb])`: `om${1}`},
	{`ę([pb])`: `em${1}`},
}

// ~60 lns of regexp. yet another mess that could be avoided with IPA
var angl = []map[string]string{
	{"ń": "n"},
	{"ą": "om"},
	{`\be\b`: `\BFF\B`},
	{`\B[ji]i`: "i"},
	{`([^pt])ch`: `${1}h`},
	{"ch": "kh"},
	{`([fk])i(e)`: `${1}${2}`},
	{`\bnie\b`: "ne"},
	{`\bmnie\b`: "mne"},
	{`w`: "v"},
	{"ś": "sz"},
	{"ź": "ż"},
	{"ć": "cz"},
	{"dź": "dż"},
	{`s([^z]|\b)`: `ss${1}`},
	{`([^cs]|\b)z`: `${1}s`},
	{`(\b)cz`: `${1}ch`},
	{"cz": "tch"},
	{"sz": "sh"},
	{"dż": "J"},
	{`([^bvzg]|\b)ż`: `${1}zsh`},
	{"ż": "sh"},
	{`c([^h]|\b)`: `ts${1}`},
	{`a([^jłr]|\b)`: `ah${1}`},
	{`([^ji])e([^j]|\b)`: `${1}eh${2}`},
	{`eh\b`: "ehh"},
	{`\bi\b`: `e`},
	{`y`: "I"},
	{`(\w{2,})iej`: `${1} yay`},
	{`(\w{2,})[ij]e`: `${1} yeah`},
	{`\bje(\w+|\p{L})([eioua])`: `yeah ${1}${2}`},
	{`(\w+|\p{L})je(\w+|\p{L})`: `${1}yeh${2}`},
	{`(\w{1})ie`: `${1}ieh`},
	{`i([^eao]|\b)`: `EE${1}`},
	{`ej\b`: "ei"},
	{`ej\B`: "ay"},
	{`EE`: "ee"},
	{`I`: "ih"},
	{`ał`: "ow"},
	{`([aeoi])(\w+|\p{L})oł`: `${1}${2} oh`},
	{"oł": "ow"},
	{"ł": "w"},
	{`(\b\w{1})aj`: `${1}ie `},
	{`(\w{1,})(\w{1})aj(?P<3W>\w+)`: `${1}${2}igh ${3}`},
	{`(\w+|\p{L})(\w{1})aj`: `${1} ${2}ie`},
	{`aj`: " i"},
	{`a\b`: "ah"},
	{"j": "y"},
	{`J`: "dj"},
	{`o([^whym]|\b)`: `aw${1}`},
	{`uw\b`: "ooo"},
	{`u([^w]|\b)`: `oo${1}`},
	{`u(w\w+)`: `oo ${1}`},
	{`g([ei])`: `gh${1}`},
	{`\bb\b`: "bhehh"},
	{`\bts\b`: "tsehh"},
	{`\bd\b`: "dehh"},
	{`\bFF\b`: "ehh"},
	{`\bf\b`: "ehf"},
	{`\bg\b`: "ghiehh"},
	{`\bh\b`: "hah"},
	{`\by\b`: "yacht"},
	{`\bk\b`: "kah"},
	{`\bp\b`: "pehh"},
	{`\bq\b`: "coo"},
	{`\br\b`: "eh"},
	{`\bss\b`: "ess"},
	{`\bt\b`: "tehh"},
	{`\bv\b`: "voo"},
	{`\bx\b`: "eeks"},
	{`\bih\b`: "eegrehk"},
	{`\bs\b`: "seht"},
}

func numToSpeech(n int) string {
	// anglicises given number n. works only for -9999 <= n <= 9999
	if n < 0 {
		return "meenoos " + numToSpeech(-n)
	}
	if n < 10 {
		return sub10[n]
	} else if n >= 10 && n < 20 {
		return teen[n]
	} else if n <= 100 {
		if n%10 == 0 {
			return dec[n]
		} else {
			return dec[n-(n%10)] + " " + sub10[n%10]
		}
	} else if n < 1000 {
		return dec[n-(n%100)] + " " + numToSpeech(n%100)
	} else if n < 10000 {
		return dec[n-(n%1000)] + " " + numToSpeech(n%1000)
	} else {
		fmt.Println("Cannot pronounce numbers with absoulute value above 199")
		return strconv.Itoa(n)
	}
}

func repolonize(s string) string {
	subbed := s
	for _, val_list := range repolon {
		for key, value := range val_list {
			subbed = regexp.MustCompile(key).ReplaceAllString(subbed, value)
		}
	}
	return subbed
}

func anglicise(s string) string {
	// anglicises simplified polish text
	subbed := []string{}
	for _, w := range strings.Split(s, " ") {
		sub := w
		if match, _ := regexp.MatchString(`(-|)\d+`, w); match {
			re := regexp.MustCompile(`(-|)\d+`)
			allMatches := re.FindStringSubmatch(w)
			i, _ := strconv.Atoi(allMatches[0])
			sub = numToSpeech(i) + strings.Join(regexp.MustCompile(`\D+`).FindAllString(w, -1), "")
		} else if match, _ := regexp.MatchString(`(\w+[^\d-])((-|)\d+)`, w); match {
			// eg "kot18"
			re := regexp.MustCompile(`(\w+[^\d-])((-|)\d+)`)
			allMatches := re.FindStringSubmatch(w)
			i, _ := strconv.Atoi(allMatches[2])
			sub = anglicise(allMatches[1]) + " " + numToSpeech(i)
		} else if match, _ := regexp.MatchString(`(.)(\d+)`, w); match {
			// eg "'-7" should be dealt with as well
			re := regexp.MustCompile(`(.)(\d+)`)
			allMatches := re.FindStringSubmatch(w)
			i, _ := strconv.Atoi(allMatches[2])
			sub = numToSpeech(i) // + strings.Join(regexp.MustCompile(`\D+`).FindAllString(w, -1), "")
		} else {
			for _, val_list := range angl {
				for key, value := range val_list {
					sub = regexp.MustCompile(key).ReplaceAllString(sub, value)
				}
			}
		}
		subbed = append(subbed, sub)
	}
	return strings.Join(subbed, " ")
}

func ConvertFromPolishToEnglishLector(text string) string {
	rep := repolonize(text)
	return anglicise(rep)
}

func ConvertSpeechText(text string) string {
	if vars.APIConfig.STT.Language == "pl-PL" {
		return ConvertFromPolishToEnglishLector(text)
	}
	return text
}
