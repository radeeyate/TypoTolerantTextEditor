package main

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
	"absence":      {"absense"},
	"acceptable":   {"acceptible", "exceptable"},
	"achieve":      {"acheive"},
	"across":       {"accross", "acrost"},
	"address":      {"adress", "addres"},
	"affect":       {"effect"},
	"aggressive":   {"agressive", "agresive"},
	"agree":        {"aggree", "agre"},
	"amateur":      {"amatuer", "amature"},
	"among":        {"amoung", "amung"},
	"apparent":     {"aparent", "apparant"},
	"appearance":   {"appearence", "appearrance"},
	"argument":     {"arguement"}, // i aint arguing about this one
	"arctic":       {"artic"},     // brr. it's cold here
	"assist":       {"asist", "assit"},
	"asthma":       {"asma", "asthama"},
	"athlete":      {"athelete", "athleat"},
	"attendance":   {"attendence", "attendense"},
	"audience":     {"audince", "audiance"},

	// B for brain tired
	"balance":   {"ballance", "balence"},
	"basically": {"basicly", "bassically"},
	"because":   {"becuase", "beacuse"},
	"before":    {"befor", "bfore"},
	"beginning": {"begining", "beggining"},
	"believe":   {"beleive", "belive"}, // believe it or not, im still going
	"benefit":   {"benifit", "beniffit"},
	"between":   {"betwen", "beetween"},
	"bicycle":   {"bycicle", "bycycle"},
	"biscuit":   {"biscut", "biskit"},
	"business":  {"buisness", "bussiness"},

	// C for can't spell. send help
	"calendar":    {"calender", "calander"},
	"campaign":    {"campain"},
	"can't":       {"cant"},
	"caribbean":   {"carribean"}, // i wish i were here instead
	"category":    {"catagory", "catagorie"},
	"cemetery":    {"cemetary", "cematery"}, // spooky spelling
	"challenge":   {"chalange"},
	"change":      {"chang"},
	"choose":      {"chose"}, // i chose to do this. i regret it now.
	"clothes":     {"cloths", "close"},
	"coming":      {"comming", "comeing"},
	"commit":      {"comitt", "comit"},
	"committee":   {"comittee", "commitee"},
	"completely":  {"completly", "completley"},
	"conscious":   {"concious", "conscous"},
	"convenience": {"conveniance", "convenince"},
	"coolly":      {"cooly"},
	"could've":    {"could of"}, // could've used a break before starting this
	"country":     {"counrty", "countrey"},
	"course":      {"corse"},

	// D for despair, this will never end
	"definitely": {"definately", "definitly"},
	"dependent":  {"dependant", "depentant"},
	"describe":   {"discribe", "decribe"},
	"desperate":  {"desparate", "desperat"},
	"develop":    {"develope", "devlope"},
	"difference": {"differance", "differrence"},
	"difficult":  {"dificult", "difficault"},
	"disappoint": {"disapoint", "dissapoint"},
	"discipline": {"disipline", "disiplin"},
	"during":     {"durring"},

	// E easy because this section isn't too long - thankgoodness!
	"easy":        {"eazy"},
	"effect":      {"affect"},
	"eighth":      {"eigth"},
	"embarrass":   {"embarass", "embarras"},
	"environment": {"enviroment", "envirnment"},
	"equipment":   {"equipmnet", "equiptment"},
	"especially":  {"expecially"},
	"exaggerate":  {"exagerate", "exagerrate"}, // not exaggerating, this is hard
	"excellent":   {"excelent", "excellant"},
	"except":      {"expect"},
	"exercise":    {"exersize", "excercise"},
	"existence":   {"existance", "existense"},
	"experience":  {"experiance", "experince"},
	"experiment":  {"expirement", "experament"},

	// F for finally halfway???
	"familiar":    {"familar", "familliar"},
	"fascinating": {"fascinateing", "facinating"},
	"february":    {"febuary", "feburary"},
	"finally":     {"finaly", "finnaly"},
	"foreign":     {"foriegn", "forreign"},
	"forgetting":  {"forgeting"}, // i am forgoring how to spell words
	"forward":     {"forword"},
	"friend":      {"freind", "frend"}, // if only had a friend to help me w/ this
	"fundamental": {"fundemental", "fundamantal"},
	"further":     {"farther"}, // further down the rabbit hole we go

	// G for gosh this is a long list
	"generally":  {"generaly", "generelly"},
	"government": {"goverment", "governmnet"},
	"grammar":    {"grammer"},  // my grammer is suffering
	"guarantee":  {"guaranty"}, // i guarantee there are more mistakes
	"guidance":   {"guidence", "guidense"},

	// H for help me!!!
	"happened":  {"hapened", "happend"},
	"harass":    {"harrass", "haras"}, // dont harass me about the mistakes...
	"have":      {"hav"},
	"having":    {"haveing"},
	"height":    {"hight", "heighth"},
	"heroes":    {"heros"},
	"hopefully": {"hopfully", "hopefuly"},  // hopefully this is iver soon
	"humorous":  {"humourous", "humerous"}, // this is not as humorous as i thought it would be

	// I for i cant believe im still doing this
	"immediately":   {"imediately", "immediatly"},
	"important":     {"importent"}, // it is generally important to spell correctly
	"incidentally":  {"incidently", "incidentallly"},
	"independent":   {"independant"},
	"indispensable": {"indispensible"},
	"influence":     {"influance", "influense"}, // i hope i don't influence anyone to do this
	"intelligence":  {"inteligence"},
	"interesting":   {"intresting", "intersting"}, // it has been interesting making this list
	"interrupt":     {"interupt", "interript"},
	"island":        {"iland", "isalnd"},
	"its":           {"it's"}, // 🤓🤓🤓🤓

	// J for just keep swimming.. swimming.. swimming...
	"jealous":    {"jelous", "jelouse"}, // am jealous of people with perfect spelling
	"judgment":   {"judgement"},         // reserve it...
	"knowledge":  {"knowlege", "knoweledge"},
	"laboratory": {"labratory", "laboritory"},

	// L for length of this list
	"leisure":   {"liesure", "leasure"}, // whaaaat
	"length":    {"lenght", "lenth"},
	"library":   {"libary", "librery"},
	"license":   {"licence"},
	"lightning": {"lightening"},
	"lose":      {"loose"},
	"lying":     {"lieing"},

	// M for my fingers hurt...
	"maintenance": {"maintainance", "maintenence"}, // maintaining sanity
	"maneuver":    {"maneouvre", "manuever"},
	"marriage":    {"marrage", "marrige"}, // marriage is a commitment, just like this project
	"mathematics": {"mathmatics", "mathematic"},
	"medicine":    {"medicin", "medcine"}, // i need some for this headache
	"millennium":  {"millenium", "millinium"},
	"miniature":   {"minature", "miniuture"}, // this list is not miniature
	"minute":      {"minuite", "minut"},
	"mischievous": {"mischevious", "mischievious"},
	"misspell":    {"mispell", "missppell"},
	"morning":     {"moring"},
	"muscle":      {"muscel", "muscule"},
	"necessary":   {"neccessary", "neccesary"},
	"neighbor":    {"neighbour"},

	// N for no more comments now
	"noticeable":  {"noticable", "noticeble"},
	"nowadays":    {"nowdays", "now a days"},
	"occasion":    {"occassion", "ocation"},
	"occurred":    {"occured"},
	"occurrence":  {"occurence", "occurance"},
	"often":       {"ofen", "offten"},
	"omission":    {"ommision"},
	"opportunity": {"oppurtunity", "oportunity"},
	"original":    {"origional", "originall"},
	"ought":       {"aught"},
	"overrated":   {"overated"},

	// P words
	"parallel":      {"paralell", "paralel"},
	"particularly":  {"particulary"},
	"pastime":       {"passtime"},
	"performance":   {"performence", "performanse"},
	"perhaps":       {"perhapps", "perhapse"},
	"personnel":     {"personel", "personnell"},
	"piece":         {"peace"},
	"playwright":    {"playwrite"},
	"possession":    {"possesion"},
	"possible":      {"possable", "possibile"},
	"potatoes":      {"potatos", "potatose"},
	"practice":      {"practise"},
	"precede":       {"preceed", "procede"},
	"preference":    {"preferance", "prefference"},
	"prejudice":     {"predjudice", "prejuduce"},
	"preparation":   {"preperation", "prepareation"},
	"principal":     {"principle"},
	"privilege":     {"privelege", "privelige"},
	"probably":      {"probly", "probaly"},
	"procedure":     {"proceedure", "proceddure"},
	"proceed":       {"precede"},
	"professor":     {"proffesor", "profesor"},
	"promise":       {"promiss", "promice"},
	"pronounce":     {"pronunce", "pronounciation"},
	"pronunciation": {"pronounciation"},
	"psychologist":  {"psycologist", "phycologist"},
	"publicly":      {"publically", "publicaly"},
	"purpose":       {"purpous"},
	"quiet":         {"quite"},

	// R words
	"really":     {"realy", "reallly"},
	"receive":    {"recieve"},
	"recommend":  {"recomend", "recommned"},
	"reference":  {"referance", "refference"},
	"relevant":   {"relevent", "relavent"},
	"religious":  {"religous", "religius"},
	"repetition": {"repitition", "repeatition"},
	"restaurant": {"resturant", "restaraunt"},
	"rhythm":     {"rythm", "rythem"},
	"ridiculous": {"rediculous", "ridiculus"},

	// S words
	"sacrifice":  {"sacrafice", "sacrifise"},
	"safety":     {"safty", "saftey"},
	"schedule":   {"schedual", "scheduele"},
	"science":    {"sciense", "sceince"},
	"secretary":  {"secratary", "secretery"},
	"sentence":   {"sentance", "sentense"},
	"separate":   {"seperate", "seperat"},
	"sergeant":   {"sergent"},
	"similar":    {"similiar", "similer"},
	"since":      {"sence"},
	"sincerely":  {"sincerly"},
	"skillful":   {"skillfull"},
	"speech":     {"speach"},
	"successful": {"succesful", "succesfull"},
	"supersede":  {"supercede", "superceed"},
	"surprise":   {"surprize", "suprize"},

	// T words
	"than":      {"then"},
	"their":     {"there", "they're"},
	"therefore": {"therfore", "therefor"},
	"third":     {"thrid", "therd"},
	"though":    {"tho"},
	"through":   {"thru"},
	"tired":     {"tird", "tierd"},
	"together":  {"togather", "togeather"},
	"tomorrow":  {"tommorrow", "tomorow"},
	"tongue":    {"tounge", "toung"},
	"toward":    {"towards"},
	"truly":     {"truely", "trully"},

	// V words
	"vacuum":    {"vacume"},
	"variety":   {"variaty"},
	"vegetable": {"vegtable", "vegeable"},
	"vehicle":   {"vehicale"},
	"villain":   {"villian"},
	"visible":   {"visable", "visibel"},

	// W words
	"wednesday": {"wendsday", "wedensday"},
	"weird":     {"wierd", "weard"},
	"welcome":   {"wellcom"},
	"which":     {"witch"},
	"women":     {"wemen"},
	"would've":  {"would of"},
	"write":     {"rite"},
	"writing":   {"writeing"},

	// 🤓🤓🤓🤓
	"your": {"you're"},
}
