package pages

type ClientDetailPage struct {
}

func (p *ClientDetailPage) GetContent() []byte {
	return []byte(p.content())
}

func (p *ClientDetailPage) content() string {
	return `
        <html>
            <head>
                <title>View Client</title>
            </head>
            <body>
                <b>Content: {{.Content}}</b>
            </body>
        </html>
    `
}
