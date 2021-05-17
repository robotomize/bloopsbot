package resource

import "github.com/enescakir/emoji"

type Bloops struct {
	Name    string
	Image   string
	Points  int
	Task    string
	Seconds int
	Weight  int
}

type Category struct {
	Text   string
	Status bool
}

type Letter struct {
	Text   string
	Status bool
}

var (
	Letters = []Letter{
		{Text: "А", Status: true},
		{Text: "Б", Status: true},
		{Text: "В", Status: true},
		{Text: "Г", Status: true},
		{Text: "Д", Status: true},
		{Text: "Е", Status: true},
		{Text: "Ж", Status: true},
		{Text: "З", Status: true},
		{Text: "И", Status: true},
		{Text: "К", Status: true},
		{Text: "Л", Status: true},
		{Text: "М", Status: true},
		{Text: "Н", Status: true},
		{Text: "О", Status: true},
		{Text: "П", Status: true},
		{Text: "Р", Status: true},
		{Text: "С", Status: true},
		{Text: "Т", Status: true},
		{Text: "У", Status: true},
		{Text: "Ф", Status: true},
		{Text: "Х", Status: true},
		{Text: "Ц"},
		{Text: "Ч"},
		{Text: "Ш"},
		{Text: "Э"},
		{Text: "Ю"},
		{Text: "Я"},
	}

	Categories = []Category{
		{Text: "Страна"},
		{Text: "Город", Status: true},
		{Text: "Овощ или фрукт", Status: true},
		{Text: "Имя", Status: true},
		{Text: "Знаменитость"},
		{Text: "Бренд", Status: true},
		{Text: "Животное", Status: true},
		{Text: "Термин"},
		{Text: "Любое слово"},
	}

	RoundsNum  = []int{1, 2, 3, 4, 5}
	RoundTimes = []int{30, 45, 60}

	Bloopses = []Bloops{
		{Name: emoji.Cinema.String() + " Артхаус режиссер", Weight: 2, Seconds: +30, Points: +10, Task: "К тебе ворвался режиссер артхаус кино и предложил помочь со своим проектом, тебе нужно заменить категории в игре на категорию *кино и актеры*\nНазывай имена фильмов, актеров или режиссеров на выпавшую букву"},
		{Name: emoji.Flamingo.String() + "Фламинго", Weight: 2, Points: +10, Task: "Так получилось, что ты стал фламинго на время, когда называешь слова, ты должен стоять на одной ноге(можно держаться за что-нибудь)"},
		{Name: emoji.Flamingo.String() + "Фламинго", Weight: 2, Points: +10, Task: "Так получилось, что ты стал фламинго на время, когда называешь слова, ты должен стоять на одной ноге(можно держаться за что-нибудь)"},
		{Name: emoji.WomanSinger.String() + " На разогреве", Weight: 2, Points: +5, Task: "Ты начинающая рок звезда и тебя попросили выступить на разогреве, каждое слово, которое ты называешь ты должен пропеть в своем любим стиле"},
		{Name: emoji.WomanSinger.String() + " На разогреве", Weight: 2, Points: +5, Task: "Ты начинающая рок звезда и тебя попросили выступить на разогреве, каждое слово, которое ты называешь ты должен пропеть в своем любим стиле"},
		{Name: emoji.Hammer.String() + " Мастерство", Weight: 1, Points: +20, Seconds: +20, Task: "Из поколения в поколение ты передавалась твоё ремесло, время показать себя на ярмарке! Тебе нужно на каждую категорию назвать не одно слово, а два"},
		{Name: emoji.ManLiftingWeights.String() + " Культурист", Weight: 2, Seconds: +8, Points: +10, Task: "Ты теперь мастер не только слова, но и тела, после каждого названного слова нужно присесть 1 раз"},
		{Name: emoji.PeopleWithBunnyEars.String() + " Командная работа", Weight: 1, Points: +8, Task: "Вы договаривает слова друг за другом и вообще непобедимы! Время поработать в команде! Твой сосед справа называет слова вместе с тобой по очереди, ты начинаешь первым"},
		{Name: emoji.PersonRunning.String() + " Флэш", Weight: 2, Points: +15, Seconds: -5, Task: "Тебя называют быстрейший из живых, в этом раунде у тебя на 5 сек меньше времени, покажи силу скорости!"},
		{Name: emoji.ManDancing.String() + " Диско", Weight: 2, Seconds: +5, Points: +20, Task: "Что ты мечтаешь услышать? Ты в танцах! Такие условия, один из игроков включает тебе песню и ты танцуешь ровно 20сек(нужно засечь), после чего нажимаешь стоп на таймере и получаешь +20очков, так что ты в танцах?"},
		//{Name: "Рекордсмен", Weight: 3, Points: 15, Task: "Тут всё не просто, поставь рекорд раунда! Если до тебя в раунде никто не сыграл, значит тебе повезло=)"},
		{Name: emoji.WaterWave.String() + " Волна удачи", Weight: 2, Points: +10, Task: "Тебя накрыла волна удачи, просто следуй за ней и делай свое дело"},
		{Name: emoji.WaterWave.String() + " Волна удачи", Weight: 2, Points: +10, Task: "Тебя накрыла волна удачи, просто следуй за ней и делай свое дело"},
		{Name: emoji.WomanGesturingNo.String() + " Неудача", Weight: 2, Points: -5, Task: "Карта не легла, встал не с той ноги, поел с ножа и уронил соль, чтобы ты не делал в этом раунде получается чуть-чуть хуже"},
		{Name: emoji.WomanGesturingNo.String() + " Неудача", Weight: 2, Points: -5, Task: "Карта не легла, встал не с той ноги, поел с ножа и уронил соль, чтобы ты не делал в этом раунде получается чуть-чуть хуже"},
		{Name: emoji.LoudlyCryingFace.String() + " Депрессия", Weight: 1, Seconds: +10, Points: -10, Task: "У тебя было много переработок по выходным, что вылилось в затяжную депрессию, но друзья приходят на помощь, в этом раунде за тебя играет человек напротив тебя, ты можешь дать ему смартфон"},
		{Name: emoji.IceHockey.String() + " Замена", Weight: 2, Task: "Ты как капитан хоккейной команды, сразу видишь слабые звенья и меняешь стратегии. Ты можешь заменить одну из сложных(на твой взгляд) категорий на другую, соответственно нужно будет назвать 2 слова заменяемую"},
		{Name: emoji.Bowling.String() + " Страйк", Weight: 2, Points: +7, Task: "После серии неудачных попыток ты наконец выбил страйк и сбил все кегли, называй слова только на одну, любую категорию"},
		{Name: emoji.Bomb.String() + "Бомба", Weight: 1, Points: +10, Task: "Бум! Что-то взорвалось, тебе нужно назвать слово на одну из категорий дважды, любую"},
		{Name: emoji.ManKneeling.String() + " Предложение", Weight: 2, Points: +5, Task: "Кажется наступил тот самый момент, которого ты ждал, называй слова, встав на одно колено!"},
		{Name: emoji.ManKneeling.String() + " Предложение", Weight: 2, Points: +5, Task: "Кажется наступил тот самый момент, которого ты ждал, называй слова, встав на одно колено!"},
		{Name: emoji.Divide.String() + " Математик", Weight: 2, Seconds: +13, Points: +15, Task: "Ты вдруг стал счетоводом среди варваров и многие считают тебя большим ученым, время подтвердить свою репутацию, каждый раз когда называешь слово, произнеси оставшееся количество секунд умноженное на 2. Например, если осталось 17 -> 34, если 23-> 46"},
		{Name: emoji.ClappingHands.String() + " Аплодисменты", Weight: 2, Points: +10, Task: "Ты наконец выступаешь на бродвее и пользуешься успехом публики, задание для остальных игроков, когда игрок произносит слово на выбранную букву нужно хлопнуть в ладоши"},
		{Name: emoji.ClappingHands.String() + " Аплодисменты", Weight: 2, Points: +10, Task: "Ты наконец выступаешь на бродвее и пользуешься успехом публики, задание для остальных игроков, когда игрок произносит слово на выбранную букву нужно хлопнуть в ладоши"},
		{Name: emoji.Ninja.String() + " Самурай", Weight: 2, Seconds: -10, Points: +20, Task: "Ты как самурай, готов ко всему, либо победа, либо смерть, у тебя будет на 10 сек меньше времени, но в награду получишь +20 очков"},
		{Name: emoji.SeeNoEvilMonkey.String() + " Блэкаут", Weight: 2, Points: +5, Task: "Мир погрузился во тьму, но ты готов к этому! Надо называть слова, закрыв глаза, вслепую, ты справишься!"},
		{Name: emoji.SeeNoEvilMonkey.String() + " Блэкаут", Weight: 2, Points: +5, Task: "Мир погрузился во тьму, но ты готов к этому! Надо называть слова, закрыв глаза, вслепую, ты справишься!"},
		{Name: emoji.Guitar.String() + " Музыкалити", Weight: 2, Points: +5, Task: "Ты идешь на звуки музыки, ты как герой фэнтези и у тебя есть свой бард, сочиняющий баллады о тебе, один из участников включает любую песню под которую вы играете раунд, конечно не на полную громкость"},
		{Name: emoji.Guitar.String() + " Музыкалити", Weight: 2, Points: +5, Task: "Ты идешь на звуки музыки, ты как герой фэнтези и у тебя есть свой бард, сочиняющий баллады о тебе, один из участников включает любую песню под которую вы играете раунд, конечно не на полную громкость"},
		{Name: emoji.MartialArtsUniform.String() + " Каратэ", Weight: 1, Points: +10, Seconds: +10, Task: "Ты долго тренировался и стал мастером боевых искусств, после каждого названного слова нужно изобразить удар каратэ с соответствующим звуком. Можешь не сильно стараться, это не экзамен"},
		{Name: emoji.FourLeafClover.String() + " Четырехлистный клевер", Weight: 2, Task: "Ты прогуливался как-то по лесу и увидел его - четырехлистный клевер. Удача! Ты можешь заменить выпавшую букву на любую другую"},
		{Name: emoji.UmbrellaWithRainDrops.String() + " Ненастье", Weight: 2, Seconds: -5, Task: "Плохая погода, или настроение, или на тебя кто-то накричал в маршрутке, короче, у тебя сгорело 5 сек, нужно выкручиваться"},
		{Name: emoji.Rainbow.String() + " Радуга", Weight: 2, Task: "Ты вышел во двор и увидел радугу, это был знак что ты на верном пути, в этом раунде ты можешь исключить одну категорию на свой выбор"},
		{Name: emoji.Unicorn.String() + " Единорог", Weight: 1, Seconds: +7, Points: +7, Task: "Как-то ты открыл дверь, а на пороге стоит единорог и требует поменять любую выпавшую букву на Е в этом раунде, ты с радостью согласился"},
		{Name: emoji.Snail.String() + " Улитка", Weight: 2, Seconds: +10, Points: -10, Task: "Бывают в жизни такие дни, что ты как улитка, вроде времени много, но толку нет, также и в этом раунде!"},
		{Name: emoji.Mage.String() + " Маг", Weight: 2, Seconds: +15, Points: +10, Task: "Из портала показался маг и говорит, чтобы выиграть нужно назвать дополнительно одно слово на магическую или фэнтези тематику на выпавшую букву"},
		{Name: emoji.RightFacingFist.String() + emoji.VictoryHand.String() + emoji.RaisedBackOfHand.String() + " Камень, ножницы, бумага", Weight: 2, Points: +10, Task: "Ты снова ощутил себя ребенком и вы опять поспорили кому играть первым, с соседом слева играете в камень, ножницы, бумага, кто побеждает, тот играет раунд. Очки достаются тебе"},
		{Name: emoji.GemStone.String() + emoji.Owl.String() + " Что? Где? Когда?", Weight: 2, Seconds: -10, Points: +20, Task: "Ты собрал команду и стал её капитаном, ты называешь слова, а остальные игроки должны подсказывать тебе, чтобы быстрее завершить раунд"},
		{Name: emoji.Ship.String() + " В одной лодке", Weight: 2, Seconds: +15, Task: "Корабль тонет! Нужно работать в команде, все игроки должны по очереди называть слова. Начинает действующий игрок, за ним игрок слева и по часовой стрелке, вперед!"},
		{Name: emoji.DivingMask.String() + " Аквалангист", Weight: 2, Points: +10, Task: "Ты погрузился с аквалангом и тут тебя застали врасплох. Нужно произносить слова зажав нос одной рукой!"},
		{Name: emoji.DivingMask.String() + " Аквалангист", Weight: 2, Points: +10, Task: "Ты погрузился с аквалангом и тут тебя застали врасплох. Нужно произносить слова зажав нос одной рукой!"},
		{Name: emoji.HighVoltage.String() + " Высокое напряжение", Weight: 1, Seconds: +15, Points: +15, Task: "Ты как участник финала интеллектуальной передачи по ТВ, тебе нужно придумать слова, и когда на таймере останется 10сек назвать все слова разом!"},
		{Name: emoji.WomanRunning.String() + " Фитнес тренер", Weight: 2, Points: +5, Task: "У тебя появился персональный тренер, который может сменить программу, пусть это будет игрок справа, вместо выпавшей буквы он может загадать свою полегче!"},
		{Name: emoji.ManBouncingBall.String() + " КМС", Weight: 2, Points: +5, Task: "Ты КМС в интеллектуальном клубе и решаешь проблемы приседаниями! Если не можешь придумать слово - приседаешь 1 раз, в следующий раз 2 и тд"},
		{Name: emoji.Brain.String() + " Цветы для Элджернона", Weight: 1, Seconds: +30, Points: +10, Task: "Твой IQ существенно вырос на короткое время и ты решаешь всех поразить, в этом раунде тебе нужно назвать ОДНО слово, но не начинающееся, а заканчивающееся на выпавшую букву!"},
		{Name: emoji.CrystalBall.String() + " Гадалка", Weight: 2, Task: "Ты решил попробовать себя в астрологии, до начала раунда игрок напротив загадывает одну из категорий(используемых в игре), если ты угадал, то выигрываешь раунд автоматом(нажми стоп на таймере сразу после начала), если нет играешь как обычно"},
		{Name: emoji.CrystalBall.String() + " Гадалка", Weight: 2, Task: "Ты решил попробовать себя в астрологии, до начала раунда игрок напротив загадывает одну из категорий(используемых в игре), если ты угадал, то выигрываешь раунд автоматом(нажми стоп на таймере сразу после начала), если нет играешь как обычно"},
		{Name: emoji.CrystalBall.String() + " Гадалка", Weight: 2, Task: "Ты решил попробовать себя в астрологии, до начала раунда игрок напротив загадывает одну из категорий(используемых в игре), если ты угадал, то выигрываешь раунд автоматом(нажми стоп на таймере сразу после начала), если нет играешь как обычно"},
		{Name: emoji.MoneyBag.String() + " Казино", Weight: 2, Task: "Ты любитель перекинуться в картишки и казино предлагает тебе сделку, подбрось монетку, если выпадет орел - ты выигрываешь раунд сразу(жми на стоп на таймере сразу после начала), если нет - проигрываешь(ждешь окончание таймера и не играешь) Что ты выбираешь?"},
		{Name: emoji.MoneyBag.String() + " Казино", Weight: 2, Task: "Ты любитель перекинуться в картишки и казино предлагает тебе сделку, подбрось монетку, если выпадет орел - ты выигрываешь раунд сразу(жми на стоп на таймере сразу после начала), если нет - проигрываешь(ждешь окончание таймера и не играешь) Что ты выбираешь?"},
		{Name: emoji.TestTube.String() + " Мамихлапинатапаи", Weight: 2, Task: "Ты нечаянно оказался на вечеринке со студентами и они предложили челлендж, ты можешь 5 раз произнести слово Мамихлапинатапаи подряд и нажать стоп на таймере или играть раунд как обычно. Выбор за тобой"},
	}

	BloopsKeys = map[string]Bloops{
		emoji.Cinema.String() + " Артхаус режиссер":              {Name: emoji.Cinema.String() + " Артхаус режиссер", Weight: 2, Seconds: +30, Points: +10, Task: "К тебе ворвался режиссер артхаус кино и предложил помочь со своим проектом, тебе нужно заменить категории в игре на категорию *кино и актеры*\nНазывай имена фильмов, актеров или режиссеров на выпавшую букву"},
		emoji.Flamingo.String() + "Фламинго":                     {Name: emoji.Flamingo.String() + "Фламинго", Weight: 2, Points: +10, Task: "Так получилось, что ты стал фламинго на время, когда называешь слова, ты должен стоять на одной ноге(можно держаться за что-нибудь)"},
		emoji.WomanSinger.String() + " На разогреве":             {Name: emoji.WomanSinger.String() + " На разогреве", Weight: 2, Points: +5, Task: "Ты начинающая рок звезда и тебя попросили выступить на разогреве, каждое слово, которое ты называешь ты должен пропеть в своем любим стиле"},
		emoji.Hammer.String() + " Мастерство":                    {Name: emoji.Hammer.String() + " Мастерство", Weight: 1, Points: +20, Seconds: +20, Task: "Из поколения в поколение ты передавалась твоё ремесло, время показать себя на ярмарке! Тебе нужно на каждую категорию назвать не одно слово, а два"},
		emoji.ManLiftingWeights.String() + " Культурист":         {Name: emoji.ManLiftingWeights.String() + " Культурист", Weight: 2, Seconds: +8, Points: +10, Task: "Ты теперь мастер не только слова, но и тела, после каждого названного слова нужно присесть 1 раз"},
		emoji.PeopleWithBunnyEars.String() + " Командная работа": {Name: emoji.PeopleWithBunnyEars.String() + " Командная работа", Weight: 1, Points: +8, Task: "Вы договаривает слова друг за другом и вообще непобедимы! Время поработать в команде! Твой сосед справа называет слова вместе с тобой по очереди, ты начинаешь первым"},
		emoji.PersonRunning.String() + " Флэш":                   {Name: emoji.PersonRunning.String() + " Флэш", Weight: 2, Points: +15, Seconds: -5, Task: "Тебя называют быстрейший из живых, в этом раунде у тебя на 5 сек меньше времени, покажи силу скорости!"},
		emoji.ManDancing.String() + " Диско":                     {Name: emoji.ManDancing.String() + " Диско", Weight: 2, Seconds: +5, Points: +20, Task: "Что ты мечтаешь услышать? Ты в танцах! Такие условия, один из игроков включает тебе песню и ты танцуешь ровно 20сек(нужно засечь), после чего нажимаешь стоп на таймере и получаешь +20очков, так что ты в танцах?"},
		//{Name: "Рекордсмен", Weight: 3, Points: 15, Task: "Тут всё не просто, поставь рекорд раунда! Если до тебя в раунде никто не сыграл, значит тебе повезло=)"},
		emoji.WaterWave.String() + " Волна удачи":                {Name: emoji.WaterWave.String() + " Волна удачи", Weight: 2, Points: +10, Task: "Тебя накрыла волна удачи, просто следуй за ней и делай свое дело"},
		emoji.WomanGesturingNo.String() + " Неудача":             {Name: emoji.WomanGesturingNo.String() + " Неудача", Weight: 2, Points: -5, Task: "Карта не легла, встал не с той ноги, поел с ножа и уронил соль, чтобы ты не делал в этом раунде получается чуть-чуть хуже"},
		emoji.LoudlyCryingFace.String() + " Депрессия":           {Name: emoji.LoudlyCryingFace.String() + " Депрессия", Weight: 1, Seconds: +10, Points: -10, Task: "У тебя было много переработок по выходным, что вылилось в затяжную депрессию, но друзья приходят на помощь, в этом раунде за тебя играет человек напротив тебя, ты можешь дать ему смартфон"},
		emoji.IceHockey.String() + " Замена":                     {Name: emoji.IceHockey.String() + " Замена", Weight: 2, Task: "Ты как капитан хоккейной команды, сразу видишь слабые звенья и меняешь стратегии. Ты можешь заменить одну из сложных(на твой взгляд) категорий на другую, соответственно нужно будет назвать 2 слова заменяемую"},
		emoji.Bowling.String() + " Страйк":                       {Name: emoji.Bowling.String() + " Страйк", Weight: 2, Points: +7, Task: "После серии неудачных попыток ты наконец выбил страйк и сбил все кегли, называй слова только на одну, любую категорию"},
		emoji.Bomb.String() + "Бомба":                            {Name: emoji.Bomb.String() + "Бомба", Weight: 1, Points: +10, Task: "Бум! Что-то взорвалось, тебе нужно назвать слово на одну из категорий дважды, любую"},
		emoji.ManKneeling.String() + " Предложение":              {Name: emoji.ManKneeling.String() + " Предложение", Weight: 2, Points: +5, Task: "Кажется наступил тот самый момент, которого ты ждал, называй слова, встав на одно колено!"},
		emoji.Divide.String() + " Математик":                     {Name: emoji.Divide.String() + " Математик", Weight: 2, Seconds: +13, Points: +15, Task: "Ты вдруг стал счетоводом среди варваров и многие считают тебя большим ученым, время подтвердить свою репутацию, каждый раз когда называешь слово, произнеси оставшееся количество секунд умноженное на 2. Например, если осталось 17 -> 34, если 23-> 46"},
		emoji.ClappingHands.String() + " Аплодисменты":           {Name: emoji.ClappingHands.String() + " Аплодисменты", Weight: 2, Points: +10, Task: "Ты наконец выступаешь на бродвее и пользуешься успехом публики, задание для остальных игроков, когда игрок произносит слово на выбранную букву нужно хлопнуть в ладоши"},
		emoji.Ninja.String() + " Самурай":                        {Name: emoji.Ninja.String() + " Самурай", Weight: 2, Seconds: -10, Points: +20, Task: "Ты как самурай, готов ко всему, либо победа, либо смерть, у тебя будет на 10 сек меньше времени, но в награду получишь +20 очков"},
		emoji.SeeNoEvilMonkey.String() + " Блэкаут":              {Name: emoji.SeeNoEvilMonkey.String() + " Блэкаут", Weight: 2, Points: +5, Task: "Мир погрузился во тьму, но ты готов к этому! Надо называть слова, закрыв глаза, вслепую, ты справишься!"},
		emoji.Guitar.String() + " Музыкалити":                    {Name: emoji.Guitar.String() + " Музыкалити", Weight: 2, Points: +5, Task: "Ты идешь на звуки музыки, ты как герой фэнтези и у тебя есть свой бард, сочиняющий баллады о тебе, один из участников включает любую песню под которую вы играете раунд, конечно не на полную громкость"},
		emoji.MartialArtsUniform.String() + " Каратэ":            {Name: emoji.MartialArtsUniform.String() + " Каратэ", Weight: 1, Points: +10, Seconds: +10, Task: "Ты долго тренировался и стал мастером боевых искусств, после каждого названного слова нужно изобразить удар каратэ с соответствующим звуком. Можешь не сильно стараться, это не экзамен"},
		emoji.FourLeafClover.String() + " Четырехлистный клевер": {Name: emoji.FourLeafClover.String() + " Четырехлистный клевер", Weight: 2, Task: "Ты прогуливался как-то по лесу и увидел его - четырехлистный клевер. Удача! Ты можешь заменить выпавшую букву на любую другую"},
		emoji.UmbrellaWithRainDrops.String() + " Ненастье":       {Name: emoji.UmbrellaWithRainDrops.String() + " Ненастье", Weight: 2, Seconds: -5, Task: "Плохая погода, или настроение, или на тебя кто-то накричал в маршрутке, короче, у тебя сгорело 5 сек, нужно выкручиваться"},
		emoji.Rainbow.String() + " Радуга":                       {Name: emoji.Rainbow.String() + " Радуга", Weight: 2, Task: "Ты вышел во двор и увидел радугу, это был знак что ты на верном пути, в этом раунде ты можешь исключить одну категорию на свой выбор"},
		emoji.Unicorn.String() + " Единорог":                     {Name: emoji.Unicorn.String() + " Единорог", Weight: 1, Seconds: +7, Points: +7, Task: "Как-то ты открыл дверь, а на пороге стоит единорог и требует поменять любую выпавшую букву на Е в этом раунде, ты с радостью согласился"},
		emoji.Snail.String() + " Улитка":                         {Name: emoji.Snail.String() + " Улитка", Weight: 2, Seconds: +10, Points: -10, Task: "Бывают в жизни такие дни, что ты как улитка, вроде времени много, но толку нет, также и в этом раунде!"},
		emoji.Mage.String() + " Маг":                             {Name: emoji.Mage.String() + " Маг", Weight: 2, Seconds: +15, Points: +10, Task: "Из портала показался маг и говорит, чтобы выиграть нужно назвать дополнительно одно слово на магическую или фэнтези тематику на выпавшую букву"},
		emoji.RightFacingFist.String() + emoji.VictoryHand.String() + emoji.RaisedBackOfHand.String() + " Камень, ножницы, бумага": {Name: emoji.RightFacingFist.String() + emoji.VictoryHand.String() + emoji.RaisedBackOfHand.String() + " Камень, ножницы, бумага", Weight: 2, Points: +10, Task: "Ты снова ощутил себя ребенком и вы опять поспорили кому играть первым, с соседом слева играете в камень, ножницы, бумага, кто побеждает, тот играет раунд. Очки достаются тебе"},
		emoji.GemStone.String() + emoji.Owl.String() + " Что? Где? Когда?":                                                         {Name: emoji.GemStone.String() + emoji.Owl.String() + " Что? Где? Когда?", Weight: 2, Seconds: -10, Points: +20, Task: "Ты собрал команду и стал её капитаном, ты называешь слова, а остальные игроки должны подсказывать тебе, чтобы быстрее завершить раунд"},
		emoji.Ship.String() + " В одной лодке":                                                                                     {Name: emoji.Ship.String() + " В одной лодке", Weight: 2, Seconds: +15, Task: "Корабль тонет! Нужно работать в команде, все игроки должны по очереди называть слова. Начинает действующий игрок, за ним игрок слева и по часовой стрелке, вперед!"},
		emoji.DivingMask.String() + " Аквалангист":                                                                                 {Name: emoji.DivingMask.String() + " Аквалангист", Weight: 2, Points: +10, Task: "Ты погрузился с аквалангом и тут тебя застали врасплох. Нужно произносить слова зажав нос одной рукой!"},
		emoji.HighVoltage.String() + " Высокое напряжение":                                                                         {Name: emoji.HighVoltage.String() + " Высокое напряжение", Weight: 1, Seconds: +15, Points: +15, Task: "Ты как участник финала интеллектуальной передачи по ТВ, тебе нужно придумать слова, и когда на таймере останется 10сек назвать все слова разом!"},
		emoji.WomanRunning.String() + " Фитнес тренер":                                                                             {Name: emoji.WomanRunning.String() + " Фитнес тренер", Weight: 2, Points: +5, Task: "У тебя появился персональный тренер, который может сменить программу, пусть это будет игрок справа, вместо выпавшей буквы он может загадать свою полегче!"},
		emoji.ManBouncingBall.String() + " КМС":                                                                                    {Name: emoji.ManBouncingBall.String() + " КМС", Weight: 2, Points: +5, Task: "Ты КМС в интеллектуальном клубе и решаешь проблемы приседаниями! Если не можешь придумать слово - приседаешь 1 раз, в следующий раз 2 и тд"},
		emoji.Brain.String() + " Цветы для Элджернона":                                                                             {Name: emoji.Brain.String() + " Цветы для Элджернона", Weight: 1, Seconds: +30, Points: +10, Task: "Твой IQ существенно вырос на короткое время и ты решаешь всех поразить, в этом раунде тебе нужно назвать ОДНО слово, но не начинающееся, а заканчивающееся на выпавшую букву!"},
		emoji.CrystalBall.String() + " Гадалка":                                                                                    {Name: emoji.CrystalBall.String() + " Гадалка", Weight: 2, Task: "Ты решил попробовать себя в астрологии, до начала раунда игрок напротив загадывает одну из категорий(используемых в игре), если ты угадал, то выигрываешь раунд автоматом(нажми стоп на таймере сразу после начала), если нет играешь как обычно"},
		emoji.MoneyBag.String() + " Казино":                                                                                        {Name: emoji.MoneyBag.String() + " Казино", Weight: 2, Task: "Ты любитель перекинуться в картишки и казино предлагает тебе сделку, подбрось монетку, если выпадет орел - ты выигрываешь раунд сразу(жми на стоп на таймере сразу после начала), если нет - проигрываешь(ждешь окончание таймера и не играешь) Что ты выбираешь?"},
		emoji.TestTube.String() + " Мамихлапинатапаи":                                                                              {Name: emoji.TestTube.String() + " Мамихлапинатапаи", Weight: 2, Task: "Ты нечаянно оказался на вечеринке со студентами и они предложили челлендж, ты можешь 5 раз произнести слово Мамихлапинатапаи подряд и нажать стоп на таймере или играть раунд как обычно. Выбор за тобой"},
	}
)
