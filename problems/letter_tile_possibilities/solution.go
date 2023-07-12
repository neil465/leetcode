var m = map[string]bool{}

func numTilePossibilities(tiles string) int {
    m = map[string]bool{}
    help([]rune(tiles), "", map[int]bool{})
    return len(m) - 1
}
func help(tiles []rune, cur string, taken map[int]bool) {
    m[cur] = true
    if len(tiles) == len(taken){
        return
    }
    for i := 0 ; i < len(tiles); i ++ {
        if !taken[i] {
            taken[i] = true
            help(tiles, cur + string(tiles[i]), taken)
            delete(taken, i)
        }
    }
    return
} 