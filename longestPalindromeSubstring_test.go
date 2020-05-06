package leetcode

import (
    "testing"
)

// Solution to https://leetcode.com/problems/longest-palindromic-substring/
func TestPalindromeLongestSubstring(t *testing.T) {
    result := longestPalindrome("babad")
    expectedResult := "bab"
    if result != expectedResult {
        t.Errorf("Got %s; expected %s", result, expectedResult)
    }
}

func TestPalindromeLongestSubstring2(t *testing.T) {
    result := longestPalindrome("a")
    expectedResult := "a"
    if result != expectedResult {
        t.Errorf("Got %s; expected %s", result, expectedResult)
    }
}

func TestPalindromeLongestSubstring3(t *testing.T) {
    result := longestPalindrome("bb")
    expectedResult := "bb"
    if result != expectedResult {
        t.Errorf("Got %s; expected %s", result, expectedResult)
    }
}

func TestPalindromeLongestSubstring4(t *testing.T) {
    result := longestPalindrome("zudfweormatjycujjirzjpyrmaxurectxrtqedmmgergwdvjmjtstdhcihacqnothgttgqfywcpgnuvwglvfiuxteopoyizgehkwuvvkqxbnufkcbodlhdmbqyghkojrgokpwdhtdrwmvdegwycecrgjvuexlguayzcammupgeskrvpthrmwqaqsdcgycdupykppiyhwzwcplivjnnvwhqkkxildtyjltklcokcrgqnnwzzeuqioyahqpuskkpbxhvzvqyhlegmoviogzwuiqahiouhnecjwysmtarjjdjqdrkljawzasriouuiqkcwwqsxifbndjmyprdozhwaoibpqrthpcjphgsfbeqrqqoqiqqdicvybzxhklehzzapbvcyleljawowluqgxxwlrymzojshlwkmzwpixgfjljkmwdtjeabgyrpbqyyykmoaqdambpkyyvukalbrzoyoufjqeftniddsfqnilxlplselqatdgjziphvrbokofvuerpsvqmzakbyzxtxvyanvjpfyvyiivqusfrsufjanmfibgrkwtiuoykiavpbqeyfsuteuxxjiyxvlvgmehycdvxdorpepmsinvmyzeqeiikajopqedyopirmhymozernxzaueljjrhcsofwyddkpnvcvzixdjknikyhzmstvbducjcoyoeoaqruuewclzqqqxzpgykrkygxnmlsrjudoaejxkipkgmcoqtxhelvsizgdwdyjwuumazxfstoaxeqqxoqezakdqjwpkrbldpcbbxexquqrznavcrprnydufsidakvrpuzgfisdxreldbqfizngtrilnbqboxwmwienlkmmiuifrvytukcqcpeqdwwucymgvyrektsnfijdcdoawbcwkkjkqwzffnuqituihjaklvthulmcjrhqcyzvekzqlxgddjoir")
    expectedResult := "gykrkyg"
    if result != expectedResult {
        t.Errorf("Got %s; expected %s", result, expectedResult)
    }
}

func longestPalindrome(s string) string {
    //fmt.Printf("%s len %d\n", s, len(s))
    var pals = make(map[string]bool)
    sLen := len(s)

    pals2 := make([][]bool, sLen + 1)
    for i := range pals2 {
        pals2[i] = make([]bool, sLen + 1)
    }

    longest := ""
    for i := 0; i < sLen; i++ {
        for j := i + 1; j <= sLen; j++ {
            if isPalindrome2(s[i:j], i, j, pals2) {
                pals[s[i:j]] = true
                pals2[i][j] = true
                if len(s[i:j]) > len(longest) {
                    longest = s[i:j]
                }
            }
            //fmt.Printf("%s pal %t\n", s[i:j], isPalindrome(s[i:j]))
        }
    }
    return longest
}

func isPalindrome2(s string, i int, j int, pals [][]bool) bool {
    if len(s) <= 1 {
        pals[i][j] = true
        return true
    } else if pals[i][j] {
        pals[i][j] = true
        return true
    } else if s[0] == s[len(s) - 1] && isPalindrome2(s[1:len(s) - 1], i +1, j - 1, pals) {
        pals[i + 1][j - 1] = true
        return true
    }

    return false
}

func isPalindrome(s string, pals map[string]bool) bool {
    if len(s) <= 1 {
        return true
    } else if pals[s] {
            return true
    } else if s[0] == s[len(s) - 1] && isPalindrome(s[1:len(s) - 1], pals) {
        pals[s] = true
        return true
    }

    return false
}
