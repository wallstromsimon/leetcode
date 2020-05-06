package leetcode

import (
    "strings"
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
    longest := ""
    for i := 0; i < len(s); i++ {
        for j := i + 1; j <= len(s); j++ {
            if isPalindrome(s[i:j]) {
                if len(s[i:j]) > len(longest) {
                    longest = s[i:j]
                }
            }
            //fmt.Printf("%s pal %t\n", s[i:j], isPalindrome(s[i:j]))
        }
    }
    return longest
}

func isPalindrome(s string) bool {
    return s == reverse(s)
}

func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func lengthOfLongestSubstring2(s string) int {
    maxLen := 0
    minIndex := 0
    maxIndex := 0
    currLen := 0

    for maxIndex < len(s) {
        if !strings.Contains(s[minIndex:maxIndex], s[maxIndex:maxIndex+1]) {
            currLen = maxIndex - minIndex + 1
            maxIndex++
        } else {
            if currLen > maxLen {
                maxLen = currLen
            }
            currLen = 0
            minIndex++
        }
    }
    if currLen > maxLen {
        maxLen = currLen
    }

    return maxLen
}
