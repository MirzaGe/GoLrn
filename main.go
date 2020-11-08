package main

import (
	"github.com/chrislusf/gleam/flow"
	"os"
)

func main() {
	flow.New().TextFile("/etc/paths").Partition(2).FlatMap(`
function(line)
return line:gmatch("%w+")
end `).Map(`
              function(word)
                     return word, 1
end `).ReduceBy(`
              function(x, y)
                     return x + y
end
`).Fprintf(os.Stdout, "%s,%dn").Run()
}
