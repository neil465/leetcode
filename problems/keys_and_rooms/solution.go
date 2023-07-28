func canVisitAllRooms(rooms [][]int) bool {
    visited := map[int]bool{0 : true}
    q, p := rooms[0], 0
    for len(q) > 0 {
        p, q = q[0], q[1:]
        if !visited[p] { q = append(q, rooms[p]...) }
        visited[p] = true
    }
    return len(visited) == len(rooms)
}