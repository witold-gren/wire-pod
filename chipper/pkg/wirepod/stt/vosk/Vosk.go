package wirepod_vosk

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	vosk "github.com/alphacep/vosk-api/go"
	"github.com/kercre123/chipper/pkg/logger"
	"github.com/kercre123/chipper/pkg/vars"
	"github.com/kercre123/chipper/pkg/wirepod/localization"
	sr "github.com/kercre123/chipper/pkg/wirepod/speechrequest"
)

var Name string = "vosk"

var model *vosk.VoskModel
var recsmu sync.Mutex
var grmRecs []ARec
var gpRecs []ARec
var modelLoaded bool

type ARec struct {
	InUse bool
	Rec   *vosk.VoskRecognizer
}

var Grammer string

func Init() error {
	if vars.APIConfig.PastInitialSetup {
		vosk.SetLogLevel(-1)
		if modelLoaded {
			logger.Println("A model was already loaded, freeing")
			model.Free()
		}
		sttLanguage := vars.APIConfig.STT.Language
		if len(sttLanguage) == 0 {
			sttLanguage = "en-US"
		}
		modelPath := "../vosk/models/" + sttLanguage + "/model"
		logger.Println("Opening VOSK model (" + modelPath + ")")
		aModel, err := vosk.NewModel(modelPath)
		if err != nil {
			log.Fatal(err)
			return err
		}
		model = aModel
		Grammer = GetGrammerList(vars.APIConfig.STT.Language)
		fmt.Println(Grammer)
		grmRecognizer, err := vosk.NewRecognizerGrm(aModel, 16000.0, Grammer)
		//aRecognizer, err := vosk.NewRecognizer(aModel, 16000.0)
		if err != nil {
			log.Fatal(err)
		}
		var grmrec ARec
		grmrec.Rec = grmRecognizer
		grmrec.InUse = false
		grmRecs = append(grmRecs, grmrec)
		gpRecognizer, err := vosk.NewRecognizer(aModel, 16000.0)
		var gprec ARec
		gprec.Rec = gpRecognizer
		gprec.InUse = false
		gpRecs = append(gpRecs, gprec)
		if err != nil {
			log.Fatal(err)
		}

		modelLoaded = true
		logger.Println("VOSK initiated successfully")
	}
	return nil
}

func getRec(withGrm bool) (*vosk.VoskRecognizer, int) {
	recsmu.Lock()
	defer recsmu.Unlock()
	if withGrm {
		for ind, rec := range grmRecs {
			if !rec.InUse {
				grmRecs[ind].InUse = true
				return grmRecs[ind].Rec, ind
			}
		}
	} else {
		for ind, rec := range gpRecs {
			if !rec.InUse {
				gpRecs[ind].InUse = true
				return gpRecs[ind].Rec, ind
			}
		}
	}
	var newrec ARec
	var newRec *vosk.VoskRecognizer
	var err error
	newrec.InUse = true
	if withGrm {
		newRec, err = vosk.NewRecognizerGrm(model, 16000.0, Grammer)
	} else {
		newRec, err = vosk.NewRecognizer(model, 16000.0)
	}
	if err != nil {
		log.Fatal(err)
	}
	newrec.Rec = newRec
	if withGrm {
		grmRecs = append(grmRecs, newrec)
		return grmRecs[len(grmRecs)-1].Rec, len(grmRecs) - 1
	} else {
		gpRecs = append(gpRecs, newrec)
		return gpRecs[len(gpRecs)-1].Rec, len(gpRecs) - 1
	}
}

func STT(req sr.SpeechRequest) (string, error) {
	logger.Println("(Bot " + strconv.Itoa(req.BotNum) + ", Vosk) Processing...")
	speechIsDone := false
	bTime := time.Now()
	var withGrm bool
	if vars.APIConfig.Knowledge.IntentGraph || req.IsKG {
		withGrm = true
	}
	rec, recind := getRec(withGrm)
	rec.SetWords(0)
	rec.AcceptWaveform(req.FirstReq)
	for {
		chunk, err := req.GetNextStreamChunk()
		if err != nil {
			return "", err
		}
		rec.AcceptWaveform(chunk)
		// has to be split into 320 []byte chunks for VAD
		speechIsDone = req.DetectEndOfSpeech()
		if speechIsDone {
			break
		}
	}
	var jres map[string]interface{}
	json.Unmarshal([]byte(rec.FinalResult()), &jres)
	if withGrm {
		grmRecs[recind].InUse = false
	} else {
		gpRecs[recind].InUse = false
	}
	fmt.Println("Process took: ", time.Now().Sub(bTime))
	transcribedText := jres["text"].(string)
	logger.Println("Bot " + strconv.Itoa(req.BotNum) + " Transcribed text: " + transcribedText)
	return transcribedText, nil
}

// more performance can be gotten via grammar

func removeDuplicates(slice []string) []string {
	seen := make(map[string]bool)
	result := []string{}
	for _, val := range slice {
		if _, ok := seen[val]; !ok {
			seen[val] = true
			result = append(result, val)
		}
	}
	return result
}

func GetGrammerList(lang string) string {
	var wordsList []string
	var grammer string
	// add words in intent json
	for _, words := range vars.MatchListList {
		for _, word := range words {
			wors := strings.Split(word, " ")
			for _, wor := range wors {
				found := model.FindWord(wor)
				if found != -1 {
					wordsList = append(wordsList, wor)
				}
			}
		}
	}
	// add words in localization
	for _, str := range localization.ALL_STR {
		text := localization.GetText(str)
		wors := strings.Split(text, " ")
		for _, wor := range wors {
			found := model.FindWord(wor)
			if found != -1 {
				wordsList = append(wordsList, wor)
			}
		}
	}
	// add custom intent matches
	for _, intent := range vars.CustomIntents {
		for _, utterance := range intent.Utterances {
			wors := strings.Split(utterance, " ")
			for _, wor := range wors {
				found := model.FindWord(wor)
				if found != -1 {
					wordsList = append(wordsList, wor)
				}
			}
		}
	}

	wordsList = removeDuplicates(wordsList)
	for i, word := range wordsList {
		if i == len(wordsList)-1 {
			grammer = grammer + `"` + word + `"`
		} else {
			grammer = grammer + `"` + word + `"` + ", "
		}
	}
	grammer = "[" + grammer + "]"
	return grammer
}
