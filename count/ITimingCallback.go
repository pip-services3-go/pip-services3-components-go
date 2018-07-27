package count

type ITimingCallback interface {
    EndTiming(name string, elapsed float32)
}
