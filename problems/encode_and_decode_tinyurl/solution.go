import "crypto/sha256"
type Codec struct {
    
    m map[string]string

}


func Constructor() Codec {
    k := make(map[string]string)
    return Codec{m:k}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
    this.m[fmt.Sprintf("%s", sha256.Sum256([]byte(longUrl)))] = longUrl
    
    return fmt.Sprintf("%s", sha256.Sum256([]byte(longUrl)))
    
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
    return this.m[shortUrl]
}


/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
