package prints

import (
	"aws-resource-inv-tool/ec2"
	"fmt"
	"os"
	"text/tabwriter"
)

func PrintInstances(instances []*ec2.Ec2Info) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tName\tState\tMachine Type")
	for _, instance := range instances {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", instance.ID, instance.Name, instance.State, instance.MachineType)
	}
	w.Flush()
}