package hashLib

func prhash (s []byte) uint32 {
    var hash uint32

//    hash = uint32(s[0]) + uint32(s[1])*256
    hash = uint32(s[0]) + uint32(s[1])*13

    return hash%16
}

func prhashAlt (s []byte) uint32 {
    var hash uint32

//	if len(s) == 0 {return 0}
//    if len(s) == 1 {return uint32(s[0])}
	if len(s) < 2 {return 0}
//  hash = uint32(s[0]) + 3*uint32(s[1] + byte(len(s)))
    hash = uint32(s[0]) + 3*uint32(s[len(s)-2]) + 7*uint32(s[len(s)-1])
//  hash = uint32(s[0]) + uint32(s[1])*uint32(s[1])
    hash = hash % uint32(256)
    return hash
}


func djb(s []byte) uint32 {
        var hash uint32 = 5381

        for _, c := range s {
                hash = ((hash << 5) + hash) + uint32(c)
                // the above line is an optimized version of the following line:
                // hash = hash * 33 + int64(c)
                // which is easier to read and understand...
        }

        return hash
}

func prhashV2 (s []byte) int {
    var hash int

    hash = int(s[0]) + 4*int(s[1])
//  hash = int(s[0]) + int(s[1])*int(s[1])
	hash = hash & 255
    return hash
}

func prhashAltV2 (s []byte) int {
    var hash int

    if len(s) < 2 {return 0}
//  hash = uint32(s[0]) + 3*uint32(s[1] + byte(len(s)))
    hash = int(s[0]) + 3*int(s[len(s)-2]) + 7*int(s[len(s)-1])
//  hash = uint32(s[0]) + uint32(s[1])*uint32(s[1])
    hash = hash % 256
    return hash
}

func djbV2(s []byte) int {
        var hash int = 5381

        for _, c := range s {
                hash = ((hash << 5) + hash) + int(c)
                // the above line is an optimized version of the following line:
                // hash = hash * 33 + int64(c)
                // which is easier to read and understand...
        }

	hash = hash & 255
        return hash
}

func prhashDir(dir []byte) int {

	var hash1 int
	if len(dir) <2 {
		hash1 = 0
	} else {
//    	hash1 = int(dir[1]) + 7*int(dir[len(dir)-2]) +13
    	hash1 = 3*int(dir[1]) + len(dir)
	}

	return hash1%16
}

func prhashCombo(dir []byte, cmd []byte) int {
    var hash int
	var hash1 int

    hash2 := int(cmd[0]) + 7*int(cmd[1])
	if len(dir) <2 {
		hash1 = 0
	} else {
    	hash1 = int(dir[1]) + 3*int(dir[len(dir)-2]) + 7*int(dir[len(dir)-1])
	}
	hash = (7*hash1 + 13*hash2)%256
	if (hash == 156) && (dir[1] == 'x') {hash = 158}
	return hash
}
