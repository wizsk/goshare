package reader

import (
	"bytes"
	"fmt"
	"time"
)

type Dir struct {
	Items []Item
}

type Item struct {
	Name         string
	LastModified time.Time
	IsDir        bool
}

func (d *Dir) String() string {
	buff := new(bytes.Buffer)

	for i, itm := range d.Items {
		// if i != 0 {
		// 	fmt.Fprintf(buff, "\nname: %s, isDir: %v, lastModified: %v", itm.Name, itm.IsDir, itm.LastModified)
		// } else {
		// 	fmt.Fprintf(buff, "name: %s, isDir: %v, lastModified: %v", itm.Name, itm.IsDir, itm.LastModified)
		// }
		if i != 0 {
			fmt.Fprintf(buff, "\n%s", itm.Name)
		} else {
			fmt.Fprintf(buff, "%s", itm.Name)
		}
	}

	return buff.String()
}
