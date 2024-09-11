package hashLib

import (
	"testing"
	"fmt"
	"math/rand"
	"log"
)


func TestHttp (t *testing.T) {

    var kw [9][]byte
	var respr[9]uint32
	var resdb[9]uint32
	var resprV2[9]int
	var resdbV2[9]int
	var resprAltV2[9]int
//	var resdbAltV2[9]int

    kw[0] = []byte("GET")
    kw[1] = []byte("POST")
    kw[2] = []byte("PUT")
    kw[3] = []byte("DELETE")
    kw[4] = []byte("HEAD")
    kw[5] = []byte("CONNECT")
    kw[6] = []byte("OPTIONS")
    kw[7] = []byte("TRACE")
    kw[8] = []byte("PATCH")

    for i, v := range kw {
        resdb[i]= djb(kw[i])
        respr[i]= prhash(kw[i])
//  for i, v := range kw {
        fmt.Printf(" --%2d: %10s %10d %4d\n", i, string(v), resdb[i], respr[i] )
    }

	for i:= 0; i<8; i++ {
		for j:= i+1; j< 9; j++ {
			if resdb[i] == resdb[j] {t.Errorf("error db %d = %d : %d \n", i, j, resdb[i]) }
			if respr[i] == respr[j] {t.Errorf("error pr %d = %d : %d \n", i, j, respr[i]) }
		}
	}

	fmt.Println("********* V2 **************")

    for i, v := range kw {
        resdbV2[i] = djbV2(kw[i])
        resprV2[i] = prhashV2(kw[i])
//  for i, v := range kw {
        fmt.Printf(" --%2d: %10s %10d %4d\n", i, string(v), resdbV2[i], resprV2[i] )
    }

	for i:= 0; i<8; i++ {
		for j:= i+1; j< 9; j++ {
			if resdbV2[i] == resdbV2[j] {t.Errorf("error dbV2 %d = %d : %d \n", i, j, resdbV2[i]) }
			if resprV2[i] == resprV2[j] {t.Errorf("error prV2 %d = %d : %d \n", i, j, resprV2[i]) }
		}
	}

	fmt.Println("********* AltV2 **************")
    for i, v := range kw {
//        resdbV2[i] = djbV2(kw[i])
        resprAltV2[i] = prhashAltV2(kw[i])
        fmt.Printf(" --%2d: %10s %4d\n", i, string(v), resprAltV2[i] )
    }

	for i:= 0; i<8; i++ {
		for j:= i+1; j< 9; j++ {
			if resprAltV2[i] == resprAltV2[j] {t.Errorf("error prAltV2 %d = %d : %d \n", i, j, resprAltV2[i]) }
		}
	}

	fmt.Println("********* end **************")
}


func TestCLI (t *testing.T) {

    var kw [8][]byte
	var respr[8]uint32
	var resdb[8]uint32
	var resprV2[8]int
	var resdbV2[8]int

    kw[0] = []byte("")
    kw[1] = []byte("/")
    kw[2] = []byte("/js/")
    kw[3] = []byte("/pdf/")
    kw[4] = []byte("/hijack")
    kw[5] = []byte("/xjson")
    kw[6] = []byte("/json/")
    kw[7] = []byte("/foo")

    for i, v := range kw {
        resdb[i]= djb(kw[i])
        respr[i]= prhashAlt(kw[i])
//  for i, v := range kw {
        fmt.Printf(" --%2d: %10s %10d %4d\n", i, string(v), resdb[i], respr[i] )
    }

	for i:= 1; i<7; i++ {
		for j:= i+1; j< 8; j++ {
			if resdb[i] == resdb[j] {t.Errorf("error db %d = %d : %d \n", i, j, resdb[i]) }
			if respr[i] == respr[j] {t.Errorf("error pr %d = %d : %d \n", i, j, respr[i]) }
		}
	}

	fmt.Println("********* V2 **************")

    for i, v := range kw {
        resdbV2[i] = djbV2(kw[i])
        resprV2[i] = prhashAltV2(kw[i])
//  for i, v := range kw {
        fmt.Printf(" --%2d: %10s %10d %4d\n", i, string(v), resdbV2[i], resprV2[i] )
    }
	fmt.Println("*********** end V2 **************")
}

