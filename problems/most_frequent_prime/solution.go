var bestval = 0
var bestcount = 0
var primes = map[int] bool {}

var cmap = map[int]int{}

func mostFrequentPrime(mat [][]int) int {
    cmap = map[int]int{}
    bestval = 0
    bestcount = 0
    if len(primes)== 0 {
        primes = SieveOfEratosthenes(1000000)
    }
    
    
    for i := range mat {
        for j := range mat[0] {
            for d := range dirs{
                re(i, j,d, 0, len(mat), len(mat[0]), mat, primes)
            }
            
        }
    }
    if bestval == 0 {
        return -1
    }
    return bestval
    
}


var dirs = [][]int{{0,1}, {0,-1}, {1, 0}, {-1, 0}, {1, 1}, {-1,1}, {1, -1}, {-1, -1}}

func re(i, j, dir int, curval int, n, m int, mat [][]int, primes map[int]bool) {
    if  primes[curval] {
        cmap[curval] ++
    }
    curcount := cmap[curval]
     if (curcount > bestcount || (curcount == bestcount && curval > bestval)) && curval > 10 && primes[curval] {
        bestval = curval
        bestcount = curcount

    }
    
    if i  < 0 ||  i  == n || j < 0 ||  j == m {
       
        return
    }
    
    re(i + dirs[dir][0], j + dirs[dir][1], dir, curval * 10 + mat[i][j], n, m , mat, primes)
    
}

func SieveOfEratosthenes(n int) map[int]bool {
	primes := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		primes[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if primes[p] == true {
			for i := p * p; i <= n; i += p {
				primes[i] = false
			}
		}
	}

    res := map[int]bool{}
	for p := 2; p <= n; p++ {
		if primes[p] {
		res[p] = true
		}
	}
	
	return res
}
