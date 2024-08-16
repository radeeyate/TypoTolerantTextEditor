package main

var nonTextChangeKeys = map[int]bool {
	9: true, // esc
	37: true, // control
	50: true, // left shift
	64: true, // right shift
}

var keyboardMap = map[string][]string{ // this was a pain to make
	"a":  {"s", "z", "q"},
	"b":  {"v", "g", "h", "n"},
	"c":  {"v", "x", "d", "f"},
	"d":  {"f", "e", "s"},
	"e":  {"w", "s", "d", "r"},
	"f":  {"d", "r", "g"},
	"g":  {"f", "t", "h"},
	"h":  {"g", "y", "j"},
	"i":  {"u", "o", "k"},
	"j":  {"h", "u", "k"},
	"k":  {"j", "i", "l", "m"},
	"l":  {"k", "o", "p"},
	"m":  {"n", "j", "k", ","},
	"n":  {"b", "h", "j", "m"},
	"o":  {"i", "p", "l", "0"},
	"p":  {"o", "l", "["},
	"q":  {"w", "a", "1"},
	"r":  {"e", "t", "f"},
	"s":  {"a", "w", "e", "d", "z", "x"},
	"t":  {"r", "y", "g"},
	"u":  {"y", "i", "j"},
	"v":  {"c", "f", "g", "b"},
	"w":  {"q", "a", "s", "e"},
	"x":  {"z", "s", "d", "c"},
	"y":  {"t", "u", "h"},
	"z":  {"a", "s", "x"},
	"1":  {"2", "q", "`", "!"},
	"2":  {"1", "3", "w", "q"},
	"3":  {"2", "4", "e", "w"},
	"4":  {"3", "5", "r", "e"},
	"5":  {"4", "6", "t", "r"},
	"6":  {"5", "7", "y", "t"},
	"7":  {"6", "8", "u", "y"},
	"8":  {"7", "9", "i", "u"},
	"9":  {"8", "0", "o", "i"},
	"0":  {"9", "-", "p", "o"},
	"-":  {"0", "=", "p", "["},
	"=":  {"-", "[", "]", "\\"},
	"`":  {"1", "t", "\\"},
	"[":  {"=", "]", "\\", ";", "'"},
	"]":  {"[", "\\", "p", ";", "'"},
	"\\": {"]", "["},
	";":  {"'", "l", "k", "."},
	"'":  {";", "/", "l", "k", "[", "]", "\\"},
	",":  {"m", ".", "/"},
	".":  {",", "/", ";"},
	"/":  {".", ",", "'", "'"},
	" ":  {" "}, // this will never be used because it cannot be replaced :)
}

