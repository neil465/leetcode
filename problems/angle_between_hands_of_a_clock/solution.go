func angleClock(hour int, minutes int) float64 {
    hd := float64((hour % 12) * 30) + (float64(minutes) / float64(60)) * 30
    md := (float64(minutes) / float64(60)) * float64(360)
    return math.Min(math.Abs(hd - md), 360.0 - math.Abs(hd - md))
}
