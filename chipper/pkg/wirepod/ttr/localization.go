package wirepod_ttr

import "github.com/kercre123/chipper/pkg/vars"

const STR_WEATHER_IN = "str_weather_in"
const STR_WEATHER_FORECAST = "str_weather_forecast"
const STR_WEATHER_TOMORROW = "str_weather_tomorrow"
const STR_WEATHER_THE_DAY_AFTER_TOMORROW = "str_weather_the_day_after_tomorrow"
const STR_WEATHER_TONIGHT = "str_weather_tonight"
const STR_WEATHER_THIS_AFTERNOON = "str_weather_this_afternoon"
const STR_EYE_COLOR_PURPLE = "str_eye_color_purple"
const STR_EYE_COLOR_BLUE = "str_eye_color_blue"
const STR_EYE_COLOR_SAPPHIRE = "str_eye_color_sapphire"
const STR_EYE_COLOR_YELLOW = "str_eye_color_yellow"
const STR_EYE_COLOR_TEAL = "str_eye_color_teal"
const STR_EYE_COLOR_TEAL2 = "str_eye_color_teal2"
const STR_EYE_COLOR_GREEN = "str_eye_color_green"
const STR_EYE_COLOR_ORANGE = "str_eye_color_orange"
const STR_ME = "str_me"
const STR_SELF = "str_self"
const STR_VOLUME_LOW = "str_volume_low"
const STR_VOLUME_QUIET = "str_volume_quiet"
const STR_VOLUME_MEDIUM_LOW = "str_volume_medium_low"
const STR_VOLUME_MEDIUM = "str_volume_medium"
const STR_VOLUME_NORMAL = "str_volume_normal"
const STR_VOLUME_REGULAR = "str_volume_regular"
const STR_VOLUME_MEDIUM_HIGH = "str_volume_medium_high"
const STR_VOLUME_HIGH = "str_volume_high"
const STR_VOLUME_LOUD = "str_volume_loud"
const STR_VOLUME_MUTE = "str_volume_mute"
const STR_VOLUME_NOTHING = "str_volume_nothing"
const STR_VOLUME_SILENT = "str_volume_silent"
const STR_VOLUME_OFF = "str_volume_off"
const STR_VOLUME_ZERO = "str_volume_zero"
const STR_NAME_IS = "str_name_is"
const STR_NAME_IS2 = "str_name_is1"
const STR_NAME_IS3 = "str_name_is2"
const STR_FOR = "str_for"

// All text must be lowercase!

var texts = map[string][]string{
	//  key                 						en-US		it-IT		es-ES		fr-FR		de-DE
	STR_WEATHER_IN:                     {" in ", " a ", " en ", " en ", " in ", " w "},
	STR_WEATHER_FORECAST:               {"forecast", "previsioni", "pronóstico", "prévisions", "wettervorhersage", "prognoza"},
	STR_WEATHER_TOMORROW:               {"tomorrow", "domani", "mañana", "demain", "morgen", "jutro"},
	STR_WEATHER_THE_DAY_AFTER_TOMORROW: {"day after tomorrow", "dopodomani", "el día después de mañana", "lendemain de demain", "am tag nach morgen", "pojutrze"},
	STR_WEATHER_TONIGHT:                {"tonight", "stasera", "esta noche", "ce soir", "heute abend", "dziś wieczorem"},
	STR_WEATHER_THIS_AFTERNOON:         {"afternoon", "pomeriggio", "esta tarde", "après-midi", "heute nachmittag", "popołudniu"},
	STR_EYE_COLOR_PURPLE:               {"purple", "lilla", "violeta", "violet", "violett", "fioletowy"},
	STR_EYE_COLOR_BLUE:                 {"blue", "blu", "azul", "bleu", "blau", "niebieski"},
	STR_EYE_COLOR_SAPPHIRE:             {"sapphire", "zaffiro", "zafiro", "saphir", "saphir", "szafir"},
	STR_EYE_COLOR_YELLOW:               {"yellow", "giallo", "amarillo", "jaune", "gelb", "żółty"},
	STR_EYE_COLOR_TEAL:                 {"teal", "verde acqua", "verde azulado", "sarcelle", "blaugrün", "morski"},
	STR_EYE_COLOR_TEAL2:                {"tell", "acquamarina", "aguamarina", "acquamarina", "acquamarina", "akwamaryn"},
	STR_EYE_COLOR_GREEN:                {"green", "verde", "verde", "vert", "grün", "zielony"},
	STR_EYE_COLOR_ORANGE:               {"orange", "arancio", "naranja", "orange", "orange", "pomarańczowy"},
	STR_ME:                             {"me", "me", "me", "moi", "mir", "mnie"},
	STR_SELF:                           {"self", "mi", "mía", "moi", "mein", "ja"},
	STR_VOLUME_LOW:                     {"low", "basso", "bajo", "bas", "niedrig", "niski"},
	STR_VOLUME_QUIET:                   {"quiet", "poco rumoroso", "tranquilo", "silencieux", "ruhig", "cichy"},
	STR_VOLUME_MEDIUM_LOW:              {"medium low", "medio basso", "medio-bajo", "moyen-doux", "mittelschwer", "średnio niski"},
	STR_VOLUME_MEDIUM:                  {"medium", "medio", "medio", "moyen", "mittel", "średni"},
	STR_VOLUME_NORMAL:                  {"normal", "normale", "normal", "normal", "normal", "normalny"},
	STR_VOLUME_REGULAR:                 {"regular", "regolare", "regular", "régulier", "regulär", "zwyczajny"},
	STR_VOLUME_MEDIUM_HIGH:             {"medium high", "medio alto", "medio-alto", "moyen-élevé", "mittelhoch", "średno wysoki"},
	STR_VOLUME_HIGH:                    {"high", "alto", "alto", "élevé", "hoch", "wysoki"},
	STR_VOLUME_LOUD:                    {"loud", "rumoroso", "fuerte", "fort", "laut", "głośny"},
	STR_VOLUME_MUTE:                    {"mute", "muto", "mudo", "", "stumm", "wyciszony"},
	STR_VOLUME_NOTHING:                 {"nothing", "nessuno", "nada", "rien", "nichts", "nic"},
	STR_VOLUME_SILENT:                  {"silent", "silenzioso", "silencio", "silencieux", "still", "cichy"},
	STR_VOLUME_OFF:                     {"off", "spento", "apagado", "éteindre", "aus", "wyłączony"},
	STR_VOLUME_ZERO:                    {"zero", "zero", "cero", "zéro", "null", "zero"},
	STR_NAME_IS:                        {" is ", " è ", " es ", " est ", " ist ", " to "},
	STR_NAME_IS2:                       {"'s", "sono ", "soy ", "suis ", "bin ", " się "},
	STR_NAME_IS3:                       {"names", " chiamo ", " llamo ", "appelle ", "werde", "imię"},
	STR_FOR:                            {" for ", " per ", " para ", " pour ", " für ", " dla "},
}

