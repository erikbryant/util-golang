package bins

import (
	"log"
	"os"
	"testing"
)

func quiet() func() {
	null, _ := os.Open(os.DevNull)
	sout := os.Stdout
	serr := os.Stderr
	os.Stdout = null
	os.Stderr = null
	log.SetOutput(null) // Also redirect log package output
	return func() {
		defer null.Close()
		os.Stdout = sout
		os.Stderr = serr
		log.SetOutput(os.Stderr) // Restore log output to stderr
	}
}

func TestMultiply(t *testing.T) {
	testCases := []struct {
		x            int
		f            int
		twos         int
		expected     int
		expectedTwos int
	}{
		// Remember that we ignore numbers w trailing zeroes
		{1, 1, 0, 1, 0},
		{1, 2, 0, 2, 0},
		{2, 1, 0, 1, 1},
		{2, 1, MaxFives, 2, MaxFives},
	}

	for _, testCase := range testCases {
		answer, answerTwos := multiply(testCase.x, testCase.f, testCase.twos)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d*%d %d expected %d, %d got %d, %d", testCase.x, testCase.f, testCase.twos, testCase.expected, testCase.expectedTwos, answer, answerTwos)
		}
		if answerTwos != testCase.expectedTwos {
			t.Errorf("ERROR: For %d*%d %d expected %d, %d got %d, %d", testCase.x, testCase.f, testCase.twos, testCase.expected, testCase.expectedTwos, answer, answerTwos)
		}
	}
}

func TestFix(t *testing.T) {
	testCases := []struct {
		f        int
		twos     int
		expected int
	}{
		// Remember that we ignore numbers w trailing zeroes
		{0, 0, 0},
		{1, 0, 1},
		{99, 0, 99},
		{0, 1, 0},
		{1, 1, 2},
		{99, 1, 198},
		{3, 3, 24},
	}

	for _, testCase := range testCases {
		answer := fix(testCase.f, testCase.twos)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d %d expected %d got %d", testCase.f, testCase.twos, testCase.expected, answer)
		}
	}
}

