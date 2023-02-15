# Juniper
Companion library for gin-gonic to provide common functionality across multiple projects

Balically im pulling out all the duplicate code from variaus web apps into here

### Cli
Command line helper components
```go
// generate register of commands
register := juniper.CommandEntries{
    {
        Key: "run-server",
        Usage: "Runs the server",
        Run: func([]string) error {
            return nil
        }
    }
}

// generate a pretty usage document
flag.Usage = juniper.CliUsage("App Name", "Description", "bin_name", register)

// check for command
cmd := flag.String("cmd", "", "Command to run")

flag.Parse()

command := register.Find(*cmd)
if command == nil {
    panic("command not found")
}

if err := command.Run(flag.Args()); err != nil {
    panic(err
}
```

### Controller
Route base controller

### Cron
Run cron tasks on a schedule, based on the cli register
```yaml
# cron.yml
- command: clean:user-sessions
  schedule: 0 9 * 1-5 *
  # args will be passed to the commands Run func 
  args:
    - "[sum_arg]"
```
```go
// setup your commands
register := juniper.CommandEntries{
    ...
}

// lod the cron config from file
schedule, err := juniper.ParseCronSchedule("/path/to/cron.yml")
if err != nil {
    panic("bad cron config")
}

// run any pending cron tasks (based on current minute of the day)
if err := juniper.RuCronTasks(schedule, register) {
    panic(err)
}

```

### Env
Dotenv wrapper

### Gorm 
custom loggers (production/debug)

### JWT
helpers for managing jwt's and bearer tokens

### Log
Custom gin log formatters

### Response
Common response shapes for generic http errors

### Test
http test helpers

### Validation
Custom struct tag extraction for validator

