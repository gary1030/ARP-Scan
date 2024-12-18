package report

import (
	"fmt"
	"os"
	"time"
)

type Row struct {
	IP     string
	Status string
}

func GenerateHTMLReport(filepath string, rows []Row) error {
	// check file path and ends with .html
	if filepath[len(filepath)-5:] != ".html" {
		return fmt.Errorf("file path must end with .html")
	}

	// create file
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	// write header
	file.WriteString("<html>\n")
	file.WriteString("<head>\n")
	file.WriteString("<title>IP Usage Report</title>\n")
	file.WriteString("<style>\n")
	file.WriteString("table, th, td {\n")
	file.WriteString("  border: 1px solid black;\n")
	file.WriteString("  border-collapse: collapse;\n")
	file.WriteString("}\n")
	file.WriteString("th, td {\n")
	file.WriteString("  padding: 10px;\n")
	file.WriteString("}\n")
	file.WriteString("</style>\n")
	file.WriteString("</head>\n")
	file.WriteString("<body>\n")
	file.WriteString("<h1>IP Usage Report</h1>\n")
	file.WriteString("<a>Generated at " + time.Now().Format("2006-01-02 15:04:05") + "</a>\n")
	file.WriteString("<h2></h2>\n")
	file.WriteString("<table border=\"1\">\n")
	file.WriteString("<tr>\n")
	file.WriteString("<th>IP</th>\n")
	file.WriteString("<th>Status</th>\n")
	file.WriteString("</tr>\n")

	// write rows
	for _, row := range rows {
		file.WriteString("<tr>\n")
		file.WriteString("<td>" + row.IP + "</td>\n")
		file.WriteString("<td>" + row.Status + "</td>\n")
		file.WriteString("</tr>\n")
	}

	// write footer
	file.WriteString("</table>\n")
	file.WriteString("</body>\n")
	file.WriteString("</html>\n")

	return nil
}
