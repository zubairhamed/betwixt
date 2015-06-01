package pages

type HomePage struct {

}

func (p *HomePage) GetContent() ([]byte) {
    return []byte(p.content())
}

func (p *HomePage) content ()(string) {
    return `
        <html>
            <head>
                <title>{{.Title}}</title>
            </head>
            <body>
                <b>Content: {{.Content}}</b>
            </body>
        </html>
    `
}