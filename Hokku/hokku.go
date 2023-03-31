// hokku
package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	lenArrow = 10
)

func main() {

	for i := ""; i != "Стоп"; {
		stringOne := defineRandom()
		stringTwo := defineDayBirth()
		stringThree := defineDayMonth()
		fmt.Println("")
		printHokku(stringOne, stringTwo, stringThree)
		if err := writeHokku(stringOne, stringTwo, stringThree); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		i = anotherHokku()
	}
	fmt.Println("Программа завершена")
}

func defineRandom() string {
	baseRandom := [lenArrow]string{
		"Целый день спят ночные цветы,",
		"Что ты видишь во взоре моём,",
		"Раскрываются тихо листы,",
		"Буря мглою небо кроет,",
		"Хочется рухнуть в траву непомятую,",
		"Мои мечты - священные чертоги,",
		"Мы разошлись на полпути,",
		"Мне страшно с тобой встречаться,",
		"Чистый ветер ели колышет,",
		"Отдыхает моя земля,",
	}
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(lenArrow)
	return baseRandom[idx]
}

func defineDayBirth() string {
	fmt.Println("Введите день вашего рождения")
	var dayBirth int
	fmt.Scan(&dayBirth)
	dayBirth %= 10
	baseDayBirth := [lenArrow]string{
		"И мглой волнистою покрыты небеса.",
		"Пришла весна - и небо прояснится,",
		"Хожу к местам заветного свиданья,",
		"Но ничего в прошедшем мне не жаль.",
		"Воспоминанье слишком давит плечи.",
		"Пальцы в рот - и весёлый свист.",
		"Захотел я уйти в себя - а там никого.",
		"И вся планета распахнулась для меня!",
		"В дом возвращаюсь.",
		"Что, если сердце бурно оборвётся?",
	}
	return baseDayBirth[dayBirth]
}

func defineDayMonth() string {
	fmt.Println("Введите номер месяца вашего рождения от 1 до 12")
	var dayBirth int
	fmt.Scan(&dayBirth)
	dayBirth %= 10
	baseDayMonth := [lenArrow]string{
		"Снова дралась во дворе?",
		"В ушах обрывки тёплого бала.",
		"Не стоит в старом рыться.",
		"Так мало пройдено дорог.",
		"При мне моя свобода.",
		"И я слышу, как сердце цветёт.",
		"Не все, кто блуждают - потеряны.",
		"Для красоты дистанция нужна.",
		"Что в имене тебе моём?",
		"Мне совершенно всё равно.",
	}
	return baseDayMonth[dayBirth]
}

func writeHokku(s1, s2, s3 string) error {
	fileName := "hokku.txt"
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	s1 = fmt.Sprintln(s1) //следующая строка в файле начнётся с новой строки
	s2 = fmt.Sprintln(s2)
	s3 = fmt.Sprintln(s3)

	if _, err = file.WriteString(s1); err != nil {
		return err
	}
	if _, err = file.WriteString(s2); err != nil {
		return err
	}
	if _, err = file.WriteString(s3); err != nil {
		return err
	}

	if _, err = file.WriteString("\n"); err != nil {
		return err
	}
	return nil
}

func printHokku(s1, s2, s3 string) {
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println("")
	return
}

func anotherHokku() string {
	fmt.Println("Ещё хокку? Введите да или нет:")
	var answer string
	fmt.Scan(&answer)
	fmt.Println("")
	if answer == "нет" {
		return "Стоп"
	}
	return ""
}
