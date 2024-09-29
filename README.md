# Pretty

Pretty printing in the terminal.

```go
fmt.Printf("%v", pretty.New(""}).WithColour(pretty.Red).WithLink("example.com") )
fmt.Printf("%v", pretty.Value{"abc"}.WithColour(pretty.Red).WithLink("example.com") )
fmt.Printf("%v", pretty.Print("abc").WithColour(pretty.Red).WithLink("example.com") )
```

```go
fmt.Printf("%v", pretty.Print("abc", WithColour(pretty.Red), WithBold(), WithLink(""))

fmt.Printf("%v", pretty.Print("abc", WithColour(pretty.Red), pretty.Bold, WithLink(""))

fmt.Printf("%v", pretty.Print("abc", "xyz", WithColour(pretty.Red), pretty.Bold, WithLink(""))

fmt.PrintLn(pretty.Print("abc", pretty.Gold))
```
