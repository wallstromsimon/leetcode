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

    longest := ""
    for i := 0; i < len(s); i++ {
        for j := i + 1; j <= len(s); j++ {
            if isPalindrome(s[i:j], pals) {
                pals[s[i:j]] = true
                if len(s[i:j]) > len(longest) {
                    longest = s[i:j]
                }
            }
            //fmt.Printf("%s pal %t\n", s[i:j], isPalindrome(s[i:j]))
        }
    }
    return longest
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
