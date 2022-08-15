package utils

import (
	"github.com/evgeniums/go-mdtopdf"
)

const helveticaFontJson string = `{"Tp":"TrueType","Name":"ArialMT","Desc":{"Ascent":728,"Descent":-210,"CapHeight":728,"Flags":32,"FontBBox":{"Xmin":-665,"Ymin":-325,"Xmax":2028,"Ymax":1037},"ItalicAngle":0,"StemV":70,"MissingWidth":750},"Up":-106,"Ut":73,"Cw":[750,750,750,750,750,750,750,750,750,750,750,750,750,750,750,750,750,750,750,750,750,750,750,750,750,750,750,750,750,750,750,750,278,278,355,556,556,889,667,191,333,333,389,584,278,333,278,278,556,556,556,556,556,556,556,556,556,556,278,278,584,584,584,556,1015,667,667,722,722,667,611,778,722,278,500,667,556,833,722,778,667,778,722,667,611,722,667,944,667,667,611,278,278,278,469,556,333,556,556,500,556,556,278,556,556,222,222,500,222,833,556,556,556,556,333,500,278,556,500,722,500,500,500,334,260,334,584,750,865,542,222,365,333,1000,556,556,556,1000,1057,333,1010,583,854,719,556,222,222,333,333,350,556,1000,750,1000,906,333,813,438,556,552,278,635,500,500,556,489,260,556,667,737,719,556,584,333,737,278,400,549,278,222,411,576,537,278,556,1073,510,556,222,667,500,278,667,656,667,542,677,667,923,604,719,719,583,656,833,722,778,719,667,722,611,635,760,667,740,667,917,938,792,885,656,719,1010,722,556,573,531,365,583,556,669,458,559,559,438,583,688,552,556,542,556,500,458,500,823,500,573,521,802,823,625,719,521,510,750,542],"Enc":"cp1251","Diff":"128 /afii10051 /afii10052 131 /afii10100 136 /Euro 138 /afii10058 140 /afii10059 /afii10061 /afii10060 /afii10145 /afii10099 152 /.notdef 154 /afii10106 156 /afii10107 /afii10109 /afii10108 /afii10193 161 /afii10062 /afii10110 /afii10057 165 /afii10050 168 /afii10023 170 /afii10053 175 /afii10056 178 /afii10055 /afii10103 /afii10098 184 /afii10071 /afii61352 /afii10101 188 /afii10105 /afii10054 /afii10102 /afii10104 /afii10017 /afii10018 /afii10019 /afii10020 /afii10021 /afii10022 /afii10024 /afii10025 /afii10026 /afii10027 /afii10028 /afii10029 /afii10030 /afii10031 /afii10032 /afii10033 /afii10034 /afii10035 /afii10036 /afii10037 /afii10038 /afii10039 /afii10040 /afii10041 /afii10042 /afii10043 /afii10044 /afii10045 /afii10046 /afii10047 /afii10048 /afii10049 /afii10065 /afii10066 /afii10067 /afii10068 /afii10069 /afii10070 /afii10072 /afii10073 /afii10074 /afii10075 /afii10076 /afii10077 /afii10078 /afii10079 /afii10080 /afii10081 /afii10082 /afii10083 /afii10084 /afii10085 /afii10086 /afii10087 /afii10088 /afii10089 /afii10090 /afii10091 /afii10092 /afii10093 /afii10094 /afii10095 /afii10096 /afii10097","File":"helvetica_1251.z","Size1":0,"Size2":0,"OriginalSize":275572,"I":0,"N":0,"DiffN":0}`

func CreatePdf(content []byte, fileName string, fontPath string) error {

	pf := mdtopdf.NewPdfRenderer("", "", fileName, "")
	pf.Pdf.SetFontLocation(fontPath)
	tr := pf.Pdf.UnicodeTranslatorFromDescriptor("cp1251")

	fontNormal := "Roboto-Regular"
	pf.Pdf.AddFont(fontNormal, "", fontNormal+".json")
	pf.Normal = mdtopdf.Styler{Font: fontNormal, Style: "", Size: 12, Spacing: 2}
	pf.TBody = mdtopdf.Styler{Font: fontNormal, Style: "", Size: 12, Spacing: 2,
		TextColor: mdtopdf.Color{0, 0, 0}, FillColor: mdtopdf.Color{240, 240, 240}}

	fontBold := "Roboto-Bold"
	pf.Pdf.AddFont("roboto-regular", "B", fontBold+".json")
	pf.Pdf.AddFont(fontBold, "", fontBold+".json")
	fontItalic := "Roboto-Italic"
	pf.Pdf.AddFont("roboto-regular", "I", fontItalic+".json")
	pf.Pdf.AddFont(fontItalic, "", fontItalic+".json")

	pf.H1 = mdtopdf.Styler{Font: fontBold, Style: "", Size: 14, Spacing: 2}
	pf.H2 = mdtopdf.Styler{Font: fontBold, Style: "", Size: 12, Spacing: 2}
	pf.H3 = mdtopdf.Styler{Font: fontBold, Style: "", Size: 12, Spacing: 2}
	pf.H4 = mdtopdf.Styler{Font: fontBold, Style: "", Size: 12, Spacing: 2}
	pf.H5 = mdtopdf.Styler{Font: fontBold, Style: "", Size: 12, Spacing: 2}
	pf.H6 = mdtopdf.Styler{Font: fontBold, Style: "", Size: 12, Spacing: 2}
	pf.THeader = mdtopdf.Styler{Font: fontBold, Style: "", Size: 12, Spacing: 2, TextColor: mdtopdf.Color{0, 0, 0}, FillColor: mdtopdf.Color{180, 180, 180}}
	pf.LoadDefaultSTyle()

	pf.Blockquote = mdtopdf.Styler{Font: fontNormal, Style: "", Size: 12, Spacing: 2,
		TextColor: mdtopdf.Color{0, 0, 0}, FillColor: mdtopdf.Color{255, 255, 255}}
	pf.Backtick = mdtopdf.Styler{Font: fontNormal, Style: "", Size: 12, Spacing: 2,
		TextColor: mdtopdf.Color{0, 0, 0}, FillColor: mdtopdf.Color{255, 255, 255}}

	err := pf.Process([]byte(tr(string(content))))
	return err
}