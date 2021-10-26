package main

import "strings"

// See LAYOUTS section in `man 7 xkeyboard-config`
func get_layout_flag(l string) string {
	if strings.Contains(l, "Albanian") {
		return "🇦🇱"
	}
	if strings.Contains(l, "Armenian") {
		return "🇦🇲"
	}
	if strings.Contains(l, "Azerbaijani") {
		return "🇦🇿"
	}
	if strings.Contains(l, "Belarusian") {
		return "🇧🇾"
	}
	if strings.Contains(l, "Belgian") {
		return "🇧🇪"
	}
	if strings.Contains(l, "Indian") ||
		strings.Contains(l, "Manipuri") ||
		strings.Contains(l, "Gujarati") ||
		strings.Contains(l, "Punjabi") ||
		strings.Contains(l, "Kannada") ||
		strings.Contains(l, "Malayalam") ||
		strings.Contains(l, "Oriya") ||
		strings.Contains(l, "Ol Chiki") ||
		strings.Contains(l, "Telugu") ||
		strings.Contains(l, "Hindi") ||
		strings.Contains(l, "Sanskrit") ||
		strings.Contains(l, "Marathi") ||
		strings.Contains(l, "Indic IPA") {
		return "🇮🇳"
	}
	if strings.Contains(l, "Bosnian") {
		return "🇧🇦"
	}
	if strings.Contains(l, "Bulgarian") {
		return "🇧🇬"
	}
	if strings.Contains(l, "Austria") {
		return "🇦🇹"
	}
	if strings.Contains(l, "Australian") {
		return "🇦🇺"
	}
	if strings.Contains(l, "Brazil") {
		return "🇧🇷"
	}
	if strings.Contains(l, "English") {
		if strings.Contains(l, "UK") {
			return "🇬🇧"
		}
		if strings.Contains(l, "South Africa") {
			return "🇿🇦"
		}
		return "🇺🇸"
	}
	if strings.Contains(l, "French") ||
		strings.Contains(l, "France") ||
		strings.Contains(l, "Occitan") {
		return "🇫🇷"
	}
	if strings.Contains(l, "Russian") ||
		strings.Contains(l, "Tatar") ||
		strings.Contains(l, "Chuvash") ||
		strings.Contains(l, "Udmurt") ||
		strings.Contains(l, "Komi") ||
		strings.Contains(l, "Yakut") ||
		strings.Contains(l, "Kalmyk") ||
		strings.Contains(l, "Bashkirian") ||
		strings.Contains(l, "Mari") {
		return "🇷🇺"
	}
	if strings.Contains(l, "Arabic") {
		return "🇸🇦"
	}
	if strings.Contains(l, "Bangla") {
		return "🇧🇩"
	}
	if strings.Contains(l, "Algeria") ||
		strings.Contains(l, "Kabyle") {
		return "🇩🇿"
	}
	if strings.Contains(l, "Morocco") {
		return "🇲🇦"
	}
	if strings.Contains(l, "Cameroon") ||
		strings.Contains(l, "Mmuock") {
		return "🇨🇲"
	}
	if strings.Contains(l, "Burmese") {
		return "🇲🇲"
	}
	if strings.Contains(l, "Canada") ||
		strings.Contains(l, "Inuktitut") {
		return "🇨🇦"
	}
	if strings.Contains(l, "Congo") {
		return "🇨🇩"
	}
	if strings.Contains(l, "Chinese") ||
		strings.Contains(l, "Tibetan") ||
		strings.Contains(l, "Uyghur") ||
		strings.Contains(l, "Hanyu Pinyin") {
		return "🇨🇳"
	}
	if strings.Contains(l, "Croatian") {
		return "🇭🇷"
	}
	if strings.Contains(l, "Czech") {
		return "🇨🇿"
	}
	if strings.Contains(l, "Danish") {
		return "🇩🇰"
	}
	if strings.Contains(l, "Dutch") {
		return "🇳🇱"
	}
	if strings.Contains(l, "Dzongkha") {
		return "🇧🇹"
	}
	if strings.Contains(l, "Estonian") {
		return "🇪🇪"
	}
	if strings.Contains(l, "Persian") ||
		strings.Contains(l, "Iran") {
		return "🇮🇷"
	}
	if strings.Contains(l, "Iraqi") ||
		strings.Contains(l, "Kurdish") {
		return "🇮🇶"
	}
	if strings.Contains(l, "Faroese") {
		return "🇫🇴"
	}
	if strings.Contains(l, "Fin") {
		return "🇫🇮"
	}
	if l == "Akan" ||
		l == "Ewe" ||
		l == "Fula" ||
		l == "Ga" ||
		l == "Avatime" {
		return "🇬🇭"
	}
	if strings.Contains(l, "Georgian") ||
		strings.Contains(l, "Ossetian") {
		return "🇬🇪"
	}
	if strings.Contains(l, "Romanian") {
		return "🇷🇴"
	}
	if strings.Contains(l, "Turkish") ||
		strings.Contains(l, "Ottoman") {
		return "🇹🇷"
	}
	if strings.Contains(l, "German") ||
		strings.Contains(l, "Lower Sorbian") {
		return "🇩🇪"
	}
	if strings.Contains(l, "Greek") {
		return "🇬🇷"
	}
	if strings.Contains(l, "Hungarian") {
		return "🇭🇺"
	}
	if strings.Contains(l, "Icelandic") {
		return "🇮🇸"
	}
	if strings.Contains(l, "Hebrew") {
		return "🇮🇱"
	}
	if strings.Contains(l, "Italian") ||
		strings.Contains(l, "Sicilian") ||
		strings.Contains(l, "Friulian") {
		return "🇮🇹"
	}
	if strings.Contains(l, "Japanese") {
		return "🇯🇵"
	}
	if strings.Contains(l, "Kyrgyz") {
		return "🇰🇬"
	}
	if strings.Contains(l, "Khmer") {
		return "🇰🇭"
	}
	if strings.Contains(l, "Kazakh") {
		return "🇰🇿"
	}
	if strings.Contains(l, "Lao") {
		return "🇱🇦"
	}
	if strings.Contains(l, "Luthuanian") ||
		strings.Contains(l, "Samogitian") {
		return "🇱🇹"
	}
	if strings.Contains(l, "Latvian") {
		return "🇱🇻"
	}
	if strings.Contains(l, "Montenegrin") {
		return "🇲🇪"
	}
	if strings.Contains(l, "Macedonian") {
		return "🇲🇰"
	}
	if strings.Contains(l, "Maltese") {
		return "🇲🇹"
	}
	if strings.Contains(l, "Mongolian") {
		return "🇲🇳"
	}
	if strings.Contains(l, "Norw") {
		return "🇳🇴"
	}
	if strings.Contains(l, "Polish") ||
		strings.Contains(l, "Kashubian") ||
		strings.Contains(l, "Silesian") {
		return "🇵🇱"
	}
	if strings.Contains(l, "Portug") {
		return "🇵🇹"
	}
	if strings.Contains(l, "Serbian") ||
		strings.Contains(l, "Pannonian Rusyn") {
		return "🇷🇸"
	}
	if strings.Contains(l, "Slovenian") {
		return "🇸🇮"
	}
	if strings.Contains(l, "Slovak") {
		return "🇸🇰"
	}
	if strings.Contains(l, "Spanish") ||
		strings.Contains(l, "Asturian") {
		return "🇪🇸"
	}
	if strings.Contains(l, "Catalan") {
		return "🇦🇩"
	}
	if strings.Contains(l, "Swedish") ||
		strings.Contains(l, "Northern Saami") {
		return "🇸🇪"
	}
	if strings.Contains(l, "Switzerland") {
		return "🇨🇭"
	}
	if strings.Contains(l, "Syriac") {
		return "🇸🇾"
	}
	if strings.Contains(l, "Tajik") {
		return "🇹🇯"
	}
	if strings.Contains(l, "Sinhala") ||
		strings.Contains(l, "Tamil") {
		return "🇱🇰"
	}
	if strings.Contains(l, "Thai") {
		return "🇹🇭"
	}
	if strings.Contains(l, "Turkmen") {
		return "🇹🇲"
	}
	if strings.Contains(l, "Taiwanese") ||
		strings.Contains(l, "Saisiyat") {
		return "🇹🇼"
	}
	if strings.Contains(l, "Ukrainian") {
		return "🇺🇦"
	}
	if strings.Contains(l, "Uzbek") {
		return "🇺🇿"
	}
	if strings.Contains(l, "Vietnamese") {
		return "🇻🇳"
	}
	if strings.Contains(l, "Korean") {
		return "🇰🇷"
	}
	if strings.Contains(l, "Irish") ||
		strings.Contains(l, "CloGaelach") ||
		strings.Contains(l, "Ogham") {
		return "🇮🇪"
	}
	if strings.Contains(l, "Urdu") {
		return "🇵🇰"
	}
	if strings.Contains(l, "Dhivehi") {
		return "🇲🇻"
	}
	if strings.Contains(l, "Nepali") {
		return "🇳🇵"
	}
	if strings.Contains(l, "Igbo") ||
		strings.Contains(l, "Yoruba") ||
		strings.Contains(l, "Hausa") {
		return "🇳🇬"
	}
	if strings.Contains(l, "Amharic") {
		return "🇪🇹"
	}
	if strings.Contains(l, "Wolof") {
		return "🇸🇳"
	}
	if strings.Contains(l, "Bambara") ||
		strings.Contains(l, "Mali") {
		return "🇲🇱"
	}
	if strings.Contains(l, "Tanzania") {
		return "🇹🇿"
	}
	if strings.Contains(l, "Togo") {
		return "🇹🇬"
	}
	if strings.Contains(l, "Kenya") {
		return "🇰🇪"
	}
	if strings.Contains(l, "Tswana") {
		return "🇧🇼"
	}
	if strings.Contains(l, "Filipino") {
		return "🇵🇭"
	}
	if strings.Contains(l, "Moldavian") {
		return "🇲🇩"
	}
	if strings.Contains(l, "Indonesian") {
		return "🇮🇩"
	}
	if strings.Contains(l, "Malay") {
		return "🇲🇾"
	}
	if strings.Contains(l, "Afghani") ||
		strings.Contains(l, "Pashto") {
		return "🇦🇫"
	}
	return ""
}