var texts_pl_PL = map[string][]string{
	//  key                 						pl-PL
	STR_WEATHER_IN:                     {" w ", " we ", " na "},
	STR_WEATHER_FORECAST:               {"prognoza", "pogoda"},
	STR_WEATHER_TOMORROW:               {"jutro"},
	STR_WEATHER_THE_DAY_AFTER_TOMORROW: {"pojutrze"},
	STR_WEATHER_TONIGHT:                {"dziś wieczorem", "wieczorem"},
	STR_WEATHER_THIS_AFTERNOON:         {"popołudniu"},
	STR_EYE_COLOR_PURPLE:               {"fioletowy"},
	STR_EYE_COLOR_BLUE:                 {"niebieski"},
	STR_EYE_COLOR_SAPPHIRE:             {"szafir"},
	STR_EYE_COLOR_YELLOW:               {"żółty"},
	STR_EYE_COLOR_TEAL:                 {"morski"},
	STR_EYE_COLOR_TEAL2:                {"akwamaryn"},
	STR_EYE_COLOR_GREEN:                {"zielony"},
	STR_EYE_COLOR_ORANGE:               {"pomarańczowy"},
	STR_ME:                             {"mi", "mnie"},
	STR_SELF:                           {"ja", "siebie"},
	STR_VOLUME_LOW:                     {"niski"},
	STR_VOLUME_QUIET:                   {"cichy"},
	STR_VOLUME_MEDIUM_LOW:              {"średnio niski"},
	STR_VOLUME_MEDIUM:                  {"średni"},
	STR_VOLUME_NORMAL:                  {"normalny"},
	STR_VOLUME_REGULAR:                 {"zwyczajny", "regularny"},
	STR_VOLUME_MEDIUM_HIGH:             {"średno wysoki"},
	STR_VOLUME_HIGH:                    {"wysoki"},
	STR_VOLUME_LOUD:                    {"głośny"},
	STR_VOLUME_MUTE:                    {"wyciszony"},
	STR_VOLUME_NOTHING:                 {"nic"},
	STR_VOLUME_SILENT:                  {"cichy"},
	STR_VOLUME_OFF:                     {"wyłączony"},
	STR_VOLUME_ZERO:                    {"zero", "nic"},
	STR_NAME_IS:                        {" to ", "jestem "},
	STR_NAME_IS2:                       {" się "},
	STR_NAME_IS3:                       {"imię"},
	STR_FOR:                            {" dla "},
}

func getText(key string) list[string] {
	var data = texts[key] 
	if data != nil {
		if vars.APIConfig.STT.Language == "it-IT" {
			return {data[1]}
		} else if vars.APIConfig.STT.Language == "es-ES" {
			return {data[2]}
		} else if vars.APIConfig.STT.Language == "fr-FR" {
			return {data[3]}
		} else if vars.APIConfig.STT.Language == "de-DE" {
			return {data[4]}
		} else if vars.APIConfig.STT.Language == "pl-PL" {
			return texts_pl_PL[key]
		}
	}
	return {data[0],}
}