func TestFactorial(t *testing.T) {
	testCases := []struct {
		c        int
		expected int
	}{
		{0, 1},                      // wolfram 1
		{1, 1},                      // wolfram 1
		{9, 36288},                  // wolfram 36288
		{10, 36288},                 // wolfram 36288
		{20, 817664},                // wolfram 90200817664
		{30, 3630848},               // wolfram 5863630848
		{31, 2556288},               // wolfram 1772556288
		{49, 7921024},               // wolfram 3137921024
		{100, 916864},               // wolfram 5210916864
		{1000, 7753472},             // wolfram 70027753472
		{1000 * 1000 * 100, 754176}, // wolfram 840754176

		// Worst case of 5^k and 2^0
		//{30517578125, 5417088}, // wolfram 5652355417088

		// Spares in case debugging is needed
		//{2000000, 4194688},                   // wolfram 7934194688
		//{6000000, 8792576},                   // wolfram 900448792576
		//{8000000, 3638144},                   // wolfram 897253638144
		//{9000000, 788096},                    // wolfram 381750788096
		//{9000001, 4788096},                   // wolfram 474614788096
		//{9500000, 1642752},                   // wolfram 132701642752
		//{9700000, 9350016},                   // wolfram 199499350016
		//{9710000, 1207168},                   // wolfram 400131207168
		//{9717000, 5531904},                   // wolfram 233985531904
		//{9721000, 7065216},                   // wolfram 556317065216
		//{9723000, 5980544},                   // wolfram 4022575980544
		//{9724000, 8702208},                   // wolfram 386628702208
		//{9724061, 8407168},                   // wolfram 316918407168
		//{9724093, 5806976},                   // wolfram 4063495806976
		//{9724065, 8691328},                   // wolfram 173598691328
		//{9724100, 9267712},                   // wolfram 755489267712
		//{9724112, 2362752},                   // wolfram 894012362752
		//{9724119, 3499264},                   // wolfram 444943499264
		//{9724122, 8761216},                   // wolfram 754298761216
		//{9724123, 2013568},                   // wolfram 932812013568
		//{9724124, 4914432},                   // wolfram 688624914432
		//{9724125, 6051072},                   // wolfram 649746051072
		//{9724250, 118784},                    // wolfram 4096320118784
		//{9724500, 4380672},                   // wolfram 938084380672
		//{9725000, 6924544},                   // wolfram 553666924544
		//{9750000, 3750016},                   // wolfram 800113750016
		//{9799999, 1900672},                   // wolfram 787151900672
		//{9800000, 6265856},                   // wolfram 140886265856
		//{9800001, 5065856},                   // wolfram 546275065856
		//{9825000, 9130624},                   // wolfram 756879130624
		//{9850000, 3322496},                   // wolfram 401353322496
		//{9900000, 2721024},                   // wolfram 521742721024
		//{13021100, 9124736},                  // wolfram 453259124736
		//{13030000, 9129344},                  // wolfram 270979129344
		//{13050000, 6526208},                  // wolfram 414866526208
		//{13056000, 3527936},                  // wolfram 423183527936
		//{13060000, 5911936},                  // wolfram 914655911936
		//{13061249, 1176832},                  // wolfram 255181176832
		//{13061250, 4589696},                  // wolfram 798514589696
		//{13062000, 6899072},                  // wolfram 296416899072
		//{13062250, 1729408},                  // wolfram 567011729408
		//{13062400, 4966144},                  // wolfram 902664966144
		//{13062450, 6141184},                  // wolfram 102266141184
		//{13062462, 7133184},                  // wolfram 456487133184
		//{13062470, 4786176},                  // wolfram 871844786176
		//{13062473, 1877376},                  // wolfram 840731877376
		//{13062474, 5188224},                  // wolfram 289195188224
		//{13062475, 2962944},                  // wolfram 549162962944
		//{13062500, 3442816},                  // wolfram 443953442816
		//{13125000, 6384256},                  // wolfram 461036384256
		//{13250000, 9667328},                  // wolfram 301789667328
		//{13500000, 3213824},                  // wolfram 82683213824
		//{1000 * 100, 7162496},                // wolfram 957162496
		//{1000 * 1000, 8412544},               // wolfram 5058412544
		//{1000 * 1000 * 10, 4194688},          // wolfram 574194688
		//{1000 * 1000 * 13, 382208},           // wolfram 53580382208
		//{1000 * 1000 * 14, 4550272},          // wolfram 14554550272
		//{1000 * 1000 * 15, 3620736},          // wolfram 49493620736
		//{1000 * 1000 * 17, 6981504},          // wolfram 95416981504
		//{1000 * 1000 * 25, 8412544},          // wolfram 90018412544
		//{1000 * 1000 * 50, 4194688},          // wolfram 774194688
		//{1000 * 1000 * 1000, 3638144},        // wolfram 933638144
		//{1000 * 1000 * 1000 * 10, 1946112},   // wolfram 36441946112
		//{1000 * 1000 * 1000 * 100, 8167808},  // wolfram 31378167808
		//{1000 * 1000 * 1000 * 1000, 3416576}, // wolfram 60283416576
	}

	for _, testCase := range testCases {
		answer := factorial(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d! expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}

func TestFactorialNoTens(t *testing.T) {
	defer quiet()()

	testCases := []struct {
		l        int
		h        int
		expected int
	}{
		// Remember that we ignore numbers w trailing zeroes
		{1, 1, 1},
		{1, 2, 2},
		{1, 5, 12},                  // wolfram 12
		{1, 99, 1744128},            // wolfram 341744128
		{10, 11, 11},                //
		{21, 29, 2450144},           // wolfram 342450144
		{20, 30, 2450144},           // wolfram 342450144
		{1, 99, 1744128},            // wolfram 341744128
		{11, 99, 749056},            // wolfram 320749056
		{101, 999, 2630016},         // wolfram 932630016
		{1001, 9999, 7715968},       // wolfram 797715968
		{10001, 99999, 8257408},     // wolfram 938257408
		{100001, 999999, 8742272},   // self-reported
		{1000001, 9999999, 1511168}, // self-reported
	}

	for _, testCase := range testCases {
		answer := factorialNoTens(testCase.l, testCase.h)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d-%d expected %d, got %d", testCase.l, testCase.h, testCase.expected, answer)
		}
	}
}

func TestPower(t *testing.T) {
	testCases := []struct {
		b        int
		e        int
		expected int
	}{
		// Remember that power() also strips trailing zeroes
		{1, 1, 1},
		{9, 2, 81},
		{10, 3, 1},
		{20, 15, 32768},
		{1000, 230, 1},
		{861511168, 1, 1511168},
	}

	for _, testCase := range testCases {
		answer := power(testCase.b, testCase.e)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d^%d expected %d, got %d", testCase.b, testCase.e, testCase.expected, answer)
		}
	}
}

func TestMakeDataset(t *testing.T) {
	testCases := []struct {
		upper             int
		expectedStages    int
		expectedStageZero int
		expectedF         int
	}{
		{10, 2, 2, 36288},
		{10000000, 8, 8, 4194688},
		{100000000, 8, 18, 754176},
		{1000000000000, 8, 111118, 3416576},
	}

	for _, testCase := range testCases {
		dataset := makeDataset(testCase.upper)
		if dataset.upper != testCase.upper {
			t.Errorf("ERROR: For %d expected upper = %d, got %d", testCase.upper, testCase.upper, dataset.upper)
		}
		if len(dataset.stages) != testCase.expectedStages {
			t.Errorf("ERROR: For %d expected len(stages) = %d, got %d", testCase.upper, testCase.expectedStages, len(dataset.stages))
		}
		if dataset.stages[0].count != testCase.expectedStageZero {
			t.Errorf("ERROR: For %d expected stage[0].count = %d, got %d", testCase.upper, testCase.expectedStageZero, dataset.stages[0].count)
		}
		if dataset.expected != testCase.expectedF {
			t.Errorf("ERROR: For %d expected expected = %d, got %d", testCase.upper, testCase.expectedF, dataset.expected)
		}
	}
}
