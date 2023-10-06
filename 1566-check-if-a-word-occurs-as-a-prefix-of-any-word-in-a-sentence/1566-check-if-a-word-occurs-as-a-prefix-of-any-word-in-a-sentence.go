func isPrefixOfWord(sentence string, searchWord string) int {
    index:=-1
    words := strings.Fields(sentence)
    for key,val:= range words {
        if strings.HasPrefix(val,searchWord){
            index=key+1
            break
        }
    }
    return index
}