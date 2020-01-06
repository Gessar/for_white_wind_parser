package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	// "github.com/crufter/goquery"
	"github.com/PuerkitoBio/goquery"
)

func inputItem(curStr *string) { //Выбор артикула, если запустили программу без аргументов
	valid := false //Валидность введенных данных
	fmt.Print("Enter item number: ")
	for {
		valid = true
		fmt.Scanln(&*curStr)        //Считаем введенную строку
		for _, v := range *curStr { //Проходим по ней
			if v < 48 || v > 57 || v == 32 { //Если там нет цифр
				valid = false //Невалидно
			}
		}
		if valid { //Если все валидно, то выходим из лупа
			break
		} else {
			fmt.Println("Error - invalid input method. Must be numerlic (like 123456)")
			fmt.Print("Enter item number: ") //Просим еще раз ввести артикул
			continue
		}
	}
	return
}

func sectionSelect() int { //Выбор сегмента
	valid := true          //Валиность
	var selectInput string //Введенный сегмент
	fmt.Println("1 - Дом & Семья")
	fmt.Println("2 - Подключайся & Общайся")
	fmt.Println("3 - Работа & Бизнес")
	fmt.Println("4 - Мультимедиа & Игровые")
	for {
		fmt.Print("Enter section: ")
		fmt.Scanln(&selectInput)
		switch selectInput {
		case "1":
			fmt.Println("Selected sector: Дом & Семья")
			valid = true
			return 1
		case "2":
			fmt.Println("Selected sector: Подключайся & Общайся")
			valid = true
			return 2
		case "3":
			fmt.Println("Selected sector: Работа & Бизнес")
			valid = true
			return 3
		case "4":
			fmt.Println("Selected sector: Мультимедиа & Игровые")
			valid = true
			return 4
		default:
			fmt.Println("Invalid input.")
			valid = false
		}
		if valid {
			break
		}
	}
	return 0
}

