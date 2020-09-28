package hw03_frequency_analysis //nolint:golint

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed
var taskWithAsteriskIsCompleted = true

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{"он", "а", "и", "что", "ты", "не", "если", "то", "его", "Кристофер", "Робин", "в"}
			require.Subset(t, expected, Top10(text))
		} else {
			expected := []string{"он", "и", "а", "что", "ты", "не", "если", "-", "то", "Кристофер"}
			require.ElementsMatch(t, expected, Top10(text))
		}
	})
}

func TestSplitCount(t *testing.T) {
	s := `раз раз раз раз раз раз раз раз раз раз
		два два два два два два два два два
		три три три три три три три три
		четыре четыре четыре четыре четыре четыре четыре
		пять пять пять пять пять пять
		шесть шесть шесть шесть шесть
		семь семь семь семь
		восемь восемь восемь 
		девять девять
		десять десять
		одиннадцать`
	expected := []string{"раз", "два", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять", "десять"}
	require.ElementsMatch(t, expected, Top10(s))
}

func TestSplitToMap(t *testing.T) {

	t.Run("First task", func(t *testing.T) {
		s := "One one, ONE"
		res, _ := splitToMap(s, `\s`)
		keys := make([]string, 0, 3)
		for key := range res {
			keys = append(keys, key)
		}
		array := []string{"One", "one,", "ONE"}
		require.ElementsMatch(t, array, keys)
	})

	t.Run("Split empty str", func(t *testing.T) {
		s := ""
		_, err := splitToMap(s, "")
		require.Error(t, err)
	})

}

func TestSplitAdvToMap(t *testing.T) {
	s := "One one 'one' one, one! two ones Кристофер Кристофером"
	array := []string{"One", "two", "ones", "Кристофер", "Кристофером"}
	res, _ := splitAdvToMap(s, `\s`)
	keys := make([]string, 0, 3)
	for key := range res {
		keys = append(keys, key)
	}
	require.ElementsMatch(t, array, keys)
}

func TestAnalyzeWord(t *testing.T) {
	list := []string{"Jhon", "Bill", "carry"}
	word := "jhon"
	array, retw, _ := analyzeWord(list, word)
	t.Run("Find existing", func(t *testing.T) {
		require.ElementsMatch(t, list, array)
		require.Equal(t, "Jhon", retw)
	})

	t.Run("Append new name", func(t *testing.T) {
		array, retw, _ = analyzeWord(list, "Larry")
		list = append(list, "Larry")
		require.ElementsMatch(t, list, array)
		require.Equal(t, "Larry", retw)
	})

	t.Run("Try empty string", func(t *testing.T) {
		_, _, err := analyzeWord(list, "")
		require.Error(t, err)
	})
}
