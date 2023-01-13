func categorizeBox(length int, width int, height int, mass int) string {
    bulky, heavy := length >= 10000 || width >= 10000 || height >= 10000 || length * width * height >= 1000000000, mass >= 100
    if bulky && heavy {
        return "Both"
    } else if bulky {
        return "Bulky"
    } else if heavy {
        return "Heavy"
    } else {
        return "Neither"
    }
    return ""
}