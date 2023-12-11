package ec2

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func PrintInstances(instances []*Ec2Info) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tName\tState\tMachine Type")
	for _, instance := range instances {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", instance.id, instance.name, instance.state, instance.machineType)
	}
	w.Flush()
}
