# Zerolog

## Log Levels

zerolog allows for logging at the following levels (from highest to lowest):

- panic (zerolog.PanicLevel, 5)
- fatal (zerolog.FatalLevel, 4)
- error (zerolog.ErrorLevel, 3)
- warn (zerolog.WarnLevel, 2)
- info (zerolog.InfoLevel, 1)
- debug (zerolog.DebugLevel, 0)
- trace (zerolog.TraceLevel, -1)

 You can set the Global logging level to any of these options using the `SetGlobalLevel` function in the zerolog package, passing in one of the given constants above, e.g. `zerolog.InfoLevel` would be the "info" level. Whichever level is chosen, all logs with a level greater than or equal to that level will be written. To turn off logging entirely, pass the `zerolog.Disabled` constant.