var wordReplacements = map[string][]string{
	// A for AUGH
	"absence":    {"absense"},
	"acceptable": {"acceptible", "exceptable"},
	"achieve":    {"acheive"},
	"across":     {"acrost"},
	"address":    {"addres"},
	"affect":     {"effect"},
	"amateur":    {"amatuer", "amature"},
	"among":      {"amung"},
	"apparent":   {"apparant"},
	"appearance": {"appearence"},
	"assist":     {"assit"},
	"asthma":     {"asthama"},
	"athlete":    {"athleat"},
	"attendance": {"attendence", "attendense"},
	"audience":   {"audiance"},

	// B for brain tired
	"balance":   {"balence"},
	"basically": {"bassically"},
	"because":   {"becuase", "beacuse"},
	"before":    {"bfore"},
	"beginning": {"beggining"},
	"believe":   {"beleive"}, // believe it or not, im still going
	"benefit":   {"benifit"},
	"between":   {"beetween"},
	"bicycle":   {"bycicle", "bycycle"},
	"biscuit":   {"biskit"},
	"business":  {"buisness"},

	// C for can't spell. send help
	"calendar":    {"calender", "calander"},
	"caribbean":   {"carribean"}, // i wish i were here instead
	"category":    {"catagory"},
	"cemetery":    {"cemetary", "cematery"}, // spooky spelling
	"clothes":     {"close"},
	"coming":      {"comeing"},
	"commit":      {"comitt"},
	"committee":   {"commitee"},
	"completely":  {"completley"},
	"conscious":   {"conscous"},
	"convenience": {"conveniance"},
	"could've":    {"could of"}, // could've used a break before starting this
	"country":     {"counrty"},

	// D for despair, this will never end
	"definitely": {"definately"},
	"dependent":  {"dependant", "depentant"},
	"describe":   {"discribe"},
	"desperate":  {"desparate"},
	"develop":    {"devlope"},
	"difference": {"differance"},
	"difficult":  {"difficault"},
	"disappoint": {"dissapoint"},
	"discipline": {"disiplin"},

	// E easy because this section isn't too long - thank goodness!
	"easy":        {"eazy"},
	"effect":      {"affect"},
	"embarrass":   {"embarras"},
	"environment": {"envirnment"},
	"equipment":   {"equipmnet"},
	"especially":  {"expecially"},
	"exaggerate":  {"exagerrate"}, // not exaggerating, this is hard
	"excellent":   {"excellant"},
	"except":      {"expect"},
	"exercise":    {"exersize"},
	"existence":   {"existance", "existense"},
	"experience":  {"experiance"},
	"experiment":  {"expirement", "experament"},

	// F for finally halfway???
	"familiar":    {"familliar"},
	"fascinating": {"facinating"},
	"february":    {"feburary"},
	"finally":     {"finnaly"},
	"foreign":     {"foriegn"},
	"forward":     {"forword"},
	"friend":      {"freind"}, // if only had a friend to help me w/ this
	"fundamental": {"fundemental", "fundamantal"},
	"further":     {"farther"}, // further down the rabbit hole we go

	// G for gosh this is a long list
	"generally":  {"generelly"},
	"government": {"governmnet"},
	"grammar":    {"grammer"}, // my grammer is suffering
	"guidance":   {"guidence", "guidense"},

	// H for help me!!!
	"happened":  {"happend"},
	"harass":    {"haras"}, // dont harass me about the mistakes...
	"height":    {"heighth"},
	"hopefully": {"hopefuly"}, // hopefully this is iver soon
	"humorous":  {"humerous"}, // this is not as humorous as i thought it would be

	// I for i cant believe im still doing this
	"immediately":   {"immediatly"},
	"important":     {"importent"}, // it is generally important to spell correctly
	"incidentally":  {"incidentallly"},
	"independent":   {"independant"},
	"indispensable": {"indispensible"},
	"influence":     {"influance", "influense"}, // i hope i don't influence anyone to do this
	"interesting":   {"intersting"}, // it has been interesting making this list
	"interrupt":     {"interript"},
	"island":        {"isalnd"},

	// J for just keep swimming.. swimming.. swimming...
	"jealous":    {"jelouse"}, // am jealous of people with perfect spelling
	"knowledge":  {"knoweledge"},
	"laboratory": {"laboritory"},

	// L for length of this list
	"leisure":   {"liesure", "leasure"}, // whaaaat
	"length":    {"lenght"},
	"library":   {"librery"},
	"license":   {"licence"},

	// M for my fingers hurt...
	"maintenance": {"maintenence"}, // maintaining sanity
	"maneuver":    {"manuever"},
	"marriage":    {"marrige"}, // marriage is a commitment, just like this project
	"mathematics": {"mathematic"},
	"medicine":    {"medcine"}, // i need some for this headache
	"millennium":  {"millinium"},
	"miniature":   {"miniuture"}, // this list is not miniature
	"minute":      {"minut"},
	"mischievous": {"mischevious"},
	"misspell":    {"missppell"},
	"muscle":      {"muscel"},
	"necessary":   {"neccesary"},

	// N for no more comments now
	"noticeable":  {"noticeble"},
	"nowadays":    {"now a days"},
	"occasion":    {"ocation"},
	"occurrence":  {"occurance"},
	"often":       {"offten"},
	"omission":    {"ommision"},
	"opportunity": {"oppurtunity"},
	"original":    {"originall"},
	"ought":       {"aught"},

	// P words
	"parallel":     {"paralell"},
	"performance":  {"performence", "performanse"},
	"perhaps":      {"perhapse"},
	"personnel":    {"personnell"},
	"piece":        {"peace"},
	"possible":     {"possable"},
	"potatoes":     {"potatose"},
	"practice":     {"practise"},
	"precede":      {"preceed", "procede"},
	"preference":   {"preferance"},
	"prejudice":    {"prejuduce"},
	"preparation":  {"preperation"},
	"principal":    {"principle"},
	"privilege":    {"privelege", "privelige"},
	"probably":     {"probaly"},
	"procedure":    {"proceddure"},
	"proceed":      {"precede"},
	"professor":    {"proffesor"},
	"promise":      {"promiss", "promice"},
	"psychologist": {"phycologist"},
	"publicly":     {"publicaly"},
	"purpose":      {"purpous"},
	"quiet":        {"quite"},

	// R words
	"really":     {"reallly"},
	"receive":    {"recieve"},
	"recommend":  {"recommned"},
	"reference":  {"referance"},
	"relevant":   {"relevent", "relavent"},
	"religious":  {"religius"},
	"repetition": {"repitition"},
	"restaurant": {"restaraunt"},
	"rhythm":     {"rythem"},
	"ridiculous": {"rediculous"},

	// S words
	"sacrifice":  {"sacrafice", "sacrifise"},
	"safety":     {"saftey"},
	"schedule":   {"schedual"},
	"science":    {"sciense", "sceince"},
	"secretary":  {"secratary", "secretery"},
	"sentence":   {"sentance", "sentense"},
	"separate":   {"seperate"},
	"similar":    {"similer"},
	"since":      {"sence"},
	"speech":     {"speach"},
	"successful": {"succesfull"},
	"supersede":  {"supercede", "superceed"},
	"surprise":   {"surprize"},

	// T words
	"than":      {"then"},
	"their":     {"there"},
	"therefore": {"therefor"},
	"third":     {"thrid", "therd"},
	"tired":     {"tierd"},
	"together":  {"togather"},
	"tomorrow":  {"tomorow"},
	"tongue":    {"tounge"},
	"truly":     {"trully"},

	// V words
	"vacuum":    {"vacume"},
	"variety":   {"variaty"},
	"vegetable": {"vegeable"},
	"villain":   {"villian"},
	"visible":   {"visable", "visibel"},

	// W words
	"wednesday": {"wedensday"},
	"weird":     {"wierd", "weard"},
	"welcome":   {"wellcom"},
	"which":     {"witch"},
	"women":     {"wemen"},
	"would've":  {"would of"},
}
