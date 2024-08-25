package pot

import "time"

func Headers() string {
	return "msgid \"\"\n" +
		"msgstr \"\"\n" +
		"Project-Id-Version: <project-name> 1.0\n" +
		"POT-Creation-Date: " + time.Now().Format("2006-01-02 15:04-0700") + "\n" +
		"PO-Revision-Date: YEAR-MO-DA HO:MI+ZONE\n" +
		"Last-Translator: Name <email@domain.com>\n" +
		"Language-Team: Translation Team <team@domain.com>\n" +
		"MIME-Version: 1.0\n" +
		"Content-Type: text/plain; charset=UTF-8\n" +
		"Content-Transfer-Encoding: 8bit\n" +
		"Plural-Forms: nplurals=2; plural=(n != 1);\n" +
		"\n"

}