func main() {
	strItem := os.Args[1:] //Взять данные из аргумента, исключая имя файла
	var section int        //Сегмент
	var currentItem string //Код товара
	var name_item string   //Тип товара
	if len(strItem) > 1 {  //Если длина аргумента больше 1, то введенно некоретное число
		fmt.Println("Error - invalid input method. Must be articul of item (like 123456)")
		return
	}

	if len(strItem) == 0 { //Если в аргументах ничего нет, то просим ввести артикул вручную
		println("Not")
		inputItem(&currentItem)
		// strItem[0] = currentItem
		// fmt.Println(currentItem)
	}

	if len(strItem) == 1 { //Если есть аргумент, то считаем его
		for _, v := range strItem[0] {
			if v < 48 || v > 57 || v == 32 {
				fmt.Println("Invalid input. Dolbayeb!")
				return
			}
		}
		currentItem = strItem[0]
	}

	println("Is")
	// return
	var strName []string
	var strValue []string
	//Проверка на кол-во принимаемых аргументов

	/////////////////////////////////////////
	res, _ := http.Get("https://shop.kz/search/?q=" + currentItem) //Взять адрес из shop.kz
	doc, err := goquery.NewDocumentFromReader(res.Body)            //Взять его тело
	if err != nil {                                                //Если не получилось
		log.Fatal(err) //Ерорим
	}
	//Парсинг и сохранение в памяти
	doc.Find(".bx_detail_chars .bx_detail_chars_i").Each(func(i int, s *goquery.Selection) {
		q := s.Find(".glossary-term").Text()
		strName = append(strName, q)
		p := s.Find(".bx_detail_chars_i_field").Text()
		strValue = append(strValue, p)
		// fmt.Printf("%s %s\n", q, p)
	})
	for _, v := range doc.Find(".bx-title").Text() { //Узнать тип товара
		// fmt.Print(v, " ")
		if v != 32 { //Пока не дойдем до пробела
			name_item += string(v) //Добавляем в строку по букве
			continue
		}
		break
	}
	fmt.Println(name_item)
	var map_out map[string]string     //Объявляем карту из стрингов
	map_out = make(map[string]string) //и иницилизируем
	for i := range strName {          //Избавляемся от лишнего
		if strName[i] == "Ссылка на сайт производителя" {
			continue
		}
		if strName[i] == "Вес с упаковкой" {
			continue
		}
		if strName[i] == "Размеры упаковки (Ш х В х Г)" {
			continue
		}
		if strName[i] == "Размеры (Ш х В х Г)" {
			continue
		}
		if strName[i] == "Дополнительно" {
			continue
		}
		if strName[i] == "Внимание" {
			continue
		}
		if strName[i] == "Безопасность" {
			continue
		}
		if strName[i] == "Особенности клавиатуры" {
			continue
		}
		if strName[i] == "Материалы отделки" {
			continue
		}
		if strName[i] == "Средства коммуникации" {
			continue
		}
		if strName[i] == "Разъемы" {
			continue
		}
		if strName[i] == "Адаптер питания" {
			continue
		}
		if strName[i] == "Емкость аккумулятора" {
			continue
		}
		if strName[i] == "Тип аккумулятора" {
			continue
		}
		if strName[i] == "Мультимедиа" {
			continue
		}
		if strName[i] == "Разрешение экрана" {
			continue
		}
		if strName[i] == "Поверхность экрана" {
			continue
		}
		if strName[i] == "Примечание" {
			continue
		}
		if strName[i] == "Количество слотов оперативной памяти" {
			continue
		}
		if strName[i] == "Конфигурация оперативной памяти" {
			continue
		}
		if strName[i] == "Частота оперативной памяти" {
			continue
		}
		if strName[i] == "Тип оперативной памяти" {
			continue
		}
		if strName[i] == "Кэш-память L3" {
			continue
		}
		if strName[i] == "Количество ядер" {
			continue
		}
		map_out[strName[i]] = strValue[i]           //Добавляем в карту пару "Ключ - Значение"
		fmt.Println(strName[i], " = ", strValue[i]) //Принтим

	}
	// fmt.Println(map_out["Вес"])
	// fmt.Println(map_out["Оптический привод"])
	// fmt.Println(doc.Find(".bx-title").Text())
	switch name_item { //Проходим по имени в название товара, и если оно отличается от ноутбука, то сразу выставляем сегмент
	case "Ультрабук":
		section = 2
		fmt.Println("Подкл")
	case "Моноблок":
		section = 3
		fmt.Println("Раб")
	case "Компьютер":
		section = 3
		fmt.Println("Раб")
	case "Результаты":
		fmt.Println("Current item is not in Shop.kz")
		return
	default:
		section = sectionSelect()
	}
	section = section
	////
	f, err := os.Create("1.ini") //Создаем файл
	if err != nil {              //Если не получилось, то Ерорим
		fmt.Println(err)
		f.Close()
		return
	}
	defer f.Close() //Отложенное закрытие файла
	var d []string  //Объявляем массив строк
	d = append(d, "")
	d = append(d, "[SETTING]")
	d = append(d, "XNBNamePos=995")
	d = append(d, "YNBNamePos=150")
	d = append(d, "XTTXPos=913")
	d = append(d, "YTTXPos=195")
	d = append(d, "XWarrantyPos=550")
	d = append(d, "YWarrantyPos=45")
	d = append(d, "XCostPos=1320")
	d = append(d, "YCostPos=670")
	d = append(d, "XCostOldPos=1320")
	d = append(d, "YCostOldPos=630")
	d = append(d, "XRatingPos=835")
	d = append(d, "YRatingPos=680")
	d = append(d, "NBNameColor=-16777187")
	d = append(d, "NBNameFontSize=22")
	d = append(d, "NBNameFontName=Arial Black")
	d = append(d, "NBNameFontStyleBold=0")
	d = append(d, "NBNameFontStyleItalic=0")
	d = append(d, "NBNameFontStyleUnderline=0")
	d = append(d, "TTXNameColor=131586")
	d = append(d, "TTXNameFontSize=19")
	d = append(d, "TTXNameFontName=Arial")
	d = append(d, "TTXNameFontStyleBold=1")
	d = append(d, "TTXNameFontStyleItalic=0")
	d = append(d, "TTXNameFontStyleUnderline=0")
	d = append(d, "TTXNameAlign=right")
	d = append(d, "TTXValueColor=394758")
	d = append(d, "TTXValueFontSize=19")
	d = append(d, "TTXValueFontName=Arial")
	d = append(d, "TTXValueFontStyleBold=0")
	d = append(d, "TTXValueFontStyleItalic=0")
	d = append(d, "TTXValueFontStyleUnderline=0")
	d = append(d, "TTXValueAlign=left")
	d = append(d, "WarrantyColor=10038784")
	d = append(d, "WarrantyFontSize=28")
	d = append(d, "WarrantyFontName=Arial")
	d = append(d, "WarrantyFontStyleBold=0")
	d = append(d, "WarrantyFontStyleItalic=0")
	d = append(d, "WarrantyFontStyleUnderline=0")
	d = append(d, "CostColor=193")
	d = append(d, "CostFontSize=50")
	d = append(d, "CostFontName=Arial")
	d = append(d, "CostFontStyleBold=1")
	d = append(d, "CostFontStyleItalic=0")
	d = append(d, "CostFontStyleUnderline=0")
	d = append(d, "CostColorOld=4")
	d = append(d, "CostFontSizeOld=26")
	d = append(d, "CostFontNameOld=Arial")
	d = append(d, "CostFontStyleBoldOld=0")
	d = append(d, "CostFontStyleItalicOld=0")
	d = append(d, "CostFontStyleStrikeOld=1")
	d = append(d, "Poradok=UID,nur_sultan,karaganda,almaty,kostanay,pavlodar,temirtau,jezkazgan,fltr_BRAND,MODEL,fltr_NOTE_PROCESSOR,a_NOTE_TYPE_PROC,a_NOTE_PROC_FREC_SLIDER,TURBO_BOOSTFR,CORES_NUMBER,NOTE_L3_CASH,NOTE_TECHNOLOGY,a_NOTE_MEMSIZE,a_HARD_DRIVE,a_SSD,a_INT_VIDEO,a_NOTE_VIDEO,a_BRAND_VIDEO,NOTE_VIDEO_SIZE,a_NOTE_SCRSIZE_SLIDER,a_NOTE_SCR_RESOL,Display_COVER_type,NOTE_OPTIC_DISC,NOTE_USB_NEW,NOTE_USB3_1,a_NOTE_OS,attention,In_addition,a_WEIGHT,WARRANTY,creditovanie,store_bv,store_post_almaty,store_post_astana,store_post_karaganda,store_post_kostanay,store_post_pavlodar")
	d = append(d, "[MAIN]")
	d = append(d, "NBName="+doc.Find(".bx-title").Text())
	d = append(d, "UID="+map_out["UID товара"])
	d = append(d, "fltr_BRAND="+map_out["Производитель"])
	d = append(d, "MODEL="+map_out["Модель"])
	d = append(d, "fltr_NOTE_PROCESSOR="+map_out["Процессор"])
	d = append(d, "a_NOTE_TYPE_PROC="+map_out["Модель процессора"])
	d = append(d, "a_NOTE_PROC_FREC_SLIDER="+map_out["Частота процессора, ГГц"])
	d = append(d, "TURBO_BOOSTFR="+map_out["Максимальная частота процессора, ГГц"])
	d = append(d, "a_NOTE_MEMSIZE="+map_out["Объем оперативной памяти"])
	d = append(d, "a_HARD_DRIVE="+map_out["Жесткий диск"])
	d = append(d, "a_SSD="+map_out["Твердотельный накопитель"])
	// d = append(d, "UID="+map_out["Слот M.2 для SSD"])
	// d = append(d, "UID="+map_out["Тип видеокарты"])//?
	d = append(d, "a_INT_VIDEO="+map_out["Интегрированная в процессор графика"])
	d = append(d, "a_NOTE_SCRSIZE_SLIDER="+map_out["Диагональ экрана, дюйм"])
	d = append(d, "NOTE_OPTIC_DISC="+map_out["Оптический привод"])
	d = append(d, "NOTE_USB_NEW="+map_out["Количество разъемов USB 2.0"])
	d = append(d, "NOTE_USB3_1="+map_out["Количество разъемов USB 3.0"])
	// d = append(d, "UID="+map_out["Цвет, используемый в оформлении"])
	d = append(d, "a_NOTE_OS="+map_out["Операционная система"])
	d = append(d, "a_WEIGHT="+map_out["Вес"])
	d = append(d, "WARRANTY="+map_out["Срок гарантии (мес.)"])

	for _, v := range d { //Построчно
		fmt.Fprintln(f, v) //Добавляем значения в файл
	}
	////
	// if map_out["Вес"] == "" {
	// 	fmt.Println("Empty")
	// }
	fmt.Println("Работа готова, хозяин!")
	fmt.Print("Press the Enter Key to terminate the console screen!")
	var pauseInput string
	fmt.Scanln(&pauseInput)
}
