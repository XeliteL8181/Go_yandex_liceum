package main

import (
	"fmt"
	"sort"
	"strings"
)

func AnalyzeText(text string) {
    lowerText := strings.ToLower(text)

    replaceText := strings.Replace(lowerText, ".", " ", -1)
    replaceText = strings.Replace(replaceText, ",", " ", -1)
    replaceText = strings.Replace(replaceText, "?", " ", -1)
    replaceText = strings.Replace(replaceText, "!", " ", -1)

    splitText := strings.Fields(replaceText)
    mapText := map[string]int{}
    for _, value := range splitText {
        _, found := mapText[value]
        if !found {
            mapText[value] = 1
        } else {
            mapText[value]++
        }
    }
    
    otv := getTopWords(mapText, 5)
    fmt.Println("Количество слов:", len(splitText))
    fmt.Println("Количество уникальных слов:", len(mapText))
    fmt.Printf(
            "Самое часто встречающееся слово: \"%s\" (встречается %d раз)\n",
            otv[0], mapText[otv[0]],
    )
    fmt.Println("Топ-5 самых часто встречающихся слов:")
    for i := range otv {
        fmt.Printf(
            "\"%s\": %d раз\n",
            otv[i], mapText[otv[i]],
    )
    }
}

func getTopWords(wordMap map[string]int, n int) []string {
    topWords := make([]string, n)
    replaceKeysAndValues := map[int]string{}
    for key, value := range wordMap {
        _, found := replaceKeysAndValues[value]
        if !found {
            replaceKeysAndValues[value] = key
        }
    }

    sortKeys := []int{}
    for k := range replaceKeysAndValues {
        sortKeys = append(sortKeys, k)
    }
    sort.Ints(sortKeys)
    s := 0

    for i := len(sortKeys) - 1; i >= 0; i-- {
        topWords[s] = replaceKeysAndValues[sortKeys[i]]
        s++
        if s == 5 {
            break
        }
    }
    return topWords[:n]
}

func main() {
    AnalyzeText("one two two three three three four four four four five five five five five")
	AnalyzeText("Go очень очень очень ОЧЕНЬ ОчЕнь очень оЧЕНь классный классный! go просто, ну просто классный. GO Классный!")
	AnalyzeText("Я так люблю море. Я на море. Я так люблю плавать. Море! Я море!!! ЛЮБЛЮ МОРЕ")
}