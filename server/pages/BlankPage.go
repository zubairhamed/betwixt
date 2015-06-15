package pages

type BlankPage struct {
}

func (p *BlankPage) GetContent() []byte {
return []byte(p.content())
}

func (p *BlankPage) content() string {
return `
        <html>
            <head>
                <title>Betwixt</title>
            </head>
            <body>
                <b>Blank!</b>
            </body>
        </html>
    `
}
