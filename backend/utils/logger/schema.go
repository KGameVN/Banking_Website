// 4. observability layer: schema validation
package logger

func validate(entry LogEntry) bool {
    if entry.Level == "" || entry.Message == "" {
        return false
    }
    return true
}