func TestDir (t *testing.T) {

    var dsl [8][]byte
	var respdb[8]int

    dsl[0] = []byte("")
    dsl[1] = []byte("/")
    dsl[2] = []byte("/js/")
    dsl[3] = []byte("/pdf/")
    dsl[4] = []byte("/hijack")
    dsl[5] = []byte("/xjson")
    dsl[6] = []byte("/foo")
    dsl[7] = []byte("/json/")

	fmt.Println("*********** Dir **************")

    for i, v1 := range dsl {
        respdb[i] = prhashDir(dsl[i])
        fmt.Printf(" --%d: %10s %10d\n", i, string(v1),respdb[i] )
    }
}

func BenchmarkHttppr(b *testing.B) {

    var kw [9][]byte

    kw[0] = []byte("GET")
    kw[1] = []byte("POST")
    kw[2] = []byte("PUT")
    kw[3] = []byte("DELETE")
    kw[4] = []byte("HEAD")
    kw[5] = []byte("CONNECT")
    kw[6] = []byte("OPTIONS")
    kw[7] = []byte("TRACE")
    kw[8] = []byte("PATCH")


	for n := 0; n < b.N; n++ {
//		b.StopTimer()
		randNum := n%9
//		randNum := rand.Intn(9)
//		b.StartTimer()
		_= prhash(kw[randNum])
	}
}

func BenchmarkHttRand(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_= rand.Intn(9)
	}
}

func BenchmarkHttdjb(b *testing.B) {

    var kw [9][]byte

    kw[0] = []byte("GET")
    kw[1] = []byte("POST")
    kw[2] = []byte("PUT")
    kw[3] = []byte("DELETE")
    kw[4] = []byte("HEAD")
    kw[5] = []byte("CONNECT")
    kw[6] = []byte("OPTIONS")
    kw[7] = []byte("TRACE")
    kw[8] = []byte("PATCH")


	for n := 0; n < b.N; n++ {
		randNum := rand.Intn(9)
		_= djb(kw[randNum])
	}

}

func BenchmarkHttpPr2(b *testing.B) {

    var kw [9][]byte

    kw[0] = []byte("GET")
    kw[1] = []byte("POST")
    kw[2] = []byte("PUT")
    kw[3] = []byte("DELETE")
    kw[4] = []byte("HEAD")
    kw[5] = []byte("CONNECT")
    kw[6] = []byte("OPTIONS")
    kw[7] = []byte("TRACE")
    kw[8] = []byte("PATCH")


	for n := 0; n < b.N; n++ {
		randNum := rand.Intn(9)
		_= prhashV2(kw[randNum])
	}

}

func BenchmarkHttpdjb2(b *testing.B) {

    var kw [9][]byte

    kw[0] = []byte("GET")
    kw[1] = []byte("POST")
    kw[2] = []byte("PUT")
    kw[3] = []byte("DELETE")
    kw[4] = []byte("HEAD")
    kw[5] = []byte("CONNECT")
    kw[6] = []byte("OPTIONS")
    kw[7] = []byte("TRACE")
    kw[8] = []byte("PATCH")


	for n := 0; n < b.N; n++ {
		randNum := rand.Intn(9)
		_= djbV2(kw[randNum])
	}

}

func BenchmarkRouter(b *testing.B) {

    var kw [8][]byte

    kw[0] = []byte("")
    kw[1] = []byte("/")
    kw[2] = []byte("/js/")
    kw[3] = []byte("/pdf/")
    kw[4] = []byte("/hijack")
    kw[5] = []byte("/xjson")
    kw[6] = []byte("/foo")
    kw[7] = []byte("/json/")


	for n := 0; n < b.N; n++ {
		randNum := rand.Intn(8)

		res := prhashDir(kw[randNum])

		switch res {
		case 0:

		case 2:

		case 4:

		case 5:

		case 6:

		case 14:

		case 15:

		default:
			log.Fatalf("error res: %d\n", res)
		}
	}
}
