package sting

/*
 * Credits @danielstjules
 * https://github.com/danielstjules/Stringy/blob/master/src/Stringy.php#L1601-L1756
 */
func Ascii() map[string][]string{

	return map[string][]string{
		"0": []string{"°", "₀", "۰"},
		"1": []string{"¹", "₁", "۱"},
		"2": []string{"²", "₂", "۲"},
		"3": []string{"³", "₃", "۳"},
		"4": []string{"⁴", "₄", "۴", "٤"},
		"5": []string{"⁵", "₅", "۵", "٥"},
		"6": []string{"⁶", "₆", "۶", "٦"},
		"7": []string{"⁷", "₇", "۷"},
		"8": []string{"⁸", "₈", "۸"},
		"9": []string{"⁹", "₉", "۹"},
		"a": []string{"à", "á", "ả", "ã", "ạ", "ă", "ắ", "ằ", "ẳ", "ẵ", "ặ", "â", "ấ", "ầ", "ẩ", "ẫ", "ậ",
		"ā", "ą", "å", "α", "ά", "ἀ", "ἁ", "ἂ", "ἃ", "ἄ", "ἅ", "ἆ", "ἇ", "ᾀ", "ᾁ", "ᾂ", "ᾃ",
		"ᾄ", "ᾅ", "ᾆ", "ᾇ", "ὰ", "ά", "ᾰ", "ᾱ", "ᾲ", "ᾳ", "ᾴ", "ᾶ", "ᾷ", "а", "أ", "အ", "ာ",
		"ါ", "ǻ", "ǎ", "ª", "ა", "अ", "ا"},
		"b": []string{"б", "β", "Ъ", "Ь", "ب", "ဗ", "ბ"},
		"c": []string{"ç", "ć", "č", "ĉ", "ċ"},
		"d": []string{"ď", "ð", "đ", "ƌ", "ȡ", "ɖ", "ɗ", "ᵭ", "ᶁ", "ᶑ", "д", "δ", "د", "ض", "ဍ", "ဒ", "დ"},
		"e": []string{"é", "è", "ẻ", "ẽ", "ẹ", "ê", "ế", "ề", "ể", "ễ", "ệ", "ë", "ē", "ę", "ě", "ĕ", "ė",
		"ε", "έ", "ἐ", "ἑ", "ἒ", "ἓ", "ἔ", "ἕ", "ὲ", "έ", "е", "ё", "э", "є", "ə", "ဧ", "ေ",
		"ဲ", "ე", "ए", "إ", "ئ"},
		"f": []string{"ф", "φ", "ف", "ƒ", "ფ"},
		"g": []string{"ĝ", "ğ", "ġ", "ģ", "г", "ґ", "γ", "ဂ", "გ", "گ"},
		"h": []string{"ĥ", "ħ", "η", "ή", "ح", "ه", "ဟ", "ှ", "ჰ"},
		"i": []string{"í", "ì", "ỉ", "ĩ", "ị", "î", "ï", "ī", "ĭ", "į", "ı", "ι", "ί", "ϊ", "ΐ", "ἰ", "ἱ",
		"ἲ", "ἳ", "ἴ", "ἵ", "ἶ", "ἷ", "ὶ", "ί", "ῐ", "ῑ", "ῒ", "ΐ", "ῖ", "ῗ", "і", "ї", "и",
		"ဣ", "ိ", "ီ", "ည်", "ǐ", "ი", "इ", "ی"},
		"j": []string{"ĵ", "ј", "Ј", "ჯ", "ج"},
		"k": []string{"ķ", "ĸ", "к", "κ", "Ķ", "ق", "ك", "က", "კ", "ქ", "ک"},
		"l": []string{"ł", "ľ", "ĺ", "ļ", "ŀ", "л", "λ", "ل", "လ", "ლ"},
		"m": []string{"м", "μ", "م", "မ", "მ"},
		"n": []string{"ñ", "ń", "ň", "ņ", "ŉ", "ŋ", "ν", "н", "ن", "န","ნ"},
		"o": []string{"ó", "ò", "ỏ", "õ", "ọ", "ô", "ố", "ồ", "ổ", "ỗ", "ộ", "ơ", "ớ", "ờ", "ở", "ỡ", "ợ",
		"ø", "ō", "ő", "ŏ", "ο", "ὀ", "ὁ", "ὂ", "ὃ", "ὄ", "ὅ", "ὸ", "ό", "о", "و", "θ", "ို",
		"ǒ", "ǿ", "º", "ო", "ओ"},
		"p": []string{"п", "π", "ပ", "პ", "پ"},
		"q": []string{"ყ"},
		"r": []string{"ŕ", "ř", "ŗ", "р", "ρ", "ر", "რ"},
		"s": []string{"ś", "š", "ş", "с", "σ", "ș", "ς", "س", "ص", "စ", "ſ", "ს"},
		"t": []string{"ť", "ţ", "т", "τ", "ț", "ت", "ط", "ဋ", "တ", "ŧ", "თ", "ტ"},
		"u": []string{"ú", "ù", "ủ", "ũ", "ụ", "ư", "ứ", "ừ", "ử", "ữ", "ự", "û", "ū", "ů", "ű", "ŭ", "ų",
		"µ", "у", "ဉ", "ု", "ူ", "ǔ", "ǖ", "ǘ", "ǚ", "ǜ", "უ", "उ"},
		"v": []string{"в", "ვ", "ϐ"},
		"w": []string{"ŵ", "ω", "ώ", "ဝ", "ွ"},
		"x": []string{"χ", "ξ"},
		"y": []string{"ý", "ỳ", "ỷ", "ỹ", "ỵ", "ÿ", "ŷ", "й", "ы", "υ", "ϋ", "ύ", "ΰ", "ي", "ယ"},
		"z": []string{"ź", "ž", "ż", "з", "ζ", "ز", "ဇ", "ზ"},
		"aa": []string{"ع", "आ", "آ"},
		"ae": []string{"ä", "æ", "ǽ"},
		"ai": []string{"ऐ"},
		"at": []string{"@"},
		"ch": []string{"ч", "ჩ", "ჭ", "چ"},
		"dj": []string{"ђ", "đ"},
		"dz": []string{"џ", "ძ"},
		"ei": []string{"ऍ"},
		"gh": []string{"غ", "ღ"},
		"ii": []string{"ई"},
		"ij": []string{"ĳ"},
		"kh": []string{"х", "خ", "ხ"},
		"lj": []string{"љ"},
		"nj": []string{"њ"},
		"oe": []string{"ö", "œ", "ؤ"},
		"oi": []string{"ऑ"},
		"oii": []string{"ऒ"},
		"ps": []string{"ψ"},
		"sh": []string{"ш", "შ", "ش"},
		"shch": []string{"щ"},
		"ss": []string{"ß"},
		"sx": []string{"ŝ"},
		"th": []string{"þ", "ϑ", "ث", "ذ", "ظ"},
		"ts": []string{"ц", "ც", "წ"},
		"ue": []string{"ü"},
		"uu": []string{"ऊ"},
		"ya": []string{"я"},
		"yu": []string{"ю"},
		"zh": []string{"ж", "ჟ", "ژ"},
		"(c)": []string{"©"},
		"A": []string{"Á", "À", "Ả", "Ã", "Ạ", "Ă", "Ắ", "Ằ", "Ẳ", "Ẵ", "Ặ", "Â", "Ấ", "Ầ", "Ẩ", "Ẫ", "Ậ", "Å",
		"Ā", "Ą", "Α", "Ά", "Ἀ", "Ἁ", "Ἂ", "Ἃ", "Ἄ", "Ἅ", "Ἆ", "Ἇ", "ᾈ", "ᾉ", "ᾊ", "ᾋ", "ᾌ", "ᾍ",
		"ᾎ", "ᾏ", "Ᾰ", "Ᾱ", "Ὰ", "Ά", "ᾼ", "А", "Ǻ", "Ǎ"},
		"B": []string{"Б", "Β", "ब"},
		"C": []string{"Ç","Ć", "Č", "Ĉ", "Ċ"},
		"D": []string{"Ď", "Ð", "Đ", "Ɖ", "Ɗ", "Ƌ", "ᴅ", "ᴆ", "Д", "Δ"},
		"E": []string{"É", "È", "Ẻ", "Ẽ", "Ẹ", "Ê", "Ế", "Ề", "Ể", "Ễ", "Ệ", "Ë", "Ē", "Ę", "Ě", "Ĕ", "Ė", "Ε",
		"Έ", "Ἐ", "Ἑ", "Ἒ", "Ἓ", "Ἔ", "Ἕ", "Έ", "Ὲ", "Е", "Ё", "Э", "Є", "Ə"},
		"F": []string{"Ф", "Φ"},
		"G": []string{"Ğ", "Ġ", "Ģ", "Г", "Ґ", "Γ"},
		"H": []string{"Η", "Ή", "Ħ"},
		"I": []string{"Í", "Ì", "Ỉ", "Ĩ", "Ị", "Î", "Ï", "Ī", "Ĭ", "Į", "İ", "Ι", "Ί", "Ϊ", "Ἰ", "Ἱ", "Ἳ", "Ἴ",
		"Ἵ", "Ἶ", "Ἷ", "Ῐ", "Ῑ", "Ὶ", "Ί", "И", "І", "Ї", "Ǐ", "ϒ"},
		"K": []string{"К", "Κ"},
		"L": []string{"Ĺ", "Ł", "Л", "Λ", "Ļ", "Ľ", "Ŀ", "ल"},
		"M": []string{"М", "Μ"},
		"N": []string{"Ń", "Ñ", "Ň", "Ņ", "Ŋ", "Н", "Ν"},
		"O": []string{"Ó", "Ò", "Ỏ", "Õ", "Ọ", "Ô", "Ố", "Ồ", "Ổ", "Ỗ", "Ộ", "Ơ", "Ớ", "Ờ", "Ở", "Ỡ", "Ợ", "Ø",
		"Ō", "Ő", "Ŏ", "Ο", "Ό", "Ὀ", "Ὁ", "Ὂ", "Ὃ", "Ὄ", "Ὅ", "Ὸ", "Ό", "О", "Θ", "Ө", "Ǒ", "Ǿ"},
		"P": []string{"П", "Π"},
		"R": []string{"Ř", "Ŕ", "Р", "Ρ", "Ŗ"},
		"S": []string{"Ş", "Ŝ", "Ș", "Š", "Ś", "С", "Σ"},
		"T": []string{"Ť", "Ţ", "Ŧ", "Ț", "Т", "Τ"},
		"U": []string{"Ú", "Ù", "Ủ", "Ũ", "Ụ", "Ư", "Ứ", "Ừ", "Ử", "Ữ", "Ự", "Û", "Ū", "Ů", "Ű", "Ŭ", "Ų", "У",
		"Ǔ", "Ǖ", "Ǘ", "Ǚ", "Ǜ"},
		"V": []string{"В"},
		"W": []string{"Ω", "Ώ", "Ŵ"},
		"X": []string{"Χ", "Ξ"},
		"Y": []string{"Ý", "Ỳ", "Ỷ", "Ỹ", "Ỵ", "Ÿ", "Ῠ", "Ῡ", "Ὺ", "Ύ", "Ы", "Й", "Υ", "Ϋ", "Ŷ"},
		"Z": []string{"Ź", "Ž", "Ż", "З", "Ζ"},
		"AE": []string{"Ä", "Æ", "Ǽ"},
		"CH": []string{"Ч"},
		"DJ": []string{"Ђ"},
		"DZ": []string{"Џ"},
		"GX": []string{"Ĝ"},
		"HX": []string{"Ĥ"},
		"IJ": []string{"Ĳ"},
		"JX": []string{"Ĵ"},
		"KH": []string{"Х"},
		"LJ": []string{"Љ"},
		"NJ": []string{"Њ"},
		"OE": []string{"Ö", "Œ"},
		"PS": []string{"Ψ"},
		"SH": []string{"Ш"},
		"SHCH": []string{"Щ"},
		"SS": []string{"ẞ"},
		"TH": []string{"Þ"},
		"TS": []string{"Ц"},
		"UE": []string{"Ü"},
		"YA": []string{"Я"},
		"YU": []string{"Ю"},
		"ZH": []string{"Ж"},
		" ": []string{"\xC2\xA0", "\xE2\x80\x80", "\xE2\x80\x81", "\xE2\x80\x82", "\xE2\x80\x83",
		"\xE2\x80\x84", "\xE2\x80\x85", "\xE2\x80\x86", "\xE2\x80\x87", "\xE2\x80\x88", 
		"\xE2\x80\x89", "\xE2\x80\x8A", "\xE2\x80\xAF", "\xE2\x81\x9F", "\xE3\x80\x80"},
	}
}
