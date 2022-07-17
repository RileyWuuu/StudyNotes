# 383 Ransom Note

Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.

Each letter in magazine can only be used once in ransomNote.

**Example 1:**

```
Input: ransomNote = "a", magazine = "b"
Output: false
```

**Example 2:**

```
Input: ransomNote = "aa", magazine = "ab"
Output: false
```

**Example 3:**

```
Input: ransomNote = "aa", magazine = "aab"
Output: true
```

**Constraints:**

- 1 <= ransomNote.length, magazine.length <= 105
- ransomNote and magazine consist of lowercase English letters.

<hr/>

## Solutions:
### 1. Use strings.Count to see whether magazine contains the(enough) letters to spell the word in ransomNote
(Runtime:4ms)
```
func canConstruct(ransomNote string, magazine string) bool{
  for _,v := range ransomNote{
    if strings.Count(ransomNote,string(v)) > strings.Count(magazine,string(v)){
      return false
     }
   }
   return true
 }
}
```
