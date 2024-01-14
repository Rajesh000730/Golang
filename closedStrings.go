import "sort"
func closeStrings(word1 string, word2 string) bool {
    m := len(word1)
    n := len(word2)
    if m!=n{
        return false
    }
    v1 := make([]int, 26)
    v2 := make([]int, 26)
    fmt.Println(v1[25])
    for i:=0; i<n;i++{
        v1[word1[i]-'a']+=1
        v2[word2[i]-'a']+=1
    }
    for j:=0; j <n;j++{
        if v1[word1[j]-'a'] == 0 || v2[word1[j]-'a'] == 0 || v1[word2[j]-'a'] == 0 || v2[word2[j]-'a'] == 0{
            return false
        }
    }
    sort.Ints(v1)
    sort.Ints(v2)
    for k:=0; k < 26;k++{
        if v1[k]!=v2[k]{
            return false
        }
    }
    return true;
}
